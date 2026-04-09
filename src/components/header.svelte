<script lang="ts">
  interface Props {
    text: string;
    variant?: "primary";
  }
  const { text, variant = "primary" }: Props = $props();
</script>

<h1 class="text-2xl">
  {text}
</h1>
