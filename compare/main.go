package main

import "fmt"

var rpc1 *FtmBridge
var rpc2 *FtmBridge

func main() {
	rpc1 = NewFtmBridge("http://xapi113.fantom.network/")
	rpc2 = NewFtmBridge("http://xapi113.fantom.network/")
	defer rpc1.Close()
	defer rpc2.Close()

	compareContracts()
	compareAccountBalances()
	compareAccountNonces()
	compareErc20Name()
	compareErc20Balance()
	fmt.Println("Done")
}
