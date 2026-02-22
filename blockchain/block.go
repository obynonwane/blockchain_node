package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/obynonwane/evoblockchain/constants"
)

type Block struct {
	PrevHash     string         `json:"prevHash"`
	Timestamp    int64          `json:"timestamp"`
	Nonce        int            `json:"nonce"`
	Transactions []*Transaction `json:"transactions"`
}

func NewBlock(prevHash string, nonce int) *Block {
	block := new(Block)
	block.PrevHash = prevHash
	block.Timestamp = time.Now().UnixNano()
	block.Nonce = nonce
	block.Transactions = []*Transaction{}

	return block
}

func (b *Block) ToJson() string {
	nb, err := json.Marshal(b)
	if err != nil {
		return err.Error()
	} else {
		return string(nb)
	}
}

func (b *Block) Hash() string {

	// byte representation of the block
	bs, _ := json.Marshal(b)

	// converts block to hash
	sum := sha256.Sum256(bs)

	// get the first 32 byte
	hexRep := hex.EncodeToString(sum[:32])

	formattedHexRep := constants.HEX_PREFIX + hexRep
	return formattedHexRep
}

func (b *Block) AddTransactionToBlock(txn *Transaction) {
	// verify the txn first
	isTxnValid := txn.VerifyTxn()

	if isTxnValid {
		txn.Status = constants.SUCCESS

	} else {
		txn.Status = constants.FAILED

	}

	b.Transactions = append(b.Transactions, txn)
}
