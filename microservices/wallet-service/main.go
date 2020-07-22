package main

import (
	"arvan.ir/app-services/wallet-service/config"
	httpEngine "arvan.ir/app-services/wallet-service/controller/http"
	"arvan.ir/app-services/wallet-service/repository"
	"flag"
	"sync"
)

func main() {
	configFlag := flag.String("config", "dev", "config flag can be dev for develop or prod for production")
	prodConfigPath := flag.String("config-path", "", "config-path production config file path")

	// init service configs
	config.Init(configFlag, prodConfigPath)

	// init repositories
	repository.Init()

	// run http and grpc servers
	wg := sync.WaitGroup{}
	wg.Add(2)
	go httpEngine.Run(config.Configs.Service.HttpPort)
	//go grpcEngine.Run(config.Configs.Service.GrpcPort)
	wg.Wait()
}
