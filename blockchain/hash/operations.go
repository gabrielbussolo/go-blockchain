package hash

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"math"
)

type hash struct {
}

type Hash interface {
	Calculate(previousProof, newProof, factor uint) uint
}

func New() *hash {
	return &hash{}
}

func (h *hash) Calculate(previousProof, newProof, factor uint) uint {
	for {
		puzzleResult := math.Pow(float64(newProof), float64(factor)) - math.Pow(float64(previousProof), float64(factor))
		newSha256 := createSha256(puzzleResult)
		const target = "0000"
		if newSha256[0:4] == target {
			return newProof
		}
		newProof++
	}
}

func createSha256(puzzleResult float64) string {
	b := new(bytes.Buffer)
	err := binary.Write(b, binary.LittleEndian, puzzleResult)
	if err != nil {
		panic("problem to convert the puzzle result to binary")
	}
	sum256 := sha256.Sum256(b.Bytes())
	return hex.EncodeToString(sum256[:])
}
