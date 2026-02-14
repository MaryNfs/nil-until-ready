#!/usr/bin/env bash
set -euo pipefail

if [[ -z "${1:-}" ]]; then
  echo "Usage: scripts/newlc.sh <name> [base_dir]"
  exit 1
fi

NAME="$1"
BASE_DIR="${2:-go/Leetcode}"
TEMPLATE_DIR="scripts/templates"

DIR="${BASE_DIR}/${NAME}"

if [[ -d "$DIR" ]]; then
  echo "❌ Directory already exists: $DIR"
  exit 1
fi

mkdir -p "$DIR"

render() {
  local template="$1"
  local output="$2"
  sed "s/{{NAME}}/${NAME}/g" "$template" > "$output"
}

render "${TEMPLATE_DIR}/solution.go.tmpl" \
       "${DIR}/${NAME}.go"

render "${TEMPLATE_DIR}/solution_test.go.tmpl" \
       "${DIR}/${NAME}_test.go"

render "${TEMPLATE_DIR}/q.md.tmpl" \
       "${DIR}/q.md"

echo "✅ Created:"
echo "  $DIR"
