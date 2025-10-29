#!/usr/bin/env bash
set -e

REPO="kurtiz/commit-feed"
INSTALL_DIR="/usr/local/bin"
APP_NAME="commitfeed"
TMP_DIR="$(mktemp -d)"
GH_API="https://api.github.com/repos/$REPO/releases/latest"

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${BLUE}🚀 Installing CommitFeed...${NC}"

# --- Detect OS and Architecture ---
OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
  Linux*)   PLATFORM="linux" ;;
  Darwin*)  PLATFORM="darwin" ;;
  MINGW*|MSYS*|CYGWIN*|Windows*) PLATFORM="windows" ;;
  *) echo -e "${YELLOW}❌ Unsupported OS: $OS${NC}"; exit 1 ;;
esac

case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo -e "${YELLOW}❌ Unsupported architecture: $ARCH${NC}"; exit 1 ;;
esac

# --- Fetch latest release info ---
echo "📦 Fetching latest release info..."
DOWNLOAD_URL=$(curl -sL "$GH_API" \
  | grep "browser_download_url" \
  | grep "$PLATFORM-$ARCH" \
  | grep -v ".sha" \
  | grep -E "commitfeed|cf" \
  | head -n 1 \
  | cut -d '"' -f 4)

if [ -z "$DOWNLOAD_URL" ]; then
  echo -e "${YELLOW}❌ Could not find a compatible binary for ${PLATFORM}-${ARCH}.${NC}"
  exit 1
fi

FILENAME="${DOWNLOAD_URL##*/}"

echo "⬇️  Downloading ${FILENAME}..."
curl -L "$DOWNLOAD_URL" -o "$TMP_DIR/$FILENAME"

# --- Extract or move binary ---
cd "$TMP_DIR"

if [[ "$PLATFORM" == "windows" && "$FILENAME" == *.exe ]]; then
  echo "💾 Detected Windows executable. Skipping extraction."
  BINARY_PATH="$TMP_DIR/$FILENAME"
elif [[ "$FILENAME" == *.zip ]]; then
  echo "📂 Extracting ZIP..."
  unzip -q "$FILENAME"
  BINARY_PATH=$(find . -type f \( -name "$APP_NAME" -o -name "cf" \) | head -n 1)
else
  echo "📂 Extracting TAR..."
  tar -xzf "$FILENAME"
  BINARY_PATH=$(find . -type f \( -name "$APP_NAME" -o -name "cf" \) | head -n 1)
fi

if [ ! -f "$BINARY_PATH" ]; then
  echo -e "${YELLOW}❌ No binary found in archive.${NC}"
  exit 1
fi

chmod +x "$BINARY_PATH"

# --- Install binary ---
echo "📦 Installing to ${INSTALL_DIR}..."
if [[ "$PLATFORM" == "windows" ]]; then
  mkdir -p "$HOME/.local/bin"
  mv "$BINARY_PATH" "$HOME/.local/bin/$APP_NAME.exe"
  INSTALL_PATH="$HOME/.local/bin/$APP_NAME.exe"
else
  sudo mv "$BINARY_PATH" "$INSTALL_DIR/$APP_NAME"
  INSTALL_PATH="$INSTALL_DIR/$APP_NAME"
fi

# --- Verify installation ---
if command -v "$APP_NAME" >/dev/null 2>&1 || [[ -f "$INSTALL_PATH" ]]; then
  echo -e "${GREEN}✅ CommitFeed installed successfully!${NC}"
  echo -e "${BLUE}Run: ${NC}${APP_NAME} --help"
else
  echo -e "${YELLOW}⚠️  Installed but not found in PATH.${NC}"
  echo "Try adding ${INSTALL_DIR} (or ~/.local/bin) to your PATH manually."
fi

# Cleanup
rm -rf "$TMP_DIR"

echo -e "${BLUE}🎉 Done!${NC}"
