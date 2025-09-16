## Zot OCI Container Registry

Runs a local container registry that mirrors selected images from Docker Hub to avoid rate limits. Mirrors only projects from `src/`.

### Usage

```shell
pnpm --filter @cinaries/zot dev
```

Available at http://localhost:5000.
Uses `/etc/zot/config.json` provided by the devcontainer.
