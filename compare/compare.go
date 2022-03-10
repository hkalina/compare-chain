package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
	"strings"
)

func compareContracts() {
	countOk, countErr := 0, 0

	log.Printf("Checking contracts...")
	readFileRows("data/contract.csv", func (row string) {
		address := common.HexToAddress(row)

		code1 := rpc1.GetCode(address, hexutil.Big(block))
		code2 := rpc2.GetCode(address, hexutil.Big(block))

		if code1 != code2 {
			log.Printf("Code different for %s:\nRPC1: %s\nRPC2: %s", address, code1, code2)
			countErr++
		} else {
			if countOk % 100 == 0 {
				log.Printf("Contracts codes - OK: %d, errors: %d", countOk, countErr)
			}
			countOk++
		}
	})

	log.Printf("Contracts codes - OK: %d, errors: %d", countOk, countErr)
	totalErrors += int64(countErr)
}

func compareAccountBalances() {
	countOk, countErr := 0, 0

	log.Printf("Checking account balances...")
	readFileRows("data/account.csv", func (row string) {
		address := common.HexToAddress(row)

		balance1 := rpc1.GetBalance(address, hexutil.Big(block))
		balance2 := rpc2.GetBalance(address, hexutil.Big(block))

		if balance1 != balance2 {
			log.Printf("Balance different for %s:\nRPC1: %s\nRPC2: %s", address, balance1, balance2)
			countErr++
		} else {
			if countOk % 100 == 0 {
				log.Printf("Account balances - OK: %d, errors: %d", countOk, countErr)
			}
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

		nonce1 := rpc1.GetNonce(address, hexutil.Big(block))
		nonce2 := rpc2.GetNonce(address, hexutil.Big(block))

		if nonce1 != nonce2 {
			log.Printf("Nonce different for %s:\nRPC1: %s\nRPC2: %s", address, nonce1, nonce2)
			countErr++
		} else {
			if countOk % 100 == 0 {
				log.Printf("Account nonces - OK: %d, errors: %d", countOk, countErr)
			}
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

		nonce1, err1 := rpc1.Erc20Name(address, block)
		nonce2, err2 := rpc2.Erc20Name(address, block)

		if err1 != nil && err2 != nil {
			log.Printf("Skip %s: %s / %s", address, err1, err2)
			countSkipped++
			return // failed on both servers
		}

		if nonce1 != nonce2 {
			log.Printf("ERC-20 name different for %s:\nRPC1: %s\nRPC2: %s", address, nonce1, nonce2)
			countErr++
		} else {
			if countOk % 100 == 0 {
				log.Printf("ERC-20 names - OK: %d, errors: %d", countOk, countErr)
			}
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

		balance1, err1 := rpc1.Erc20BalanceOf(token, owner, block)
		balance2, err2 := rpc2.Erc20BalanceOf(token, owner, block)

		if err1 != nil && err2 != nil {
			log.Printf("Skip %s/%s: %s / %s", token, owner, err1, err2)
			countSkipped++
			return // failed on both servers
		}

		if balance1.String() != balance2.String() {
			log.Printf("ERC-20 balance different for %s/%s:\nRPC1: %s\nRPC2: %s", token, owner, balance1.String(), balance2.String())
			countErr++
		} else {
			if countOk % 100 == 0 {
				log.Printf("ERC-20 balances - OK: %d, errors: %d", countOk, countErr)
			}
			countOk++
		}
	})

	log.Printf("ERC-20 balances checked - OK: %d, skipped: %d, errors: %d", countOk, countSkipped, countErr)
	totalErrors += int64(countErr)
}
