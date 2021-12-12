package blockchain

import (
	"crypto/sha256"
	"go-chain/blockchain/block"
	"testing"
)

func TestSave(t *testing.T) {
	sum := sha256.Sum256([]byte("sha256 from previous block"))
	Save(1, sum)
	if len(chain) == 0 {
		t.Errorf("size of the chain should be 1")
	}
}

func TestGetPreviousBlock(t *testing.T) {
	chain = []block.Block{
		{
			Index: 0,
		},
		{
			Index: 1,
		},
		{
			Index: 2,
		},
	}

	previousBlock := GetPreviousBlock()

	if previousBlock.Index != 2 {
		t.Errorf("the index of the block should be 2")
	}
}
