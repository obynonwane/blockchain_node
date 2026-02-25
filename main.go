package main

import (
	"log"
	"sync"

	"github.com/obynonwane/evoblockchain/blockchain"
	"github.com/obynonwane/evoblockchain/constants"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ":")
}

func main() {

	var wg sync.WaitGroup

	genesisBlock := blockchain.NewBlock("0x0", 0)
	blockchain := blockchain.NewBlockchain(*genesisBlock)
	log.Println(blockchain.ToJson())
	log.Println("Starting Mining", "\n\n")

	wg.Add(1)

	//simulates proof or work continuously running and adding blocks
	go blockchain.ProofOfWorkMinning("alice")
	wg.Wait()
}
