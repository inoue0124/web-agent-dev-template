# Web Agent Dev Template

Web アプリ開発を Claude Code（AI エージェント）と共に進めるためのスターターテンプレート。

このリポジトリをクローンしてセットアップスクリプトを実行するだけで、Next.js + Go/Gin の Web 開発環境と AI エージェント開発をすぐに始められる環境が整います。

## 概要

このテンプレートは以下を提供します。

- **ビルド可能な Next.js + Go/Gin のベースアプリ** — サンプル CRUD 付きで、セットアップ直後にビルド・実行できる
- **開発ツールの自動セットアップ** — ESLint、Prettier、golangci-lint、Docker Compose を一括セットアップ
- **[web-development-claude-plugins](https://github.com/inoue0124/web-claude-plugins) の導入** — コードレビュー、規約統一、テスト生成、オンボーディングなど Web チーム開発を包括サポートするプラグイン群
- **MCP サーバーの設定** — Figma MCP / PostgreSQL MCP / Playwright MCP で AI エージェントからデザイン参照・DB 操作・ブラウザ操作を直接実行可能に
- **チーム開発の基盤** — GitHub テンプレート、CI ワークフロー、pre-commit hook、コーディング規約設定

## クイックスタート

### 事前に必要なもの

以下は setup.sh では自動インストールされません。事前にインストールしてください。

| ツール | インストール方法 |
|---|---|
| Node.js 20+ | https://nodejs.org |
| Go 1.23+ | https://go.dev/dl |
| Docker / Docker Compose | https://docs.docker.com/get-docker |
| Claude Code | `npm install -g @anthropic-ai/claude-code`（[公式ドキュメント](https://docs.anthropic.com/en/docs/claude-code)） |

### セットアップ手順

```bash
# 1. テンプレートをクローン
git clone https://github.com/inoue0124/web-agent-dev-template.git <your-project-name>
cd <your-project-name>

# 2. セットアップスクリプトを実行
./scripts/setup.sh

# 3. 開発サーバーを起動
docker compose up -d            # PostgreSQL + Redis + Meilisearch
cd backend && go run ./cmd/server &  # Go API サーバー
cd frontend && npm run dev      # Next.js 開発サーバー

# 4. AI エージェントと開発スタート
claude
```

## セットアップスクリプトが行うこと

`scripts/setup.sh` は以下を順に実行します。

1. **前提条件の確認** — Node.js / Go / Docker / gh CLI
2. **フロントエンド**
   - `cd frontend && npm install`
3. **バックエンド**
   - `cd backend && go mod download`
4. **インフラ**
   - `docker compose up -d`（PostgreSQL + Redis + Meilisearch）
   - DB マイグレーションの実行
5. **MCP サーバーの自動セットアップ**
   - `.mcp.json` に Figma MCP / PostgreSQL MCP / Playwright MCP を設定
6. **web-development-claude-plugins のインストール**
   - マーケットプレース登録 + 全 7 プラグインを自動インストール
7. **Git hooks のインストール** — pre-commit / commit-msg

## セットアップ後の開発ワークフロー

セットアップ完了後、`claude` コマンドで AI エージェントと対話しながら開発を進められます。

### 作るものがまだ決まっていない場合 — スペック駆動開発

アイデアはあるが詳細が固まっていない段階では、スペック駆動開発スキルを使ってドキュメントから先に作成します。

```
1. docs/ideas/ にアイデアメモを置く（箇条書き・雑なメモでOK）
2. /prd-writing でプロダクト要求定義書を作成（ユーザー承認あり）
   ↓ 承認後、以下は自動で連鎖生成
3. /functional-design で機能設計書を作成
4. /architecture-design でアーキテクチャ設計書を作成
5. /glossary-gen で用語集を作成
```

生成されたドキュメントは `docs/` 配下に配置されます。

### 作るものが決まっている場合 — フィーチャー実装

`/implement-feature` で要件定義から実装までをスペック駆動で一気通貫に進められます。

```
1. /implement-feature で実装ワークフローを開始
   ↓ 要件定義書の生成 — 機能要件・非機能要件・受け入れ条件を整理
   ↓ 詳細設計書の生成 — コンポーネント設計・API 設計・レイヤー設計
   ↓ タスクリストの生成 — 詳細設計から実装タスクを分解
2. 各タスクを順に実装
   ↓ アーキテクチャ検査が Server/Client Component 分離や Go 3 層を自動チェック
   ↓ コード品質チェックが lint / format を実行
   ↓ テスト生成が Vitest / Go test のテストを自動生成
3. コミット時に pre-commit hook が最終チェック
4. PR 作成時にレビュー支援が差分を分析
```

### 日常の開発でよく使うコマンド

| やりたいこと | Claude への指示例 |
|---|---|
| 新機能をスペック駆動で実装 | `/implement-feature` |
| テストを生成 | `/unit-test-gen` |
| テストを実行 | `/test-run` |
| コードレビュー | `/pr-review` |
| PR を作成 | `/pr-create` |
| Issue を作成 | `/issue-create` |
| 規約チェック | `/lint-check`, `/format-check` |
| アーキテクチャ検査 | `/component-check`, `/go-layer-check` |
| 新メンバー向けガイド | `/codebase-overview`, `/architecture-map` |

## ディレクトリ構成

```
<your-project-name>/
├── frontend/
│   ├── src/
│   │   ├── app/
│   │   │   ├── layout.tsx              # Root Layout (Server Component)
│   │   │   ├── page.tsx                # トップページ
│   │   │   ├── items/                  # サンプル Feature Module
│   │   │   │   ├── page.tsx            # 一覧 (Server Component)
│   │   │   │   ├── [id]/page.tsx       # 詳細 (Server Component)
│   │   │   │   └── new/page.tsx        # 作成 (Client Component)
│   │   │   └── api/                    # BFF (API Routes)
│   │   ├── components/
│   │   │   └── ui/                     # shadcn/ui コンポーネント
│   │   ├── hooks/                      # カスタム hooks
│   │   ├── lib/                        # ユーティリティ・API クライアント
│   │   ├── types/                      # 型定義
│   │   └── styles/                     # グローバルスタイル
│   ├── public/
│   ├── next.config.ts
│   ├── tsconfig.json
│   ├── eslint.config.js
│   └── package.json
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go                 # エントリーポイント + DI
│   ├── internal/
│   │   ├── handler/                    # HTTP ハンドラ（Gin）
│   │   ├── service/                    # ビジネスロジック
│   │   ├── repository/                 # データアクセス
│   │   ├── model/                      # ドメインモデル + DTO
│   │   ├── middleware/                 # 認証・ログ・CORS
│   │   └── config/                     # 環境変数読み込み
│   ├── migrations/                     # DB マイグレーション
│   ├── go.mod
│   └── .env.example
├── docker-compose.yml                  # PostgreSQL + Redis + Meilisearch
├── scripts/
│   ├── setup.sh                        # 初回セットアップ
│   └── hooks/
│       ├── pre-commit                  # コミット前の自動 lint
│       └── commit-msg                  # Conventional Commits チェック
├── .github/
│   ├── ISSUE_TEMPLATE/
│   │   ├── bug_report.md
│   │   ├── feature_request.md
│   │   └── task.md
│   ├── PULL_REQUEST_TEMPLATE.md
│   ├── workflows/
│   │   └── ci.yml                      # PR 時の自動 lint・build・test
│   └── dependabot.yml                  # npm + Go modules 自動更新
├── .mcp.json                           # MCP サーバー設定
├── CLAUDE.md                           # AI エージェントへの指示書
├── .editorconfig
└── .gitignore
```

> サンプル Feature Module（`frontend/src/app/items/` + `backend/internal/handler/item_handler.go`）は CRUD の実装例です。新機能追加時の参考にしてください。

## 導入されるプラグイン

[web-development-claude-plugins](https://github.com/inoue0124/web-claude-plugins) から以下のプラグインが導入されます。

### Tier 0: プロジェクト立ち上げ時に最初だけ使うもの

| プラグイン | 説明 |
|---|---|
| spec-driven-dev | 仕様駆動開発 — PRD・設計・ガイドライン生成から実装まで一気通貫 |

### Tier 1: 日常の開発サイクルで毎日使うもの

| プラグイン | 説明 |
|---|---|
| conventions | 規約統一 — lint・format・命名規則・ファイル配置の検査と自動注入 |
| web-architecture | アーキテクチャ検査 — コンポーネント設計・Go 3 層・レイヤー分離 |
| testing | テスト — Vitest / Go testing / Playwright によるテスト生成・実行 |
| github-workflow | GitHub ワークフロー — Issue・PR・リリースノートの構造化作成 |
| code-review-assist | レビュー支援 — 差分分析・影響範囲特定・セキュリティチェック |

### Tier 2: チーム拡大・新メンバー参入時に使うもの

| プラグイン | 説明 |
|---|---|
| onboarding | オンボーディング — プロジェクト構造・アーキテクチャ・変更履歴の解説 |

## Git hooks

| フック | 内容 |
|---|---|
| `pre-commit` | ESLint + Prettier + golangci-lint を自動実行（ステージング済みファイルのみ） |
| `commit-msg` | コミットメッセージのフォーマットチェック（Conventional Commits） |

## ライセンス

MIT
