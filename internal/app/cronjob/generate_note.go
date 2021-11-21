package cronjob

import (
	"context"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/chaosi-zju/daily-problem/internal/pkg/consts"
	"github.com/chaosi-zju/daily-problem/internal/pkg/model"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
)

func GenerateNote(ctx context.Context) {
	problems := make([]model.Problem, 0)
	err := mysqld.Db.Raw(consts.SelectAllProblems, 1).Scan(&problems).Error
	if err != nil {
		log.Errorf("%+v", err)
		return
	}

	problemMap := make(map[string]map[string][]model.Problem)
	for _, p := range problems {
		if _, ok := problemMap[p.Type]; !ok {
			problemMap[p.Type] = make(map[string][]model.Problem)
		}
		problemMap[p.Type][p.SubType] = append(problemMap[p.Type][p.SubType], p)
	}

	key1 := make([]string, 0)
	key2 := make(map[string][]string, 0)
	for k1, subMap := range problemMap {
		key1 = append(key1, k1)
		for k2 := range subMap {
			key2[k1] = append(key2[k1], k2)
		}
	}

	sort.Strings(key1)
	for k := range key2 {
		sort.Strings(key2[k])
	}

	path := viper.GetString("path.note")

	for i, k1 := range key1 {
		path1 := path + "nav." + strconv.Itoa(i+1) + "." + k1 + "/"
		for j, k2 := range key2[k1] {
			path2 := path1 + strconv.Itoa(j+1) + "-" + k2 + "/"
			err = os.MkdirAll(path2, os.ModePerm)
			if err != nil {
				log.Errorf("%+v", err)
				continue
			}
			var ok bool
			var f *os.File
			var arr []model.Problem
			if arr, ok = problemMap[k1][k2]; !ok {
				log.Errorf("k1: %s, k2: %s, not exist", k1, k2)
				continue
			}
			for _, p := range arr {
				filename := path2 + p.Name + ".md"
				f, err = os.Create(filename)
				if err != nil {
					log.Errorf("%+v", err)
					continue
				}
				content := "### " + p.Name + "\n"
				if p.Link != "" && !strings.Contains(p.Link, "chaosi-zju.com") {
					content += "[OJ链接](" + p.Link + ")\n\n"
				}
				if p.Content != "" && p.Type == "interview" {
					content += "### 一句话总结\n"
				}
				content += p.Content + "\n### 解答\n" + p.Result
				_, err = io.WriteString(f, content)
				if err != nil {
					log.Errorf("k1: %s, k2: %s, not exist", k1, k2)
				}
			}
		}
		_, err = os.Create(path1 + "README.md")
		if err != nil {
			log.Errorf("%+v", err)
		}
	}

	log.Infof("write to %s success", path)
}
