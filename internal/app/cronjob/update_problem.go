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
	UnPickedCount        = `select count(*) from problems where id not in (select problem_id from user_problems where user_id = ?)`
	SelectRandUnPicked   = `select id from (select * from problems where id not in (select problem_id from user_problems where user_id = ?)) as unpicked limit ?,1`
	UnFinishedCount      = `select count(*) from user_problems where user_id = ? and finished = true and should_redo = true`
	SelectRandUnFinished = `select * from (select * from user_problems where user_id = ? and finished = true and should_redo = true) as picked limit ?,1`
)

func UpdateProblem(ctx context.Context) {
	users, err := model.GetAllUsers()
	if err != nil {
		panic("get all users failed: " + err.Error())
	}

	dailyNum := viper.GetInt("problem.daily_num")
	rand.Seed(time.Now().UnixNano())

	// 遍历每一个用户
	for _, user := range users {
		// 遍历出题次数
		for i := 0; i < dailyNum; i++ {
			// 查询有多少道没选过的题
			unpickCnt := 0
			err := mysqld.Db.Raw(UnPickedCount, user.ID).Row().Scan(&unpickCnt)
			if err == nil && unpickCnt > 0 {
				// 从没选过的题中随机选一道
				var pickId uint = 0
				offset := rand.Intn(unpickCnt)
				err = mysqld.Db.Raw(SelectRandUnPicked, user.ID, offset).Row().Scan(&pickId)
				if err == nil {
					// 将该题插入user_problem表中
					up := model.UserProblem{
						UserId:     user.ID,
						ProblemId:  pickId,
						PickTime:   time.Now(),
						Finished:   false,
						ShouldRedo: true,
						Times:      0,
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
			err = mysqld.Db.Raw(UnFinishedCount, user.ID).Row().Scan(&unfinishCnt)
			if err == nil && unfinishCnt > 0 {
				// 再选一道没做完的
				var up model.UserProblem
				offset := rand.Intn(unfinishCnt)
				err = mysqld.Db.Raw(SelectRandUnFinished, user.ID, offset).Row().Scan(&up)
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
				log.Errorf("select picked problem of %s failed: %+v", user.Name, err)
			}
		}
	}
}
