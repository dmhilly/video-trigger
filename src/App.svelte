<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";

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
  let status = $state("Connecting...");
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
      machineName = machine.name;
      status = "Connected";
    } catch (err) {
      status = `Error: ${err}`;
    }
  }

  connect();
</script>

<h1>Video Trigger Dashboard</h1>
<p>Status: {status}</p>
{#if machineName}
  <p>Machine: {machineName}</p>
{/if}
