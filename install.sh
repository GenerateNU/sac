#!/bin/zsh
set -e
set -o pipefail

# Get the absolute path to the current script
SCRIPT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Build the Go CLI tool
go build -o "sac-cli" "cli/main.go"

# Identify the user's shell
SHELL_NAME=$(basename "$SHELL")

COMMAND_NAME="sac-cli"

if command -v "$COMMAND_NAME" >/dev/null 2>&1; then
    exit 1
fi

# Add sac-cli to the user's PATH
if [[ "$SHELL_NAME" == "zsh" ]]; then
    echo "export PATH=\"$SCRIPT_PATH:\$PATH\"" >> ~/.zshrc
    source ~/.zshrc
elif [[ "$SHELL_NAME" == "bash" ]]; then
    echo "export PATH=\"$SCRIPT_PATH:\$PATH\"" >> ~/.bashrc
    source ~/.bashrc    
else
    echo "Unsupported shell: $SHELL_NAME"
    exit 1
fi

# Inform the user
echo "Installation complete. You can now run 'sac-cli' from anywhere."
