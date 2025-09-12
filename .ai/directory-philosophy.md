# Filesystem Hierarchy Standard

### Adapted for the **Cinaries** project

---

## `/etc`

**Official FHS:** Host-specific system-wide configuration files.  
**Adapted:** Centralized, project-wide configuration for the monorepo’s toolchain and policies. Reusable, shared configuration packages that affect the entire codebase.

- **Examples:** `eslint-config`, `tsconfig`, `renovate`

---

## `/lib`

**Official FHS:** Essential shared libraries and kernel modules.  
**Adapted:** Shared, reusable helper libraries and utilities meant to be consumed by other apps and packages in the monorepo.

- **Examples:** Shared UI component library, common data-fetching utility, types package

---

## `/var`

**Official FHS:** Variable data (logs, spool files, temporary data) that may change.  
**Adapted:** Logical place for build artifacts, cache files, and logs. Stores the output of your build processes.

- **Examples:** `.turbo`, `.pnpm-store`

---

## `/tmp`

**Official FHS:** Temporary files.  
**Adapted:** Short-lived temporary files and directories used during the build process. Often handled automatically by tools like Turborepo’s caching.

---

## `/out`

**Official FHS:** _Not part of the standard hierarchy._ Common in C codebases.
**Adapted:** A dedicated directory for production-ready output artifacts. Unlike `/var`, which may hold intermediate or temporary build data, `/out` is meant for finalized bundles, deployable packages, or static exports.

- **Examples:** Compiled TypeScript files
