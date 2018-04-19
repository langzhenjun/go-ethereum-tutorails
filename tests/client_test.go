package tests

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

var globalAccountForTesting accounts.Account

func TestNewAccount(t *testing.T) {
	ks := keystore.NewKeyStore("../.geth/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount("123456")
	if err != nil {
		t.Errorf("Failed to create a new account: %v", err)
	}

	_, err = ks.Export(account, "123456", "123456")
	if err != nil {
		t.Errorf("Failed to export a new account: %v", err)
	}

	t.Logf("Successed created account: %v", account.Address.Hex())

	client, err := ethclient.Dial("../.geth/geth.ipc")
	if err != nil {
		t.Errorf("Failed to connect to ethereum: %v", err)
	}

	balance, err := client.BalanceAt(context.TODO(), account.Address, nil)
	if err != nil {
		t.Errorf("Failed to get balance of account: %v", err)
	}

	if balance.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Exception of BalanceAt address: %v", account.Address.Hex())
	}

	t.Logf("Successed get balance at address: %v\r\n", account.Address.Hex())
}
