<script lang="ts">
  import type { ClassValue } from "svelte/elements";
  import { Size } from "../lib/types";
  import { IconByName } from "./icon";

  interface Props {
    name: string;
    size?: Size;
    class?: ClassValue;
  }
  const { name, size = Size.Medium, class: extraClasses }: Props = $props();

  const icon = $derived(IconByName[name]);

  const sizeMap: Record<Size, string> = {
    sm: "16",
    md: "24",
    lg: "36",
    xl: "48",
  };
</script>

<svg
  width={sizeMap[size]}
  height={sizeMap[size]}
  viewBox="0 0 24 24"
  class={extraClasses}
>
  {@html icon.body}
</svg>
