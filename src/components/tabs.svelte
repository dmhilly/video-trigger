<script lang="ts">
  import Tab from "./tab.svelte";

  const tabs = ["camera"];

  let selected = $state(tabs[0]);
</script>

<div class="flex gap-4 self-center">
  <Tab text="camera" icon="camera" />
</div>
