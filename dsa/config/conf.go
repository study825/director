package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"sync"
)

var (
	conf *Config
	once sync.Once
)

type redisStruct struct {
	Address  string `toml:"address"`
	Username string `toml:"username"`
}

type Config struct {
	RedisConf redisStruct `toml:"redis"`
}

func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	cfg, err := os.ReadFile("./config/test.toml")
	if err != nil {
		log.Printf("read err=%s", err.Error())
		return
	}

	_, err = toml.Decode(string(cfg), &conf)
	if err != nil {
		log.Printf("decode err=%s", err.Error())
		return
	}
}
