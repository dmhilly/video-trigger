<script lang="ts">
  import Section from "../section.svelte";
  import { addRemoveDay, toDateString } from "../../lib/helpers";
  import Timeline from "./timeline.svelte";

  interface Props {
    wakeUpTimes: Date[];
  }
  const { wakeUpTimes }: Props = $props();

  const wakeUpTimesByPeriod = $derived.by(() => {
    const grouped: Record<string, Date[]> = {};
    for (const date of wakeUpTimes) {
      const hour = date.getHours();
      let periodKey: string;

      if (hour < 7) {
        const prevDay = addRemoveDay(date, true);
        periodKey = `${toDateString(prevDay)} - ${toDateString(date)} night`;
      } else if (hour < 19) {
        periodKey = `${toDateString(date)} day`;
      } else {
        const nextDay = addRemoveDay(date);
        periodKey = `${toDateString(date)} - ${toDateString(nextDay)} night`;
      }

      if (!grouped[periodKey]) {
        grouped[periodKey] = [];
      }
      grouped[periodKey].push(date);
    }
    return grouped;
  });
</script>

{#if Object.keys(wakeUpTimesByPeriod).length === 0}
  <span>No wake up times recorded.</span>
{/if}

{#each Object.entries(wakeUpTimesByPeriod) as [period, times]}
  <Section title={period}>
    {#snippet content()}
      <Timeline {times} />
    {/snippet}
  </Section>
{/each}
