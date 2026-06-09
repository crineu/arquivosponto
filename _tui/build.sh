#!/usr/bin/env bash
# =============================================================================
# build.sh - Build script for the Stow TUI
# =============================================================================
# Gera um binário estático, universal e otimizado para Linux.
# Funciona em diversas distribuições sem dependências externas.
# =============================================================================

set -euo pipefail

cd "$(dirname "$0")"

BINARY_NAME="tui"
OUTPUT="../${BINARY_NAME}"
VERSION="${VERSION:-$(date +%Y%m%d)}"

echo "🚀 Building Stow TUI..."

# Limpeza
rm -f "${OUTPUT}" "${OUTPUT}-static" 2>/dev/null || true

# Build principal (estaticamente linkado)
echo "📦 Building fully static binary..."
CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64 \
go build \
  -trimpath \
  -ldflags="-s -w -buildid= -X main.buildVersion=${VERSION} -extldflags=-static" \
  -tags=netgo,osusergo \
  -o "${OUTPUT}" \
  .

# Informações do build
echo "✅ Build completed successfully!"
echo "📍 Output: ${OUTPUT}"
echo "📏 Size: $(ls -lh "${OUTPUT}" | awk '{print $5}')"

# Verifica se realmente está estático
if command -v ldd >/dev/null 2>&1; then
    if ldd "${OUTPUT}" 2>&1 | grep -q "not a dynamic executable"; then
        echo "📋 Dependencies: none (fully static ✓)"
    else
        echo "⚠️  Warning: Binary is not fully static"
    fi
fi

echo ""
echo "🎯 Usage:"
echo "   cd .. && ./tui                    # Run the TUI"
echo "   cd _tui && ./build.sh             # Rebuild"
echo ""
echo "💡 Tip: Run with HELP_DEBUG=1 ./tui to enable debug logging."
