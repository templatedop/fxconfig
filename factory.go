package fxconfig

import (
	"fmt"
	"github.com/spf13/viper"
)

type ConfigFactory interface {
	Create(options ...ConfigOption) (Econfig, error)
}

type DefaultConfigFactory struct{}

func NewDefaultConfigFactory() ConfigFactory {
	return &DefaultConfigFactory{}
}

func (f *DefaultConfigFactory) Create(options ...ConfigOption) (Econfig, error) {

	appliedOptions := DefaultConfigOptions
	for _, opt := range options {
		opt(&appliedOptions)
	}
	v := viper.New()
	v.SetConfigName(appliedOptions.FileName)
	v.AddConfigPath(appliedOptions.FilePath)
	//v.AddConfigPath(".")
	v.SetConfigType(appliedOptions.FileExt)

	if err := v.ReadInConfig(); err != nil {
		return Econfig{}, err

	}
	var config Econfig
	if err := v.Unmarshal(&config); err != nil {
		return Econfig{}, err

	}
	fmt.Println("config create", config)
	return config, nil
}
