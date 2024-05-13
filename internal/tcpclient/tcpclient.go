package tcpclient

import (
	"errors"
	"io"
	"log/slog"
	"net"
	"strconv"
	"strings"
)

type TCPClient struct {
	addr string
}

func New(addr string) *TCPClient {
	return &TCPClient{addr: addr}
}

func (c *TCPClient) GetQuote(challenge string, nonce string) (string, error) {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		slog.Error("Error:", err)
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(challenge + " " + nonce + "\n"))
	if err != nil {
		slog.Error("Error:", err)
		return "", err
	}

	b, err := io.ReadAll(conn)
	if err != nil {
		slog.Error("Error:", err)
		return "", err
	}

	return string(b), nil
}

func (c *TCPClient) GetChallenge() (string, int, error) {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		slog.Error("Error:", err)
		return "", 0, err
	}
	defer conn.Close()

	_, err = conn.Write([]byte("GET\n"))
	if err != nil {
		slog.Error("Error:", err)
		return "", 0, err
	}

	b, err := io.ReadAll(conn)
	if err != nil {
		slog.Error("Error:", err)
		return "", 0, err
	}

	payload := strings.Split(string(b), " ")
	if len(payload) != 2 {
		slog.Error("Invalid payload")
		return "", 0, errors.New("invalid payload")
	}

	challenge := payload[0]

	difficulty, err := strconv.Atoi(payload[1])
	if err != nil {
		slog.Error("Error:", err)
		return "", 0, err
	}

	return challenge, difficulty, nil
}
