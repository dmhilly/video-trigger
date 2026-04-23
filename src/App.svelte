<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";

  import Header from "./components/header.svelte";
  import MachineStatus from "./components/machine-status.svelte";
  import CameraFeedView from "./components/views/camera-feed-view.svelte";
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
  let mediaStream = $state<MediaStream | undefined>(undefined);
  let motionClassifications = $state<VIAM.Detection[]>([]);
  let awakeClassifications = $state<VIAM.Classification[]>([]);

  let pollInterval: ReturnType<typeof setInterval> | undefined;

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
          [motionClassifications, awakeClassifications] = await Promise.all([
            motionDetector.getClassificationsFromCamera(cameraName, 1),
            awakeClassifier.getClassificationsFromCamera(cameraName, 2),
          ]);
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

<div class="p-8 flex flex-col gap-8">
  <div class="flex gap-4 items-center justify-between">
    <Header text="Baby Dashboard" />
    <MachineStatus name={machineName} {status} {error} />
  </div>

  <ViewsTabSection onselect={selectTab} selected={view} />

  {#if view === Views.Camera}
    <CameraFeedView
      {mediaStream}
      {motionClassifications}
      {awakeClassifications}
    />
  {:else if view === Views.Test}
    <p>Test View</p>
  {:else if view === Views.Videos}
    <VideosView />
  {/if}
</div>
