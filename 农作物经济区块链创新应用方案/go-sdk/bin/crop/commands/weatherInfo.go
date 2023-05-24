package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"log"
	"math/big"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/contracts"
	"github.com/ethereum/go-ethereum/common"
)

var WeatherInfoCommand = &cli.Command{
	Name:  "weather",
	Usage: "look up weather info",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "index",
			Aliases:  []string{"i"},
			Value:    0,
			Usage:    "the index of weather infos",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "address",
			Aliases:  []string{"addr"},
			Value:    "",
			Usage:    "the contract address",
			Required: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		address := ctx.String("address")
		index := ctx.Int64("index")

		configs, err := conf.ParseConfigFile("config.toml")
		if err != nil {
			log.Fatalf("ParseConfigFile failed, err: %v", err)
		}
		client, err := client.Dial(&configs[0])
		if err != nil {
			log.Fatal(err)
		}
		//load the contract
		contractAddress := common.HexToAddress(address)
		instance, err := insurances.NewInsurances(contractAddress, client)
		if err != nil {
			log.Fatal(err)
		}
		insuranceSession := &insurances.InsurancesSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

		info, err := insuranceSession.WeatherInfos(big.NewInt(index))
		if err != nil {
			log.Fatal("why?", err)
		}
		fmt.Printf("%+v", info)
		return nil
	},
}
