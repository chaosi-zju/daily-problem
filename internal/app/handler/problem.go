package handler

import (
	"github.com/chaosi-zju/daily-problem/internal/pkg/model"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
	"github.com/chaosi-zju/daily-problem/internal/pkg/util"
	"github.com/gin-gonic/gin"
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

func FinishProblem(c *gin.Context) {

}
