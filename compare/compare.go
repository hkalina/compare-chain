package main

import (
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
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
	totalErrors += int64(countErr)
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
	totalErrors += int64(countErr)
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
	totalErrors += int64(countErr)
}

// compareErc20Name not used - too much
func compareErc20Name() {
	countOk, countErr, countSkipped := 0, 0, 0

	log.Printf("Checking ERC-20 name...")
	readFileRows("data/erc20.csv", func (row string) {
		address := common.HexToAddress(row)

		nonce1, err1 := rpc1.Erc20Name(address)
		nonce2, err2 := rpc2.Erc20Name(address)

		if err1 != nil && err2 != nil {
			log.Printf("Skip %s: %s / %s", address, err1, err2)
			countSkipped++
			return // failed on both servers
		}

		if nonce1 != nonce2 {
			log.Printf("ERC-20 name different for %s:\nRPC1: %s\nRPC2: %s", address, nonce1, nonce2)
			countErr++
		} else {
			//log.Printf("%s OK (%d)", address, len(code1))
			countOk++
		}
	})

	log.Printf("ERC-20 name checked - OK: %d, skipped: %d, errors: %d", countOk, countSkipped, countErr)
	totalErrors += int64(countErr)
}

func compareErc20Balance() {
	countOk, countErr, countSkipped := 0, 0, 0

	log.Printf("Checking ERC-20 balances...")
	readFileRows("data/erc20disp.csv", func (row string) {
		rowParts := strings.Split(row, ",")
		token := common.HexToAddress(rowParts[0])
		owner := common.HexToAddress(rowParts[1])

		balance1, err1 := rpc1.Erc20BalanceOf(token, owner)
		balance2, err2 := rpc2.Erc20BalanceOf(token, owner)

		if err1 != nil && err2 != nil {
			log.Printf("Skip %s/%s: %s / %s", token, owner, err1, err2)
			countSkipped++
			return // failed on both servers
		}

		if balance1.String() != balance2.String() {
			log.Printf("ERC-20 balance different for %s/%s:\nRPC1: %s\nRPC2: %s", token, owner, balance1.String(), balance2.String())
			countErr++
		} else {
			//log.Printf("%s/%s OK (%s)", token, owner, balance1.String())
			countOk++
		}
	})

	log.Printf("ERC-20 balances checked - OK: %d, skipped: %d, errors: %d", countOk, countSkipped, countErr)
	totalErrors += int64(countErr)
}
