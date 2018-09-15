package main

import (
	"core"
)

func main() {
	//初始化区块链
	bc := core.CreateNewBlcokchain()
	defer bc.DB.Close()

	cli := core.CLI{bc}
	cli.Run()
}
