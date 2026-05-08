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
    "h-9 w-full p-4 rounded-xl flex gap-1 items-center justify-start hover:cursor-pointer",
    { "bg-gray-700 text-white": selected },
    { "bg-white text-gray-700": !selected },
  ]}
  {onclick}
>
  <Icon name={icon} /> <span>{text}</span>
</button>
