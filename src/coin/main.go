package main

import (
	"core"
	"fmt"
	"strconv"
)

func main() {
	//初始化区块链
	bc := core.CreateNewBlcokchain()
	//想区块链中添加区块
	bc.AddBlock("This is the first Block!")
	bc.AddBlock("This is the second Block!")

	//打印区块内容并验证
	for _,v := range bc.Blcok {
		fmt.Printf("区块创建时间 ：%d\n",v.Timestamp)
		fmt.Printf("前驱Hash散列值 ：%x\n", v.PreBlockHash)
		fmt.Printf("当前Hash散列值 ：%x\n", v.Hash)
		fmt.Printf("当前区块数据 : %s\n", v.Data)

		pow := core.NewProofOfWork(v)
		fmt.Printf("Pow:%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
