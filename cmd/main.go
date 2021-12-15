package main

import (
	"go-chain/cmd/gochain"
	"net/http"
)

func main() {
	gochain.Build()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
