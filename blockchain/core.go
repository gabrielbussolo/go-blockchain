package blockchain

import (
	"crypto/sha256"
	"go-chain/blockchain/block"
	"time"
)

var chain []block.Block

func Save(proof uint, previousHash [sha256.Size]byte) {
	newBlock := block.Block{
		Index:        uint(len(chain) + 1),
		Timestamp:    time.Now(),
		Proof:        proof,
		PreviousHash: previousHash,
	}
	chain = append(chain, newBlock)
}

func GetPreviousBlock() block.Block {
	lastBlockIndex := len(chain) - 1
	return chain[lastBlockIndex]
}
