package config

type MongoDBConfig struct {
	Addresses string `yaml:"addresses"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}
