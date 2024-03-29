package model

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"

	"github.com/chaosi-zju/daily-problem/internal/pkg/consts"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
)

type Problem struct {
	gorm.Model
	Name      string `gorm:"column:name" json:"name" validate:"required"` //题目标题
	Content   string `gorm:"column:content;type:text" json:"content"`     //题解内容
	Result    string `gorm:"column:result;type:text" json:"result"`       //题目答案
	Link      string `gorm:"column:link" json:"link"`                     //题目链接
	Type      string `gorm:"column:type" json:"type"`                     //题目类别 algorithm、sql...
	SubType   string `gorm:"column:sub_type" json:"sub_type"`             //题目子类别 图、树、数组...
	IsPublic  bool   `gorm:"column:is_public" json:"is_public"`           //题目是否公开
	CreatorID uint   `gorm:"column:creator_id" json:"creator_id"`         //题目创建者ID
	Creator   User   `gorm:"foreignKey:CreatorID" json:"-"`               //题目创建者
}

// 将题目加入某用户的学习计划
func (p Problem) AddToUserStudyPlan(userId uint) error {
	// 先判断该题是否在学习计划，即被软删除过
	var up UserProblem
	err := mysqld.Db.Unscoped().Where(&UserProblem{UserID: userId, ProblemID: p.ID}).First(&up).Error

	if err == gorm.ErrRecordNotFound {
		// 不曾在学习计划，将该题插入user_problem表中
		loc, _ := time.LoadLocation("Local")
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2006-01-02 15:04:05", loc)
		up = UserProblem{
			UserID:      userId,
			ProblemID:   p.ID,
			ProblemType: p.Type,
			Picked:      false,
			PickTime:    t,
			Finished:    false,
			Times:       0,
		}
		return mysqld.Db.Create(&up).Error
	} else if err == nil {
		// 曾在学习计划，恢复
		up.DeletedAt.Valid = false
		return mysqld.Db.Save(&up).Error
	} else {
		return err
	}
}

// 创建题目的日志
func (p Problem) LogCreateProblem() error {
	log := UserProblemLog{
		UserID:      p.CreatorID,
		ProblemID:   p.ID,
		ProblemType: p.Type,
		Action:      CREATE,
		ActionTime:  time.Now(),
	}
	return mysqld.Db.Create(&log).Error
}

// GetProblemByID 获取指定ID的题目
func GetProblemByID(id uint) (Problem, error) {
	var problem Problem
	if err := mysqld.Db.First(&problem, id).Error; err == nil {
		return problem, nil
	}
	return problem, fmt.Errorf("problem not exist")
}

// PickRandonUnPicked 随机选一道没选过的题
func PickRandonUnPicked(userId uint, ttype string) (bool, error) {
	// 查询有多少道没选过的题
	unpickCnt := 0
	err := mysqld.Db.Raw(consts.CountUnPickedProblemSQL, userId, ttype).Row().Scan(&unpickCnt)
	if err == nil && unpickCnt > 0 {
		// 从没选过的题中随机选一道
		var up UserProblem
		offset := rand.Intn(unpickCnt)
		err = mysqld.Db.Raw(consts.SelectRandUnPickedProblemSQL, userId, ttype, offset).Scan(&up).Error
		if err == nil {
			up.Picked = true
			up.PickTime = time.Now()
			up.Finished = false
			return true, mysqld.Db.Save(&up).Error
		}
	}
	return false, err
}

// PickRandonFinished 随机选一道做过的题
func PickRandonFinished(userId uint, ttype string) (bool, error) {
	// 查询做过的题的数量
	finishCnt := 0
	err := mysqld.Db.Raw(consts.CountFinishedProblemSQL, userId, ttype).Row().Scan(&finishCnt)
	if err == nil && finishCnt > 0 {
		// 从做过的题中随机选一道
		var up UserProblem
		offset := rand.Intn(finishCnt)
		err = mysqld.Db.Raw(consts.SelectRandFinishedProblemSQL, userId, ttype, offset).Scan(&up).Error
		if err == nil {
			up.PickTime = time.Now()
			up.Finished = false
			return true, mysqld.Db.Save(&up).Error
		}
	}
	return false, err
}
