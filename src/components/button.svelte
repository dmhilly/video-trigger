<script lang="ts">
  import { Size } from "../lib/types";
  import Icon from "./icon.svelte";

  interface Props {
    text: string;
    icon?: string;
    class?: string;
    onclick: () => void;
  }
  const { text, icon, class: extraClasses, onclick }: Props = $props();
</script>

<button
  class={[
    "bg-black text-white px-8 py-2 rounded-xl hover:cursor-pointer",
    extraClasses,
  ]}
  {onclick}
>
  <div class="flex gap-1 items-center justify-center">
    <span>{text}</span>
    {#if icon}
      <Icon name={icon} size={Size.Small} />
    {/if}
  </div>
</button>
