<script lang="ts">
  import { Size } from "../lib/types";
  import Icon from "./icon.svelte";

  interface Props {
    icon: string;
    onclick?: () => void;
    size?: Size;
    class?: string;
  }
  const {
    icon,
    onclick,
    size = Size.Small,
    class: extraClasses,
  }: Props = $props();
</script>

<button {onclick} class={["hover:cursor-pointer", extraClasses]}>
  <Icon name={icon} {size} />
</button>
