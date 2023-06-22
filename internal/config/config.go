package config

/**
TODO: implement loading config from flags
*/

import (
	"errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"strings"
)

var (
	InvalidEnvPrefix = errors.New("invalid env prefix")
)

type AppConfig struct {
	envPrefix string
	defaults  []DefaultValue
}

func NewAppConfig(envPrefix string, defaults []DefaultValue) *AppConfig {
	return &AppConfig{
		envPrefix: envPrefix,
		defaults:  defaults,
	}
}

type DefaultValue struct {
	Name  string
	Flag  string
	Value interface{}
}

func (r *AppConfig) InitializeAppConfig() error {
	if len(r.envPrefix) == 0 {
		return InvalidEnvPrefix
	}

	viper.SetEnvPrefix(r.envPrefix)
	viper.AutomaticEnv()

	pflag.Parse()
	if r.defaults != nil {
		for _, value := range r.defaults {
			viper.SetDefault(value.Name, value.Value)
		}
	}

	return nil
}

func (r *AppConfig) GetConfigValues(s interface{}) error {
	return viper.Unmarshal(s)
}

func ExtractConfigValues(config interface{}) ([]DefaultValue, error) {
	confType := reflect.TypeOf(config).Elem()
	numFields := confType.NumField()
	res := make([]DefaultValue, 0, numFields)
	for i := 0; i < numFields; i++ {
		field := confType.Field(i)
		if field.Anonymous {
			continue
		}
		envVariableName := field.Tag.Get("mapstructure")
		defaultValue := field.Tag.Get("default")
		flagName := field.Tag.Get("flag")
		res = append(res, DefaultValue{
			Name:  envVariableName,
			Flag:  flagName,
			Value: defaultValue,
		})
	}
	return res, nil
}

func IsDevelopment() bool {
	de := os.Getenv("development")
	if strings.Compare(strings.ToLower(de), "true") == 0 {
		return true
	}
	return false
}
