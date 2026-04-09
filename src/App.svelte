<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";
  import MachineStatus from "./components/machine-status.svelte";
  import Header from "./components/header.svelte";
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

  // Reactive state for the UI
  let status = $state(Status.Connecting);
  let error = $state("");
  let machineName = $state("");

  // Connect to Viam when the component loads
  async function connect() {
    try {
      const client = await VIAM.createViamClient({
        serviceHost: "https://app.viam.com",
        credentials: {
          type: "api-key",
          payload: apiKeySecret,
          authEntity: apiKeyId,
        },
      });
      const machine = await client.appClient.getRobot(machineId);
      machineName = machine?.name ?? "";
      status = Status.Connected;
    } catch (err) {
      status = Status.Error;
      error = `${err}`;
    }
  }

  connect();
</script>

<div class="p-4">
  <div class="flex gap-4 items-center">
    <Header text="Baby Dashboard" />
    <MachineStatus name={machineName} {status} {error} />
  </div>
</div>
