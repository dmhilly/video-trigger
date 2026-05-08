<script lang="ts">
  import type { DataClient } from "@viamrobotics/sdk";
  import Section from "../section.svelte";
  import { get7AM, get7PM, toDateString } from "../../lib/helpers";
  import Timeline from "./timeline.svelte";
  import type { WakeUpEvent } from "./times";
  import type { Video } from "./video";

  interface WakeUpPeriod {
    events: WakeUpEvent[];
    start: Date;
    end: Date;
  }

  interface Props {
    dataClient: DataClient | undefined;
    wakeUpTimes: Date[];
    videos: Video[];
  }
  const { dataClient, wakeUpTimes, videos }: Props = $props();

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
        end: new Date(date.getTime() + 15 * 60 * 60),
      });
    }
    return grouped;
  });
</script>

{#if Object.keys(wakeUpTimesByPeriod).length === 0}
  <span>No wake up times recorded.</span>
{/if}

{#each Object.entries(wakeUpTimesByPeriod) as [name, period]}
  {@const isNight = name.includes("night")}
  <Section
    title={name}
    class={isNight ? "bg-indigo-200!" : "bg-yellow-50!"}
    icon={isNight ? "moon" : "sun"}
    iconClass={isNight ? "text-indigo-700" : "text-yellow-500"}
  >
    {#snippet content()}
      <Timeline
        {dataClient}
        events={period.events}
        start={period.start}
        end={period.end}
        {videos}
      />
    {/snippet}
  </Section>
{/each}
