<script lang="ts">
  import ViewTab from "./view-tab.svelte";
  import { Views } from "./views";

  interface Props {
    onselect: (view: Views) => void;
    selected: Views;
  }
  const { onselect, selected }: Props = $props();
</script>

<div class="flex flex-col gap-2 bg-white p-2 h-fit rounded-2xl">
  <ViewTab
    text="Info"
    icon="info"
    selected={selected === Views.Info}
    onclick={() => onselect(Views.Info)}
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
