<script lang="ts">
  import { toDateString, toTimeString } from "../lib/helpers";

  interface Props {
    date: Date | undefined;
    class?: string;
  }
  const { date, class: extraClasses }: Props = $props();

  const displayDate = $derived(toDateString(date));
  const displayTime = $derived(toTimeString(date));
</script>

<span class={extraClasses}>
  {displayDate} - <span class="font-mono">{displayTime}</span>
</span>
