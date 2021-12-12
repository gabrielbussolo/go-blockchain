package block

import (
	"time"
)

type Block struct {
	Index        uint
	Timestamp    time.Time
	Proof        uint
	PreviousHash []byte
}
