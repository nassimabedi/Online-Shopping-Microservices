package config

type Microservices struct {
	UserService microserviceInfo `yaml:"user_service"`
}

type microserviceInfo struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}