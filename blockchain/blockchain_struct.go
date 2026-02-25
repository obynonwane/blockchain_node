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

	// get key since using single key
	exists, _ := KeyExists()
	if exists {

		// if key exist get the entire blockchain
		blockchainStruct, err := GetBlockchain()
		if err != nil {
			panic(err.Error())
		}

		// then return it
		return blockchainStruct

	} else {

		// if key do not exist create new blockchain
		blockchainStruct := new(BlockchainStruct)
		blockchainStruct.TransactionPool = []*Transaction{}                      // initialise emtyp tx pool
		blockchainStruct.Blocks = []*Block{}                                     // initialise empty block
		blockchainStruct.Blocks = append(blockchainStruct.Blocks, &genesisBlock) // append genesis block to bc struct
		err := PutIntoDb(*blockchainStruct)                                      // save to levelDB
		if err != nil {
			panic(err.Error())
		}
		return blockchainStruct // return saved blockchainstruct, containing genesis stuff
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
	//save the block to database
	err := PutIntoDb(*bc)
	if err != nil {
		panic(err.Error())
	}
}

// Add transaction to transaction pool
func (bc *BlockchainStruct) AddTransactionTotransactionPool(transaction Transaction) {
	bc.TransactionPool = append(bc.TransactionPool, &transaction)
	//save the block to database
	err := PutIntoDb(*bc)
	if err != nil {
		panic(err.Error())
	}
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
			// log.Println(bc.ToJson(), "\n\n")
			log.Println("TOTAL CRYPTO", bc.CalculateTotalCrypto("bob"))

			prevHash = bc.Blocks[len(bc.Blocks)-1].Hash() // extract the last block
			nonce = 0                                     // reset nonce mining has been done by miner
			continue                                      // jumps back to the top of the loop, start afresh
		}
		nonce++
	}

}

func (bc BlockchainStruct) CalculateTotalCrypto(address string) uint64 {

	sum := int64(0)

	for _, blocks := range bc.Blocks {
		for _, txns := range blocks.Transactions {
			if txns.Status == constants.SUCCESS {
				if txns.To == address {
					sum += int64(txns.Value)
				} else if txns.From == address {
					sum -= int64(txns.Value)
				}
			}

		}
	}

	return uint64(sum)
}
