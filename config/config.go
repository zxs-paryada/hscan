package config

import (
	"log"

	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

// Config wraps all config
type Config struct {
	Node  NodeConfig  `yaml:"node"`
	Web   WebConfig   `yaml:"web"`
	Mysql MysqlConfig `yaml:"mysql"`
	Total TotalConfig `yaml:"total"`
}

// NodeConfig wraps all node endpoints that are used in this project
type NodeConfig struct {
	NodeServerEndPoint string `yaml:"node_server_endpoint"`
	LCDServerEndpoint  string `yaml:"lcd_server_endpoint"`
	PriServerEndpoint  string `yaml:"pri_server_endpoint"`
}

// WebConfig wraps all required paramaters for boostraping web server
type WebConfig struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type MysqlConfig struct {
	MysqlRes string `yaml:"mysql_res"`
}

type TotalConfig struct {
	Supplement string `yaml:"supplement"`
}

// ParseConfig attempts to read and parse config.yaml from the given path
// An error reading or parsing the config results in a panic.
func ParseConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(errors.Wrap(err, "failed to read config"))
	}

	cfg := Config{}

	if viper.GetString("active") == "" {
		log.Fatal("define active param in your config file.")
	}

	switch viper.GetString("active") {
	case "mainnet":
		cfg.Node = NodeConfig{
			NodeServerEndPoint: viper.GetString("mainnet.node.node_server_endpoint"),
			LCDServerEndpoint:  viper.GetString("mainnet.node.lcd_server_endpoint"),
			PriServerEndpoint:  viper.GetString("mainnet.node.pri_server_endpoint"),
		}

		cfg.Web = WebConfig{
			Ip:   viper.GetString("mainnet.web.ip"),
			Port: viper.GetString("mainnet.web.port"),
		}

		cfg.Mysql = MysqlConfig{
			MysqlRes: viper.GetString("mainnet.mysql.mysql_res"),
		}

		cfg.Total = TotalConfig{
			Supplement: viper.GetString("mainnet.total.supplement"),
		}

	case "testnet":
		cfg.Node = NodeConfig{
			NodeServerEndPoint: viper.GetString("testnet.node.node_server_endpoint"),
			LCDServerEndpoint:  viper.GetString("testnet.node.lcd_server_endpoint"),
			PriServerEndpoint:  viper.GetString("testnet.node.pri_server_endpoint"),
		}

		cfg.Web = WebConfig{
			Ip:   viper.GetString("testnet.web.ip"),
			Port: viper.GetString("testnet.web.port"),
		}

		cfg.Mysql = MysqlConfig{
			MysqlRes: viper.GetString("testnet.mysql.mysql_res"),
		}

		cfg.Total = TotalConfig{
			Supplement: viper.GetString("testnet.total.supplement"),
		}

	default:
		log.Fatalf("active parameter in config.yaml cannot be set as '%s'", viper.GetString("active"))
	}

	return &cfg
}
