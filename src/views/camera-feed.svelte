<script lang="ts">
  import InfoList from "../components/info-list.svelte";
  import Section from "../components/section.svelte";

  interface Props {
    mediaStream: MediaStream | undefined;
  }
  const { mediaStream }: Props = $props();

  let videoElement: HTMLVideoElement;

  // Example data
  const babyStatus = $state<"awake" | "asleep">("awake");
  const lastAwake = new Date(2026, 3, 9, 19, 41, 0);
  const wokeUp = new Date(2026, 3, 10, 4, 20, 12);

  const lastAwakeLocal = lastAwake.toLocaleString();
  const wokeUpLocal = wokeUp.toLocaleString();

  $effect(() => {
    if (videoElement && mediaStream) videoElement.srcObject = mediaStream;
  });
</script>

<Section title="Camera feed:">
  {#snippet content()}
    <video
      bind:this={videoElement}
      autoplay
      playsinline
      muted
      class="flex-1 min-w-0"
    ></video>
    <InfoList
      items={{
        Status: babyStatus,
        "Last Awake": lastAwakeLocal,
        "Noise Level": 0,
      }}
    />
  {/snippet}
</Section>

<Section title="Activity:">
  {#snippet content()}
    <InfoList
      items={babyStatus === "asleep"
        ? {
            "Fell Asleep at": lastAwakeLocal,
          }
        : { "Awoke at": wokeUpLocal }}
    />
  {/snippet}
</Section>
