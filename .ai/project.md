# 🌀 Cinaries — Container Image Binaries

> Pronounced: /ˈsaɪ.nə.riːz/ (“sajnaris”)

**Cinaries** is an open-source initiative to reduce container image bloat, download times, and carbon footprint by distributing **standalone binaries** as **minimal OCI images**.

These images contain **only the binary**, fully verified, versioned, and tagged using a consistent pattern — no layers, no extras, no surprises.

---

## 🎯 Goals

- 🌍 Reduce network load and carbon emissions
- 🧠 Improve developer experience with slim, predictable binary containers
- 🔒 Secure distribution with GPG + checksum validation
- 📦 Consistent tagging and discovery for all supported tools
- ⚙️ Automation-first with powerful metadata and changelog tracking

---

## 🧩 Key Features

✅ Minimal OCI images (based on `SCRATCH`)  
✅ Only one binary per image (e.g. `/usr/bin/composer`)  
✅ Verified via upstream **SHA256 and GPG**  
✅ Tagging: `cinaries/<name>:<version>` — supports `major`, `minor`, `patch`, and `latest`  
✅ Deterministic build system with changelogs  
✅ Optional container image signing (cosign, Sigstore)  
✅ Automation CLI to fetch metadata, generate Dockerfiles, validate checksums, and push images  
✅ JSON response verification using JWS for trustable metadata

---

## 🚀 Initial Supported Tools

| Tool       | Description                             |
|------------|-----------------------------------------|
| **Composer** | PHP dependency manager (single binary) |
| **Node.js**  | Popular JavaScript runtime             |
| **npm**      | Node package manager                   |
| **pnpm**     | Fast & disk-efficient npm alternative  |
| **Yarn**     | Modern Node package manager            |
| **Bun**      | JS runtime with fast tooling           |
| **kubectl**  | Kubernetes CLI                         |
| **jq**       | Lightweight JSON processor             |
| **wget**     | Classic file downloader                |
| **mkcert**   | Local CA for development certificates  |

> Want to suggest a binary? [Open an issue](https://github.com/your-org/cinaries/issues)

---

## 📊 Why It Matters

```dockerfile
# Even this...
COPY --from=composer /usr/bin/composer /usr/bin/composer
```

...forces Docker to pull **hundreds of MBs**.  
With **Cinaries**, you only pull a ~2MB image — verified, fast, and cache-friendly.

---

## 🔐 Security

- Every binary is verified:
  - ✅ Against official **SHA256 checksums**
  - 🔐 Via **GPG signatures** (optional)
- Optional signed JSON metadata (JWS)
- Future support for **container image signing**

---

## 📁 Project Structure

```
dist/
  composer/
    2/
      2.2.24/
        Dockerfile
        README.md (changelog)
packages/
  composer/
    Dockerfile (base)
    repository.txt
    metadata.json
```

---

## 📈 Metrics (coming soon)

- Image size comparison
- Bandwidth savings
- CO₂ savings estimate per pull
- Cache performance stats

---

## ⚒️ Under Development

- ✅ Composer prototype (2.2.24: ~2.42MB vs 186MB official)
- ⚙️ Full CLI automation system
- 🖊️ JWS metadata signature service
- 📦 GitHub release/tag integration
- 🔐 GPG & checksum validators

---

## 📜 License

[MIT License](./LICENSE)

---

## 🙌 Contributing

PRs and issues welcome! Let’s reduce container waste together.
