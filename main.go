package painterconfig

import (
	"github.com/just-arun/painter-config/config"
)

var (
	Init func(path string, name string, fileType string)                                          = config.InitConfig
	Load func(path string, name string, fileType string) (configuration config.Config, err error) = config.LoadVars
)

func main() {
	
}