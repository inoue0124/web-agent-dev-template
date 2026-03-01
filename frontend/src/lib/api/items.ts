import { Item, CreateItemRequest } from "@/types/item";

const API_BASE = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export async function getItems(): Promise<Item[]> {
  const res = await fetch(`${API_BASE}/api/v1/items`, { cache: "no-store" });
  if (!res.ok) throw new Error("Failed to fetch items");
  return res.json();
}

export async function getItem(id: string): Promise<Item> {
  const res = await fetch(`${API_BASE}/api/v1/items/${id}`, { cache: "no-store" });
  if (!res.ok) throw new Error("Failed to fetch item");
  return res.json();
}

export async function createItem(data: CreateItemRequest): Promise<Item> {
  const res = await fetch(`${API_BASE}/api/v1/items`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error("Failed to create item");
  return res.json();
}

export async function deleteItem(id: string): Promise<void> {
  const res = await fetch(`${API_BASE}/api/v1/items/${id}`, { method: "DELETE" });
  if (!res.ok) throw new Error("Failed to delete item");
}
