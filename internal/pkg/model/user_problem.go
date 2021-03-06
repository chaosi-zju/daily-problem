package model

import (
	"github.com/chaosi-zju/daily-problem/internal/pkg/consts"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
	"gorm.io/gorm"
	"time"
)

type UserProblem struct {
	gorm.Model
	UserId      uint      `gorm:"column:user_id;primaryKey"`    //用户ID
	ProblemId   uint      `gorm:"column:problem_id;primaryKey"` //题目ID
	ProblemType string    `gorm:"column:problem_type"`          //题目类型
	Picked      bool      `gorm:"column:picked"`                //是否被选
	PickTime    time.Time `gorm:"column:pick_time"`             //选题时间
	Finished    bool      `gorm:"column:finished"`              //是否完成
	Times       int       `gorm:"column:times"`                 //已做次数
}

func (up UserProblem) RemoveFromUserPlan() error {
	// 移出学习计划默认你完成了该题
	up.Finished = true
	up.Times++
	up.DeletedAt.Time = time.Now()
	up.DeletedAt.Valid = true
	return mysqld.Db.Save(&up).Error
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
