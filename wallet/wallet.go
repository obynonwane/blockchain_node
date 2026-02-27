package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/obynonwane/evoblockchain/blockchain"
	"github.com/obynonwane/evoblockchain/constants"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey `json:"private_key"`
	PublicKey  *ecdsa.PublicKey  `json:"public_key"`
}

func NewWallet() (*Wallet, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	wallet := new(Wallet)
	wallet.PrivateKey = privateKey
	wallet.PublicKey = &privateKey.PublicKey

	return wallet, nil
}

// get the privatekey  hex string format
func (w *Wallet) GetPrivateKeyHex() string {
	return fmt.Sprintf("0x%x", w.PrivateKey.D)
}

func (w *Wallet) GetPublicKeyHex() string {
	return fmt.Sprintf("0x%x%x", w.PublicKey.X, w.PublicKey.Y)
}

func (w *Wallet) GetAddress() string {
	// slice it  to remove the 0x appended in the method GetPublicKeyHex
	hash := sha256.Sum256([]byte(w.GetPublicKeyHex()[2:]))

	// convert the byte slice to hexadecimal hash
	hex := fmt.Sprintf("%x", hash[:])

	// take the last 40 characters - which will now become the wallet address
	// added ADDRESS_PREFIX  to make the address fancy
	address := constants.ADDRESS_PREFIX + hex[len(hex)-40:]

	return address
}

// Function to sign transaction
func (w *Wallet) GetSignedTxn(unsignedTxn blockchain.Transaction) (*blockchain.Transaction, error) {

	// get the bytes of the unsignedTxn
	bs, err := json.Marshal(unsignedTxn)
	if err != nil {
		return nil, err
	}

	// get the hash of the bytes of unsignedTxn
	hash := sha256.Sum256(bs)

	// sign the unsignedTxn hash
	sig, err := ecdsa.SignASN1(rand.Reader, w.PrivateKey, hash[:])
	if err != nil {
		return nil, err
	}

	// return the signed signed transaction
	var signedTxn blockchain.Transaction
	signedTxn.From = unsignedTxn.From
	signedTxn.To = unsignedTxn.To
	signedTxn.Data = unsignedTxn.Data
	signedTxn.Status = unsignedTxn.Status
	signedTxn.Value = unsignedTxn.Value
	signedTxn.TransactionHash = unsignedTxn.TransactionHash
	signedTxn.Signature = sig
	signedTxn.PublicKey = w.GetPublicKeyHex()

	return &signedTxn, nil

}

// convert the private key hex to actuall wallet
// also basically converting the private key hex back to ecdsa.Private key format
func NewWalletFromPrivateKeyHex(privateKeyHex string) *Wallet {
	// start from second index ignoring or skiping 0x i.e the 0th and 1st index
	pk := privateKeyHex[2:] // skips the first 2 index i.e (0x in the privateKeyHex)

	// convert private to to big int
	d := new(big.Int)   // aloocates memory for a bigint
	d.SetString(pk, 16) // interpret the string hex as base 16 input, because it actually a big number represented as Hex

	var npk ecdsa.PrivateKey
	npk.D = d
	npk.PublicKey.Curve = elliptic.P256()
	npk.PublicKey.X, npk.PublicKey.Y = npk.PublicKey.Curve.ScalarBaseMult(d.Bytes())

	wallet := new(Wallet)
	wallet.PrivateKey = &npk
	wallet.PublicKey = &npk.PublicKey

	return wallet
}
