<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";
  import InfoList from "../info-list.svelte";
  import { Status, statusToVariant } from "../../lib/types";
  import Section from "../section.svelte";
  import Heading from "../heading.svelte";
  import Badge from "../badge.svelte";

  interface Props {
    status: Status;
    mediaStream: MediaStream | undefined;
    awakeData: VIAM.JsonValue;
    movementDetected: boolean;
    lastMotionLocal: Date | undefined;
  }
  const {
    status,
    mediaStream,
    awakeData,
    movementDetected,
    lastMotionLocal,
  }: Props = $props();

  let videoElement: HTMLVideoElement | undefined = $state(undefined);
  $effect(() => {
    if (videoElement && mediaStream) videoElement.srcObject = mediaStream;
  });

  let motionStatus = $derived(
    movementDetected
      ? "Motion detected within 2 minutes"
      : "No motion detected",
  );

  const babyStatus = $derived(awakeData as Record<string, any>);
</script>

<Section title="Camera Feed">
  {#snippet content()}
    {#if status === Status.Connected}
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
    {:else}
      <div
        class="bg-gray-200 py-10 rounded-2xl w-full flex flex-col gap-2 items-center"
      >
        <Heading text="Camera feed not available." variant="secondary" />
        <div>
          <span>Machine Status:</span>
          <Badge text={status} variant={statusToVariant[status]} />
        </div>
      </div>
    {/if}
  {/snippet}
</Section>
