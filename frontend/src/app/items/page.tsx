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
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { getItems } from "@/lib/api/items";
import { Item } from "@/types/item";

const mockItems: Item[] = [
  {
    id: "mock-1",
    name: "Sample Item 1",
    description: "This is a sample item for development",
    created_at: "2025-01-01T00:00:00Z",
    updated_at: "2025-01-01T00:00:00Z",
  },
  {
    id: "mock-2",
    name: "Sample Item 2",
    description: "Another sample item",
    created_at: "2025-01-02T00:00:00Z",
    updated_at: "2025-01-02T00:00:00Z",
  },
  {
    id: "mock-3",
    name: "Sample Item 3",
    description: null,
    created_at: "2025-01-03T00:00:00Z",
    updated_at: "2025-01-03T00:00:00Z",
  },
];

export default async function ItemsPage() {
  let items: Item[];
  let isUsingMock = false;

  try {
    items = await getItems();
  } catch {
    items = mockItems;
    isUsingMock = true;
  }

  return (
    <main className="min-h-screen bg-background">
      <div className="mx-auto max-w-6xl px-6 py-10">
        {/* Header */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold tracking-tight">Items</h1>
            <p className="mt-1 text-muted-foreground">Manage your items</p>
          </div>
          <Link href="/items/new">
            <Button>New Item</Button>
          </Link>
        </div>

        <Separator className="my-6" />

        {isUsingMock && (
          <div className="mb-4 rounded-md border border-yellow-200 bg-yellow-50 p-3 text-sm text-yellow-800 dark:border-yellow-800 dark:bg-yellow-950 dark:text-yellow-200">
            API is unavailable. Showing mock data.
          </div>
        )}

        {/* Items Table */}
        <Card>
          <CardHeader>
            <CardTitle>All Items</CardTitle>
            <CardDescription>
              A list of all items. Click on an item to view details.
            </CardDescription>
          </CardHeader>
          <CardContent>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-[200px]">ID</TableHead>
                  <TableHead>Name</TableHead>
                  <TableHead>Description</TableHead>
                  <TableHead className="w-[180px]">Created</TableHead>
                  <TableHead className="w-[100px] text-right">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {items.length === 0 ? (
                  <TableRow>
                    <TableCell colSpan={5} className="text-center text-muted-foreground">
                      No items found.
                    </TableCell>
                  </TableRow>
                ) : (
                  items.map((item) => (
                    <TableRow key={item.id}>
                      <TableCell className="font-mono text-sm">{item.id}</TableCell>
                      <TableCell className="font-medium">{item.name}</TableCell>
                      <TableCell className="text-muted-foreground">
                        {item.description ? (
                          item.description
                        ) : (
                          <Badge variant="outline">No description</Badge>
                        )}
                      </TableCell>
                      <TableCell className="text-sm text-muted-foreground">
                        {new Date(item.created_at).toLocaleDateString("ja-JP")}
                      </TableCell>
                      <TableCell className="text-right">
                        <Link href={`/items/${item.id}`}>
                          <Button variant="ghost" size="sm">
                            View
                          </Button>
                        </Link>
                      </TableCell>
                    </TableRow>
                  ))
                )}
              </TableBody>
            </Table>
          </CardContent>
        </Card>
      </div>
    </main>
  );
}
