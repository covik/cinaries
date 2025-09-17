## OCI Layout Tools

`layoci` — a single CLI for handling local [OCI Image Layouts](https://specs.opencontainers.org/image-spec/image-layout/), backed by the internal `oci-layout` library.

Usage

- Build: `pnpm --filter @repo/oci-layout build`
- Run: `./lib/oci-layout/out/layoci --help`

Notes

- Project/module remains `oci-layout` for clarity and spec alignment.
- CLI binary name is `layoci` for brevity (Layout + OCI).
