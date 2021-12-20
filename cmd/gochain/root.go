package gochain

import (
	"go-chain/cmd/gochain/handlers"
	"go-chain/internal/blockchain"
	"go-chain/internal/blockchain/block"
	"go-chain/internal/blockchain/hash"
	"go-chain/internal/blockchain/transaction"
	"net/http"
)

func Build() {
	var chainWithGenesis = []block.Block{block.GetGenesis()}
	h := hash.New()
	mempool := make([]transaction.Transaction, 0)
	b := blockchain.New(h, chainWithGenesis, mempool)
	handler := handlers.NewBlockchainHandler(b)
	http.HandleFunc("/mine", handler.HandleMine)
	http.HandleFunc("/chain", handler.HandleChain)
	http.HandleFunc("/chain/valid", handler.HandleValidChain)
}
