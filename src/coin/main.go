package main

import (
	"core"
	"fmt"
)

func main() {
	//初始化
	bc := core.CreateNewBlcokchain()
	bc.AddBlock("This is the first Block!")
	bc.AddBlock("This is the second Block!")
	for _,v := range bc.Blcok {
		fmt.Printf("区块创建时间 ：%d\n",v.Timestamp)
		fmt.Printf("前驱Hash散列值 ：%x\n", v.PreBlockHash)
		fmt.Printf("当前Hash散列值 ：%x\n", v.Hash)
		fmt.Printf("当前区块数据 : %s\n", v.Data)
		fmt.Println()
	}
}
