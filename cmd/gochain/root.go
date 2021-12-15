package gochain

import (
	"go-chain/cmd/gochain/handlers"
	"go-chain/internal/blockchain"
	"go-chain/internal/blockchain/block"
	"go-chain/internal/blockchain/hash"
	"net/http"
)

func Build() {
	var chainWithGenesis = []block.Block{block.GetGenesis()}
	h := hash.New()
	b := blockchain.New(h, chainWithGenesis)
	handler := handlers.NewBlockchainHandler(b)
	http.HandleFunc("/mine", handler.HandleMine)
	http.HandleFunc("/chain", handler.HandleChain)
	http.HandleFunc("/chain/valid", handler.HandleValidChain)
}
