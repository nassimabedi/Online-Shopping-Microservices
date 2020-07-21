package main

import (
	"flag"

	"arvan.ir/app-services/discount-service/config"

	httpEngine "arvan.ir/app-services/discount-service/controller/http"
	"arvan.ir/app-services/discount-service/repository"
	"sync"
)

func main() {
	configFlag := flag.String("config", "dev", "config flag can be dev for develop or prod for production")
	prodConfigPath := flag.String("config-path", "", "config-path production config file path")

	// init service configs
	config.Init(configFlag, prodConfigPath)

	// init repositories
	repository.Init()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go httpEngine.Run(config.Configs.Service.HttpPort)
	//go grpcEngine.Run(config.Configs.Service.GrpcPort)
	wg.Wait()
}
