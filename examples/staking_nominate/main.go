package main

import (
	"github.com/qazxcvio/avail-go-sdk/src/config"
	"github.com/qazxcvio/avail-go-sdk/src/sdk"
	"github.com/qazxcvio/avail-go-sdk/src/sdk/tx"

	"fmt"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("cannot load config:%v", err)
	}
	api, err := sdk.NewSDK(config.ApiURL)
	if err != nil {
		fmt.Printf("cannot create api:%v", err)
	}
	WaitFor := sdk.BlockInclusion

	stash := []string{"5GNJqTPyNqANBkUVMN1LPPrxXnFouWXoe2wNSmmEoLctxiZY", "5FHneW46xGXgs5mUiveU4sbTyGBzmstUspZC92UhjJM694ty"}
	BlockHash, txHash, err := tx.Nominate(api, config.Seed, WaitFor, stash)
	if err != nil {
		fmt.Printf("cannot submit Transaction:%v", err)
	}
	fmt.Printf("Transaction submitted successfully with block hash: %v\n and ext hash:%v", BlockHash.Hex(), txHash.Hex())
}
