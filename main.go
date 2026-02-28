package main

import (
	"log"
	"sync"

	"github.com/obynonwane/evoblockchain/blockchain"
	"github.com/obynonwane/evoblockchain/constants"
	"github.com/obynonwane/evoblockchain/wallet"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ":")
}

func main() {

	var wg sync.WaitGroup
	wallet2, _ := wallet.NewWallet()
	genesisBlock := blockchain.NewBlock("0x0", 0)
	blockchain1 := blockchain.NewBlockchain(*genesisBlock)
	log.Println(blockchain1.ToJson())
	log.Println("Starting Mining", "\n\n")
	wg.Add(1)
	//simulates proof or work continuously running and adding blocks
	go blockchain1.ProofOfWorkMinning(wallet2.GetAddress())
	wg.Wait()

	// wallet1, _ := wallet.NewWallet()
	// log.Println("Private Key Hex :", wallet1.GetPrivateKeyHex())
	// log.Println("Public Key Hex :", wallet1.GetPublicKeyHex())
	// log.Println("Address :", wallet1.GetAddress())

	// wallet2 := wallet.NewWalletFromPrivateKeyHex(wallet1.GetPrivateKeyHex())
	// log.Println("Private Key Hex :", wallet2.GetPrivateKeyHex())
	// log.Println("Public Key Hex :", wallet2.GetPublicKeyHex())
	// log.Println("Address :", wallet2.GetAddress())

	// log.Println("checking equals-----")
	// log.Println("Private Key Hex :", wallet1.GetPrivateKeyHex() == wallet2.GetPrivateKeyHex())
	// log.Println("Public Key Hex :", wallet1.GetPublicKeyHex() == wallet2.GetPublicKeyHex())
	// log.Println("Address :", wallet1.GetAddress() == wallet2.GetAddress())

	// wallet1, _ := wallet.NewWallet()
	// wallet2, _ := wallet.NewWallet()

	// uTxn := blockchain.NewTransaction(wallet1.GetAddress(), wallet2.GetAddress(), 100, []byte{})
	// sTxn, _ := wallet1.GetSignedTxn(*uTxn)
	// blockchain1.AddTransactionTotransactionPool(*sTxn)

	// // log.Println(blockchain1)
	// log.Println(blockchain1.ToJson())
}
