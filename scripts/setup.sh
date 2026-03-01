#!/usr/bin/env bash
set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

log_info() { echo -e "${BLUE}[INFO]${NC} $1"; }
log_success() { echo -e "${GREEN}[OK]${NC} $1"; }
log_warn() { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }

echo ""
echo "=========================================="
echo "  Web Agent Dev Template - Setup"
echo "=========================================="
echo ""

# Step 1: Check prerequisites
log_info "Checking prerequisites..."

check_command() {
  if command -v "$1" &> /dev/null; then
    log_success "$1 is installed ($(command -v $1))"
    return 0
  else
    log_error "$1 is not installed."
    echo "  Install: $2"
    return 1
  fi
}

MISSING=0
check_command "node" "https://nodejs.org" || MISSING=1
check_command "go" "https://go.dev/dl" || MISSING=1
check_command "docker" "https://docs.docker.com/get-docker" || MISSING=1
check_command "gh" "https://cli.github.com" || MISSING=1

# Check Node.js version >= 20
if command -v node &> /dev/null; then
  NODE_VERSION=$(node -v | sed 's/v//' | cut -d. -f1)
  if [ "$NODE_VERSION" -lt 20 ]; then
    log_error "Node.js 20+ is required (found v$(node -v))"
    MISSING=1
  fi
fi

if [ "$MISSING" -eq 1 ]; then
  log_error "Please install missing prerequisites and re-run this script."
  exit 1
fi
log_success "All prerequisites met!"
echo ""

# Step 2: Frontend setup
log_info "Setting up frontend..."
cd frontend
npm install
log_success "Frontend dependencies installed."
cd ..
echo ""

# Step 3: Backend setup
log_info "Setting up backend..."
cd backend
go mod download
log_success "Backend dependencies downloaded."
cd ..
echo ""

# Step 4: Infrastructure
log_info "Starting infrastructure services..."
if docker compose version &> /dev/null; then
  docker compose up -d
  log_success "Docker services started (PostgreSQL, Redis, Meilisearch)."
else
  log_info "Docker Compose not available. Start services manually: docker compose up -d"
fi
echo ""

# Step 5: Database migration
log_info "Running database migrations..."

# Ensure GOPATH/bin is in PATH
GOPATH_BIN="$(go env GOPATH)/bin"
if [[ ":$PATH:" != *":$GOPATH_BIN:"* ]]; then
  export PATH="$PATH:$GOPATH_BIN"
fi

# Install golang-migrate if not available
if ! command -v migrate &> /dev/null; then
  log_info "Installing golang-migrate..."
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  log_success "golang-migrate installed."
fi

if command -v migrate &> /dev/null; then
  cd backend
  if make migrate-up; then
    log_success "Database migrations applied."
  else
    log_info "Migration skipped. Ensure PostgreSQL is running and retry: cd backend && make migrate-up"
  fi
  cd ..
else
  log_info "golang-migrate installation failed. Run manually: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
fi
echo ""

# Step 6: MCP server setup
log_info "Checking MCP server configuration..."

# Load .env file if it exists
if [ -f ".env" ]; then
  export $(grep -v '^\s*#' .env | grep -v '^\s*$' | xargs)
  log_success ".env file loaded."
fi

if [ -f ".mcp.json" ]; then
  log_success ".mcp.json already exists."
else
  log_info ".mcp.json not found. It should be in the project root."
fi

# Validate FIGMA_TOKEN against Figma API
if [ -n "${FIGMA_TOKEN:-}" ]; then
  HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -H "X-Figma-Token: $FIGMA_TOKEN" "https://api.figma.com/v1/files/placeholder?depth=1" 2>/dev/null || echo "000")
  if [ "$HTTP_STATUS" = "404" ] || [ "$HTTP_STATUS" = "200" ]; then
    log_success "FIGMA_TOKEN is valid. Figma MCP server is configured."
  else
    log_info "FIGMA_TOKEN is set but may be invalid (HTTP $HTTP_STATUS). Verify your token at:"
    echo "  https://www.figma.com/settings → Personal access tokens"
  fi
else
  log_info "FIGMA_TOKEN is not set. Set it in .env for Figma MCP integration:"
  echo "  FIGMA_TOKEN=your_token_here"
fi
echo ""

# Step 7: Claude plugins
log_info "Installing web-development-claude-plugins..."
if command -v claude &> /dev/null; then
  claude plugin marketplace add inoue0124/web-claude-plugins &>/dev/null || true

  for plugin in spec-driven-dev conventions web-architecture testing github-workflow code-review-assist onboarding; do
    claude plugin install "$plugin" --scope project &>/dev/null || true
  done
  log_success "Claude plugins are ready."
else
  log_info "Claude Code CLI not installed. Install: npm install -g @anthropic-ai/claude-code"
  log_info "Then run this script again to install plugins."
fi
echo ""

# Step 8: Git hooks
log_info "Setting up Git hooks..."
git config core.hooksPath scripts/hooks/
chmod +x scripts/hooks/* 2>/dev/null || true
log_success "Git hooks configured (scripts/hooks/)."
echo ""

# Done
echo "=========================================="
echo -e "  ${GREEN}Setup complete!${NC}"
echo "=========================================="
echo ""
echo "Quick start:"
echo "  docker compose up -d            # Start services"
echo "  cd backend && go run ./cmd/server &  # Start API server"
echo "  cd frontend && npm run dev      # Start frontend"
echo "  claude                          # Start AI agent"
echo ""
