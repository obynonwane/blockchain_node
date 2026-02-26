package main

import (
	"log"

	"github.com/obynonwane/evoblockchain/constants"
	"github.com/obynonwane/evoblockchain/wallet"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ":")
}

func main() {

	// var wg sync.WaitGroup

	// genesisBlock := blockchain.NewBlock("0x0", 0)
	// blockchain := blockchain.NewBlockchain(*genesisBlock)
	// log.Println(blockchain.ToJson())
	// log.Println("Starting Mining", "\n\n")

	// wg.Add(1)

	// //simulates proof or work continuously running and adding blocks
	// go blockchain.ProofOfWorkMinning("alice")
	// wg.Wait()
	wallet1, _ := wallet.NewWallet()
	log.Println("Private Key Hex :", wallet1.GetPrivateKeyHex())
	log.Println("Public Key Hex :", wallet1.GetPublicKeyHex())
	log.Println("Address :", wallet1.GetAddress())

	wallet2 := wallet.NewWalletFromPrivateKeyHex(wallet1.GetPrivateKeyHex())
	log.Println("Private Key Hex :", wallet2.GetPrivateKeyHex())
	log.Println("Public Key Hex :", wallet2.GetPublicKeyHex())
	log.Println("Address :", wallet2.GetAddress())

	log.Println("checking equals-----")
	log.Println("Private Key Hex :", wallet1.GetPrivateKeyHex() == wallet2.GetPrivateKeyHex())
	log.Println("Public Key Hex :", wallet1.GetPublicKeyHex() == wallet2.GetPublicKeyHex())
	log.Println("Address :", wallet1.GetAddress() == wallet2.GetAddress())
}
