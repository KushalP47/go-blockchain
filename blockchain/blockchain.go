package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

// Blockchain struct
type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

func (chain *BlockChain) AddBlock(data string) {
	// prevBlock := chain.Blocks[len(chain.Blocks)-1]
	// newBlock := CreateBlock(data, prevBlock.Hash)
	// chain.Blocks = append(chain.Blocks, newBlock)
}

func InitBlockchain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Printf("Genesis Proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash
			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			err = item.Value(func(val []byte) error {
				lastHash = val
				return nil
			})
			return err
		}
	})

	Handle(err)

	blockchain := &BlockChain{lastHash, db}
	return blockchain
}
