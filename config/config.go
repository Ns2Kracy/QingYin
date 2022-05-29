package config

import (
	"encoding/json"
	"os"
)

type QingYin struct {
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode       string `json:"mode"`
	MySQL      MySQL  `json:"mysql"`
	Redis      Redis  `json:"redis"`
}

func InitConfig() *QingYin {
	var config *QingYin
	file, err := os.Open("./config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err.Error())
	}
	return config
}
