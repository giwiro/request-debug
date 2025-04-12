export function join(url: string, path: string): string {
  const u = new URL(path, url);
  return u.href;
}
