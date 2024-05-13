package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	clientConfig "github.com/wawan93/faraway-test/internal/config/client"
	"github.com/wawan93/faraway-test/internal/service/pow"
	"github.com/wawan93/faraway-test/internal/tcpclient"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Error:", slog.Any("error", err))
		os.Exit(1)
	}
}

func run() error {
	cfg, err := clientConfig.FromEnv()
	if err != nil {
		return fmt.Errorf("cannot load config: %w", err)
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(cfg.LogLevel),
	})))

	client := tcpclient.New(cfg.ServerAddr)

	challenge, difficulty, err := client.GetChallenge()
	if err != nil {
		return fmt.Errorf("cannot get challenge: %w", err)
	}

	slog.Debug("Challenge:", slog.Any("challenge", challenge), slog.Any("difficulty", difficulty))

	ps := pow.New(difficulty)

	i := 0
	for {
		if ps.VerifyChallenge(challenge, strconv.Itoa(i)) {
			break
		}
		i++
	}

	nonce := strconv.Itoa(i)

	quote, err := client.GetQuote(challenge, nonce)
	if err != nil {
		return fmt.Errorf("cannot get quote: %w", err)
	}

	fmt.Println(quote)

	return nil
}
