package main

import (
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func compareContracts() {
	countOk, countErr := 0, 0

	log.Printf("Checking contracts...")
	readFileRows("data/contract.csv", func (row string) {
		address := common.HexToAddress(row)

		code1 := rpc1.GetCode(address)
		code2 := rpc2.GetCode(address)

		if code1 != code2 {
			log.Printf("Code different for %s:\nRPC1: %s\nRPC2: %s", address, code1, code2)
			countErr++
		} else {
			//log.Printf("%s OK (%d)", address, len(code1))
			countOk++
		}
	})

	log.Printf("Contracts checked - OK: %d, errors: %d", countOk, countErr)
}

func compareAccountBalances() {
	countOk, countErr := 0, 0

	log.Printf("Checking account balances...")
	readFileRows("data/account.csv", func (row string) {
		address := common.HexToAddress(row)

		balance1 := rpc1.GetBalance(address)
		balance2 := rpc2.GetBalance(address)

		if balance1 != balance2 {
			log.Printf("Balance different for %s:\nRPC1: %s\nRPC2: %s", address, balance1, balance2)
			countErr++
		} else {
			//log.Printf("%s OK (%d)", address, len(code1))
			countOk++
		}
	})

	log.Printf("Account balances checked - OK: %d, errors: %d", countOk, countErr)
}

func compareAccountNonces() {
	countOk, countErr := 0, 0

	log.Printf("Checking account nonces...")
	readFileRows("data/account.csv", func (row string) {
		address := common.HexToAddress(row)

		nonce1 := rpc1.GetNonce(address)
		nonce2 := rpc2.GetNonce(address)

		if nonce1 != nonce2 {
			log.Printf("Nonce different for %s:\nRPC1: %s\nRPC2: %s", address, nonce1, nonce2)
			countErr++
		} else {
			//log.Printf("%s OK (%d)", address, len(code1))
			countOk++
		}
	})

	log.Printf("Account nonces checked - OK: %d, errors: %d", countOk, countErr)
}
