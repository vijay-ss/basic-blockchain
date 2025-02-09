package main

import (
	"fmt"
	"github.com/vijay-ss/basic-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("previous hash: %x\n", block.PrevHash)
		fmt.Printf("data in block: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
	}
}