<script lang="ts">
  import ViewTab from "./view-tab.svelte";
  import { Views } from "./views";

  interface Props {
    onselect: (view: Views) => void;
    selected: Views;
  }
  const { onselect, selected }: Props = $props();
</script>

<div class="flex gap-16 self-center">
  <ViewTab
    text="camera"
    icon="camera"
    selected={selected === Views.Camera}
    onclick={() => onselect(Views.Camera)}
  />
  <ViewTab
    text="Info"
    icon="info"
    selected={selected === Views.Info}
    onclick={() => onselect(Views.Info)}
  />
  <ViewTab
    text="videos"
    icon="video"
    selected={selected === Views.Videos}
    onclick={() => onselect(Views.Videos)}
  />
</div>
