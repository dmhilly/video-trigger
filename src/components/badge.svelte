<script lang="ts">
  import type { ClassValue } from "svelte/elements";
  import { Status } from "../lib/types";

  interface Props {
    text: string;
    variant: "warning" | "success" | "error" | "neutral";
  }
  const { text, variant }: Props = $props();

  const classes: Record<string, ClassValue> = {
    warning: "bg-yellow-100 text-yellow-600 border border-yellow-600 ",
    success: "bg-green-100 text-green-600 border border-green-600 ",
    error: "bg-red-100 text-red-600 border border-red-600 ",
    neutral: "bg-gray-100 text-gray-600 border border-gray-600 ",
  };
</script>

<span class={["px-2 rounded-xl", classes[variant]]}>{text}</span>
