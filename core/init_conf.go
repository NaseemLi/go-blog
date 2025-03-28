package core

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var confPath = "settings.yaml"

type Config struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

func ReadConf() {
	byteData, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
	}
	var config Config

	err = yaml.Unmarshal(byteData, &config)
	if err != nil {
		panic(fmt.Sprintf("yaml配置文件格式错误%s", err))
	}
	fmt.Println(config)
}
