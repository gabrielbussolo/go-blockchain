package hash

import (
	"crypto/sha256"
	"fmt"
)

type hash struct {
}

type Hash interface {
	Calculate(previousProof, newProof, factor int) int
	Create(obj interface{}) string
}

func New() *hash {
	return &hash{}
}

func (h *hash) Calculate(previousProof, newProof, factor int) int {
	return previousProof * newProof * factor
}

func (h *hash) Create(obj interface{}) string {
	sha := sha256.New()
	sha.Write([]byte(fmt.Sprintf("%v", obj)))
	return fmt.Sprintf("%X", sha.Sum(nil))
}
