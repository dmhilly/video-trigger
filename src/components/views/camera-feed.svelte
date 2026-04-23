<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";
  import InfoList from "../info-list.svelte";
  import Section from "../section.svelte";
  import { untrack } from "svelte";

  interface Props {
    mediaStream: MediaStream | undefined;
    motionDetections: VIAM.Classification[];
    awakeClassifications: VIAM.Classification[];
  }
  const { mediaStream, motionDetections, awakeClassifications }: Props =
    $props();

  let videoElement: HTMLVideoElement;
  $effect(() => {
    if (videoElement && mediaStream) videoElement.srcObject = mediaStream;
  });

  let moving = $state(false);
  let motionStatus = $derived(
    moving ? "Motion detected within 2 minutes" : "No motion detected",
  );
  let lastMotionTime: Date | undefined = undefined;
  let lastMotionLocal = $state("None");
  const TWO_MINUTES = 2 * 60 * 1000;

  $effect(() => {
    if (motionDetections.length > 0 && motionDetections[0].confidence > 0.001) {
      lastMotionTime = new Date();
      lastMotionLocal = lastMotionTime.toLocaleString();
      console.log(lastMotionTime, motionDetections[0].confidence);
      moving = true;
    } else if (
      untrack(() => moving) &&
      lastMotionTime &&
      Date.now() - lastMotionTime.getTime() >= TWO_MINUTES
    ) {
      console.log("two minutes since movement");
      moving = false;
    }
  });

  const babyStatus = $derived(awakeClassifications);
  $effect(() => {
    // TODO - remove when awake classifier works
    if (babyStatus.length > 0 && babyStatus[0].confidence > 0)
      console.log("awake: ", babyStatus);
  });

  // TODO - replace with actual data
  const lastAwake = new Date(2026, 3, 9, 19, 41, 0);
  const lastAwakeLocal = lastAwake.toLocaleString();
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

<Section title="Recent Activity:">
  {#snippet content()}
    <InfoList items={{ "Fell Asleep at": lastAwakeLocal }} />
  {/snippet}
</Section>
