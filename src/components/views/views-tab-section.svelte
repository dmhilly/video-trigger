<script lang="ts">
  import ViewTab from "./view-tab.svelte";
  import { Views } from "./views";

  interface Props {
    onselect: (view: Views) => void;
    selected: Views;
  }
  const { onselect, selected }: Props = $props();
</script>

<div class="flex flex-col gap-2 bg-white p-2 h-fit rounded-2xl w-45">
  <ViewTab
    text="Dashboard"
    icon="info"
    selected={selected === Views.Dashboard}
    onclick={() => onselect(Views.Dashboard)}
  />
  <ViewTab
    text="Camera"
    icon="camera"
    selected={selected === Views.Camera}
    onclick={() => onselect(Views.Camera)}
  />
  <ViewTab
    text="Videos"
    icon="video"
    selected={selected === Views.Videos}
    onclick={() => onselect(Views.Videos)}
  />
</div>
