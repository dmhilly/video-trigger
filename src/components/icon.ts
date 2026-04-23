import IC from "@iconify-json/ic/icons.json";

export type Icon = {
  body: string;
  width?: number;
  height?: number;
};

export const IconByName: Record<string, Icon> = {
  camera: (IC as { icons: Record<string, Icon> }).icons["baseline-camera"],
  info: (IC as { icons: Record<string, Icon> }).icons["round-info"],
};
