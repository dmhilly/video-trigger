<script lang="ts">
  import { toTimeString } from "../../lib/helpers";
  import type { WakeUpEvent } from "./times";

  interface Props {
    events: WakeUpEvent[];
    start: Date;
    end: Date;
  }
  const { events, start, end }: Props = $props();

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
</script>

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
      x={x1}
      y={BAR_Y}
      width={Math.max(x2 - x1, 4)}
      height={BAR_HEIGHT}
      rx="2"
      class="fill-white hover:cursor-pointer"
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
