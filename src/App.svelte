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
  let machineName = $state("");
  let robotClient = $state<VIAM.RobotClient | undefined>(undefined);
  let streamClient = $state<VIAM.StreamClient | undefined>(undefined);
  let dataClient = $state<VIAM.DataClient | undefined>(undefined);
  let mediaStream = $state<MediaStream | undefined>(undefined);
  let viamClient: Awaited<ReturnType<typeof VIAM.createViamClient>>;

  // machine status
  let status = $state(Status.Connecting);
  let error = $state("");
  // reconnect
  const MAX_DELAY = 60000;
  let retryDelay = 5000;
  let currentRetryTimeout: ReturnType<typeof setTimeout> | undefined;
  let retryListenersActive = false;

  // motion detection
  const TWO_MINUTES = 2 * 60 * 1000;
  let movementDetected = $state(false);
  let lastMotionTime: Date | undefined = undefined;
  let lastMotionLocal = $state<Date | undefined>(undefined);
  let awakeData = $state<VIAM.JsonValue>({});
  let pollInterval: ReturnType<typeof setInterval> | undefined;

  // video data
  const MAX_VIDEO_COUNT = 50;
  let videos = $state<Video[]>([]);
  let wakeUpTimes = $state<Date[]>([]);
  const filter = new VIAM.dataApi.Filter({
    mimeType: ["video/mp4"],
    tagsFilter: { tags: ["awake"] },
  });

  async function fetchVideos() {
    if (dataClient) {
      const { data } = await dataClient.binaryDataByFilter(
        filter,
        MAX_VIDEO_COUNT,
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
    }
  }

  function startPollingForData(robotClient: VIAM.RobotClient) {
    const motionDetector = new VIAM.VisionClient(
      robotClient,
      "motion-detector",
    );
    const awakeClassifier = new VIAM.SensorClient(
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
          awakeClassifier.getReadings(),
        ]);
        awakeData = awakeResult;

        if (motionResults.length > 0 && motionResults[0].confidence > 0.001) {
          lastMotionTime = new Date();
          lastMotionLocal = lastMotionTime;
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
  }

  // Set retry listeners for when user is moving tabs or changes connection
  function setRetryListeners(active: boolean) {
    if (active === retryListenersActive) return;
    if (active) {
      window.addEventListener("online", handleRetryFromListener);
      document.addEventListener("visibilitychange", handleRetryFromListener);
    } else {
      window.removeEventListener("online", handleRetryFromListener);
      document.removeEventListener("visibilitychange", handleRetryFromListener);
    }
    retryListenersActive = active;
  }

  // Attempt connecting after resetting the retry delay
  // Cancel background reconnect attempts when user's not connected or on tab
  function handleRetryFromListener() {
    if (document.hidden) {
      if (currentRetryTimeout) clearTimeout(currentRetryTimeout);
      return;
    }
    if (currentRetryTimeout) clearTimeout(currentRetryTimeout);
    retryDelay = 5000;
    connect();
  }

  async function connect() {
    try {
      status = Status.Connecting;

      const machine = await viamClient.appClient.getRobot(machineId);
      machineName = machine?.name ?? "";

      if (machine?.onlineState === 2) {
        status = Status.Offline;
        currentRetryTimeout = setTimeout(connect, retryDelay);
        retryDelay = Math.min(retryDelay * 2, MAX_DELAY);
        setRetryListeners(true);
        return;
      }

      retryDelay = 5000;
      setRetryListeners(false);

      // Connect directly to the machine (not the app API)
      robotClient = await viamClient.connectToMachine({ host });
      status = Status.Connected;

      robotClient?.on(VIAM.MachineConnectionEvent.DISCONNECTED, () => {
        status = Status.Offline;
      });

      robotClient?.on(
        VIAM.MachineConnectionEvent.CONNECTED,
        () => (status = Status.Connected),
      );
      startPollingForData(robotClient);

      streamClient = new VIAM.StreamClient(robotClient);
      mediaStream = await streamClient.getStream(cameraName);
    } catch (err) {
      status = Status.Error;
      error = `${err}`;
      currentRetryTimeout = setTimeout(connect, retryDelay);
      retryDelay = Math.min(retryDelay * 2, MAX_DELAY);
      setRetryListeners(true);
    }
  }

  async function init() {
    viamClient = await VIAM.createViamClient({
      serviceHost: "https://app.viam.com",
      credentials: {
        type: "api-key",
        payload: apiKeySecret,
        authEntity: apiKeyId,
      },
    });

    dataClient = viamClient.dataClient;
    await fetchVideos();
    setInterval(fetchVideos, 60000);

    connect();
  }

  init();

  let view: Views = $state(Views.Info);
  const selectTab = (newView: Views) => (view = newView);
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
        {status}
        {mediaStream}
        {awakeData}
        {movementDetected}
        {lastMotionLocal}
      />
    {:else if view === Views.Info}
      <InfoView {dataClient} {wakeUpTimes} {videos} />
    {:else if view === Views.Videos}
      <VideosView {dataClient} {videos} maxVideoCount={MAX_VIDEO_COUNT} />
    {/if}
  </div>
</div>
