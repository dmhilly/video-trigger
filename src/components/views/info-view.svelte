<script lang="ts">
  import Section from "../section.svelte";
  import { get7AM, get7PM, toDateString } from "../../lib/helpers";
  import Timeline from "./timeline.svelte";
  import type { WakeUpEvent } from "./times";

  interface WakeUpPeriod {
    events: WakeUpEvent[];
    start: Date;
    end: Date;
  }

  interface Props {
    wakeUpTimes: Date[];
  }
  const { wakeUpTimes }: Props = $props();

  const wakeUpTimesByPeriod: Record<string, WakeUpPeriod> = $derived.by(() => {
    const grouped: Record<string, WakeUpPeriod> = {};
    for (const date of wakeUpTimes) {
      const hour = date.getHours();
      let periodKey: string;

      let start: Date;
      let end: Date;

      if (hour < 7) {
        start = get7PM(date, -1);
        end = get7AM(date);
        periodKey = `${toDateString(start)} - ${toDateString(date)} night`;
      } else if (hour < 19) {
        periodKey = `${toDateString(date)} day`;
        start = get7AM(date);
        end = get7PM(date);
      } else {
        start = get7PM(date);
        end = get7AM(date, 1);
        periodKey = `${toDateString(date)} - ${toDateString(end)} night`;
      }

      if (!grouped[periodKey]) {
        grouped[periodKey] = { events: [], start, end };
      }
      grouped[periodKey].events.push({
        start: date,
        end: new Date(date.getTime() + 2 * 60 * 60),
      });
    }
    return grouped;
  });
</script>

{#if Object.keys(wakeUpTimesByPeriod).length === 0}
  <span>No wake up times recorded.</span>
{/if}

{#each Object.entries(wakeUpTimesByPeriod) as [name, period]}
  <Section title={name}>
    {#snippet content()}{/snippet}
  </Section>
{/each}
