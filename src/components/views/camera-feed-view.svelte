<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";
  import InfoList from "../info-list.svelte";
  import Section from "../section.svelte";

  interface Props {
    mediaStream: MediaStream | undefined;
    awakeData: VIAM.JsonValue;
    movementDetected: boolean;
    lastMotionLocal: Date | undefined;
  }
  const { mediaStream, awakeData, movementDetected, lastMotionLocal }: Props =
    $props();

  let videoElement: HTMLVideoElement;
  $effect(() => {
    if (videoElement && mediaStream) videoElement.srcObject = mediaStream;
  });

  let motionStatus = $derived(
    movementDetected
      ? "Motion detected within 2 minutes"
      : "No motion detected",
  );

  const babyStatus = $derived(awakeData);
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
        "Is Awake": babyStatus["is_awake"],
        Activity: motionStatus,
        "Last Motion": lastMotionLocal,
        "Eyes Detected": babyStatus["eyes_detected"],
        "Sound Detected": babyStatus["sound_detected"],
      }}
    />
  {/snippet}
</Section>
