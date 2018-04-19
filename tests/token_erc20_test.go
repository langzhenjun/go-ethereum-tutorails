package tests

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/langzhenjun/go-ethereum-tutorials/contracts"
	"github.com/langzhenjun/go-ethereum-tutorials/utils"
)

func TestTokenERC20(t *testing.T) {
	//
	client, err := ethclient.Dial("../.geth/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v\r\n", err)
	}

	//
	configs := utils.LoadConfigs("../configs.json")
	mainAccountAddressHex := configs.MainAccount.AddressHex
	mainAccountKeyJSON := configs.MainAccount.KeyJSON
	mainAccountPassword := configs.MainAccount.Password

	//
	mainAccountKeyJSONBytes, err := json.Marshal(mainAccountKeyJSON)
	mainAccountKeyJSONString := string(mainAccountKeyJSONBytes)
	if err != nil {
		t.Fatalf("Failed to get key json: %v\r\n", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(mainAccountKeyJSONString), mainAccountPassword)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v\r\n", err)
	}

	//
	ERC20AddressHex := configs.Contracts["ERC20"]
	var ERC20 *contracts.TokenERC20
	if ERC20AddressHex == "" {
		// 发布合约
		ERC20Address, _, token, err := contracts.DeployTokenERC20(auth, client, big.NewInt(100), "TokenERC20", "ERC20")
		if err != nil {
			t.Fatalf("Failed to deploy TokenERC20: %v\r\n", err)
		}
		ERC20 = token
		configs.Contracts["ERC20"] = ERC20Address.Hex()
		configs.Save()
	} else {
		// 获取合约
		ERC20, err = contracts.NewTokenERC20(common.HexToAddress(ERC20AddressHex), client)
		if err != nil {
			t.Fatalf("Failed to deploy contract: %v\r\n", err)
		}
	}

	name, err := ERC20.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		t.Fatalf("Failed to retrieve pending name: %v\r\n", err)
	}

	t.Logf("Pending Contract: %v\r\n", name)

	ctx := context.Background()
	transferC := make(chan *contracts.TokenERC20Transfer)
	// froms := []common.Address{common.HexToAddress(mainAccountAddressHex)}
	_, err = ERC20.WatchTransfer(&bind.WatchOpts{Context: ctx}, transferC, nil, nil)
	if err != nil {
		log.Fatalf("Failed to watch transfer: %v\r\n", err)
	}
	// defer sub.Unsubscribe()

	const toAddressHex = "0xdc4bb8b33c0aa1eb028c73b27a988c4e0b56140a"

	balanceFromBefore, _ := ERC20.BalanceOf(nil, common.HexToAddress(mainAccountAddressHex))
	balanceToBefore, _ := ERC20.BalanceOf(nil, common.HexToAddress(toAddressHex))
	t.Logf("Before transfere: FROM: %d, TO: %d\n", balanceFromBefore, balanceToBefore)

	//
	tx, err := ERC20.Transfer(auth, common.HexToAddress(toAddressHex), big.NewInt(123456789))
	if err != nil {
		t.Fatalf("Failed to request token transfer: %v", err)
	}

	t.Logf("Pending Transfer : 0x%x\n", tx.Hash())

	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("tx mining error:%v\n", err)
	}

	transfer := <-transferC
	t.Logf("FROM: %v, TO: %v, VALUE: %v\r\n", transfer.From.Hex(), transfer.To.Hex(), transfer.Value)
	valFrom, _ := ERC20.BalanceOf(nil, common.HexToAddress(mainAccountAddressHex))
	valTo, _ := ERC20.BalanceOf(nil, common.HexToAddress(toAddressHex))
	t.Logf("After transfere: FROM: %d, TO: %d\n", valFrom, valTo)
}
