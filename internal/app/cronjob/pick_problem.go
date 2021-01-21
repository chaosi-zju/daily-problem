package cronjob

import (
	"context"
	"github.com/chaosi-zju/daily-problem/internal/pkg/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

func UpdateProblem(ctx context.Context) {

	begin := time.Now()
	rand.Seed(time.Now().UnixNano())
	defaultDailyNum := viper.GetInt("problem.default_daily_num")

	users, err := model.GetAllUsers()
	if err != nil {
		panic("get all users failed: " + err.Error())
	}

	// 遍历每一个用户
	for _, user := range users {
		// 查询该用户的所有题类型
		types, err := model.GetUserProblemType(user.ID)
		if err != nil {
			log.Errorf("get problem types of user %d failed: %+v", user.ID, err)
			continue
		}
		// 遍历该用户的每一种题类型
		for _, ttype := range types {

			// 如果用户自定义了各类型的题目数量
			dailyNum := defaultDailyNum
			if num, ok := user.Config.ProblemNum[ttype]; ok {
				dailyNum = num
			}
			// 遍历出题次数
			for i := 0; i < dailyNum; i++ {

				// 优先从没选过的题中选
				if picked, err := model.PickRandonUnPicked(user.ID, ttype); err != nil {
					log.Errorf("select unpicked problem for %s failed: %+v", user.Name, err)
				} else if picked {
					continue
				}

				// 再从做过的题中选
				if picked, err := model.PickRandonFinished(user.ID, ttype); err != nil {
					log.Errorf("select finished problem for %s failed: %+v", user.Name, err)
				} else if picked {
					continue
				}

				log.Errorf("no problem left for %s to pick", user.Name)
			}
		}
	}
	log.Infof("cronjob update_problem cost %+v", time.Now().Sub(begin))
}
