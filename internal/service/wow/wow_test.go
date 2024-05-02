package wow_test

import (
	"testing"

	"github.com/wawan93/faraway-test/internal/service/wow"
)

func TestService_Quote(t *testing.T) {
	svc := wow.New()
	quote := svc.Quote()

	if quote == "" {
		t.Error("Expected a string, got empty string")
	}
}
