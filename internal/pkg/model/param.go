package model

type LoginParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterParam struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type AddProblemParam struct {
	Name    string `gorm:"column:name"`              //题目标题
	Content string `gorm:"column:content;type:text"` //题解内容
	Result  string `gorm:"column:result;type:text"`  //题目答案
	Link    string `gorm:"column:link"`              //题目链接
	Type    string `gorm:"column:type"`              //题目类别 algorithm、sql...
	SubType string `gorm:"column:sub_type"`          //题目子类别 图、树、数组...
}
