#!/bin/bash

echo "Installing SkibidiLang for Linux/macOS..."

TARGET_DIR="$HOME/.skibidi"
EXECUTABLE="skibidi"

# Create directory
mkdir -p "$TARGET_DIR"

# Move the binary
cp "$EXECUTABLE" "$TARGET_DIR/$EXECUTABLE"
chmod +x "$TARGET_DIR/$EXECUTABLE"

# Add to PATH via shell profile
PROFILE=""
if [ -n "$ZSH_VERSION" ]; then
  PROFILE="$HOME/.zshrc"
elif [ -n "$BASH_VERSION" ]; then
  PROFILE="$HOME/.bashrc"
elif [ -f "$HOME/.profile" ]; then
  PROFILE="$HOME/.profile"
fi

if ! grep -q "$TARGET_DIR" "$PROFILE"; then
  echo "export PATH=\"$TARGET_DIR:\$PATH\"" >> "$PROFILE"
  echo "Added SkibidiLang to PATH in $PROFILE"
else
  echo "SkibidiLang already in PATH"
fi

echo "Installation complete. Please restart your terminal or run: source $PROFILE"
