<script lang="ts">
  import Section from "../section.svelte";
  import { toDateString, toTimeString } from "../../lib/helpers";

  interface Props {
    wakeUpTimes: Date[];
  }
  const { wakeUpTimes }: Props = $props();

  const wakeUpTimesByDate = $derived.by(() => {
    const grouped: Record<string, string[]> = {};
    for (const date of wakeUpTimes) {
      const dateKey = toDateString(date);
      const timeValue = toTimeString(date);
      if (!grouped[dateKey]) {
        grouped[dateKey] = [];
      }
      grouped[dateKey].push(timeValue);
    }
    return grouped;
  });
</script>

<Section title="Wake Up Times">
  {#snippet content()}
    <div class="flex flex-col gap-4 w-full">
      {#each Object.entries(wakeUpTimesByDate) as [date, times]}
        <div class="flex gap-4">
          <span class="text-lg w-40">{date}</span>
          <ul class="flex flex-col gap-1 text-gray-700 font-mono">
            {#each times as time}
              <li>{time}</li>
            {/each}
          </ul>
        </div>
      {/each}
      {#if Object.keys(wakeUpTimesByDate).length === 0}
        <span>No wake up times recorded.</span>
      {/if}
    </div>
  {/snippet}
</Section>
