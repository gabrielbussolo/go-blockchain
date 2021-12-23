package blockchain

import (
	"encoding/json"
	"fmt"
	"go-chain/internal/blockchain/block"
	"go-chain/internal/blockchain/hash"
	"go-chain/internal/blockchain/node"
	"go-chain/internal/blockchain/transaction"
	"io/ioutil"
	"net/http"
	"time"
)

const hashGenesisBlock = "0000000000000000000000000000000000000000000000000000000000000000"

type Blockchain interface {
	Save(block block.Block)
	GetPreviousBlock() block.Block
	MineBlock(time time.Time) block.Block
	IsChainValid([]block.Block) bool
	GetChain() []block.Block
}

type blockchain struct {
	chain   []block.Block
	mempool []transaction.Transaction
	hash    hash.Hash
	nodes   *node.Nodes
}

func New(hash hash.Hash, chain []block.Block, transaction []transaction.Transaction, nodes *node.Nodes) *blockchain {
	return &blockchain{
		chain:   chain,
		mempool: transaction,
		hash:    hash,
		nodes:   nodes,
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
			Transactions: b.mempool,
		}
		create := b.hash.Create(b2)
		if create[0:4] == target {
			b.cleanMemPool()
			return b2
		}
		proof++
	}
}

func (b *blockchain) cleanMemPool() {
	b.mempool = make([]transaction.Transaction, 0)
}

func (b *blockchain) IsChainValid(chain []block.Block) bool {
	previousBlock := b.chain[0]
	for i, chainBlock := range chain {
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

func (b *blockchain) AddTransaction(sender, receiver string, amount float64) {
	b.mempool = append(b.mempool, transaction.New(sender, receiver, amount))
}
func (b *blockchain) AddNode(address string) {
	b.nodes.Add(address)
}

//TODO: Refactor this part about replace chain to a better segregation of responsibility and improve testability
type chainResponse struct {
	Chain  []block.Block
	Length int
}

func (b *blockchain) ReplaceChain() {
	actualLength := len(b.chain)
	var longestValidChain []block.Block
	const baseUrl string = "http://%s/chain"
	for _, url := range b.nodes.GetAll() {
		get, err := http.Get(fmt.Sprintf(baseUrl, url))
		if err != nil {
			panic("error to call other nodes")
		}
		all, err := ioutil.ReadAll(get.Body)
		if err != nil {
			panic("error to call other nodes")
		}
		var chainResp *chainResponse
		err = json.Unmarshal(all, chainResp)
		if err != nil {
			panic("error to call other nodes")
		}
		if chainResp.Length > actualLength && b.IsChainValid(chainResp.Chain) {
			longestValidChain = chainResp.Chain
		}
	}
	if longestValidChain != nil {
		b.chain = longestValidChain
	}
}
