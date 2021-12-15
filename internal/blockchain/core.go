package blockchain

import (
	"fmt"
	"go-chain/internal/blockchain/block"
	"go-chain/internal/blockchain/hash"
	"time"
)

const hashGenesisBlock = "0000000000000000000000000000000000000000000000000000000000000000"

type blockchain struct {
	chain []block.Block
	hash  hash.Hash
}

func New(hash hash.Hash, chain []block.Block) *blockchain {
	return &blockchain{
		chain: chain,
		hash:  hash,
	}
}

func (b *blockchain) Save(block block.Block) {
	b.chain = append(b.chain, block)
}

func (b *blockchain) GetPreviousBlock() block.Block {
	lastBlockIndex := len(b.chain) - 1
	return b.chain[lastBlockIndex]
}

func (b *blockchain) MineBlock(time time.Time) block.Block {
	const target = "0000"
	const factor int = 2
	proof := 1
	previousBlock := b.GetPreviousBlock()
	for {
		nonce := b.hash.Calculate(previousBlock.Proof, proof, factor)
		b2 := block.Block{
			Index:        uint(len(b.chain) + 1),
			Timestamp:    time,
			Proof:        nonce,
			PreviousHash: b.hash.Create(previousBlock),
		}
		create := b.hash.Create(b2)
		if create[0:4] == target {
			fmt.Printf("%d", proof)
			return b2
		}
		proof++
	}
}

func (b *blockchain) IsChainValid() bool {
	previousBlock := b.chain[0]
	for i, chainBlock := range b.chain {
		if i == 0 && chainBlock.PreviousHash != hashGenesisBlock {
			return false
		}
		if i != 0 && chainBlock.PreviousHash != b.hash.Create(previousBlock) {
			return false
		}
		previousBlock = chainBlock
	}
	return true
}

func (b *blockchain) GetChain() []block.Block {
	return b.chain
}
