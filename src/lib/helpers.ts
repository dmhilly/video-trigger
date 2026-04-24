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

export function toTimeString(date: Date | undefined): string {
  if (date === undefined) return "";
  return date.toLocaleTimeString("en-US", {
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
  });
}

export function addRemoveDay(date: Date, previous = false): Date {
  const multiplier = previous ? -1 : 1;
  return new Date(
    date.getFullYear(),
    date.getMonth(),
    date.getDate() + 1 * multiplier,
  );
}
