package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/wawan93/faraway-test/internal/service/pow"
	"github.com/wawan93/faraway-test/internal/service/wow"
	"github.com/wawan93/faraway-test/internal/tcpserver"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	ps := pow.New(3)
	ws := wow.New()

	s := tcpserver.New("localhost:8080", ps, ws) // TODO: move addr to config

	ctx, cancel := context.WithCancel(context.Background())

	go s.Start(ctx)

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	<-interrupt
	slog.Info("Shutting down...")
	cancel()
	<-time.After(1 * time.Second)
	slog.Info("Goodbye!")
}
