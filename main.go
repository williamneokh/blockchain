package main

import (
	"crypto/sha256"
	"encoding/json"
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

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int      `json:"nonce"`
		PreviousHash string   `json:"previous_hash"`
		Timestamp    int64    `json:"timestamp"`
		Transaction  []string `json:"transaction"`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Transaction:  b.transaction,
	})
}
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
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
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.PrintAll()
		fmt.Printf("%s\n", strings.Repeat("*", 25))
	}
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	block := &Block{1, "", 0, nil}
	block.Hash()

}
