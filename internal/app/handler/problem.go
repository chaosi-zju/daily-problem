package handler

import (
	"github.com/gin-gonic/gin"
)

const (
	SQL = `select * from problems where finished = true and should_redo = true order by pick_time, times limit 3 group by type
           union 
           `

//`select * from problems where type = %s and finished = false`
)

func GetDailyProblem(c *gin.Context) {

	c.JSON(200, 2)
}
