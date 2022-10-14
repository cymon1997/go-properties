package provider

import (
	"sync"

	"github.com/cymon1997/go-properties/internal/config"
)

var (
	mainConfig     *config.MainConfig
	syncMainConfig sync.Once
)

func GetMainConfig() *config.MainConfig {
	syncMainConfig.Do(func() {
		mainConfig = config.New()
	})
	return mainConfig
}
