package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
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
