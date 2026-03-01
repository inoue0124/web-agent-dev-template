"use client";

import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { z } from "zod";

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Separator } from "@/components/ui/separator";
import { createItem } from "@/lib/api/items";

const createItemSchema = z.object({
  name: z.string().min(1, "Name is required").max(100, "Name must be 100 characters or less"),
  description: z.string().max(500, "Description must be 500 characters or less").optional(),
});

type FieldErrors = {
  name?: string;
  description?: string;
};

export default function NewItemPage() {
  const router = useRouter();
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [errors, setErrors] = useState<FieldErrors>({});
  const [submitError, setSubmitError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setErrors({});
    setSubmitError(null);

    const result = createItemSchema.safeParse({
      name,
      description: description || undefined,
    });

    if (!result.success) {
      const fieldErrors: FieldErrors = {};
      for (const issue of result.error.issues) {
        const field = issue.path[0] as keyof FieldErrors;
        fieldErrors[field] = issue.message;
      }
      setErrors(fieldErrors);
      return;
    }

    setIsSubmitting(true);
    try {
      await createItem(result.data);
      router.push("/items");
    } catch {
      setSubmitError("Failed to create item. Please make sure the API server is running.");
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <main className="bg-background min-h-screen">
      <div className="mx-auto max-w-2xl px-6 py-10">
        {/* Header */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold tracking-tight">New Item</h1>
            <p className="text-muted-foreground mt-1">Create a new item</p>
          </div>
          <Link href="/items">
            <Button variant="outline">Back to List</Button>
          </Link>
        </div>

        <Separator className="my-6" />

        {submitError && (
          <div className="mb-4 rounded-md border border-red-200 bg-red-50 p-3 text-sm text-red-800 dark:border-red-800 dark:bg-red-950 dark:text-red-200">
            {submitError}
          </div>
        )}

        {/* Form */}
        <Card>
          <CardHeader>
            <CardTitle>Item Details</CardTitle>
            <CardDescription>Fill in the information below to create a new item.</CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-6">
              <div className="space-y-2">
                <Label htmlFor="name">Name</Label>
                <Input
                  id="name"
                  placeholder="Enter item name"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  aria-invalid={!!errors.name}
                />
                {errors.name && <p className="text-destructive text-sm">{errors.name}</p>}
              </div>

              <div className="space-y-2">
                <Label htmlFor="description">Description (optional)</Label>
                <Input
                  id="description"
                  placeholder="Enter item description"
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                  aria-invalid={!!errors.description}
                />
                {errors.description && (
                  <p className="text-destructive text-sm">{errors.description}</p>
                )}
              </div>

              <div className="flex gap-3">
                <Button type="submit" disabled={isSubmitting}>
                  {isSubmitting ? "Creating..." : "Create Item"}
                </Button>
                <Link href="/items">
                  <Button type="button" variant="outline">
                    Cancel
                  </Button>
                </Link>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </main>
  );
}
