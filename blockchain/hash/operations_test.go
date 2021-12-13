package hash

import (
	"testing"
)

func TestHash_Calculate(t *testing.T) {
	var expected uint = 3545
	h := &hash{}
	got := h.Calculate(2, 3540, 2)
	if got != expected {
		t.Errorf("calculated new proof as %d should be %d", got, expected)
	}
}

func TestCreateSha256(t *testing.T) {
	expected := "0000daf1cda7ccad6fb4b641a48867e85bf3c6f0f26753806cb35f562d4df839"
	const puzzleResult = 52928523843
	got := createSha256(puzzleResult)
	const target = "0000"
	if got[0:4] != target {
		t.Errorf("sha256 should start with %s", target)
	}
	if expected != got {
		t.Errorf("generated sha256 was %s\nit was expected %s", got, expected)
	}
}
