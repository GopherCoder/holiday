package configs

import (
	"holiday/pkg/logger"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func (c *Config) initConfig() {

	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("configs")
		viper.SetConfigName("settings")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	var info = logger.Info{
		Package: "config",
		Action:  "Watch Config",
		Message: "Config file changed",
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.ErrorLog(info)
	})
}
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	c.initConfig()
	c.watchConfig()

	return nil
}
