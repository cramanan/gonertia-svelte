package main

import (
	"cmp"
	"context"
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gravel/internal/http/middlewares"
	"gravel/internal/http/routes"
	"gravel/internal/http/view"

	"github.com/romsar/gonertia/v3"
)

//go:embed resources/views/index.html
var rootTemplate []byte

//go:embed all:public
var publicFS embed.FS

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	if err := run(ctx, os.Getenv, os.Stdout); err != nil {
		slog.Error("run", "error", err)
		os.Exit(1)
	}
}

// run is the whole program; main only hands it the process (signals,
// environment, stdout), so a test can run it with its own.
func run(ctx context.Context, getenv func(string) string, stdout io.Writer) error {
	slog.SetDefault(slog.New(slog.NewTextHandler(stdout, nil)))

	root, err := gonertia.NewFromBytes(rootTemplate)
	if err != nil {
		return fmt.Errorf("parse root template: %w", err)
	}
	vite, err := gonertia.NewViteFromFS(root, publicFS, gonertia.WithAggressivePreload())
	if err != nil {
		return fmt.Errorf("load vite manifest: %w", err)
	}
	v := view.New(vite)
	public, err := fs.Sub(publicFS, "public")
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	routes.Routes(mux, v, http.FileServerFS(public))
	var handler http.Handler = mux
	handler = middlewares.Error(v)(handler)
	handler = v.Middleware(handler)

	srv := &http.Server{
		Addr:    cmp.Or(getenv("LISTEN_ADDR"), ":8000"),
		Handler: handler,
	}
	listenErr := make(chan error, 1)
	go func() {
		slog.Info("listening", "addr", srv.Addr)
		listenErr <- srv.ListenAndServe()
	}()

	select {
	case err := <-listenErr:
		return fmt.Errorf("listen: %w", err)
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
