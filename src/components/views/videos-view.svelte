<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";

  import Button from "../button.svelte";
  import { formatFileSize } from "../../lib/helpers";

  interface Video {
    name: string | undefined;
    url: string;
    timestamp: VIAM.Timestamp | undefined;
    size: bigint | undefined;
    loaded: boolean;
  }

  interface Props {
    dataClient: VIAM.DataClient | undefined;
  }
  const { dataClient }: Props = $props();

  let videos = $state<Video[]>([]);
  let videoCount = $state("0");

  async function loadVideos() {
    if (!dataClient) return;

    const filter = new VIAM.dataApi.Filter({
      mimeType: ["video/mp4"],
      tagsFilter: { tags: ["awake"] },
    });

    const { data, count } = await dataClient.binaryDataByFilter(
      filter,
      undefined,
      undefined,
      undefined,
      false,
    );
    videoCount = count.toString();

    videos = await Promise.all(
      data.map(async (item) => ({
        name: item.metadata?.fileName,
        url: await dataClient.createBinaryDataSignedURL(
          item.metadata?.binaryDataId ?? "",
        ),
        timestamp: item.metadata?.timeRequested,
        size: item.metadata?.fileSizeBytes,
        loaded: false,
      })),
    );
  }

  loadVideos();
</script>

{#if dataClient}
  <span>Found {videoCount} video{videoCount === "1" ? "" : "s"}</span>
  <div class="flex flex-col gap-4">
    {#each videos as video}
      <div
        class="flex flex-col bg-white gap-4 border rounded-xl p-4 items-start"
      >
        <div class="flex flex-col gap-1">
          <span class="font-lg font-semibold">{video.name}</span>
          <span>
            {video.timestamp?.toDate().toLocaleString()} ·
            {formatFileSize(video.size)}
          </span>
        </div>
        {#if video.loaded}
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
            onclick={() => (video.loaded = true)}
            class="items-end"
          />
        {/if}
      </div>
    {/each}
  </div>
{/if}
