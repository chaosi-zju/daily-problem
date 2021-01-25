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
	conf = flag.String("conf", "conf/config_dev.yaml", "path to config directory")
)

func Init(ctx context.Context) (err error) {
	flag.Parse()

	viper.SetConfigFile(*conf)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config, path: %s, err: %+v", *conf, err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("config file changed: %s", e.Name)
	})

	return nil
}
