package main

import (
	"log"
	"os"

	"github.com/satriaprayoga/capruk/config"
	capruk "github.com/satriaprayoga/capruk/framework"
	"github.com/spf13/viper"
)

func setup(arg1, arg2 string) {
	if arg1 != "new" && arg1 != "version" && arg1 != "help" {
		path, err := os.Getwd()
		if err != nil {
			exitCLI(err)
		}
		capruk.RootPath = path
		var conf = &config.AppConfig{}
		viper.SetConfigFile("config.json")
		err = viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Fail to parse config file")
		}
		err = viper.Unmarshal(conf)
		if err != nil {
			log.Fatalf("Fail to Unmarshall 'config.json': %v", err)
		}
		capruk.Config = conf
	}

}
