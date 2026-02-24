package utils

import (
	"encoding/json"

	"github.com/obynonwane/evoblockchain/blockchain"
	"github.com/obynonwane/evoblockchain/constants"
	"github.com/syndtr/goleveldb/leveldb"
)

// using level db = add blockchain
func PutIntoDb(bc blockchain.BlockchainStruct) error {

	db, err := leveldb.OpenFile(constants.BLOCKCHAIN_DB_PATH, nil)
	if err != nil {
		return err
	}

	defer db.Close()

	value, err := json.Marshal(bc)
	if err != nil {
		return err
	}
	//save into the database
	err = db.Put([]byte(constants.BLOCKCHAIN_KEY), value, nil)
	if err != nil {
		return err
	}

	return nil
}

func GetBlockchain() (*blockchain.BlockchainStruct, error) {
	db, err := leveldb.OpenFile(constants.BLOCKCHAIN_DB_PATH, nil)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	data, err := db.Get([]byte(constants.BLOCKCHAIN_KEY), nil)
	if err != nil {
		return nil, err
	}

	var bc blockchain.BlockchainStruct
	err = json.Unmarshal(data, &bc)
	if err != nil {
		return nil, err
	}

	return &bc, nil
}

func KeyExists() (bool, error) {

	db, err := leveldb.OpenFile(constants.BLOCKCHAIN_DB_PATH, nil)
	if err != nil {
		return false, err
	}

	defer db.Close()

	exists, err := db.Has([]byte(constants.BLOCKCHAIN_KEY), nil)
	return exists, err
}
