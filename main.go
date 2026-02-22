package main

import (
	"log"

	"github.com/obynonwane/evoblockchain/blockchain"
	"github.com/obynonwane/evoblockchain/constants"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ":")
}

func main() {

	transaction := blockchain.NewTransaction("0x1", "0x2", 1, []byte{})
	log.Println("TRANSACTION", transaction.ToJson())

	block := blockchain.NewBlock("0x", 9)
	log.Println("BLOCK", block.ToJson())

	transaction1 := blockchain.NewTransaction("0x1", "0x2", 1, []byte{})
	genesisBlock := blockchain.NewBlock("0x", 100)
	genesisBlock.Transactions = append(genesisBlock.Transactions, transaction1)
	blockchain := blockchain.NewBlockchain(*genesisBlock)
	log.Println("BLOCKCHAIN", blockchain.ToJson())
	log.Println("HASH OF BLOCK", block.Hash())

}
