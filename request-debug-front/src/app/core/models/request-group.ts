export interface Request {
  id: string;
  method: string;
  host: string;
  date: string;
  queryParams: Record<string, string>;
  ip: string;
}

export interface RequestGroup {
  id: string;
  requests: Request[];
  createdAt: string;
  updatedAt: string;
}
