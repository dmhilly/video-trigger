<script lang="ts">
  import type { Snippet } from "svelte";
  import Header from "./header.svelte";

  interface Props {
    title: string;
    content: Snippet;
    class?: string;
  }
  const { title, content, class: extraClasses }: Props = $props();
</script>

<div
  class={["flex flex-col gap-4 border p-4 bg-white rounded-xl", extraClasses]}
>
  <Header text={title} variant="secondary" />
  <div class="flex gap-2">
    {@render content()}
  </div>
</div>
