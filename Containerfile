FROM node:24-alpine AS node-builder

RUN corepack enable

WORKDIR /app

COPY package.json pnpm-lock.yaml /app/
RUN --mount=type=cache,target=/pnpm-store \
    pnpm install --frozen-lockfile --store-dir /pnpm-store

COPY vite.config.ts /app
COPY public /app/public
COPY resources /app/resources
RUN pnpm build

FROM golang:1.27-alpine AS go-builder

RUN apk add --no-cache git

ENV WORKDIR=/app
WORKDIR /app

COPY go.mod go.sum /app/
RUN --mount=type=cache,target=/root/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=secret,id=netrc,dst=/root/.netrc \
    go mod download

COPY main.go /app/
COPY vite /app/vite
COPY --from=node-builder /app/vite/dist /app/vite/dist
COPY --from=node-builder /app/resources/views/index.html /app/resources/views/index.html

RUN --mount=type=cache,target=/root/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -ldflags="-s -w -buildid=" -trimpath -o /app/server /app/main.go


FROM gcr.io/distroless/static:nonroot AS web

WORKDIR /app
COPY --from=go-builder /app/server /app/server

ENTRYPOINT [ "/app/server" ]
