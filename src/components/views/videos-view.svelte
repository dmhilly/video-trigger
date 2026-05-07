<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";

  import type { Video } from "./video";
  import Button from "../button.svelte";
  import Section from "../section.svelte";
  import { formatFileSize } from "../../lib/helpers";
  import Datetime from "../datetime.svelte";
  import IconButton from "../icon-button.svelte";
  import Tooltip from "../tooltip.svelte";

  interface Props {
    dataClient: VIAM.DataClient | undefined;
    videos: Video[];
    maxVideoCount: number;
    lastVideoFetch: Date | undefined;
  }
  const { dataClient, videos, maxVideoCount, lastVideoFetch }: Props = $props();

  let videoCount = $derived(videos.length);

  async function loadVideo(video: Video) {
    if (!dataClient) return;
    if (!video.url) {
      video.url = await dataClient.createBinaryDataSignedURL(
        video.binaryDataId,
      );
    }
  }
  async function openVideo(video: Video) {
    loadVideo(video);
    video.loaded = true;
  }

  async function downloadVideo(video: Video) {
    await loadVideo(video);
    if (!video.url) return;
    window.open(video.url, "_blank");
  }
</script>

{#if dataClient}
  <Section title="Wake Time Videos" contentClasses="justify-between">
    {#snippet content()}
      Found {videoCount} video{videoCount === 1 ? "" : "s"}
      <span class="text-subtle">
        Last fetched at: <Datetime date={lastVideoFetch} />
      </span>
    {/snippet}
  </Section>
  <div class="flex flex-col gap-4">
    {#each videos as video}
      <div
        class="flex flex-col bg-white gap-4 border rounded-xl p-4 items-start"
      >
        <div class="flex gap-1 w-full justify-between items-center">
          <span>
            <Datetime
              date={video.timestamp?.toDate()}
              class="text-lg font-semibold"
            />
            · {formatFileSize(video.size)}
          </span>
          <Tooltip text={video.name ?? ""}>
            <IconButton icon="download" onclick={() => downloadVideo(video)} />
          </Tooltip>
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
            onclick={() => openVideo(video)}
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
