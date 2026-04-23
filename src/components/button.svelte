<script lang="ts">
  interface Props {
    text: string;
    class?: string;
    onclick: () => void;
  }
  const { text, class: extraClasses, onclick }: Props = $props();
</script>

<button
  class={[
    "bg-black text-white px-8 py-2 rounded-xl hover:cursor-pointer",
    extraClasses,
  ]}
  {onclick}
>
  {text}
</button>
