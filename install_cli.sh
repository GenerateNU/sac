#!/bin/zsh

set -e
set -o pipefail

# Get the absolute path to the current script
SCRIPT_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

cd "$SCRIPT_PATH/cli"

# Check for Windows or Unix-like system
if [[ $(uname -s) == "MINGW64" || $(uname -s) == "CYGWIN" || $(uname -s) == "MSYS" ]]; then
  export GOPATH="%USERPROFILE%\go"
else
  export GOPATH="$HOME/go"
fi

# Build the executable
go build -o "sac"

mkdir -p "$GOPATH/bin"

mv "sac" "$GOPATH/bin"

echo "Installation complete. You can now run 'sac'"
