#!/usr/bin/env bash
set -euo pipefail

# Defaults can be overridden via env vars
: "${IMAGE:=docker.io/bitnami/kubectl}"
: "${TAG:=1.30.0}"
: "${DIGEST:=sha256:09c66fb4719def5c6b495dd7fe69dc580867b905c9470e4354d5a862310e6ffd}"
: "${OUT:=kubectl-src}"

OUTDIR="${PWD}/${OUT}"

echo "Fetching ${IMAGE}:${TAG}@${DIGEST} to: $OUTDIR"
skopeo copy \
  --format=oci \
  --multi-arch=all \
  --dest-oci-accept-uncompressed-layers \
  "docker://${IMAGE}@${DIGEST}" "oci:${OUTDIR}:${TAG}"

git status --short -- "$OUTDIR" || true
