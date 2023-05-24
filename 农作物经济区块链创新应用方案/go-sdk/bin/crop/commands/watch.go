package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"strconv"

	"net/http"
	"time"

	insurances "github.com/FISCO-BCOS/go-sdk/contracts"

	"github.com/urfave/cli/v2"

	"log"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
)

type WeatherInfo struct {
	Status   string `json:"status"`
	Count    string `json:"count"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	Lives    []struct {
		Province      string `json:"province"`
		City          string `json:"city"`
		Adcode        string `json:"adcode"`
		Weather       string `json:"weather"`
		Temperature   string `json:"temperature"`
		Winddirection string `json:"winddirection"`
		Windpower     string `json:"windpower"`
		Humidity      string `json:"humidity"`
		Reporttime    string `json:"reporttime"`
	} `json:"lives"`
}

var WatchCommand = &cli.Command{
	Name:  "watch",
	Usage: "watch the data of Iot and weather.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "api",
			Aliases: []string{"k"},
			Value:   "https://restapi.amap.com/v3/weather/weatherInfo?key=f80ceb04d1e696388879aafe9c152ee4&city=510100&extensions=base&output=JSON",
			Usage:   "the api of weather",
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

		// 间隔24小时获取一次数据
		//ticker := time.NewTicker(24 * time.Hour)
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:

				//获取Iot设备检测的土壤等数据（暂未实现）
				/*
					xxxx
					xxxx
					xxxx
				*/

				// 模拟获取天气数据
				resp, err := http.Get("https://restapi.amap.com/v3/weather/weatherInfo?key=f80ceb04d1e696388879aafe9c152ee4&city=510100&extensions=base&output=JSON")
				if err != nil {
					return err
				}
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(string(body))

				var wi WeatherInfo
				err = json.Unmarshal(body, &wi)
				if err != nil {
					return err
				}
				temperature, err := strconv.ParseInt(wi.Lives[0].Temperature, 10, 8)
				if err != nil {
					return err
				}
				humidity, err := strconv.ParseInt(wi.Lives[0].Humidity, 10, 8)
				if err != nil {
					return err
				}

				// 上传天气数据
				tx, _, err := insuranceSession.UploadWeathers(wi.Lives[0].Province, wi.Lives[0].City, wi.Lives[0].Weather, int8(temperature), wi.Lives[0].Winddirection, wi.Lives[0].Windpower, int8(humidity))
				if err != nil {
					return err
				}
				fmt.Println("transaction hash: ", tx.Hash().Hex())

				// 如果Iot数据异常，执行相应操作
				//insuranceSession.CensorInsurance()

				// 如果天气恶劣，则执行保险合约
				//insuranceSession.InsuranceIds()
				//insuranceSession.ExecuteInsurance()
			}
		}
	},
}
