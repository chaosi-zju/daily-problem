package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/chaosi-zju/daily-problem/internal/app"
	"github.com/chaosi-zju/daily-problem/internal/app/cronjob"
	"github.com/chaosi-zju/daily-problem/internal/pkg/conf"
	"github.com/chaosi-zju/daily-problem/internal/pkg/cron"
	"github.com/chaosi-zju/daily-problem/internal/pkg/model"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
	"github.com/chaosi-zju/daily-problem/internal/pkg/util"
)

func init() {
	log.Info("init start...")
	defer log.Info("init finish...")

	runInit(context.Background(), conf.Init, mysqld.Init, model.Init, cron.Init, cronjob.Init)
}

func main() {

	r := app.SetupRoutes()
	businessCh := make(chan error, 1)
	go func() {
		port := viper.GetString("gin.port")
		if err := r.Run(":" + port); err != nil {
			businessCh <- err
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	log.Info("listening for signals ...")
	select {
	case <-signalCh:
		log.Info("signal received, gracefully shutdown...")
	case err := <-businessCh:
		log.WithField("Error", err).Fatal("failed to start gin service")
	}

}

type funcType func(ctx context.Context) error

func runInit(ctx context.Context, fs ...funcType) {
	for _, f := range fs {
		funcName := util.FuncName(f)
		log.Infof("init [%s] start...", funcName)

		if err := f(ctx); err != nil {
			panic(fmt.Sprintf("init [%s] failed: %+v", funcName, err))
		}

		log.Infof("init [%s] finish...", funcName)
	}
}
