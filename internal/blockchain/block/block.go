package block

import (
	"time"
)

type Block struct {
	Index        uint
	Timestamp    time.Time
	Proof        int
	PreviousHash string
}

func GetGenesis() Block {
	return Block{
		Index:        1,
		Timestamp:    time.UnixMilli(854697300000),
		Proof:        587686291618922,
		PreviousHash: "0000000000000000000000000000000000000000000000000000000000000000",
	}
}
