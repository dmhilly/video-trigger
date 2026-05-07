<script lang="ts">
  import type { DataClient } from "@viamrobotics/sdk";

  import { toTimeString } from "../../lib/helpers";
  import { Size } from "../../lib/types";
  import Button from "../button.svelte";
  import DownloadButton from "./download-button.svelte";
  import type { WakeUpEvent } from "./times";
  import type { Video } from "./video";
  import IconButton from "../icon-button.svelte";
  import Tooltip from "../tooltip.svelte";

  interface Props {
    dataClient: DataClient | undefined;
    events: WakeUpEvent[];
    start: Date;
    end: Date;
    videos: Video[];
  }
  const { dataClient, events, start, end, videos }: Props = $props();

  const WIDTH = 600;
  const HEIGHT = 60;
  const PAD = { left: 10, right: 10 };
  const chartW = WIDTH - PAD.left - PAD.right;
  const BAR_HEIGHT = 20;
  const BAR_Y = 10;

  const minT = $derived(start.getTime());
  const maxT = $derived(end.getTime());

  function timeToX(date: Date): number {
    return PAD.left + ((date.getTime() - minT) / (maxT - minT)) * chartW;
  }

  const ticks = $derived.by(() => {
    const count = 5;
    const step = (maxT - minT) / (count - 1);
    return Array.from({ length: count }, (_, i) => new Date(minT + step * i));
  });

  let video = $state<Video | undefined>(undefined);
  async function setVideo(eventStart: Date) {
    if (!dataClient) return;
    const match = videos.find((v) => {
      if (!v.timestamp) return false;
      return v.timestamp.toDate().getTime() === eventStart.getTime();
    });
    if (!match) {
      console.warn("No video found for event at", eventStart);
      return;
    }
    if (!match.url) {
      match.url = await dataClient.createBinaryDataSignedURL(
        match.binaryDataId,
      );
    }

    video = match;
  }

  let hoveredEvent = $state<
    { event: WakeUpEvent; xPercent: number } | undefined
  >(undefined);
</script>

<div class="w-full flex flex-col gap-4">
  <div class="relative">
    <svg viewBox="0 0 {WIDTH} {HEIGHT}" class="w-full">
      <rect
        x={PAD.left}
        y={BAR_Y}
        width={chartW}
        height={BAR_HEIGHT}
        rx="4"
        class="fill-gray-500"
      />

      {#each events as event}
        {@const x1 = timeToX(event.start)}
        {@const x2 = timeToX(event.end)}
        <rect
          role="button"
          x={x1}
          y={BAR_Y}
          width={Math.max(x2 - x1, 4)}
          height={BAR_HEIGHT}
          rx="2"
          class="fill-white hover:cursor-pointer"
          onclick={() => setVideo(event.start)}
          onmouseenter={() =>
            (hoveredEvent = {
              event,
              xPercent: ((x1 + Math.max(x2 - x1, 4) / 2) / WIDTH) * 100,
            })}
          onmouseleave={() => (hoveredEvent = undefined)}
        />
      {/each}

      {#each ticks as tick, i}
        {@const x = timeToX(tick)}
        <line
          x1={x}
          y1={BAR_Y + BAR_HEIGHT}
          x2={x}
          y2={BAR_Y + BAR_HEIGHT + 4}
          stroke="#666"
        />
        <text
          {x}
          y={HEIGHT - 2}
          fill="#aaa"
          font-size="10"
          text-anchor={i === 0
            ? "start"
            : i === ticks.length - 1
              ? "end"
              : "middle"}
        >
          {toTimeString(tick)}
        </text>
      {/each}
    </svg>

    {#if hoveredEvent}
      <div
        class="absolute pointer-events-none top-3 translate-x-1/2"
        style="left:{hoveredEvent.xPercent}%;"
      >
        <Tooltip
          text="{toTimeString(hoveredEvent.event.start)} – {toTimeString(
            hoveredEvent.event.end,
          )}"
          visible={true}
          position="top"
          class="whitespace-nowrap"
        />
      </div>
    {/if}
  </div>

  {#if video && dataClient}
    <div class="flex flex-col gap-1">
      <div class="self-end flex items-center">
        <DownloadButton {dataClient} {video} />
        <IconButton
          icon="close"
          size={Size.Medium}
          onclick={() => (video = undefined)}
        />
      </div>
      <video src={video.url} controls class="w-full rounded">
        <track kind="captions" />
      </video>
    </div>
  {/if}
</div>
