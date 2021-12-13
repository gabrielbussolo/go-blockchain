package blockchain

import (
	"crypto/sha256"
	"go-chain/blockchain/block"
	"go-chain/blockchain/hash"
	"testing"
)

type hashMock struct {
	calculateMock func(previousProof, newProof, factor uint) uint
}

func (h *hashMock) Calculate(previousProof, newProof, factor uint) uint {
	return h.calculateMock(previousProof, newProof, factor)
}

func TestSave(t *testing.T) {
	blockchain := New(hash.New())
	sum := sha256.Sum256([]byte("sha256 from previous block"))
	blockchain.Save(1, sum)
	if len(blockchain.chain) == 0 {
		t.Errorf("size of the chain should be 1")
	}
}

func TestGetPreviousBlock(t *testing.T) {
	blockchain := New(hash.New())
	blockchain.chain = []block.Block{
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

	previousBlock := blockchain.GetPreviousBlock()

	if previousBlock.Index != 2 {
		t.Errorf("the index of the block should be 2")
	}
}

func TestProofOfWork(t *testing.T) {
	hashMock := &hashMock{calculateMock: func(previousProof, newProof, factor uint) uint {
		return 3545
	}}
	blockchain := New(hashMock)

	t.Run("fail when previous proof is 0", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("it was expecting to panic when the previousProof is %d", 0)
			}
		}()
		expected := 0
		work := blockchain.ProofOfWork(0)
		if work != 0 {
			t.Errorf("proof of work should not be different from %d", expected)
		}
	})

	t.Run("process the proof", func(t *testing.T) {
		var expected uint = 3545
		got := blockchain.ProofOfWork(2)
		if got != expected {
			t.Errorf("proof of work should not be different from %d", expected)
		}
	})
}
