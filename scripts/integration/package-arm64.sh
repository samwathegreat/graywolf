#!/usr/bin/env bash
# Compatibility entry point for the original ARM64 packaging command.
set -euo pipefail
exec "$(dirname "$0")/package-linux.sh" arm64 "$@"
