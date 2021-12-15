package hash

import (
	"testing"
)

func TestHash_Calculate(t *testing.T) {
	expected := 1071600
	h := New()
	got := h.Calculate(53580, 10, 2)
	if got != expected {
		t.Errorf("calculated new proof as %d should be %d", got, expected)
	}
}

func TestHash_Create(t *testing.T) {
	h := New()
	expected := "25D8807DAB2CC30598ECC45AB63ACD9BF8FAF57ECC0C6774ABC8A55310DECFC0"
	const puzzleResult = 52928523843
	got := h.Create(puzzleResult)
	if expected != got {
		t.Errorf("generated sha256 was %s\nit was expected %s", got, expected)
	}
}
