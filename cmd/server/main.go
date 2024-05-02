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
	ps := pow.New(3)
	ws := wow.New()

	s := tcpserver.New("localhost:8080", ps, ws)

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
