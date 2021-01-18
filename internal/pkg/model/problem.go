package model

import (
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
	"gorm.io/gorm"
	"time"
)

type Problem struct {
	gorm.Model
	Name    string `gorm:"column:name"`              //题目标题
	Content string `gorm:"column:content;type:text"` //题解内容
	Result  string `gorm:"column:result;type:text"`  //题目答案
	Link    string `gorm:"column:link"`              //题目链接
	Type    string `gorm:"column:type"`              //题目类别 algorithm、sql...
	SubType string `gorm:"column:sub_type"`          //题目子类别 图、树、数组...
}

type UserProblem struct {
	gorm.Model
	UserId      uint      `gorm:"column:user_id;primaryKey"`    //用户ID
	ProblemId   uint      `gorm:"column:problem_id;primaryKey"` //题目ID
	ProblemType string    `gorm:"column:problem_type"`          //题目类型
	PickTime    time.Time `gorm:"column:pick_time"`             //选题时间
	Finished    bool      `gorm:"column:finished"`              //是否完成
	ShouldRedo  bool      `gorm:"column:should_redo"`           //是否需要重做
	Times       int       `gorm:"column:times"`                 //已做次数
}

func GetAllProblemType() ([]string, error) {
	sql := `select distinct(type) from problems`
	rows, err := mysqld.Db.Raw(sql).Rows()
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	var types []string
	for rows.Next() {
		tmp := ""
		rows.Scan(&tmp)
		types = append(types, tmp)
	}

	return types, nil
}
