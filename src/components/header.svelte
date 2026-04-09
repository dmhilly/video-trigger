<script lang="ts">
  interface Props {
    text: string;
    variant?: "primary" | "secondary";
  }
  const { text, variant = "primary" }: Props = $props();
</script>

{#if variant === "primary"}
  <h1 class="text-2xl">{text}</h1>
{:else if variant === "secondary"}
  <h2 class="text-xl">{text}</h2>
{/if}
