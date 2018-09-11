package main

import "core"

func main() {
	//创建新区块链
	blockchain := core.CreateNewBlockchain()
	//向区块链中存入值
	blockchain.SendData("PHPerJiang")
	blockchain.SendData("GoerJiang")
	//打印区块链内容
	blockchain.ShowBlockchain()
}
