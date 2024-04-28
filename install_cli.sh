#!/bin/zsh

set -e
set -o pipefail

# Get the absolute path to the current script
SCRIPT_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

cd "$SCRIPT_PATH/cli"

go build -o "sac"
go install

rm sac

echo "Installation complete. You can now run 'sac' from anywhere."
