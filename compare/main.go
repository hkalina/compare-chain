package main

import (
	"fmt"
	"log"
	"os"
)

var rpc1 *FtmBridge
var rpc2 *FtmBridge
var totalErrors int64

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s [all/contracts/balances/nonces/ercnames/ercbalances] http://rpc1/ http://rpc2/\n", os.Args[0])
		return
	}

	task := os.Args[1]
	rpc1 = NewFtmBridge(os.Args[2])
	rpc2 = NewFtmBridge(os.Args[3])
	defer rpc1.Close()
	defer rpc2.Close()

	if task == "contracts" || task == "all" {
		compareContracts()
	}
	if task == "balances" || task == "all" {
		compareAccountBalances()
	}
	if task == "nonces" || task == "all" {
		compareAccountNonces()
	}
	if task == "ercnames" || task == "all" {
		compareErc20Name()
	}
	if task == "ercbalances" || task == "all" {
		compareErc20Balance()
	}

	log.Printf("Done, total errors: %d", totalErrors)
	if totalErrors != 0 {
		os.Exit(1)
	}
}
