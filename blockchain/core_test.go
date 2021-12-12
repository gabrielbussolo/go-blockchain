package blockchain

import (
	"crypto/sha256"
	"testing"
)

func TestSave(t *testing.T) {
	sum := sha256.Sum256([]byte("sha256 from previous block"))
	Save(1, sum)
	if len(chain) == 0 {
		t.Errorf("size of the chain should be 1")
	}
}
