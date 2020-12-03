package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	WatchNs          string
	WatchKey         string
	DownLoadDst      string
	DownLoadUrl      string
	DownLoadFileName string
}

var Conf Config

func InitConfig(configFilePath string) {
	if configFilePath == "" {
		configFilePath = os.Getenv(configFilePath)
	}
	file, err := os.Open(configFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	Conf = Config{}
	err = json.NewDecoder(file).Decode(&Conf)
	if err != nil {
		panic(err)
	}
}
