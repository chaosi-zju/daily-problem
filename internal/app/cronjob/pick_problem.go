package cronjob

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/chaosi-zju/daily-problem/internal/pkg/model"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
)

// PickProblem the detail of cornjob, pick daily problem for every user
func PickProblem(ctx context.Context) {

	begin := time.Now()

	users, err := model.GetAllUsers()
	if err != nil {
		panic("get all users failed: " + err.Error())
	}

	// 遍历每一个用户
	for _, user := range users {
		if err := PickProblemByUser(user); err != nil {
			log.Errorf("%+v", err)
		}
	}

	log.Infof("cronjob pick_problem cost %+v", time.Now().Sub(begin))
}

func PickProblemByUser(user model.User) error {
	rand.Seed(time.Now().UnixNano())

	// 查询该用户的所有题类型
	types, err := model.GetUserProblemType(user.ID)
	if err != nil {
		return fmt.Errorf("get problem types of user %d failed: %+v", user.ID, err)
	}
	// 遍历该用户的每一种题类型
	complete := true
	for _, ttype := range types {

		// 1. 如果用户自定义了各类型的题目数量
		dailyNum := viper.GetInt("problem.default_daily_num")
		if num, ok := user.Config.ProblemNum[ttype]; ok {
			dailyNum = num
		}
		// 2. 对于昨天新出的题目，如果完成的题数小于要求的题数，则不达标 (移出学习计划的也算完成)
		if num, err := GetDailyDoneNum(user.ID, ttype); err == nil && num < dailyNum {
			complete = false
		}
		// 3. 出今天的新题
		if err := PickProblemByTypeAndNum(user.ID, ttype, dailyNum); err != nil {
			log.Errorf("pick problem error: %+v", err)
		}
	}
	// 依据用户是否完成昨天的任务，更新用户相应字段
	if err := user.UpdateComplete(complete); err != nil {
		log.Errorf("update user error: %+v", err)
	}

	return nil
}

func PickProblemByTypeAndNum(userid uint, ttype string, dailyNum int) error {
	// 遍历出题次数
	for i := 0; i < dailyNum; i++ {

		// 优先从没选过的题中选
		if picked, err := model.PickRandonUnPicked(userid, ttype); err != nil {
			log.Errorf("select unpicked problem for user(%d) failed: %+v", userid, err)
		} else if picked {
			continue
		}

		// 再从做过的题中选
		if picked, err := model.PickRandonFinished(userid, ttype); err != nil {
			log.Errorf("select finished problem for user(%d) failed: %+v", userid, err)
		} else if picked {
			continue
		}

		return fmt.Errorf("no %s problem left for user(%d) to pick", ttype, userid)
	}
	return nil
}

func GetDailyDoneNum(userId uint, ttype string) (int, error) {
	up := model.UserProblem{
		UserId:      userId,
		ProblemType: ttype,
		Picked:      true,
		Finished:    true,
	}

	var cnt int64 = 0
	err := mysqld.Db.Unscoped().Model(&up).Where(&up).Where("TO_DAYS(NOW()) - TO_DAYS(pick_time) = 1").Count(&cnt).Error

	return int(cnt), err
}
