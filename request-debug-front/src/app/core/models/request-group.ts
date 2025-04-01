export interface Request {
  id: string;
}

export interface RequestGroup {
  id: string;
  requests: Request[];
  createdAt: string;
  updatedAt: string;
}
