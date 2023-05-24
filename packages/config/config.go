package config

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type ConsumerKeyConfigurations struct {
	ConsumerKey       string //`mapstructure:"consumer_key" validate:"required"`
	ConsumerSecretKey string //`mapstructure:"consumer_secret_key" validade:"required"`
}

type AccessKeyConfigurations struct {
	AcessKey       string //`mapstructure:"acess_key" validade:"required"`
	AcessSecretKey string //`mapstructure:"acess_secret_key" validade:"required"`
}

// Configurations exported
type Configurations struct {
	AccessKey   AccessKeyConfigurations
	ConsumerKey ConsumerKeyConfigurations
}

var instance *Configurations
var once sync.Once

func Load() *Configurations {
	if instance == nil {
		if os.Args[0][len(os.Args[0])-5:] == ".test" {
			instance = &Configurations{}
		} else {
			once.Do(func() {
				instance = loadConfig()
			})
		}
	}
	return instance

}

func loadConfig() (config *Configurations) {

	// set defaults
	config = &Configurations{}

	// Get all mapstructure
	//registerVariables(config, nil)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.Unmarshal(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = validator.New().Struct(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}
