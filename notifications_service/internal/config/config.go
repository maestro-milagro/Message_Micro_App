package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env                  string `yaml:"env" env-default:"local"`
	KafkaBootstrapServer string `yaml:"kafka_bootstrap_server" env-default:"localhost:9095"`
	Secret               string `yaml:"secret"`
	DB                   `yaml:"db"`
	HTTPServer           `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DB struct {
	DataBase string `yaml:"database"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBname   string `yaml:"dbname"`
	SSLmode  string `yaml:"sslmode"`
}

func MustLoad() *Config {
	configPath := "config/prod.yaml"

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}

//func MustLoad() *Config {
//	path := fetchConfigPath()
//	if path == "" {
//		panic("config path is empty")
//	}
//
//	return MustLoadByPath(path)
//}
//
//func MustLoadByPath(configPath string) *Config {
//	if _, err := os.Stat(configPath); os.IsNotExist(err) {
//		panic("config file does not exist: " + configPath)
//	}
//
//	var cfg Config
//
//	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
//		panic("failed to read config: " + err.Error())
//	}
//
//	return &cfg
//}
//
//func fetchConfigPath() string {
//	var res string
//
//	flag.StringVar(&res, "config", "", "path to config file")
//	flag.Parse()
//
//	if res == "" {
//		res = os.Getenv("CONFIG_PATH")
//	}
//
//	return res
//}
