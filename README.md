# Discovery — Go + Svelte + Inertia, single-binary app

Go server-rendered Svelte app, glued via Inertia. Single static binary — Vite assets embedded at build time, no separate frontend server in production.

## Stack

- **Go** (module `govite`) + [Gonertia](https://github.com/romsar/gonertia) v3 — Inertia.js protocol adapter for Go
- **Svelte 5** + TypeScript, via `@inertiajs/svelte`
- **Vite 8** — dev server + production asset bundling, embedded into the Go binary with `go:embed` (see `vite/`)

## Why this setup

- **Vite embedding** — production assets (`vite/dist`) get `go:embed`-ed straight into the Go binary. One artifact to ship, no CDN/nginx/static-host needed, no path/asset drift between server and build.
- **Inertia** — server-rendered routing/data, SPA-like page transitions client-side. No REST/GraphQL API layer to hand-build; Go controllers return page components + props directly.
- **air hot reload** — rebuilds and restarts the Go binary on `.go`/`.html` change while dev keeps Vite HMR for Svelte/CSS. Full-stack loop stays fast without manual restarts.
- **Distroless final image** — `gcr.io/distroless/static:nonroot` has no shell, no package manager, minimal attack surface, small image size.
- **Multi-stage container build** — node stage only exists to produce assets; final image has no node/npm/pnpm at all, just the static Go binary.
- **Single static binary** (`CGO_ENABLED=0`) — no runtime deps, trivial to deploy/copy anywhere Linux runs.

## Layout

```
main.go                  HTTP server, wires Inertia + Vite middleware
vite/
  vite.go                production: serves embedded dist via manifest
  vite_dev.go            dev: proxies to Vite dev server (build tag: hot)
resources/
  views/index.html       Inertia root template
  js/
    app.ts               Inertia app entry
    layouts/layout.svelte
    pages/index.svelte
  css/app.css
public/                  static assets (favicon, icons)
```

## Dev

```bash
pnpm install
air              # hot-reload Go server (tags: hot) + Vite dev middleware
```

`air` rebuilds on `.go`/`.html` changes per `.air.toml`. Vite serves assets in dev mode via `vite_dev.go`.

Frontend only:

```bash
pnpm dev         # vite dev server
pnpm check       # svelte-check + tsc
```

## Build

```bash
pnpm build                 # vite build -> resources/views + build manifest
go build -o server .       # embeds dist, produces single binary
```

Or via container:

```bash
podman build -t discovery .
```

Multi-stage: `pnpm build` in a node stage, then `go build` embeds the resulting `vite/dist` and rendered `index.html`, landing on `gcr.io/distroless/static:nonroot`.

## Run

```bash
./server        # listens on :8080
```
