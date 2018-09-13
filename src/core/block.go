package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}
/**
 * 创建区块函数
 */
func CreateNewBlock(preBlockHash []byte, data string) *Block{
	//初始化数组并赋值
	newBlock := &Block{time.Now().Unix(),preBlockHash,[]byte{},[]byte(data)}
	//调用hash加密函数为新区快的hash赋值
	newBlock.setHash()
	return newBlock
}
/**
 * hash加密函数用于加密
 */
func (b *Block)setHash(){
	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	headers := bytes.Join([][]byte{b.PreBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
/**
 *构建创世区块函数，由于创世区块没有前驱，所以preblockhash传空
 */
func CreateGenesisBlcok() *Block{
	return CreateNewBlock([]byte{}, "This is the genesisblock!")
}

