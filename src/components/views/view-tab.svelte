<script lang="ts">
  import Icon from "../icon.svelte";

  interface Props {
    icon: string;
    text: string;
    selected: boolean;
    onclick: () => void;
  }
  const { icon, text, selected, onclick }: Props = $props();
</script>

<button
  class={[
    "h-9 w-36 p-4 rounded-xl flex gap-1 items-center justify-start hover:cursor-pointer",
    { "bg-black text-white": selected },
    { "bg-white": !selected },
  ]}
  {onclick}
>
  <Icon name={icon} /> <span>{text}</span>
</button>
