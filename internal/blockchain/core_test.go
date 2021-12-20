package blockchain

import (
	"go-chain/internal/blockchain/block"
	"go-chain/internal/blockchain/hash"
	"go-chain/internal/blockchain/transaction"
	"testing"
	"time"
)

var chainWithGenesis = []block.Block{block.GetGenesis()}
var mempool = make([]transaction.Transaction, 0)

func TestSave(t *testing.T) {
	blockchain := New(hash.New(), chainWithGenesis, mempool)
	b := block.Block{
		Index:        2,
		Timestamp:    time.Now(),
		Proof:        999999,
		PreviousHash: "",
	}
	blockchain.Save(b)
	if len(blockchain.chain) != 2 {
		t.Errorf("size of the chain should be 2")
	}
}

func TestGetPreviousBlock(t *testing.T) {
	blockchain := New(hash.New(), chainWithGenesis, mempool)
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

func TestBlockchain_MineBlock(t *testing.T) {
	h := hash.New()
	blockchain := New(h, chainWithGenesis, mempool)
	mineBlock := blockchain.MineBlock(time.Unix(1639509419, 0))
	if mineBlock.PreviousHash != h.Create(blockchain.GetPreviousBlock()) {
		t.Errorf("mined a block with wrong previous hash")
	}
	create := h.Create(mineBlock)
	if create[0:4] != "0000" {
		t.Errorf("invalid hash")
	}
}

func TestBlockchain_IsChainValid(t *testing.T) {
	h := hash.New()
	b := New(h, chainWithGenesis, mempool)

	t.Run("valid chain", func(t *testing.T) {
		mineBlock := b.MineBlock(time.Now())
		b.Save(mineBlock)
		valid := b.IsChainValid()
		if !valid {
			t.Errorf("chain should be valid by mining a new block")
		}
	})
	t.Run("invalid chain", func(t *testing.T) {
		b.Save(block.Block{
			Index:        2,
			Timestamp:    time.Now(),
			Proof:        1231231,
			PreviousHash: h.Create("invalidhash"),
		})
		valid := b.IsChainValid()
		if valid {
			t.Errorf("chain should be invalid")
		}
	})
	t.Run("invalid genesis block on chain", func(t *testing.T) {
		chainWithoutGenesis := []block.Block{{
			Index:        1,
			Timestamp:    time.Now(),
			Proof:        21321321,
			PreviousHash: h.Create("invalidhash"),
		}}
		b2 := New(hash.New(), chainWithoutGenesis, mempool)
		mineBlock := b2.MineBlock(time.Now())
		b2.Save(mineBlock)
		valid := b2.IsChainValid()
		if valid {
			t.Errorf("chain with invalid genesis should be invalid")
		}
	})
}

func TestBlockchain_GetChain(t *testing.T) {
	b := New(hash.New(), chainWithGenesis, mempool)
	mineBlock := b.MineBlock(time.Now())
	b.Save(mineBlock)

	chain := b.GetChain()

	if len(chain) != 2 {
		t.Errorf("chain with invalid size")
	}
	if chain[0].Index != block.GetGenesis().Index {
		t.Errorf("chain not starting with genesis")
	}
	if chain[1].Index != mineBlock.Index {
		t.Errorf("chain not with updated data")
	}
}

func TestBlockchain_AddTransaction(t *testing.T) {
	b := New(hash.New(), chainWithGenesis, mempool)
	tx := transaction.New("", "", 0.5)
	b.AddTransaction(tx.Sender, tx.Receiver, tx.Amount)

	if b.mempool[0].Sender != tx.Sender {
		t.Errorf("added tx isnt the same, expected %v got %v", tx.Sender, b.mempool[0].Sender)
	}
}
