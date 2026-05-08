export function formatFileSize(bytes: bigint | undefined): string {
  if (!bytes) return `-`;
  const size = Number(bytes);
  if (size < 1024) return `${size} B`;
  if (size < 1024 * 1024) return `${(size / 1024).toFixed(1)} KB`;
  return `${(size / (1024 * 1024)).toFixed(1)} MB`;
}

export function toDateTimeString(date: Date | undefined): string {
  if (date === undefined) return "";
  return date.toLocaleString("en-US", {
    weekday: "short",
    month: "short",
    day: "2-digit",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}

export function toDateString(date: Date | undefined): string {
  if (date === undefined) return "";
  return date.toLocaleDateString("en-US", {
    weekday: "short",
    month: "short",
    day: "2-digit",
    year: "numeric",
  });
}

export function toTimeHourString(date: Date | undefined): string {
  if (date === undefined) return "";
  return date.toLocaleTimeString("en-US", { hour: "numeric" });
}

export function toTimeString(date: Date | undefined): string {
  if (date === undefined) return "";
  return date.toLocaleTimeString("en-US", {
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
  });
}

// basically
export function get7PM(date: Date, dateChange: number = 0): Date {
  return new Date(
    date.getFullYear(),
    date.getMonth(),
    date.getDate() + dateChange,
    19,
  );
}

export function get7AM(date: Date, dateChange: number = 0): Date {
  return new Date(
    date.getFullYear(),
    date.getMonth(),
    date.getDate() + dateChange,
    7,
  );
}
