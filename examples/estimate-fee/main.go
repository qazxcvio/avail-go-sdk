package main

import (
	"encoding/json"
	"log"
	"math"
	"math/big"

	"github.com/availproject/avail-go-sdk/src/config"
	"github.com/availproject/avail-go-sdk/src/sdk"
	"github.com/availproject/avail-go-sdk/src/sdk/types"
	"github.com/vedhavyas/go-subkey"

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

	amount := uint64(math.Pow(10, 18)) * 10 // send amount 10 AVAIL
	dest := "5FHneW46xGXgs5mUiveU4sbTyGBzmstUspZC92UhjJM694ty"
	keyringPair, err := sdk.KeyringFromSeed(config.Seed)
	if err != nil {
		panic(fmt.Sprintf("cannot create KeyPair:%v", err))
	}
	_, pubkeyBytes, _ := subkey.SS58Decode(dest)
	hexString := subkey.EncodeHex(pubkeyBytes)
	destAddr, err := sdk.NewMultiAddressFromHexAccountID(hexString)
	if err != nil {
		log.Fatalf("Failed to create address from given hex: %v", err)
	}
	bondAmountUCompact := types.NewUCompactFromUInt(amount)

	ext, err := sdk.CreateExtrinsic(api, "Balances.transfer_keep_alive", keyringPair, 0, destAddr, bondAmountUCompact)
	if err != nil {
		log.Fatalf("Failed to create extrinsic: %v", err)
	}

	encodedExt, err := sdk.EncodeToHex(ext)
	if err != nil {
		log.Fatalf("Failed to encode extrinsic: %v", err)
	}
	fmt.Println("Encoded Extrinsic:", encodedExt)
	var paymentInfo map[string]interface{}
	err = api.Client.Call(&paymentInfo, "payment_queryInfo", encodedExt, nil)
	if err != nil {
		log.Fatalf("Failed to get payment info: %v", err)
	}

	// Format the weight
	weight := paymentInfo["weight"].(map[string]interface{})
	weightJSON, err := json.Marshal(map[string]interface{}{
		"refTime":   weight["ref_time"],
		"proofSize": weight["proof_size"],
	})
	if err != nil {
		log.Fatalf("Failed to marshal weight: %v", err)
	}

	// Format the partial fee
	partialFee, ok := new(big.Int).SetString(paymentInfo["partialFee"].(string), 10)
	if !ok {
		log.Fatalf("Failed to parse partialFee")
	}
	mAVAIL := new(big.Float).Quo(new(big.Float).SetInt(partialFee), big.NewFloat(1e15))
	fmt.Printf("Transaction Fee for Balance Transfer:\n")
	fmt.Printf("    class=%s,\n", paymentInfo["class"])
	fmt.Printf("    weight=%s,\n", string(weightJSON))
	fmt.Printf("    partialFee=%.4f mAVAIL\n", mAVAIL)
}
