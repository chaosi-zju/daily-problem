package model

import (
	"time"

	"gorm.io/gorm"

	"github.com/chaosi-zju/daily-problem/internal/pkg/consts"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
)

const (
	CREATE = "create"
	FINISH = "finish"
)

type UserProblem struct {
	gorm.Model
	UserID      uint      `gorm:"column:user_id;primaryKey"`    //用户ID
	ProblemID   uint      `gorm:"column:problem_id;primaryKey"` //题目ID
	ProblemType string    `gorm:"column:problem_type"`          //题目类型
	Picked      bool      `gorm:"column:picked"`                //是否被选
	PickTime    time.Time `gorm:"column:pick_time"`             //选题时间
	Finished    bool      `gorm:"column:finished"`              //是否完成
	Times       int       `gorm:"column:times"`                 //已做次数
}

type UserProblemLog struct {
	gorm.Model
	UserID      uint      `gorm:"column:user_id;primaryKey"`    //用户ID
	ProblemID   uint      `gorm:"column:problem_id;primaryKey"` //题目ID
	ProblemType string    `gorm:"column:problem_type"`          //题目类型
	Action      string    `gorm:"column:action"`                //操作
	ActionTime  time.Time `gorm:"column:action_time"`           //操作时间
}

func (up UserProblem) RemoveFromUserPlan() error {
	// 移出学习计划默认你完成了该题
	up.Finished = true
	up.Times++
	up.DeletedAt.Time = time.Now()
	up.DeletedAt.Valid = true
	return mysqld.Db.Save(&up).Error
}

func (up UserProblem) FinishProblem() error {
	up.Finished = true
	up.Times++
	if err := mysqld.Db.Save(&up).Error; err != nil {
		return err
	}
	log := UserProblemLog{
		UserID:      up.UserID,
		ProblemID:   up.ProblemID,
		ProblemType: up.ProblemType,
		Action:      FINISH,
		ActionTime:  time.Now(),
	}
	return mysqld.Db.Create(&log).Error
}

func GetUserProblemType(userId uint) ([]string, error) {
	rows, err := mysqld.Db.Raw(consts.SelectUserProblemTypeSQL, userId).Rows()
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	var types []string
	for rows.Next() {
		tmp := ""
		if err = rows.Scan(&tmp); err != nil {
			return types, err
		}
		types = append(types, tmp)
	}

	return types, nil
}
