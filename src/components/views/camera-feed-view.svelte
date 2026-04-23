<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";
  import InfoList from "../info-list.svelte";
  import Section from "../section.svelte";

  interface Props {
    mediaStream: MediaStream | undefined;
    awakeClassifications: VIAM.Classification[];
    movementDetected: boolean;
    lastMotionTime: string;
  }
  const {
    mediaStream,
    awakeClassifications,
    movementDetected,
    lastMotionLocal,
  }: Props = $props();

  let videoElement: HTMLVideoElement;
  $effect(() => {
    if (videoElement && mediaStream) videoElement.srcObject = mediaStream;
  });

  let motionStatus = $derived(
    movementDetected
      ? "Motion detected within 2 minutes"
      : "No motion detected",
  );

  const babyStatus = $derived(awakeClassifications);
  $effect(() => {
    // TODO - remove when awake classifier works
    if (babyStatus.length > 0 && babyStatus[0].confidence > 0)
      console.log("awake: ", babyStatus);
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
        Activity: motionStatus,
        "Last Motion": lastMotionLocal,
      }}
    />
  {/snippet}
</Section>
