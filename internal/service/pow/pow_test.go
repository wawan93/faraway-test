package pow_test

import (
	"testing"

	"github.com/wawan93/faraway-test/internal/service/pow"
)

func TestService_GenerateChallenge(t *testing.T) {
	str := pow.New(3).GenerateChallenge()
	if str == "" {
		t.Error("Expected a string, got empty string")
	}
}

func TestService_VerifyChallenge(t *testing.T) {
	svc := pow.New(3)
	challenge := "722528"
	nonce := "601"

	if !svc.VerifyChallenge(challenge, nonce) {
		t.Errorf("expected nonce %s to be valid for challenge %s", nonce, challenge)
	}
}
