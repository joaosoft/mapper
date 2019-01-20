package mapper

import (
	"fmt"

	gomanager "github.com/joaosoft/go-manager"
	logger "github.com/joaosoft/logger"
)

// Mapper ...
type Mapper struct {
	config        *MapperConfig
	pm            *gomanager.Manager
	isLogExternal bool
}

// NewMapper ...
func NewMapper(options ...MapperOption) *Mapper {
	pm := gomanager.NewManager(gomanager.WithRunInBackground(false))

	mapper := &Mapper{
		pm:     pm,
		config: &MapperConfig{},
	}

	if mapper.isLogExternal {
		pm.Reconfigure(gomanager.WithLogger(log))
	}

	// load configuration file
	appConfig := &AppConfig{}
	if simpleConfig, err := gomanager.NewSimpleConfig(fmt.Sprintf("/config.%s.json", GetEnv()), appConfig); err != nil {
		log.Error(err.Error())
	} else if appConfig.Mapper != nil {
		pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(appConfig.Mapper.Log.Level)
		log.Debugf("setting log level to %s", level)
		log.Reconfigure(logger.WithLevel(level))
		mapper.config = appConfig.Mapper
	}

	mapper.Reconfigure(options...)

	return mapper
}
