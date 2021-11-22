package cronjob

import (
	"bufio"
	"context"
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/chaosi-zju/daily-problem/internal/pkg/conf"
	"github.com/chaosi-zju/daily-problem/internal/pkg/model"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
)

var fileMap = map[string]string{}

func TestMdToMysql(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "conf.path", "../../../conf/config_dev.yaml")

	conf.Init(ctx)
	mysqld.Init(ctx)
	model.Init(ctx)

	for name, path := range fileMap {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		var title, result string
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "## ") {
				recordToMysql("interview", name, title, result)
				title = line
				result = ""
			} else {
				result += line + "\n"
			}
		}
		file.Close()
	}

}

func TestGenerateNote(t *testing.T) {
	GenerateNote(context.Background())
}

func recordToMysql(typ, subType, title, result string) {
	if title == "" || result == "" {
		return
	}
	if strings.HasPrefix(title, "## ") {
		title = strings.Split(title, "## ")[1]
	} else if strings.HasPrefix(title, "### ") {
		title = strings.Split(title, "### ")[1]
	}

	title = strings.ReplaceAll(title, "/", "、")

	user := model.User{
		Model: gorm.Model{ID: 1},
	}
	problem := model.Problem{
		Name:      title,
		Content:   "",
		Result:    result,
		Link:      "",
		Type:      typ,
		SubType:   subType,
		IsPublic:  true,
		CreatorID: user.ID,
		Creator:   user,
	}
	if err := mysqld.Db.Create(&problem).Error; err != nil {
		log.Errorf("add problem failed: " + err.Error())
	} else {
		if err = problem.AddToUserStudyPlan(user.ID); err != nil {
			// 只做log记录
			log.Errorf("add problem %d to user %d study plan falied: %+v", problem.ID, user.ID, err)
		}
		if err = problem.LogCreateProblem(); err != nil {
			// 只做log记录
			log.Errorf("log user %d created problem %d falied: %+v", user.ID, problem.ID, err)
		}
	}
}
