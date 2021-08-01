package conf

import (
	"context"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var (
	conf = flag.String("conf", "conf/config_dev.yaml", "path to config directory")
)

func Init(ctx context.Context) (err error) {
	flag.Parse()

	viper.SetDefault("log.path", "/tmp/daily_problem.log")
	viper.SetConfigFile(*conf)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config, path: %s, err: %+v", *conf, err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("config file changed: %s", e.Name)
	})

	if err := initLogger(); err != nil {
		return err
	}

	return nil
}

func initLogger() error {

	path := viper.GetString("log.path")

	writer, err := rotatelogs.New(
		path+".%Y%m%d%H",
		rotatelogs.WithLinkName("path"),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		return fmt.Errorf("init logger failed: %+v", err)
	}

	logrus.New().Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.ErrorLevel: writer,
		},
		&logrus.JSONFormatter{},
	))

	return nil
}
