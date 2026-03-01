import { NextResponse } from "next/server";

const API_BASE = process.env.API_URL || "http://localhost:8080";

export async function GET() {
  try {
    const res = await fetch(`${API_BASE}/api/v1/items`, { cache: "no-store" });
    const data = await res.json();
    return NextResponse.json(data);
  } catch {
    return NextResponse.json([], { status: 200 });
  }
}

export async function POST(request: Request) {
  const body = await request.json();
  const res = await fetch(`${API_BASE}/api/v1/items`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body),
  });
  const data = await res.json();
  return NextResponse.json(data, { status: res.status });
}
