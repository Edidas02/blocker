package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/Edidas02/blocker/blockchain"
)


func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("first block (not including genesis)")
	chain.AddBlock("second block")
	chain.AddBlock("third block")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}