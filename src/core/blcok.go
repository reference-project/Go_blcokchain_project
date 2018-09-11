package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

/**
 *区块结构体
 */
type Block struct {
	Index int64   		//区块号
	Timestamp int64			//当前时间戳
	PreBlockHash string		//前一个区块的hash值
	BlockHash string		//当前区块的hash值
	data string				//区块携带的数据
}

/**
 *计算区块hash值的函数
 */
func CalculateHash (b Block) string {
	//拼装加密前的字符串，由区块号，区块创建时间，前一个区块的hash值组成
	blockdata := string(b.Index) + string(b.Timestamp) + b.PreBlockHash;
	//加密拼装好的字符串，使用sha256加密，参数为byte字节数组
	calculateBlock := sha256.Sum256([]byte(blockdata))
	//返回是字符串所以使用hex函数转换类型
	return hex.EncodeToString(calculateBlock[:])
}

/**
 *创建一个新区块
 */
func CreateNewBlock(preBlock Block, data string) Block {
	//初始化一个区块结构体
	newBlock := Block{}
	//为新区块赋值
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Index = preBlock.Index + 1
	newBlock.PreBlockHash = preBlock.BlockHash
	newBlock.BlockHash = CalculateHash(newBlock)
	newBlock.data = data
	return newBlock
}

/**
 *创建创世区块
 */
func CreateGenesisBlock() Block {
	//初始化区块
	newBlock := Block{}
	//为了能让创世区块的编号为0 所以此值为-1
	newBlock.Index = -1
	//创世区块为第一个区块，没有前一个区块，所以前一个区块的hash值为空
	newBlock.PreBlockHash = ""
	return CreateNewBlock(newBlock, "This is the GenesisBlock!")
}


