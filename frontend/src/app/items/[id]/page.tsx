import Link from "next/link";

import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { getItem } from "@/lib/api/items";
import { Item } from "@/types/item";

function createMockItem(id: string): Item {
  return {
    id,
    name: `Sample Item (${id})`,
    description: "This is mock data displayed because the API is unavailable.",
    created_at: "2025-01-01T00:00:00Z",
    updated_at: "2025-01-01T00:00:00Z",
  };
}

export default async function ItemDetailPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;

  let item: Item;
  let isUsingMock = false;

  try {
    item = await getItem(id);
  } catch {
    item = createMockItem(id);
    isUsingMock = true;
  }

  return (
    <main className="min-h-screen bg-background">
      <div className="mx-auto max-w-3xl px-6 py-10">
        {/* Header */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold tracking-tight">Item Detail</h1>
            <p className="mt-1 text-muted-foreground">View item information</p>
          </div>
          <Link href="/items">
            <Button variant="outline">Back to List</Button>
          </Link>
        </div>

        <Separator className="my-6" />

        {isUsingMock && (
          <div className="mb-4 rounded-md border border-yellow-200 bg-yellow-50 p-3 text-sm text-yellow-800 dark:border-yellow-800 dark:bg-yellow-950 dark:text-yellow-200">
            API is unavailable. Showing mock data.
          </div>
        )}

        {/* Item Detail Card */}
        <Card>
          <CardHeader>
            <CardTitle>{item.name}</CardTitle>
            <CardDescription>
              <span className="font-mono text-xs">{item.id}</span>
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div>
              <p className="text-sm font-medium text-muted-foreground">Description</p>
              {item.description ? (
                <p className="mt-1">{item.description}</p>
              ) : (
                <Badge variant="outline" className="mt-1">
                  No description
                </Badge>
              )}
            </div>
            <Separator />
            <div className="grid grid-cols-2 gap-4">
              <div>
                <p className="text-sm font-medium text-muted-foreground">Created</p>
                <p className="mt-1 text-sm">
                  {new Date(item.created_at).toLocaleString("ja-JP")}
                </p>
              </div>
              <div>
                <p className="text-sm font-medium text-muted-foreground">Updated</p>
                <p className="mt-1 text-sm">
                  {new Date(item.updated_at).toLocaleString("ja-JP")}
                </p>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </main>
  );
}
