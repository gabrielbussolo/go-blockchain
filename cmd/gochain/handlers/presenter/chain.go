package presenter

import "go-chain/internal/blockchain/block"

type ValidChain struct {
	Message string `json:"message"`
}

type Chain struct {
	Chain  []block.Block `json:"chain"`
	Length int           `json:"length"`
}
