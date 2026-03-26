# Module video-trigger 

Provide a description of the purpose of the module and any relevant information.

## Models

This module provides the following model(s):

- [`devin-hilly:video-trigger:generic_service`](devin-hilly_video-trigger_generic_service.md) - Provide a brief description of the model

Use the following configuration:
```json
{
  "vision_service": <name of vision service>,
  "video_service": <name of video service>,
  "camera": <name of camera>,
  "threshold": <pct confidence of classification,
  "capture_padding_secs": <seconds of video to include before and after the triggering event>
}
```
