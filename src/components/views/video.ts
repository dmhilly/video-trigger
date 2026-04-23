import * as VIAM from "@viamrobotics/sdk";

export interface Video {
  name: string | undefined;
  binaryDataId: string;
  url: string | undefined;
  timestamp: VIAM.Timestamp | undefined;
  size: bigint | undefined;
  loaded: boolean;
}
