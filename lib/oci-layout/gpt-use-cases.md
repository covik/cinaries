Love the “local-only, producer-first” stance. Here are **strong real-world use cases** for a **pure OCI layout** library—no registries involved.

# Where this shines

## Dev/CI & hermetic builds

- **Air-gapped pipelines:** preload OCI layouts; all steps read from local CAS only.
- **Reproducible builds:** lock by digest; verify content, fail on drift.
- **Test fixtures:** tiny, deterministic layouts for unit/integration tests (no network).
- **Layer surgery:** flatten, rebase, or strip layers to craft minimal artifacts.

## Supply chain / compliance

- **Referrers graph offline:** enumerate & attach SBOMs, attestations, signatures.
- **Policy gates:** allow/deny by mediaType/digest; prove provenance offline.
- **Forensics:** snapshot layouts, diff digests, audit unexpected bytes.

## Packaging / distribution (offline media)

- **Portable bundles:** USB/DVD export/import of full repos as OCI layout.
- **Relocation:** rewrite names/annotations without touching content blobs.
- **Format bridges:** docker-archive ⇆ OCI layout; tarballs ⇆ layout.

## Runtime prep / extraction

- **Selective export:** grab `/usr/bin/tool` (handle whiteouts/opaque dirs correctly).
- **Whole rootfs merge:** apply layers locally to a directory or tar stream.
- **Multi-platform pick:** resolve `os/arch/variant` manifests offline.

## Performance & storage

- **Dedup metrics:** compute byte savings across images (shared blob accounting).
- **GC/pruning:** drop unreachable blobs; compact layouts safely.
- **Recompression/optimization:** convert to estargz/zstd-seekable for faster loads (still local).

## Edge / embedded / enterprise

- **Golden images:** curate internal catalogs as OCI layouts (no registry dependency).
- **Field updates:** ship delta layers to sites with poor connectivity.
- **Signed drops:** distribute signed, verified bundles to regulated environments.

# Core capabilities these imply (no code—just shape)

- **Index & manifest traversal:** resolve tag/digest → platform manifest → config.
- **CAS ops:** list, stat, copy, hard-link/clone blobs between layouts.
- **Rootfs apply/extract:** correct whiteout semantics; deterministic merges.
- **Referrers API:** list/add/remove signatures, SBOMs, SLSA, custom artifacts.
- **Normalization:** validate layout, fix mediaTypes, canonicalize JSON.
- **Conversion:** docker-archive/tar ⇆ oci-layout; optional stargz/zstd.
- **Integrity:** sha256 verify, size match, provenance record (digests, paths).
- **GC & compaction:** mark-and-sweep unreachable; pack for transport.

# Nice non-goals (to keep it tight)

- No registry pulls/pushes.
- No runtime (no containerd/dockerd integration).
- No signing backends (provide hooks, not implementations).
- No format zoo beyond OCI + docker-archive + tar.

If you want, I can draft a one-pager “API surface” (types + function names only) reflecting these—still zero code, just the contract.
