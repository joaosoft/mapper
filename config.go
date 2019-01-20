package mapper

import (
	"fmt"

	gomanager "github.com/joaosoft/go-manager"
)

// AppConfig ...
type AppConfig struct {
	Mapper *MapperConfig `json:"mapper"`
}

// MapperConfig ...
type MapperConfig struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
}

// NewConfig ...
func NewConfig() (*MapperConfig, error) {
	appConfig := &AppConfig{}
	if _, err := gomanager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig); err != nil {
		log.Error(err.Error())

		return &MapperConfig{}, err
	}

	return appConfig.Mapper, nil
}
