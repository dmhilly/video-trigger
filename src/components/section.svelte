<script lang="ts">
  import type { Snippet } from "svelte";
  import Heading from "./heading.svelte";

  interface Props {
    title: string;
    content: Snippet;
    class?: string;
    contentClasses?: string;
  }
  const {
    title,
    content,
    class: extraClasses,
    contentClasses,
  }: Props = $props();
</script>

<div class={["flex flex-col gap-4  p-4 bg-white rounded-xl", extraClasses]}>
  <Heading text={title} variant="secondary" />
  <div class={["flex gap-2", contentClasses]}>
    {@render content()}
  </div>
</div>
