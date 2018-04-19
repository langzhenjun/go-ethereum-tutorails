package main

import (
	"log"

	"github.com/langzhenjun/xiuhu/utils"
)

// const adminAddress = "0xf4cf445afe8945f76dea4cbcb80e82d18a4940ed"
// const adminKey = `{"address":"f4cf445afe8945f76dea4cbcb80e82d18a4940ed","crypto":{"cipher":"aes-128-ctr","ciphertext":"5ea7cc5184f29e108b9d8c59656479f5dd9430e372930e43795d6786d01d8b00","cipherparams":{"iv":"dce98618e1481ac4b55ce7f84efca685"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"cd892a78881b14d05ba90e266716c9b4784aa865a8fd895b709ffeab45183b8e"},"mac":"df384ac124e77d0f2b6aca4e775a641c82530485962a21a817e74e051aea8d9e"},"id":"63caf746-be4e-400d-84f9-ac275992a04e","version":3}`

func main() {
	configs := utils.LoadConfigs("./configs.json")
	log.Printf("%v", configs.MainAccount.KeyJSON)

	// 创建一个到本地节点的连接
	// client, err := ethclient.Dial(".geth/geth.ipc")
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the Ethereum client: %v\r\n", err)
	// }

	// //
	// auth, err := bind.NewTransactor(strings.NewReader(adminKey), "")
	// if err != nil {
	// 	log.Fatalf("Failed to create authorized transactor: %v", err)
	// }

	// ERC20, err := examples.NewTokenERC20(common.HexToAddress("0x4b933105503fbb8d806c5bc64645a3a14979aa7f"), client)
	// // 发布合约
	// // contractAddress, _, tokenDeployed, err := contracts.DeployTokenERC20(auth, client, big.NewInt(100), "TEST", "TEST")
	// if err != nil {
	// 	log.Fatalf("Failed to deploy contract: %v", err)
	// }

	// log.Printf("Deployed contract: 0x%x\n", contractAddress)

	// ctx := context.Background()

	// transferChan := make(chan *examples.TokenERC20Transfer)
	// froms := []common.Address{common.HexToAddress("0xf4cf445afe8945f76dea4cbcb80e82d18a4940ed")}
	// sub, err := ERC20.WatchTransfer(&bind.WatchOpts{Context: ctx}, transferChan, froms, nil)
	// if err != nil {
	// 	log.Fatalf("Failed to watch transfer: %v\r\n", err)
	// }

	// defer sub.Unsubscribe()

	// tx, err := ERC20.Transfer(auth, common.HexToAddress("0xdc4bb8b33c0aa1eb028c73b27a988c4e0b56140a"), big.NewInt(123456789))
	// if err != nil {
	// 	log.Fatalf("Failed to request token transfer: %v", err)
	// }

	// log.Printf("Transfer pending: 0x%x\n", tx.Hash())

	// _, err = bind.WaitMined(ctx, client, tx)
	// if err != nil {
	// 	log.Fatalf("tx mining error:%v\n", err)
	// }

	// valFrom, _ := ERC20.BalanceOf(nil, common.HexToAddress("0xf4cf445afe8945f76dea4cbcb80e82d18a4940ed"))
	// valTo, _ := ERC20.BalanceOf(nil, common.HexToAddress("0xdc4bb8b33c0aa1eb028c73b27a988c4e0b56140a"))

	// fmt.Printf("after transfere:%d, %d\n", valFrom, valTo)

	// name, err := ERC20.Name(&bind.CallOpts{Pending: true})
	// if err != nil {
	// 	log.Fatalf("Failed to retrieve pending name: %v", err)
	// }
	// fmt.Println("Pending name:", name)

	// for tran := range transferChan {
	// 	log.Printf("%v, %v, %v\r\n", tran.From.Hex(), tran.To.Hex(), tran.Value)
	// }

}
