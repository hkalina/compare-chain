package main

import (
	"fmt"
	"math/big"
	"os"
)

var rpc1 *FtmBridge
var rpc2 *FtmBridge
var totalErrors int64
var block big.Int
var skipRows int64

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s [blockNumber] http://rpc1/\n", os.Args[0])
		return
	}

	block.SetString(os.Args[1], 10)
	rpc1 = NewFtmBridge(os.Args[2])
	defer rpc1.Close()

	generateBalances()
}
