<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";

  import type { Video } from "./video";
  import Button from "../button.svelte";
  import Section from "../section.svelte";
  import {
    formatFileSize,
    toDateString,
    toTimeString,
  } from "../../lib/helpers";

  interface Props {
    dataClient: VIAM.DataClient | undefined;
    videos: Video[];
    maxVideoCount: number;
  }
  const { dataClient, videos, maxVideoCount }: Props = $props();

  let videoCount = $derived(videos.length);

  async function loadVideos() {
    if (!dataClient) return;
  }

  async function loadVideo(video: Video) {
    if (!dataClient) return;
    if (!video.url) {
      video.url = await dataClient.createBinaryDataSignedURL(
        video.binaryDataId,
      );
    }
    video.loaded = true;
  }

  loadVideos();
</script>

{#if dataClient}
  <Section title="Wake Time Videos">
    {#snippet content()}
      Found {videoCount} video{videoCount === 1 ? "" : "s"}
    {/snippet}
  </Section>
  <div class="flex flex-col gap-4">
    {#each videos as video}
      <div
        class="flex flex-col bg-white gap-4 border rounded-xl p-4 items-start"
      >
        <div class="flex flex-col gap-1">
          <span>
            <span class="text-lg font-semibold">
              {toDateString(video.timestamp?.toDate())} -
              <span class="font-mono">
                {toTimeString(video.timestamp?.toDate())}
              </span>
            </span>
            · {formatFileSize(video.size)}
          </span>
          <span>{video.name}</span>
        </div>
        {#if video.loaded && video.url}
          <div class="flex flex-col gap-1">
            <Button
              text="Close"
              icon="close"
              onclick={() => (video.loaded = false)}
              class="self-end"
            />
            <video
              src={video.url}
              controls
              class="w-full"
              oncanplay={() => {
                video.loaded = true;
              }}
            >
              <track kind="captions" />
            </video>
          </div>
        {:else}
          <Button
            text="Load video"
            onclick={() => loadVideo(video)}
            class="items-end"
          />
        {/if}
      </div>
    {/each}
    {#if videoCount === maxVideoCount}
      <div class="bg-gray border rounded-xl py-3 px-8 self-center">
        <span>Max number of videos ({maxVideoCount}) shown</span>
      </div>
    {/if}
  </div>
{/if}
