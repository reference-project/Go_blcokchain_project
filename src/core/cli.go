package core

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	Bc *Blcokchain
}

func (cli *CLI)showUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA     *add a blcok to blockchain")
	fmt.Println("  showchain                     *show all the blocks of the blcokchain")
}

func (cli *CLI)validateArgs() {
	if len(os.Args) < 2 {
		cli.showUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string){
	cli.Bc.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) showBlcokchain(){

}

func (cli *CLI)Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addBlcok", flag.ExitOnError)
	showBlockchainCmd := flag.NewFlagSet("showBlockchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
		case "addBlock":
			err := addBlockCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "showBlockchain":
			err := showBlockchainCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		default:
			cli.showUsage()
			os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if showBlockchainCmd.Parsed() {
		cli.showUsage()
	}

}