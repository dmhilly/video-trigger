<script lang="ts">
  import type { Snippet } from "svelte";

  interface Props {
    text: string;
    position?: "top" | "bottom" | "left" | "right";
    delay?: number;
    visible?: boolean;
    class?: string;
    children?: Snippet;
  }
  const {
    text,
    position = "bottom",
    delay = 500,
    visible: visibleProp,
    class: extraClasses,
    children,
  }: Props = $props();

  let internalVisible = $state(false);
  let timeout: ReturnType<typeof setTimeout>;

  function show() {
    timeout = setTimeout(() => {
      internalVisible = true;
    }, delay);
  }

  function hide() {
    clearTimeout(timeout);
    internalVisible = false;
  }

  const positionClasses: Record<string, string> = {
    top: "bottom-full left-1/2 -translate-x-1/2 mb-1",
    bottom: "top-full left-1/2 -translate-x-1/2 mt-1",
    left: "right-full top-1/2 -translate-y-1/2 mr-1",
    right: "left-full top-1/2 -translate-y-1/2 ml-1",
  };

  const isVisible = $derived(
    visibleProp !== undefined ? visibleProp : internalVisible,
  );
</script>

<div
  class="relative flex items-center"
  role="tooltip"
  onmouseenter={show}
  onmouseleave={hide}
>
  {@render children?.()}
  {#if isVisible}
    <span
      class={[
        "absolute z-10 bg-gray-900 text-white text-xs rounded p-2 pointer-events-none",
        positionClasses[position],
        extraClasses,
      ]}
    >
      {text}
    </span>
  {/if}
</div>
