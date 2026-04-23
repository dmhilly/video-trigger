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
    "border h-18 w-18 rounded-xl flex flex-col gap-1 items-center justify-center hover:cursor-pointer shadow-lg",
    { "bg-black text-white": selected },
    { "bg-white": !selected },
  ]}
  {onclick}
>
  <Icon name={icon} />
  <span>{text}</span>
</button>
