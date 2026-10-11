#!/usr/bin/env bash
# Regenerate the Go client from spec/swagger.json.
# Uses the same openapi-generator image as CI, so local output matches CI output.
# Requires: docker, python3, go (for post-generation tidy/vet; CI does its own build).
set -euo pipefail

cd "$(dirname "$0")/.."

IMAGE="openapitools/openapi-generator-cli:v7.12.0"

if ! docker image inspect "$IMAGE" >/dev/null 2>&1; then
  echo "Pulling $IMAGE ..."
  docker pull "$IMAGE"
fi

# swag leaves a few operations without responses; openapi-generator refuses
# those. Backfill before generating (no-op once the backend annotations land).
python3 scripts/normalize-spec.py spec/swagger.json

# The generator writes but never cleans; stale files from earlier specs would
# linger and break the build. Wipe generated output first (README.md, LICENSE,
# scripts/, spec/, .github/ and the generator-ignore file are hand-maintained
# or protected and stay).
rm -f ./*.go go.mod go.sum
rm -rf api docs test

docker run --rm \
  -u "$(id -u):$(id -g)" \
  -v "$PWD:/work" \
  -w /work \
  "$IMAGE" \
  generate \
    -i /work/spec/swagger.json \
    -g go \
    -o /work \
    -c /work/openapi-generator-config.json

# The generator writes placeholder module paths; pin ours everywhere.
grep -rl "GIT_USER_ID/GIT_REPO_ID" --include="*.go" --include="go.mod" . 2>/dev/null |
  while read -r f; do
    sed -i.bak 's|github.com/GIT_USER_ID/GIT_REPO_ID|github.com/bare-metal-cloud/client-go|g' "$f"
    rm -f "$f.bak"
  done

# Generated tests import testify; keep the module graph honest.
go mod tidy

echo
echo "Regenerated. Review with: git status && git diff"
