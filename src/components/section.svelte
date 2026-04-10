<script lang="ts">
  import type { Snippet } from "svelte";
  import Header from "./header.svelte";

  interface Props {
    title: string;
    content: Snippet;
  }
  const { title, content }: Props = $props();
</script>

<div class="flex flex-col gap-4 border p-4">
  <Header text={title} variant="secondary" />
  <div class="flex gap-2">
    {@render content()}
  </div>
</div>
