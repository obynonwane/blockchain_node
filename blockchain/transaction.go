package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math"

	"github.com/obynonwane/evoblockchain/constants"
)

type Transaction struct {
	From            string `json:"from"`
	To              string `json:"to"`
	Value           uint64 `json:"value"`
	Data            []byte `json:"data"`
	Status          string `json:"status"`
	TransactionHash string `json:"transaction_hash"`
}

func NewTransaction(from, to string, value uint64, data []byte) *Transaction {

	t := new(Transaction)
	t.From = from
	t.To = to
	t.Value = value
	t.Data = data
	t.Status = constants.PENDING
	t.TransactionHash = t.Hash()
	return t
}

func (t *Transaction) ToJson() string {
	nb, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	} else {
		return string(nb)
	}
}

// verify transactiion before adding to block
func (t Transaction) VerifyTxn() bool {
	if t.Value == 0 {
		return false
	}

	if t.Value > math.MaxUint64 {
		return false
	}

	//TODO: check the signature

	return true
}

func (t Transaction) Hash() string {

	// byte representation of the block
	bs, _ := json.Marshal(t)

	// converts block to hash
	sum := sha256.Sum256(bs)

	// get the first 32 byte
	hexRep := hex.EncodeToString(sum[:32])

	formattedHexRep := constants.HEX_PREFIX + hexRep
	return formattedHexRep
}
