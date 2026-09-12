FROM node:24-alpine AS node-builder

RUN corepack enable

WORKDIR /app

COPY package.json pnpm-lock.yaml /app/
RUN --mount=type=cache,target=/pnpm-store \
    pnpm install --frozen-lockfile --store-dir /pnpm-store

COPY vite.config.ts svelte.config.js /app/
COPY gravel-vite-plugin gravel-vite-plugin
COPY public public
COPY resources resources
RUN pnpm build

FROM golang:1.27-alpine AS go-builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum /app/
RUN --mount=type=cache,target=/root/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=secret,id=netrc,dst=/root/.netrc \
    go mod download

COPY main.go /app/
COPY --from=node-builder /app/public public
COPY --from=node-builder /app/resources/views/index.html resources/views/index.html

RUN --mount=type=cache,target=/root/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -ldflags="-s -w -buildid=" -trimpath -o /app/server /app/main.go


FROM gcr.io/distroless/static:nonroot

WORKDIR /app
COPY --from=go-builder /app/server /app/server

ENTRYPOINT [ "/app/server" ]
