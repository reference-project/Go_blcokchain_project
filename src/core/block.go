package core

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

/**
 * 区块结构体
 */
type Block struct {
	Timestamp int64
	PreBlockHash []byte
	Hash []byte
	Data []byte
	Nonce int
}

func (b *Block)Serialize() []byte{
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}
/**
 * 创建区块函数
 */
func CreateNewBlock(preBlockHash []byte, data string) *Block{
	//初始化数组并赋值
	newBlock := &Block{time.Now().Unix(),preBlockHash,[]byte{},[]byte(data),0}
	//创建工作量证明
	pow := NewProofOfWork(newBlock)
	//运行工作量证明
	nonce, hash := pow.Run()
	newBlock.Hash = hash[:]
	newBlock.Nonce = nonce
	//newBlock.setHash()
	return newBlock
}

/**
 *构建创世区块函数，由于创世区块没有前驱，所以preblockhash传空
 */
func CreateGenesisBlcok() *Block{
	return CreateNewBlock([]byte{}, "This is the genesisblock!")
}

