package blockchain

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/obynonwane/evoblockchain/constants"
)

type BlockchainStruct struct {
	TransactionPool []*Transaction `json:"transaction_pool"`
	Blocks          []*Block       `json:"block_chain"`
}

func NewBlockchain(genesisBlock Block) *BlockchainStruct {
	exists, _ := KeyExists()
	if exists {

		blockchainStruct, err := GetBlockchain()
		if err != nil {
			panic(err.Error())
		}
		return blockchainStruct

	} else {
		blockchainStruct := new(BlockchainStruct)
		blockchainStruct.TransactionPool = []*Transaction{}
		blockchainStruct.Blocks = []*Block{}
		blockchainStruct.Blocks = append(blockchainStruct.Blocks, &genesisBlock)
		err := PutIntoDb(*blockchainStruct)
		if err != nil {
			panic(err.Error())
		}
		return blockchainStruct
	}
}

func (bc BlockchainStruct) ToJson() string {
	nb, err := json.Marshal(bc)
	if err != nil {
		return err.Error()
	} else {
		return string(nb)
	}
}

func (bc *BlockchainStruct) AddBlock(b *Block) {
	//TODO: add a block to the blockchain
	m := map[string]bool{}

	//  range over the transactions in the given block
	for _, txn := range b.Transactions {
		// push the hashes of the transaction to above map
		m[txn.Hash()] = true
	}

	// loop through the TransactionPool
	for idx, txn := range bc.TransactionPool {
		// check if transaction exist in the above map
		_, ok := m[txn.TransactionHash]
		if ok {
			// remove the transaction with the hash from transaction pool
			bc.TransactionPool = append(bc.TransactionPool[:idx], bc.TransactionPool[idx+1:]...)
		}
	}

	// add block to blockchain
	bc.Blocks = append(bc.Blocks, b)
}

// Add transaction to transaction pool
func (bc *BlockchainStruct) AddTransactionTotransactionPool(transaction Transaction) {
	bc.TransactionPool = append(bc.TransactionPool, &transaction)
}

func (bc *BlockchainStruct) ProofOfWorkMinning(minersAddress string) {
	//calculate the prevHash
	prevHash := bc.Blocks[len(bc.Blocks)-1].Hash() // extract the last block
	//start with a nonce - tracks how many time it took me to arrive at the correct hash
	nonce := 0
	for {

		// create a new Block
		guessBlock := NewBlock(prevHash, nonce)

		//copy the transaction pool
		for _, txn := range bc.TransactionPool {
			newTxn := NewTransaction(txn.From, txn.To, txn.Value, txn.Data)

			// add transaction to the crreated block
			guessBlock.AddTransactionToBlock(newTxn)
		}
		// guess the Hash
		guesHash := guessBlock.Hash()
		desiredHash := strings.Repeat("0", constants.MINING_DIFFICULTY)
		// extract the giving contant begining from index 2 skipping 0 & 1 index
		ourSolutionHash := guesHash[2 : 2+constants.MINING_DIFFICULTY]

		if ourSolutionHash == desiredHash {
			// reward the miner
			rewardTxn := NewTransaction(constants.BLOCKCHAIN_ADDRESS, minersAddress, constants.MINING_REWARD, []byte{})
			rewardTxn.Status = constants.SUCCESS
			// add the reward transaction to block - custom way, so it those not go through validation
			guessBlock.Transactions = append(guessBlock.Transactions, rewardTxn)

			// the the Block
			bc.AddBlock(guessBlock)
			log.Println(bc.ToJson(), "\n\n")

			// utilities.PrettyLog("Blockchain", bc)
			// log.Println("\n\n\n")
			prevHash = bc.Blocks[len(bc.Blocks)-1].Hash() // extract the last block
			nonce = 0                                     // reset nonce mining has been done by miner
			continue                                      // jumps back to the top of the loop, start afresh
		}
		nonce++
	}
	//compare this hash with the mining difficulty
}
