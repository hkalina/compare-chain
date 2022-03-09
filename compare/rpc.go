package main

import (
	"compare-chain/compare/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

type FtmBridge struct {
	rpc *rpc.Client
	eth *ethclient.Client
}

func NewFtmBridge(rpcUrl string) *FtmBridge {
	rpcClient, err := rpc.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	ethClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	return &FtmBridge{
		rpc: rpcClient,
		eth: ethClient,
	}
}

func (ftm *FtmBridge) Close() {
	if ftm.rpc != nil {
		ftm.rpc.Close()
		ftm.eth.Close()
	}
}

func (ftm *FtmBridge) GetCode(contract common.Address) string {
	var code string
	if err := ftm.rpc.Call(&code, "eth_getCode", contract, "latest"); err != nil {
		log.Printf("failed eth_getCode: %s", err)
		return ""
	}
	return code
}

func (ftm *FtmBridge) GetBalance(address common.Address) string {
	var code string
	if err := ftm.rpc.Call(&code, "eth_getBalance", address, "latest"); err != nil {
		log.Printf("failed eth_getBalance: %s", err)
		return ""
	}
	return code
}

func (ftm *FtmBridge) GetNonce(address common.Address) string {
	var nonce string
	if err := ftm.rpc.Call(&nonce, "eth_getTransactionCount", address, "latest"); err != nil {
		log.Printf("failed eth_getTransactionCount: %s", err)
		return ""
	}
	return nonce
}

func (ftm *FtmBridge) Erc20Name(token common.Address) (string, error) {
	// connect the contract
	contract, err := contracts.NewERCTwenty(token, ftm.eth)
	if err != nil {
		return "", err
	}

	// get the token name
	symbol, err := contract.Name(nil)
	if err != nil {
		return "", err
	}

	return symbol, err
}

func (ftm *FtmBridge) Erc20BalanceOf(token common.Address, owner common.Address) (hexutil.Big, error) {
	// connect the contract
	contract, err := contracts.NewERCTwenty(token, ftm.eth)
	if err != nil {
		return hexutil.Big{}, err
	}

	// get the balance
	val, err := contract.BalanceOf(nil, owner)
	if err != nil {
		return hexutil.Big{}, err
	}

	// make sur we always have a value; at least zero
	// this should always be the case since the contract should
	// return zero even for unknown owners, but let's be sure here
	if val == nil {
		val = new(big.Int)
	}

	// return the account balance
	return hexutil.Big(*val), nil
}
