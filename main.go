package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"rsc.io/quote"
)

// Blockchain struct
type BlockChain struct {
	blocks []*Block
}

// Block struct
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// Hash Derivation
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// Block Creation
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func InitBlockchain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	fmt.Println(quote.Hello())
	chain := InitBlockchain()

	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block After Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
