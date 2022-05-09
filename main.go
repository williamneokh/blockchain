package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transaction  []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b

}

func (b *Block) PrintAll() {
	fmt.Printf("Nonce: %d\n", b.nonce)
	fmt.Printf("Time Stamp: %d\n", b.timestamp)
	fmt.Printf("Previous Hash: %s\n", b.previousHash)
	fmt.Printf("Transaction: %v\n", b.transaction)
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) PrintBlock() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d%s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.PrintAll()
		fmt.Printf("%s\n", strings.Repeat("*", 25))
	}
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	blockChain := NewBlockchain()
	blockChain.PrintBlock()
	blockChain.CreateBlock(5, "hash 1")
	blockChain.PrintBlock()
	blockChain.CreateBlock(3, "hash 2")
	blockChain.PrintBlock()

}
