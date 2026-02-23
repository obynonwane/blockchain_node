package main

import (
	"log"
	"sync"
	"time"

	"github.com/obynonwane/evoblockchain/blockchain"
	"github.com/obynonwane/evoblockchain/constants"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ":")
}

func main() {

	var wg sync.WaitGroup

	genesisBlock := blockchain.NewBlock("0x0", 0)
	transaction1 := blockchain.NewTransaction("0x1", "0x2", 12, []byte{})
	blockchain := blockchain.NewBlockchain(*genesisBlock)
	log.Println(blockchain.ToJson())
	wg.Add(1)

	//simulates proof or work continuously running and adding blocks
	go blockchain.ProofOfWorkMinning("alice")
	time.Sleep(2000)
	// simulates user sending transaction
	blockchain.AddTransactionTotransactionPool(*transaction1)
	log.Println("Transaction Pool", blockchain.ToJson())
	wg.Wait()
}
