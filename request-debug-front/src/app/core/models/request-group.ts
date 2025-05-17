export interface RequestFile {
  filename: string;
  size: number;
}

export interface Request {
  id: string;
  method: string;
  host: string;
  url: string;
  bodySize: number;
  bodyRaw: string;
  date: string;
  ip: string;
  form: Record<string, string[]>;
  files: Record<string, RequestFile[]>;
  queryParams: Record<string, string>;
  headers: Record<string, string>;
}

export interface RequestGroup {
  id: string;
  requests: Request[];
  createdAt: string;
  updatedAt: string;
}
