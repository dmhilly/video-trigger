<script lang="ts">
  import type { Snippet } from "svelte";

  interface Props {
    text: string;
    position?: "top" | "bottom" | "left" | "right";
    delay?: number;
    children: Snippet;
  }
  const { text, position = "bottom", delay = 500, children }: Props = $props();

  let visible = $state(false);
  let timeout: ReturnType<typeof setTimeout>;

  function show() {
    timeout = setTimeout(() => {
      visible = true;
    }, delay);
  }

  function hide() {
    clearTimeout(timeout);
    visible = false;
  }

  const positionClasses: Record<string, string> = {
    top: "bottom-full left-1/2 -translate-x-1/2 mb-1",
    bottom: "top-full left-1/2 -translate-x-1/2 mt-1",
    left: "right-full top-1/2 -translate-y-1/2 mr-1",
    right: "left-full top-1/2 -translate-y-1/2 ml-1",
  };
</script>

<div
  class="relative group inline-block"
  role="tooltip"
  onmouseenter={show}
  onmouseleave={hide}
>
  {@render children()}
  {#if visible}
    <span
      class={[
        "absolute z-10 bg-gray-900 text-white text-xs rounded p-2 pointer-events-none",
        positionClasses[position],
      ]}
    >
      {text}
    </span>
  {/if}
</div>
