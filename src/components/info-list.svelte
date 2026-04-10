<script lang="ts">
  interface Props {
    items: Record<string, string | number>;
  }
  const { items }: Props = $props();
</script>

<div class="flex flex-col gap-2 flex-1">
  {#each Object.entries(items) as [key, value]}
    <span>{key}: {value}</span>
  {/each}
</div>
