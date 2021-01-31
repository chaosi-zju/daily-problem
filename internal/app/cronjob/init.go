package cronjob

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
	spec := viper.GetString("cron.pick_problem")
	log.Infof("cron pick_problem spec: %s", spec)

	if _, err := cron.AddFunc(spec, PickProblem, "pick_problem"); err != nil {
		return fmt.Errorf("add cron func pick_problem failed: %+v", err)
	}

	return nil
}
