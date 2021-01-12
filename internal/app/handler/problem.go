package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

type Param struct {
	Number string
}

func Test(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	var p Param
	_ = json.Unmarshal(data, &p)
	number, _ := strconv.Atoi(p.Number)
	fmt.Println(number)
	c.JSON(200, number*2)
}
