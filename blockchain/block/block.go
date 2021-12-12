package block

import (
	"crypto/sha256"
	"time"
)

type Block struct {
	Index        uint
	Timestamp    time.Time
	Proof        uint
	PreviousHash [sha256.Size]byte
}
