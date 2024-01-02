package configuration

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func NewConfiguration[T any](input *T) (T, error) {
	if input == nil {
		return *new(T), fmt.Errorf("input is nil")
	}

	readFile(input, "config", "app", "")

	return *input, nil
}

func readFile(cfg interface{}, confPath, confFileName, envPrefix string) {
	ymlConfig := viper.New()
	ymlConfig.AddConfigPath(confPath)
	ymlConfig.SetConfigName(confFileName)
	replacer := strings.NewReplacer(".", "_", "-", "_")
	ymlConfig.SetEnvKeyReplacer(replacer)
	ymlConfig.SetEnvPrefix(envPrefix)
	ymlConfig.AutomaticEnv()

	err := ymlConfig.ReadInConfig()
	if err != nil {
		processError(err)
	}
	err = ymlConfig.Unmarshal(&cfg)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
