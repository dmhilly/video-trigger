<script lang="ts">
  import { Status, statusToVariant } from "../lib/types";
  import Badge from "./badge.svelte";

  interface Props {
    name: string;
    status: Status;
    error?: string;
  }
  const { name, status, error }: Props = $props();
</script>

<div class="flex gap-3 bg-white py-2 px-3 rounded-xl w-fit items-center">
  <div class="flex gap-1 items-center">
    <span>Machine: </span>
    <Badge text={name == "" ? "-" : name} variant="neutral" />
  </div>
  <div class="flex gap-1 items-center">
    <span>Status: </span>
    <Badge text={status} variant={statusToVariant[status]} />
    {#if status === Status.Error}
      <span class="text-red-600">{error}</span>
    {/if}
  </div>
</div>
