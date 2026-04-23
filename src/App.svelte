<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";
  import MachineStatus from "./components/machine-status.svelte";
  import Header from "./components/header.svelte";
  import { Status } from "./lib/types";
  import CameraFeed from "./views/camera-feed.svelte";
  import Tabs from "./components/tabs.svelte";
  import Tab from "./components/tab.svelte";

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

  let status = $state(Status.Connecting);
  let error = $state("");
  let machineName = $state("");
  let robotClient = $state<VIAM.RobotClient | undefined>(undefined);
  let streamClient = $state<VIAM.StreamClient | undefined>(undefined);
  let mediaStream = $state<MediaStream | undefined>(undefined);

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
      mediaStream = await streamClient.getStream("replay-camera");
    } catch (err) {
      status = Status.Error;
      error = `${err}`;
      setTimeout(connect, 5000);
    }
  }

  connect();
</script>

<div class="p-8 flex flex-col gap-8">
  <div class="flex gap-4 items-center justify-between">
    <Header text="Baby Dashboard" />
    <MachineStatus name={machineName} {status} {error} />
  </div>

  <Tabs />

  <CameraFeed {mediaStream} />
</div>
