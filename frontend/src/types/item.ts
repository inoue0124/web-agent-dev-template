export type Item = {
  id: string;
  name: string;
  description: string | null;
  created_at: string;
  updated_at: string;
};

export type CreateItemRequest = {
  name: string;
  description?: string;
};

export type UpdateItemRequest = {
  name?: string;
  description?: string;
};
