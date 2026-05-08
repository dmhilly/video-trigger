<script lang="ts">
  import type { Snippet } from "svelte";
  import Heading from "./heading.svelte";
  import Icon from "./icon.svelte";

  interface Props {
    title: string;
    content: Snippet;
    class?: string;
    icon?: string;
    iconClass?: string;
    contentClasses?: string;
  }
  const {
    title,
    content,
    class: extraClasses,
    icon,
    iconClass,
    contentClasses,
  }: Props = $props();
</script>

<div class={["flex flex-col gap-4 p-4 bg-white rounded-xl", extraClasses]}>
  <div class="flex gap-2 items-center">
    {#if icon}<Icon name={icon} class={iconClass} />{/if}
    <Heading text={title} variant="secondary" />
  </div>
  <div class={["flex gap-2", contentClasses]}>
    {@render content()}
  </div>
</div>
