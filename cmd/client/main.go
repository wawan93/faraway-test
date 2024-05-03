package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/wawan93/faraway-test/internal/service/pow"
	"github.com/wawan93/faraway-test/internal/tcpclient"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	addr := "localhost:8080" // TODO: move to config

	client := tcpclient.New(addr)

	challenge, difficulty, err := client.GetChallenge()
	if err != nil {
		slog.Error("Error:", err)
		return
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
		slog.Error("Error:", err)
		return
	}

	fmt.Println(quote)
}
