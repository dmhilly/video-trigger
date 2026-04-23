<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";

  import Header from "./components/header.svelte";
  import MachineStatus from "./components/machine-status.svelte";
  import CameraFeedView from "./components/views/camera-feed-view.svelte";
  import InfoView from "./components/views/info-view.svelte";
  import type { Video } from "./components/views/video";
  import VideosView from "./components/views/videos-view.svelte";
  import { Views } from "./components/views/views";
  import ViewsTabSection from "./components/views/views-tab-section.svelte";
  import { Status } from "./lib/types";

  // Receive the credentials passed from main.ts
  let {
    apiKeyId,
    apiKeySecret,
    machineId,
    host,
  }: {
    apiKeyId: string;
    apiKeySecret: string;
    machineId: string;
    host: string;
  } = $props();

  const cameraName = "replay-camera";

  let status = $state(Status.Connecting);
  let error = $state("");
  let machineName = $state("");
  let robotClient = $state<VIAM.RobotClient | undefined>(undefined);
  let streamClient = $state<VIAM.StreamClient | undefined>(undefined);
  let dataClient = $state<VIAM.DataClient | undefined>(undefined);
  let mediaStream = $state<MediaStream | undefined>(undefined);

  let awakeClassifications = $state<VIAM.Classification[]>([]);

  let movementDetected = $state(false);
  let lastMotionTime: Date | undefined = undefined;
  let lastMotionLocal = $state("None");
  const TWO_MINUTES = 2 * 60 * 1000;

  const maxVideoCount = 50;
  let videos = $state<Video[]>([]);
  let wakeUpTimes = $state<Date[]>([]);

  let pollInterval: ReturnType<typeof setInterval> | undefined;

  const filter = new VIAM.dataApi.Filter({
    mimeType: ["video/mp4"],
    tagsFilter: { tags: ["awake"] },
  });

  async function connect() {
    try {
      status = Status.Connecting;

      const viamClient = await VIAM.createViamClient({
        serviceHost: "https://app.viam.com",
        credentials: {
          type: "api-key",
          payload: apiKeySecret,
          authEntity: apiKeyId,
        },
      });

      dataClient = viamClient.dataClient;
      const { data } = await dataClient.binaryDataByFilter(
        filter,
        maxVideoCount,
        VIAM.dataApi.Order.DESCENDING,
        undefined,
        false,
      );
      videos = data.map((item) => ({
        name: item.metadata?.fileName,
        binaryDataId: item.metadata?.binaryDataId ?? "",
        url: undefined,
        timestamp: item.metadata?.timeRequested,
        size: item.metadata?.fileSizeBytes,
        loaded: false,
      }));
      wakeUpTimes = videos
        .map((v) => v.timestamp?.toDate())
        .filter((d): d is Date => d !== undefined);

      const machine = await viamClient.appClient.getRobot(machineId);
      machineName = machine?.name ?? "";

      if (machine?.onlineState === 2) {
        status = Status.Offline;
        setTimeout(connect, 5000);
        return;
      }

      // Connect directly to the machine (not the app API)
      robotClient = await viamClient.connectToMachine({ host });
      status = Status.Connected;

      robotClient?.on(VIAM.MachineConnectionEvent.DISCONNECTED, () => {
        setTimeout(connect, 5000);
        status = Status.Offline;
      });

      robotClient?.on(
        VIAM.MachineConnectionEvent.CONNECTED,
        () => (status = Status.Connected),
      );

      streamClient = new VIAM.StreamClient(robotClient);
      mediaStream = await streamClient.getStream(cameraName);

      const motionDetector = new VIAM.VisionClient(
        robotClient,
        "motion-detector",
      );
      const awakeClassifier = new VIAM.VisionClient(
        robotClient,
        "awake-classifier",
      );

      let polling = false;
      if (pollInterval) clearInterval(pollInterval);
      pollInterval = setInterval(async () => {
        if (polling) return;
        polling = true;
        try {
          const [motionResults, awakeResult] = await Promise.all([
            motionDetector.getClassificationsFromCamera(cameraName, 1),
            awakeClassifier.getClassificationsFromCamera(cameraName, 1),
          ]);
          awakeClassifications = awakeResult;

          if (motionResults.length > 0 && motionResults[0].confidence > 0.001) {
            lastMotionTime = new Date();
            lastMotionLocal = lastMotionTime.toLocaleString();
            movementDetected = true;
          } else if (
            movementDetected &&
            lastMotionTime &&
            Date.now() - lastMotionTime.getTime() >= TWO_MINUTES
          ) {
            movementDetected = false;
          }
        } catch (err) {
          console.error("Detection error:", err);
        } finally {
          polling = false;
        }
      }, 500);
    } catch (err) {
      status = Status.Error;
      error = `${err}`;
      setTimeout(connect, 5000);
    }
  }

  connect();

  let view: Views = $state(Views.Camera);

  const selectTab = (newView: Views) => {
    view = newView;
  };
</script>

<div class="min-h-screen bg-linear-to-b from-blue-50 to-pink-50">
  <div class="p-8 mx-auto flex flex-col gap-8 max-w-228">
    <div class="flex gap-4 items-center justify-between">
      <Header text="Baby Dashboard" />
      <MachineStatus name={machineName} {status} {error} />
    </div>

    <ViewsTabSection onselect={selectTab} selected={view} />

    {#if view === Views.Camera}
      <CameraFeedView
        {mediaStream}
        {awakeClassifications}
        {movementDetected}
        {lastMotionLocal}
      />
    {:else if view === Views.Info}
      <InfoView {wakeUpTimes} />
    {:else if view === Views.Videos}
      <VideosView {dataClient} {videos} {maxVideoCount} />
    {/if}
  </div>
</div>
