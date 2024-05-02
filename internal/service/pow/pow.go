package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Service provides a Proof of Work challenge and verifies the nonce
type Service struct {
	difficulty int
}

// New creates a new Service
func New(difficulty int) *Service {
	return &Service{
		difficulty: difficulty,
	}
}

// GenerateChallenge generates a random challenge
func (s Service) GenerateChallenge() string {
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(gen.Intn(1_000_000))
}

// VerifyChallenge verifies the nonce against the challenge
func (s Service) VerifyChallenge(challenge, nonce string) bool {
	hash := sha256.Sum256([]byte(challenge + nonce))
	hexHash := hex.EncodeToString(hash[:])
	return hexHash[:s.difficulty] == strings.Repeat("0", s.difficulty)
}

// Difficulty returns the required difficulty
func (s Service) Difficulty() int {
	return s.difficulty
}
