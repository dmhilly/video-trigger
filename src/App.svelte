<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";
  import MachineStatus from "./components/machine-status.svelte";
  import Header from "./components/header.svelte";
  import { Status } from "./lib/types";
  import Section from "./components/section.svelte";

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
  let cameraClient = $state<VIAM.CameraClient | undefined>(undefined);

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

      // Connect directly to the machine (not the app API)
      robotClient = await viamClient.connectToMachine({ host });
      status = Status.Connected;

      cameraClient = new VIAM.CameraClient(robotClient, "cam");
    } catch (err) {
      status = Status.Error;
      error = `${err}`;
    }
  }

  connect();
</script>

<div class="p-8 flex flex-col gap-8">
  <div class="flex gap-4 items-center">
    <Header text="Baby Dashboard" />
    <MachineStatus name={machineName} {status} {error} />
  </div>

  <Section title="Current:">
    {#snippet content()}{/snippet}
  </Section>
</div>
