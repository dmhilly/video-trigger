<script lang="ts">
  import { Size } from "../lib/types";
  import Icon from "./icon.svelte";

  interface Props {
    icon: string;
    onclick?: () => void;
    class?: string;
  }
  const { icon, onclick, class: extraClasses }: Props = $props();
</script>

<button {onclick} class={["hover:cursor-pointer", extraClasses]}>
  <Icon name={icon} size={Size.Small} />
</button>
