<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";
  import MachineStatus from "./components/machine-status.svelte";
  import Header from "./components/header.svelte";
  import { Status } from "./lib/types";
  import Section from "./components/section.svelte";
  import InfoList from "./components/info-list.svelte";

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
  let videoElement: HTMLVideoElement;

  async function connect() {
    try {
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
        return;
      }

      // Connect directly to the machine (not the app API)
      robotClient = await viamClient.connectToMachine({ host });
      status = Status.Connected;

      streamClient = new VIAM.StreamClient(robotClient);
      const mediaStream = await streamClient.getStream("tapo-camera");

      videoElement.srcObject = mediaStream;
    } catch (err) {
      status = Status.Error;
      error = `${err}`;
    }
  }

  connect();

  // Example data
  const babyStatus = $state<"awake" | "asleep">("awake");
  const lastAwake = new Date(2026, 3, 9, 19, 41, 0);
  const wokeUp = new Date(2026, 3, 10, 4, 20, 12);

  const lastAwakeLocal = lastAwake.toLocaleString();
  const wokeUpLocal = wokeUp.toLocaleString();
</script>

<div class="p-8 flex flex-col gap-8">
  <div class="flex gap-4 items-center">
    <Header text="Baby Dashboard" />
    <MachineStatus name={machineName} {status} {error} />
  </div>

  <Section title="Camera feed:">
    {#snippet content()}
      <video
        bind:this={videoElement}
        autoplay
        playsinline
        muted
        class="flex-1 min-w-0"
      ></video>
      <InfoList
        items={{
          Status: babyStatus,
          "Last Awake": lastAwakeLocal,
          "Noise Level": 0,
        }}
      />
    {/snippet}
  </Section>

  <Section title="Activity:">
    {#snippet content()}
      <InfoList
        items={babyStatus === "asleep"
          ? {
              "Fell Asleep at": lastAwakeLocal,
            }
          : { "Awoke at": wokeUpLocal }}
      />
    {/snippet}
  </Section>
</div>
