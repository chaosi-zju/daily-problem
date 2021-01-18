package cronjob

import (
	"context"
	"github.com/chaosi-zju/daily-problem/internal/pkg/model"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

const (
	UnPickedCount        = `select count(*) from problems where type = ? and id not in (select problem_id from user_problems where user_id = ?)`
	SelectRandUnPicked   = `select id from (select * from problems where type = ? and id not in (select problem_id from user_problems where user_id = ?)) as unpicked limit ?,1`
	UnFinishedCount      = `select count(*) from user_problems where user_id = ? and problem_type = ? and finished = true and should_redo = true`
	SelectRandUnFinished = `select * from (select * from user_problems where user_id = ? and problem_type = ? and finished = true and should_redo = true) as picked limit ?,1`
)

func UpdateProblem(ctx context.Context) {
	begin := time.Now()

	users, err := model.GetAllUsers()
	if err != nil {
		panic("get all users failed: " + err.Error())
	}

	types, err := model.GetAllProblemType()
	if err != nil {
		panic("get all problem types failed: " + err.Error())
	}

	defaultDailyNum := viper.GetInt("problem.default_daily_num")
	rand.Seed(time.Now().UnixNano())

	// 遍历每一个用户
	for _, user := range users {
		// 遍历每一种题类型
		for _, ttype := range types {
			// 如果用户自定义了各类型的题目数量
			dailyNum := defaultDailyNum
			if num, ok := user.Config.ProblemNum[ttype]; ok {
				dailyNum = num
			}
			// 遍历出题次数
			for i := 0; i < dailyNum; i++ {
				// 查询有多少道没选过的题
				unpickCnt := 0
				err := mysqld.Db.Raw(UnPickedCount, ttype, user.ID).Row().Scan(&unpickCnt)
				if err == nil && unpickCnt > 0 {
					// 从没选过的题中随机选一道
					var pickId uint = 0
					offset := rand.Intn(unpickCnt)
					err = mysqld.Db.Raw(SelectRandUnPicked, ttype, user.ID, offset).Row().Scan(&pickId)
					if err == nil {
						// 将该题插入user_problem表中
						up := model.UserProblem{
							UserId:      user.ID,
							ProblemId:   pickId,
							ProblemType: ttype,
							PickTime:    time.Now(),
							Finished:    false,
							ShouldRedo:  true,
							Times:       0,
						}
						if err = mysqld.Db.Create(&up).Error; err == nil {
							// 选成功了则返回，继续选
							continue
						}
					}
				}

				if err != nil {
					log.Errorf("select unpicked problem of %s failed: %+v", user.Name, err)
				}

				// 如果没选过的题选失败了，就选没做完的，先查没做完的数量
				unfinishCnt := 0
				err = mysqld.Db.Raw(UnFinishedCount, user.ID, ttype).Row().Scan(&unfinishCnt)
				if err == nil && unfinishCnt > 0 {
					// 再选一道没做完的
					var up model.UserProblem
					offset := rand.Intn(unfinishCnt)
					err = mysqld.Db.Raw(SelectRandUnFinished, user.ID, ttype, offset).Scan(&up).Error
					if err == nil {
						up.PickTime = time.Now()
						up.Finished = false
						if err = mysqld.Db.Save(&up).Error; err == nil {
							continue
						}
					}
				}

				// 如果没做完的题也选失败了，报错
				if err != nil {
					log.Errorf("select unfinished problem of %s failed: %+v", user.Name, err)
				} else {
					log.Infof("neither unpicked nor unfinished problems for %s", user.Name)
				}
			}
		}
	}
	log.Infof("cronjob update_problem cost %+v", time.Now().Sub(begin))
}
