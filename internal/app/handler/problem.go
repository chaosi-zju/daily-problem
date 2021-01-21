package handler

import (
	"github.com/chaosi-zju/daily-problem/internal/pkg/model"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
	"github.com/chaosi-zju/daily-problem/internal/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

const (
	GetDailyProblemSQL = `select * from problems where id in (select problem_id from user_problems where user_id = ? and (pick_time > CURDATE() or finished = false))`
)

func GetDailyProblem(c *gin.Context) {
	userId, err := util.GetUserIdFromContext(c)
	if err != nil {
		util.ResponseError(c, 500, err.Error())
	}

	var problems []model.Problem
	err = mysqld.Db.Raw(GetDailyProblemSQL, userId).Scan(&problems).Error
	if err != nil {
		util.ResponseError(c, 500, "db error")
		return
	}

	util.ResponseSuccess(c, problems)
}

func AddProblem(c *gin.Context) {
	var err error
	var user model.User
	var problem model.Problem

	if err := c.BindJSON(&problem); err != nil {
		util.ResponseError(c, 500, "param invalid")
		return
	}

	if err := validator.New().Struct(problem); err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	if user, err = util.GetUserFromContext(c); err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	problem.Creator = user
	if err := mysqld.Db.Create(&problem).Error; err != nil {
		util.ResponseError(c, 500, "add problem failed: "+err.Error())
	} else {
		util.ResponseSuccess(c, problem)
	}

}

func UpdateProblem(c *gin.Context) {
	var param model.UpdateProblemParam

	// 1. 参数解析
	if err := c.BindJSON(&param); err != nil {
		util.ResponseError(c, 500, "param invalid")
		return
	}

	// 2. 参数校验
	if err := validator.New().Struct(param); err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	// 3. 判断该题是否存在
	problem, err := model.GetProblemByID(param.ID)
	if err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	// 4. 判断当前用户是否有修改的权限 (创建者才有权限)
	if userId, err := util.GetUserIdFromContext(c); err != nil || userId != problem.CreatorID {
		util.ResponseError(c, 500, "您不是该题创建者，没有权限修改")
		return
	}

	// 5. 执行修改
	if err = mysqld.Db.Model(&problem).Updates(param.ToMap()).Error; err != nil {
		util.ResponseError(c, 500, "update problem failed: "+err.Error())
	} else {
		util.ResponseSuccess(c, problem)
	}
}

func FinishProblem(c *gin.Context) {
	userId, err := util.GetUserIdFromContext(c)
	if err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	problemId, err := strconv.Atoi(c.Query("problem_id"))
	if err != nil || problemId == 0 {
		util.ResponseError(c, 500, "no problem_id specificed in url")
		return
	}

	var up model.UserProblem
	err = mysqld.Db.Where("user_id = ? and problem_id = ? and finished = ?", userId, problemId, false).First(&up).Error
	if err == nil {
		up.Finished = true
		up.Times++
		if err = mysqld.Db.Save(&up).Error; err == nil {
			util.ResponseSuccess(c, nil)
			return
		}
	} else if err == gorm.ErrRecordNotFound {
		util.ResponseError(c, 500, "problem_id invalid or has been finished")
		return
	}

	log.Errorf("finish problem error: %+v", err)
	util.ResponseError(c, 500, "db error")
}

func ShouldNotRedo(c *gin.Context) {
	userId, err := util.GetUserIdFromContext(c)
	if err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	problemId, err := strconv.Atoi(c.Query("problem_id"))
	if err != nil || problemId == 0 {
		util.ResponseError(c, 500, "no problem_id specificed in url")
		return
	}

	up := model.UserProblem{UserId: userId, ProblemId: uint(problemId)}
	if err = mysqld.Db.Where(&up).First(&up).Error; err == nil {
		up.ShouldRedo = false
		if err = mysqld.Db.Save(&up).Error; err == nil {
			util.ResponseSuccess(c, nil)
			return
		}
	}

	log.Errorf("set problem not redo error: %+v", err)
	util.ResponseError(c, 500, "db error")
}
