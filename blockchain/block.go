package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/big"
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

// To verify the transaction we need to convert the
// public key hex  back to ecdsa format
func GetPublicKeyFromhex(publicKeyHex string) *ecdsa.PublicKey {
	rpk := publicKeyHex[2:] //skips the first 2 characters (0x) attached to the public key
	xHex := rpk[:64]        // starts from begining to 64 index
	yHex := rpk[64:]        // starts from 64th index to the end

	// to ge the X & Y on the eliptive curve - they are Big integer
	x := new(big.Int)     // aloocates memory for a bigint
	x.SetString(xHex, 16) // interpret the string hex as base 16 input, because it actually a big number represented as Hex

	y := new(big.Int)
	y.SetString(yHex, 16)

	var npk ecdsa.PublicKey

	npk.Curve = elliptic.P256()
	npk.X = x
	npk.Y = y

	return &npk

}
