package core

import (
	"fmt"
	"log"
)

/**
 *区块链结构体
 */
type Blockchain struct {
	Block []*Block
}

/**
 *验证函数，用于新添加的区块链是否合法
 */
func isValid(oldBlock Block, newBlock Block) bool {
	//验证区块序列号是否为连续
	if oldBlock.Index != newBlock.Index-1 {
		return false
	}
	//验证当前区块的prehash值是否为上一个区块的hash值
	if oldBlock.BlockHash != newBlock.PreBlockHash {
		return false
	}
	//验证当前区块的hash值是否与加密后的一致
	if CalculateHash(newBlock) != newBlock.BlockHash {
		return false
	}
	return true
}

/**
 *向区块链中添加新区块，使用指针类型，指针类型可以改变原本结构体值
 */
func (currentBlock *Blockchain)AppendBlock(newBlock *Block){
	//如果当前区块链的长度为0则直接将新区块插入
	if len(currentBlock.Block) == 0 {
		currentBlock.Block = append(currentBlock.Block, newBlock)
		return
	}
	//验证区块链的最后一个块与将要插入的块
	if isValid(*currentBlock.Block[len(currentBlock.Block)-1], *newBlock) {
		currentBlock.Block = append(currentBlock.Block, newBlock)
	}else {
		log.Fatal("This block has problem!")
	}
}

//创建新区块链，并将创世区块追加到区块链中
func CreateNewBlockchain() *Blockchain {
	//创建创世区块
	genessisBlock := CreateGenesisBlock()
	//初始化区块链
	newBlockchain := Blockchain{}
	//将创世区块追加到区块链中
	newBlockchain.AppendBlock(&genessisBlock)
	return &newBlockchain
}

//创建新区块，并给区块赋值后追加到区块链中
func (currentkchain *Blockchain)SendData(data string){
	preBlock := currentkchain.Block[len(currentkchain.Block)-1]
	newBlock := CreateNewBlock(*preBlock, data)
	currentkchain.AppendBlock(&newBlock)
}

/**
 * 打印区块函数，用于方便的查看当前区块链
 */
 func (currentBlockchain *Blockchain)ShowBlockchain(){
 	for _,v := range currentBlockchain.Block{
 		fmt.Println("Index", v.Index)
 		fmt.Println("CreateTime:", v.Timestamp)
 		fmt.Println("PreBlockHash:", v.PreBlockHash)
 		fmt.Println("CurrentBlockHash:", v.BlockHash)
 		fmt.Println("Nonce:",v.Nonce)
 		NewProofOfWork(v)
 		fmt.Println("Data:", v.data)
	}
 }





