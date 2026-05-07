<script lang="ts">
  import * as VIAM from "@viamrobotics/sdk";

  import IconButton from "../icon-button.svelte";
  import Tooltip from "../tooltip.svelte";
  import type { Video } from "./video";
  import { Size } from "../../lib/types";

  interface Props {
    dataClient: VIAM.DataClient;
    video: Video;
  }
  const { dataClient, video }: Props = $props();

  async function downloadVideo(video: Video) {
    if (!video.url) {
      video.url = await dataClient.createBinaryDataSignedURL(
        video.binaryDataId,
      );
    }
    if (!video.url) return;
    window.open(video.url, "_blank");
  }
</script>

<Tooltip text={video.name ?? ""}>
  <IconButton
    icon="download"
    size={Size.Medium}
    onclick={() => downloadVideo(video)}
  />
</Tooltip>
