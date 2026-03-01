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
  log_warn "Docker Compose not available. Start services manually: docker compose up -d"
fi
echo ""

# Step 5: Database migration
log_info "Running database migrations..."
if command -v migrate &> /dev/null; then
  cd backend
  make migrate-up || log_warn "Migration failed. Ensure PostgreSQL is running."
  cd ..
  log_success "Database migrations applied."
else
  log_warn "golang-migrate not installed. Install: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
  log_warn "Then run: cd backend && make migrate-up"
fi
echo ""

# Step 6: MCP server setup
log_info "Checking MCP server configuration..."
if [ -f ".mcp.json" ]; then
  log_success ".mcp.json already exists."
else
  log_warn ".mcp.json not found. It should be in the project root."
fi

# Check for FIGMA_TOKEN
if [ -z "${FIGMA_TOKEN:-}" ]; then
  log_warn "FIGMA_TOKEN is not set. Set it for Figma MCP integration:"
  echo "  export FIGMA_TOKEN=your_token_here"
fi
echo ""

# Step 7: Claude plugins
log_info "Installing web-development-claude-plugins..."
if command -v claude &> /dev/null; then
  claude plugin marketplace add inoue0124/web-claude-plugins 2>/dev/null || log_warn "Marketplace registration failed (may already be registered)"

  for plugin in spec-driven-dev conventions web-architecture testing github-workflow code-review-assist onboarding; do
    claude plugin install "$plugin" --scope project 2>/dev/null || log_warn "Plugin $plugin installation skipped (may already be installed)"
  done
  log_success "Claude plugins installed."
else
  log_warn "Claude Code CLI not installed. Install: npm install -g @anthropic-ai/claude-code"
  log_warn "Then run this script again to install plugins."
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
