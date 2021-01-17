package cron

import (
	"context"
	"fmt"
	"github.com/chaosi-zju/daily-problem/internal/pkg/util"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

var (
	c *cron.Cron
)

func Init(_ context.Context) error {
	c = cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.DelayIfStillRunning(cron.DefaultLogger),
		),
	)
	c.Start()
	return nil
}

type jobFunc func(ctx context.Context)

func AddFunc(spec string, f jobFunc, name string) (cron.EntryID, error) {
	job := jobWithRecovery(f, name)
	return c.AddFunc(spec, job)
}

func jobWithRecovery(f jobFunc, name string) func() {
	return func() {
		ctx := context.Background()
		defer func() {
			if err := recover(); err != nil {
				errMsg := util.PrintStackTrace(fmt.Errorf("%s failed: %v", name, err))
				log.Errorf("cronjob_failed: %s", errMsg)
			}
		}()

		log.Infof("cronjob %s start...", name)
		defer log.Infof("cronjob %s finished...", name)

		f(ctx)
	}
}
