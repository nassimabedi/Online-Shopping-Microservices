/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"Online-Shopping-Microservices/microservices/wallet-service/config"
	"Online-Shopping-Microservices/microservices/wallet-service/console/cmd"
	"Online-Shopping-Microservices/microservices/wallet-service/repository"
	"flag"
)

func init() {
	configFlag := flag.String("config", "dev", "config flag can be dev for develop or prod for production")
	prodConfigPath := flag.String("config-path", "", "config-path production config file path")
	// init service configs
	config.Init(configFlag, prodConfigPath)

	// init repositories
	repository.Init()

}
func main() {
	cmd.Execute()
}
