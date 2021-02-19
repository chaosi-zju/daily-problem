package model

type LoginParam struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterParam struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password" validate:"required"`
}

type UpdateProblemParam struct {
	ID       uint   `json:"id" validate:"required"`   //题目ID
	Name     string `json:"name" validate:"required"` //题目标题
	Content  string `json:"content"`                  //题解内容
	Result   string `json:"result"`                   //题目答案
	Link     string `json:"link"`                     //题目链接
	Type     string `json:"type"`                     //题目类别 algorithm、sql...
	SubType  string `json:"sub_type"`                 //题目子类别 图、树、数组...
	IsPublic bool   `json:"is_public"`                //题目是否公开
}

func (param *UpdateProblemParam) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Name":     param.Name,
		"Content":  param.Content,
		"Result":   param.Result,
		"Link":     param.Link,
		"Type":     param.Type,
		"SubType":  param.SubType,
		"IsPublic": param.IsPublic,
	}
}

type OverviewRes struct {
	PersistDay   int        `json:"persist_day"`
	InterruptDay int        `json:"interrupt_day"`
	PersistNum   int        `json:"persist_num"`
	PersistTimes int        `json:"persist_times"`
	Todulist     []TodoItem `json:"todulist"`
}

type TodoItem struct {
	Done    bool   `json:"done"`
	Content string `json:"content"`
}
