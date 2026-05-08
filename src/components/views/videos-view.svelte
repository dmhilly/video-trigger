<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";

  import type { Video } from "./video";
  import Button from "../button.svelte";
  import Section from "../section.svelte";
  import { formatFileSize } from "../../lib/helpers";
  import Datetime from "../datetime.svelte";
  import DownloadButton from "./download-button.svelte";
  import IconButton from "../icon-button.svelte";
  import { Size } from "../../lib/types";

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
</script>

{#if dataClient}
  <Section title="Wake Time Videos" contentClasses="justify-between">
    {#snippet content()}
      <span class="text-subtle">
        Found {videoCount} video{videoCount === 1 ? "" : "s"}
      </span>
      <span class="text-subtle">
        Last fetched at: <Datetime date={lastVideoFetch} />
      </span>
    {/snippet}
  </Section>
  <div class="flex flex-col gap-4">
    {#each videos as video}
      <div class="flex flex-col bg-white gap-4 rounded-xl p-4 items-start">
        <div class="flex gap-1 w-full justify-between items-center">
          <span>
            <Datetime
              date={video.timestamp?.toDate()}
              class="text-lg font-semibold"
            />
            · {formatFileSize(video.size)}
          </span>
          <div class="flex items-center">
            <DownloadButton {dataClient} {video} />
            {#if video.loaded && video.url}
              <IconButton
                icon="close"
                size={Size.Medium}
                onclick={() => (video.loaded = false)}
              />
            {/if}
          </div>
        </div>
        {#if video.loaded && video.url}
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
