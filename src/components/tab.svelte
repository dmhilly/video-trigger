<script lang="ts">
  import Icon from "./icon.svelte";

  interface Props {
    icon: string;
    text: string;
  }
  const { icon, text }: Props = $props();
</script>

<div
  class="border h-18 w-18 rounded-xl flex flex-col gap-1 items-center justify-center"
>
  <Icon name={icon} />
  <span>{text}</span>
</div>
