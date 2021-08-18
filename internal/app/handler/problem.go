package handler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/chaosi-zju/daily-problem/internal/app/cronjob"
	"github.com/chaosi-zju/daily-problem/internal/pkg/consts"
	"github.com/chaosi-zju/daily-problem/internal/pkg/model"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
	"github.com/chaosi-zju/daily-problem/internal/pkg/util"
)

func GetDailyProblem(c *gin.Context) {
	userId, err := util.GetUserIdFromContext(c)
	if err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	problems := make([]model.Problem, 0)
	err = mysqld.Db.Raw(consts.GetDailyProblemSQL, userId).Scan(&problems).Error
	if err != nil {
		util.ResponseError(c, 500, "db error")
		return
	}

	util.ResponseSuccess(c, problems)
}

func GetCommonDailyProblem(c *gin.Context) {
	hour := "0"
	specArr := strings.Split(viper.GetString("cron.pick_problem"), " ")
	if len(specArr) >= 3 {
		hour = specArr[2]
	}

	problems := make([]model.Problem, 0)
	err := mysqld.Db.Raw(consts.GetCommonDailyProblemSQL, hour).Scan(&problems).Error
	if err != nil {
		util.ResponseError(c, 500, "db error")
		return
	}

	util.ResponseSuccess(c, problems)
}

func PickMoreProblem(c *gin.Context) {
	userId, err := util.GetUserIdFromContext(c)
	if err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	ttype := c.Query("type")
	if ttype == "" {
		util.ResponseError(c, 500, "no type specificed")
		return
	}

	num, err := strconv.Atoi(c.Query("num"))
	if err != nil || num <= 0 {
		util.ResponseError(c, 500, "no valid num specificed")
		return
	}

	// pick problem
	if err := cronjob.PickProblemByTypeAndNum(userId, ttype, num); err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	// get problem
	problems := make([]model.Problem, 0)
	err = mysqld.Db.Raw(consts.GetDailyProblemSQL, userId).Scan(&problems).Error
	if err != nil {
		util.ResponseError(c, 500, "选题成功，但刷新失败，请重新刷新")
		return
	}

	util.ResponseSuccess(c, problems)
}

