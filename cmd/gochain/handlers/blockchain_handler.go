package handlers

import (
	"encoding/json"
	"go-chain/cmd/gochain/handlers/presenter"
	"go-chain/internal/blockchain"
	"net/http"
	"time"
)

type blockchainHandler struct {
	blockchain blockchain.Blockchain
}

func NewBlockchainHandler(blockchain blockchain.Blockchain) *blockchainHandler {
	return &blockchainHandler{blockchain: blockchain}
}

func (bh *blockchainHandler) HandleMine(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
	block := bh.blockchain.MineBlock(time.Now())
	bh.blockchain.Save(block)
	jsonBytes, err := json.Marshal(block)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (bh *blockchainHandler) HandleChain(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
	chain := bh.blockchain.GetChain()
	chainPresenter := &presenter.Chain{Chain: chain, Length: len(chain)}
	jsonBytes, err := json.Marshal(chainPresenter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (bh *blockchainHandler) HandleValidChain(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
	valid := bh.blockchain.IsChainValid(bh.blockchain.GetChain())
	var err error
	var jsonBytes []byte
	jsonBytes, err = json.Marshal(presenter.ValidChain{Message: "The Blockchain is not valid."})
	if valid {
		jsonBytes, err = json.Marshal(presenter.ValidChain{Message: "All good. The Blockchain is valid."})
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
