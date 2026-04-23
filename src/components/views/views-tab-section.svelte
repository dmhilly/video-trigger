<script lang="ts">
  import ViewTab from "./view-tab.svelte";
  import { Views } from "./views";

  interface Props {
    onselect: (view: Views) => void;
    selected: Views;
  }
  const { onselect, selected }: Props = $props();
</script>

<div class="flex gap-4 self-center">
  <ViewTab
    text="camera"
    icon="camera"
    selected={selected === Views.Camera}
    onclick={() => onselect(Views.Camera)}
  />
  <ViewTab
    text="test"
    icon="info"
    selected={selected === Views.Test}
    onclick={() => onselect(Views.Test)}
  />
</div>
