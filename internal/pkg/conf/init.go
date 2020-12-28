package conf

import (
	"context"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	mode    = flag.String("mode", "dev", "operational mode")
	confDir = flag.String("conf_dir", "conf", "path to config directory")
)

func Init(ctx context.Context) (err error) {
	flag.Parse()
	configPath := fmt.Sprintf("%s/config_%s.yaml", *confDir, *mode)

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config, path: %s, err: %+v", configPath, err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("config file changed: %s", e.Name)
	})

	return nil
}