func AddProblem(c *gin.Context) {
	var err error
	var user model.User
	var problem model.Problem

	if err = c.BindJSON(&problem); err != nil {
		util.ResponseError(c, 500, "param invalid")
		return
	}

	if err = validator.New().Struct(problem); err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	if user, err = util.GetUserFromContext(c); err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	problem.Creator = user
	if err = mysqld.Db.Create(&problem).Error; err != nil {
		util.ResponseError(c, 500, "add problem failed: "+err.Error())
	} else {
		if err = problem.AddToUserStudyPlan(user.ID); err != nil {
			// 只做log记录
			log.Errorf("add problem %d to user %d study plan falied: %+v", problem.ID, user.ID, err)
		}
		if err = problem.LogCreateProblem(); err != nil {
			// 只做log记录
			log.Errorf("log user %d created problem %d falied: %+v", user.ID, problem.ID, err)
		}
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

func GetProblemByID(c *gin.Context) {
	problemId, err := strconv.Atoi(c.Query("problem_id"))
	if err != nil || problemId == 0 {
		util.ResponseError(c, 500, "no problem_id specificed in url")
		return
	}

	if problem, err := model.GetProblemByID(uint(problemId)); err == nil {
		util.ResponseSuccess(c, problem)
		return
	}

	util.ResponseError(c, 500, "problem not found")
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
	if err == gorm.ErrRecordNotFound {
		util.ResponseError(c, 500, "problem_id invalid or has been finished")
		return
	}

	if err == nil {
		if err = up.FinishProblem(); err == nil {
			util.ResponseSuccess(c, nil)
			return
		}
	}

	log.Errorf("finish problem error: %+v", err)
	util.ResponseError(c, 500, "db error")
}

func AddToDaily(c *gin.Context) {
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
	err = mysqld.Db.Where("user_id = ? and problem_id = ?", userId, problemId).First(&up).Error
	if err == nil {
		if up.Picked == true && up.Finished == false {
			util.ResponseError(c, 500, "该题已在今日任务，不需重复加入")
			return
		}
		up.Picked = true
		up.PickTime = time.Now()
		up.Finished = false
		if err = mysqld.Db.Save(&up).Error; err == nil {
			util.ResponseSuccess(c, nil)
			return
		}
	}

	util.ResponseError(c, 500, "db error")
}

func AddToUserPlan(c *gin.Context) {
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

	var problem model.Problem
	if problem, err = model.GetProblemByID(uint(problemId)); err == nil {
		if err = problem.AddToUserStudyPlan(userId); err == nil {
			util.ResponseSuccess(c, nil)
			return
		}
	}

	log.Errorf("add problem %d to user %d study plan falied: %+v", problemId, userId, err)
	util.ResponseError(c, 500, "add to user plan failed")
}

func RemoveFromPlan(c *gin.Context) {
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

	up := model.UserProblem{UserID: userId, ProblemID: uint(problemId)}
	if err = mysqld.Db.Where(&up).First(&up).Error; err == nil {
		if err = up.RemoveFromUserPlan(); err == nil {
			util.ResponseSuccess(c, nil)
			return
		}
	}

	log.Errorf("remove problem: %d from user: %d 's plan error: %+v", problemId, userId, err)
	util.ResponseError(c, 500, "db error")
}

func GetAllUnPlanned(c *gin.Context) {
	userId, err := util.GetUserIdFromContext(c)
	if err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	problems := make([]model.Problem, 0)
	err = mysqld.Db.Raw(consts.SelectProblemUnplannedSQL, userId, userId).Scan(&problems).Error
	if err != nil {
		util.ResponseError(c, 500, "db error")
		return
	}

	util.ResponseSuccess(c, problems)
}

func GetAllPlanned(c *gin.Context) {
	userId, err := util.GetUserIdFromContext(c)
	if err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	problems := make([]model.Problem, 0)
	err = mysqld.Db.Raw(consts.SelectProblemPlannedSQL, userId).Scan(&problems).Error
	if err != nil {
		util.ResponseError(c, 500, "db error")
		return
	}

	util.ResponseSuccess(c, problems)
}

func GetAllTypes(c *gin.Context) {
	userId, err := util.GetUserIdFromContext(c)
	if err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	if types, err := model.GetUserProblemType(userId); err != nil {
		log.Errorf("get user(%d) problem types failed: %+v", userId, err)
		util.ResponseError(c, 500, "get user problem types failed")
	} else {
		util.ResponseSuccess(c, types)
	}
}

func GetTodayOverview(c *gin.Context) {
	user, err := util.GetUserFromContext(c)
	if err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	num, times := 0, 0
	if err = mysqld.Db.Raw(consts.SelectDoneProblemSQL, user.ID).Row().Scan(&num, &times); err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	type res struct {
		Cnt int `json:"cnt"`
	}

	numArr := make([]res, 2)
	hour := "0"
	specArr := strings.Split(viper.GetString("cron.pick_problem"), " ")
	if len(specArr) >= 3 {
		hour = specArr[2]
	}
	err = mysqld.Db.Raw(consts.SelectTodayWorkloadSQL, user.ID, user.ID, hour).Scan(&numArr).Error
	if err != nil || len(numArr) < 2 {
		util.ResponseError(c, 500, fmt.Sprintf("查询今日任务完成量失败, err: %+v, numArr: %+v", err, numArr))
		return
	}

	// 如果今日没做完的题数>0，则还未完成今日任务
	hasDone := true
	if numArr[0].Cnt > 0 {
		hasDone = false
	}

	util.ResponseSuccess(c, model.OverviewRes{
		PersistDay:   user.PersistDay,
		InterruptDay: user.InterruptDay,
		PersistNum:   num,
		PersistTimes: times,
		Todulist: []model.TodoItem{
			{Done: hasDone, Content: fmt.Sprintf("今日应做 %d 道题，已完成 %d 道题", numArr[0].Cnt+numArr[1].Cnt, numArr[1].Cnt)},
		},
	})

}

func GetFinishInfo(c *gin.Context) {
	user, err := util.GetUserFromContext(c)
	if err != nil {
		util.ResponseError(c, 500, err.Error())
		return
	}

	type item struct {
		Date   string `json:"date"`
		User   string `json:"user"`
		Amount string `json:"amount"`
	}

	finishInfoList := make([]item, 0)
	hour := 0
	specArr := strings.Split(viper.GetString("cron.pick_problem"), " ")
	if len(specArr) >= 3 {
		hour, _ = strconv.Atoi(specArr[2])
	}

	err = mysqld.Db.Raw(consts.SelectFinishInfoSQL, hour).Scan(&finishInfoList).Error
	if err != nil {
		util.ResponseError(c, 500, "查询各用户完成量失败"+err.Error())
		return
	}

	type item2 struct {
		Date        string `json:"date"`
		ProblemType string `json:"problem_type"`
		Count       int    `json:"count"`
	}

	finishInfoList2 := make([]item2, 0)
	err = mysqld.Db.Raw(consts.SelectFinishChartSQL, hour, user.ID).Scan(&finishInfoList2).Error
	if err != nil {
		util.ResponseError(c, 500, "查询该用户历史完成量失败"+err.Error())
		return
	}

	xData := make([]string, 30)
	startDate := time.Now()
	for i := 29; i >= 0; i-- {
		xData[i] = startDate.Format("01-02")
		startDate = startDate.AddDate(0, 0, -1)

	}
	infoMap := make(map[string]map[string]int)
	for _, v := range finishInfoList2 {
		if _, ok := infoMap[v.ProblemType]; !ok {
			infoMap[v.ProblemType] = make(map[string]int)
		}
		infoMap[v.ProblemType][v.Date] = v.Count
	}
	yData := make(map[string][]int)
	for ttype, v := range infoMap {
		yData[ttype] = make([]int, 0)
		for _, date := range xData {
			if cnt, ok := v[date]; ok {
				yData[ttype] = append(yData[ttype], cnt)
			} else {
				yData[ttype] = append(yData[ttype], 0)
			}
		}
	}
	finishInfoChart := map[string]interface{}{
		"xData": xData,
		"yData": yData,
	}

	util.ResponseSuccess(c, map[string]interface{}{
		"list_info":  finishInfoList,
		"chart_info": finishInfoChart,
	})
}
