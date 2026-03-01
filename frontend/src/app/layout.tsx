import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Web Agent Dev Template",
  description: "Next.js + Go/Gin で始める AI エージェント開発",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="ja">
      <body>{children}</body>
    </html>
  );
}
