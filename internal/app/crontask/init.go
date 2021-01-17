package crontask

import (
	"context"
	"fmt"
	"github.com/chaosi-zju/daily-problem/internal/pkg/cron"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init(ctx context.Context) error {
	if err := initUpdateProblem(ctx); err != nil {
		return fmt.Errorf("init update_problem failed: %+v", err)
	}

	return nil
}

func initUpdateProblem(ctx context.Context) error {
	spec := viper.GetString("cron.update_problem")
	log.Infof("cron update_problem spec: %s", spec)

	if _, err := cron.AddFunc(spec, UpdateProblem, "update_problem"); err != nil {
		return fmt.Errorf("add cron func update_problem failed: %+v", err)
	}

	return nil
}
