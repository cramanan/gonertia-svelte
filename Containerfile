FROM ghcr.io/voidzero-dev/vite-plus:1.5.0 AS node-builder

# Build image, not a clone: keep the `prepare` script (vp config) from installing git hooks.
ENV VP_GIT_HOOKS=0

WORKDIR /app

COPY --chown=vp:vp package.json pnpm-lock.yaml pnpm-workspace.yaml .node-version /app/
RUN --mount=type=cache,target=/pnpm-store,uid=1000,gid=1000 \
    npm_config_store_dir=/pnpm-store vp install --frozen-lockfile

COPY --chown=vp:vp vite.config.ts svelte.config.js /app/
COPY --chown=vp:vp public /app/public
COPY --chown=vp:vp resources /app/resources
RUN vp build

FROM golang:1.27-alpine AS go-builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum /app/
RUN --mount=type=cache,target=/root/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=secret,id=netrc,dst=/root/.netrc \
    go mod download

COPY main.go /app/main.go
COPY internal /app/internal
COPY --from=node-builder /app/public public
COPY --from=node-builder /app/resources/views/index.html resources/views/index.html

RUN --mount=type=cache,target=/root/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -ldflags="-s -w -buildid=" -trimpath -o /app/server .


FROM gcr.io/distroless/static:nonroot

WORKDIR /app
COPY --from=go-builder /app/server /app/server

# With THP on "always", the kernel backs the Go heap with 2 MB huge pages:
# idle RSS jumps between ~5 and ~7 MB per start instead of ~2.5 MB.
ENV GODEBUG=disablethp=1

ENTRYPOINT [ "/app/server" ]
