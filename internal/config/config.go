package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Addr struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type Config struct {
	Name   string `json:"name"`
	Server Addr   `json:"addr"`
}

var GlobalConfig = &Config{}

func LoadConfig(fileName *string) *Config {
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("loadConfig: file open error, fileName=%s, err=%v.", fileName, err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("loadConfig: ioutil.ReadAll error, fileName=%s, err=%v.", fileName, err)
	}

	err = json.Unmarshal(data, GlobalConfig)
	if err != nil {
		log.Fatalf("loadConfig: json.Unmarshal error, fileName=%s, err=%v.", fileName, err)
	}
	return GlobalConfig
}
