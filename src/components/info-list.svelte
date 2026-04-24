<script lang="ts">
  import Datetime from "./datetime.svelte";

  interface Props {
    items: Record<string, string | number | Date | undefined>;
  }
  const { items }: Props = $props();
</script>

<div class="flex flex-col gap-2 flex-1">
  {#each Object.entries(items) as [key, value]}
    {#if value instanceof Date}
      <span>
        <span class="font-semibold">{key}</span>: <Datetime date={value} />
      </span>
    {:else}
      <span><span class="font-semibold">{key}</span>: {value}</span>
    {/if}
  {/each}
</div>
