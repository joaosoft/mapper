package mapper

import (
	gomanager "github.com/joaosoft/go-manager"
	"github.com/joaosoft/logger"
)

// Mapper ...
type Mapper struct {
	config        *MapperConfig
	pm            *gomanager.Manager
	isLogExternal bool
}

// NewMapper ...
func NewMapper(options ...MapperOption) *Mapper {
	config, simpleConfig, err := NewConfig()
	pm := gomanager.NewManager(gomanager.WithRunInBackground(false))

	mapper := &Mapper{
		pm:     pm,
		config: &MapperConfig{},
	}

	if mapper.isLogExternal {
		pm.Reconfigure(gomanager.WithLogger(log))
	}

	if err == nil {
		mapper.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Mapper.Log.Level)
		log.Debugf("setting log level to %s", level)
		log.Reconfigure(logger.WithLevel(level))
	}

	mapper.Reconfigure(options...)

	return mapper
}
