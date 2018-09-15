package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const dbFile  = "blockchain.db"
const blockBucket  = "blcoks"
/**
 *区块链结构体，以文件存储的形式存放区块，理论上可以建立无数个区块
 */
type Blcokchain struct {
	tip []byte
	DB *bolt.DB
}
/**
 *添加区块函数
 */
func (bc *Blcokchain)AddBlock(data string){

}
/**
 * 创建区块链，并创建创世函数
 */
func CreateNewBlcokchain() *Blcokchain{
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))

		if b == nil {
			fmt.Println("No existing	blockchain found. Create a new one ...")
			genesis := CreateGenesisBlcok()

			b, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("1"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}

			tip = genesis.Hash
		}else {
			tip = b.Get([]byte("1"))
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	bc := Blcokchain{tip, db}
	return  &bc
}
