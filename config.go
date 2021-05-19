package main

import "github.com/spf13/viper"

var (
	AppConfig Config
)

// for loading environmental variable from location
func LoadVars(path, name, fileType string) (configuration Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(fileType)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		return
	}
	return
}

func InitConfig(path, name, fileType string) {
	var err error
	AppConfig, err = LoadVars(path, name, fileType)
	if err != nil {
		panic(err)
	}
}

// func GetVar() config {
// 	conf, _ := LoadVars("./../")
// 	return conf
// }
