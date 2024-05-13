package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/wawan93/faraway-test/internal/config"
	"github.com/wawan93/faraway-test/internal/service/pow"
	"github.com/wawan93/faraway-test/internal/service/wow"
	"github.com/wawan93/faraway-test/internal/tcpserver"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Error:", slog.Any("error", err))
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.FromEnv()
	if err != nil {
		return err
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(cfg.LogLevel),
	})))

	ps := pow.New(3)
	ws := wow.New()

	addr := fmt.Sprintf(":%d", cfg.ListenPort)

	s := tcpserver.New(addr, ps, ws)

	ctx, cancel := context.WithCancel(context.Background())

	go s.Start(ctx)

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	<-interrupt
	slog.Info("Shutting down...")
	cancel()
	<-time.After(1 * time.Second)

	return nil
}
