package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	ProvHash []byte
}
type BlockChain struct {
	blocks []*Block
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")
	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash:%x\n",block.ProvHash)
		fmt.Printf("Data Hash:%s\n",block.Data)
		fmt.Printf("Hash:%x\n",block.Hash)
	}
}

func (b *Block) DeriverHash() {
	info := bytes.Join([][]byte{b.Data, b.ProvHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriverHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
