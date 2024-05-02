package tcpserver

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net"
	"strconv"
	"sync"
	"time"
)

// PowService is an interface for a service that provides and verifies a Proof of Work challenge
type PowService interface {
	// GenerateChallenge generates a random challenge
	GenerateChallenge() string

	// VerifyChallenge verifies the nonce against the challenge
	VerifyChallenge(challenge, nonce string) bool

	// Difficulty returns the required difficulty
	Difficulty() int
}

// WoWService is an interface for a service that provides a random Word of Wisdom quote
type WoWService interface {
	// Quote returns a random quote
	Quote() string
}

// Server is a TCP server that provides a Proof of Work challenge and responds with a random Word of Wisdom quote
type Server struct {
	addr       string
	challenges sync.Map
	pow        PowService
	wow        WoWService
}

// New creates a new Server
func New(addr string, pow PowService, wow WoWService) *Server {
	return &Server{
		addr:       addr,
		challenges: sync.Map{},
		pow:        pow,
		wow:        wow,
	}
}

// Start starts the TCP server
func (s *Server) Start(ctx context.Context) {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		slog.Error("Error:", err)
		return
	}

	go func() {
		select {
		case <-ctx.Done():
			if err := listener.Close(); err != nil {
				slog.Error("failed to close listener: ", slog.Any("error", err))
			}
			slog.Info("Listener closed")
		}
	}()

	slog.Info("Server is listening on", slog.Any("addr", s.addr))

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}

			slog.Error("cannot accept connection", slog.Any("error", err))
			continue
		}

		slog.Info("New connection: ", slog.Any("remote_addr", conn.RemoteAddr()))

		go s.handleClient(conn)
	}

	slog.Info("Server stopped")
}

func (s *Server) handleClient(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			slog.Error("cannot close connection", slog.Any("error", err))
		}
	}()

	// Read and process data from the client
	const MaxBodyBytes = int64(1 << 16) // 65536 bytes
	r := io.LimitReader(conn, MaxBodyBytes)

	payload, err := io.ReadAll(r)
	if err != nil {
		slog.Error("cannot read request body", err)
		return
	}

	// if there isn't a challenge, generate one
	if len(payload) == 0 {
		s.generateAndWriteChallenge(conn)
		return
	}

	val, ok := s.challenges.Load(conn.RemoteAddr().String())
	if !ok {
		slog.Error("challenge not found")
		s.generateAndWriteChallenge(conn)
		return
	}

	challenge, ok := val.(string)
	if !ok {
		slog.Error("challenge from map has wrong type")
		s.generateAndWriteChallenge(conn)
		return
	}

	// if challenge wrong, generate new one
	if !s.pow.VerifyChallenge(challenge, string(payload)) {
		slog.Error("challenge verification failed")
		s.generateAndWriteChallenge(conn)
		return
	}

	// if everything ok, respond with a quote
	quote := s.wow.Quote()
	if _, err := conn.Write([]byte(quote)); err != nil {
		slog.Error("cannot write response", err)
	}
}

func (s *Server) generateAndWriteChallenge(conn net.Conn) {
	challenge := s.pow.GenerateChallenge()

	s.challenges.Store(conn.RemoteAddr().String(), challenge)
	go func() {
		<-time.After(time.Minute)
		s.challenges.Delete(conn.RemoteAddr().String())
	}()

	difficulty := s.pow.Difficulty()

	if _, err := conn.Write([]byte(challenge + "\n" + strconv.Itoa(difficulty))); err != nil {
		slog.Error("cannot write challenge", err)
	}
}
