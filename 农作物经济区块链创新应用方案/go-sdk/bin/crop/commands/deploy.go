package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"

	"log"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/contracts"
)

var DeployCommand = &cli.Command{
	Name:  "deploy",
	Usage: "deploy the contract",

	Action: func(ctx *cli.Context) error {

		configs, err := conf.ParseConfigFile("config.toml")
		if err != nil {
			log.Fatalf("ParseConfigFile failed, err: %v", err)
		}
		client, err := client.Dial(&configs[0])
		if err != nil {
			log.Fatal(err)
		}
		contractAddress, tx, _, err := insurances.DeployInsurances(client.GetTransactOpts(), client)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("contract address: ", contractAddress.Hex())
		fmt.Println("transaction hash: ", tx.Hash().Hex())

		return nil
	},
}
