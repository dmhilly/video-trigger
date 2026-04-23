<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";

  import Button from "../button.svelte";
  import { formatFileSize } from "../../lib/helpers";
  import Section from "../section.svelte";

  interface Video {
    name: string | undefined;
    binaryDataId: string;
    url: string | undefined;
    timestamp: VIAM.Timestamp | undefined;
    size: bigint | undefined;
    loaded: boolean;
  }

  interface Props {
    dataClient: VIAM.DataClient | undefined;
  }
  const { dataClient }: Props = $props();

  const maxCount = 50;

  let videos = $state<Video[]>([]);
  let videoCount = $state(0);

  async function loadVideos() {
    if (!dataClient) return;

    const filter = new VIAM.dataApi.Filter({
      mimeType: ["video/mp4"],
      tagsFilter: { tags: ["awake"] },
    });

    const { data, count } = await dataClient.binaryDataByFilter(
      filter,
      maxCount,
      VIAM.dataApi.Order.DESCENDING,
      undefined,
      false,
    );
    videoCount = Number(count);

    videos = data.map((item) => ({
      name: item.metadata?.fileName,
      binaryDataId: item.metadata?.binaryDataId ?? "",
      url: undefined,
      timestamp: item.metadata?.timeRequested,
      size: item.metadata?.fileSizeBytes,
      loaded: false,
    }));
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
              {video.timestamp?.toDate().toLocaleString()}
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
    {#if videoCount === maxCount}
      <div class="bg-gray border rounded-xl py-3 px-8 self-center">
        <span>Max number of videos ({maxCount}) shown</span>
      </div>
    {/if}
  </div>
{/if}
