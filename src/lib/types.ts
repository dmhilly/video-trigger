export enum Status {
  Connecting = "Connecting...",
  Connected = "Connected",
  Error = "Error",
  Offline = "Offline",
}

export const statusToVariant: Record<
  Status,
  "warning" | "success" | "error" | "neutral"
> = {
  [Status.Connecting]: "warning",
  [Status.Connected]: "success",
  [Status.Error]: "error",
  [Status.Offline]: "neutral",
};

export enum Size {
  Small = "sm",
  Medium = "md",
  Large = "lg",
  XLarge = "xl",
}
