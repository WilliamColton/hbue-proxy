package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Version       string
	ServerAddress string
	Key           string //这里不能是[]byte QAQ,只能是string，不然会报错 {  []} illegal base64 data at input byte 4
}

func ReadConfig(configSrc string) (Config, error) {
	config, err := os.ReadFile(configSrc)
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = json.Unmarshal(config, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}
