# CLAUDE.md

このプロジェクトは Web Agent Dev Template で生成された Next.js + Go/Gin Web アプリケーションです。

## 技術スタック

### フロントエンド

- Next.js 15+ (App Router)
- React 19
- TypeScript strict mode
- Tailwind CSS v4
- shadcn/ui（コンポーネントライブラリ）

### バックエンド

- Go 1.23+
- Gin（HTTP フレームワーク）
- GORM（ORM）
- PostgreSQL 16（メイン DB）
- Redis 7（セッション・キャッシュ）
- Meilisearch（全文検索）

## アーキテクチャ

### フロントエンド

- **App Router** を使用（Pages Router は使わない）
- **Server Component** をデフォルトとし、`"use client"` は最小限に
- **BFF パターン** — `src/app/api/` の API Routes でバックエンドへプロキシ
- Server Actions は使わず BFF 経由でデータを取得・更新する

### バックエンド

- **3 層アーキテクチャ** — Handler → Service → Repository
- Handler: HTTP リクエスト/レスポンスの処理（Gin）
- Service: ビジネスロジック
- Repository: データアクセス（GORM）
- interface は consumer 側に定義（handler に Service interface、service に Repository interface）
- DI はコンストラクタインジェクション（`cmd/server/main.go` で組み立て）

### レイヤー依存方向

```
[Next.js Client Component] → [Next.js API Routes (BFF)] → [Go Handler] → [Service] → [Repository] → [DB]
```

上位層から下位層への一方向のみ許可。逆方向の依存は禁止。

## ディレクトリ構成

- `frontend/src/app/` — ページ・レイアウト・API Routes
- `frontend/src/components/` — 共有コンポーネント（`ui/` は shadcn/ui）
- `frontend/src/hooks/` — カスタム hooks
- `frontend/src/lib/` — ユーティリティ・API クライアント
- `frontend/src/types/` — 型定義
- `backend/cmd/server/` — エントリーポイント
- `backend/internal/handler/` — HTTP ハンドラ
- `backend/internal/service/` — ビジネスロジック
- `backend/internal/repository/` — データアクセス
- `backend/internal/model/` — ドメインモデル・DTO
- `backend/internal/middleware/` — ミドルウェア
- `backend/internal/config/` — 環境変数
- `backend/migrations/` — DB マイグレーション

## コーディング規約

### フロントエンド

- コンポーネント: PascalCase（`UserProfile.tsx`）
- 変数・関数: camelCase
- 型: PascalCase（`type UserProfile = {...}`）
- ESLint + Prettier の設定に従う
- import 順序: React → ライブラリ → 内部モジュール → 型

### バックエンド

- ファイル名: snake_case（`item_handler.go`）
- Exported: PascalCase（`ItemHandler`）
- Unexported: camelCase
- error は wrap して返す（`fmt.Errorf("...: %w", err)`）
- `err.Error()` をクライアントに返さない（内部情報漏洩防止）
- panic 禁止
- golangci-lint の設定に従う
- ログは `log/slog` を使用

## ブランチ戦略

- main: 安定ブランチ
- feature/<issue-number>-<description>: 機能開発

## コミットメッセージ

Conventional Commits に準拠:

```
<type>(<scope>): <subject>

type: feat, fix, docs, style, refactor, test, chore, build, ci, perf, revert
scope: frontend, backend, infra, ci, docs（省略可）
```

## MCP サーバー

利用可能な場合は MCP ツールを優先して使用する:

- **Figma MCP** — デザインデータ参照・デザイントークン取得・コード生成
- **PostgreSQL MCP (DBHub)** — DB スキーマ参照・クエリ実行・データ分析
- **Playwright MCP** — ブラウザ操作・スクリーンショット・E2E テスト

## テスト

- フロントエンド: Vitest + React Testing Library
- バックエンド: Go testing（`*_test.go`）
- E2E: Playwright
- テストファイルは対象ファイルと同じディレクトリに `__tests__/` または `_test.go` で配置
