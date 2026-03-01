import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

const stats = [
  { title: "Total Items", value: "128", description: "All registered items" },
  { title: "Active", value: "96", description: "Currently active items" },
  { title: "Pending", value: "24", description: "Awaiting review" },
  { title: "Completed", value: "8", description: "Processed this week" },
];

const items = [
  { id: "ITEM-001", name: "Data Pipeline Setup", status: "active" },
  { id: "ITEM-002", name: "API Integration", status: "active" },
  { id: "ITEM-003", name: "Authentication Module", status: "completed" },
  { id: "ITEM-004", name: "Dashboard UI", status: "pending" },
  { id: "ITEM-005", name: "Database Migration", status: "active" },
  { id: "ITEM-006", name: "Monitoring Setup", status: "pending" },
];

function StatusBadge({ status }: { status: string }) {
  const variant =
    status === "active" ? "default" : status === "completed" ? "secondary" : "outline";
  return <Badge variant={variant}>{status}</Badge>;
}

export default function HomePage() {
  return (
    <main className="bg-background min-h-screen">
      <div className="mx-auto max-w-6xl px-6 py-10">
        {/* Header */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold tracking-tight">Web Agent Dev Template</h1>
            <p className="text-muted-foreground mt-1">
              Next.js + Go/Gin で始める AI エージェント開発
            </p>
          </div>
          <Button>New Item</Button>
        </div>

        <Separator className="my-6" />

        {/* Stats Cards */}
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
          {stats.map((stat) => (
            <Card key={stat.title}>
              <CardHeader className="pb-2">
                <CardDescription>{stat.title}</CardDescription>
                <CardTitle className="text-2xl">{stat.value}</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-muted-foreground text-xs">{stat.description}</p>
              </CardContent>
            </Card>
          ))}
        </div>

        {/* Items Table */}
        <Card className="mt-8">
          <CardHeader>
            <CardTitle>Items</CardTitle>
            <CardDescription>A list of sample items managed by the agent.</CardDescription>
          </CardHeader>
          <CardContent>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-[120px]">ID</TableHead>
                  <TableHead>Name</TableHead>
                  <TableHead className="w-[120px]">Status</TableHead>
                  <TableHead className="w-[100px] text-right">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {items.map((item) => (
                  <TableRow key={item.id}>
                    <TableCell className="font-mono text-sm">{item.id}</TableCell>
                    <TableCell>{item.name}</TableCell>
                    <TableCell>
                      <StatusBadge status={item.status} />
                    </TableCell>
                    <TableCell className="text-right">
                      <Button variant="ghost" size="sm">
                        Edit
                      </Button>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </CardContent>
        </Card>
      </div>
    </main>
  );
}
