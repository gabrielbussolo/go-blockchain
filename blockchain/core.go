package blockchain

import (
	"crypto/sha256"
	"go-chain/blockchain/block"
	"go-chain/blockchain/hash"
	"time"
)

type blockchain struct {
	chain []block.Block
	hash  hash.Hash
}

func New(hash hash.Hash) *blockchain {
	return &blockchain{
		chain: make([]block.Block, 0),
		hash:  hash,
	}
}

func (b *blockchain) Save(proof uint, previousHash [sha256.Size]byte) {
	newBlock := block.Block{
		Index:        uint(len(b.chain) + 1),
		Timestamp:    time.Now(),
		Proof:        proof,
		PreviousHash: previousHash,
	}
	b.chain = append(b.chain, newBlock)
}

func (b *blockchain) GetPreviousBlock() block.Block {
	lastBlockIndex := len(b.chain) - 1
	return b.chain[lastBlockIndex]
}

func (b *blockchain) ProofOfWork(previousProof uint) uint {
	if previousProof == 0 {
		panic("previous proof should be bigger than 0")
	}
	var newProof uint = 1
	const factor uint = 2
	return b.hash.Calculate(previousProof, newProof, factor)
}
