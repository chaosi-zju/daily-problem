# daily-problem

日常做题工具

#### 目录结构
```
|-- .github              // github action config
|-- cmd
|   |-- main.go          // main entry
|-- conf                 // local config
|-- internal             // main backend logic
|   |-- app
|   |   |-- cronjob    
|   |   |-- handler
|   |   |-- middleware
|   |   |-- routes.go
|   |-- pkg
|   |   |-- ......       // self library
|-- web                  // front page by Vue
|-- wx                   // front page by weixin app
|-- makefile             // how to build
```

#### 本地运行
```
// 导入数据库
conf/daily_problem.sql

// 启动server端
在goland中运行main函数

// 启动web前端
cd web
npm install
npm run serve
```


- [x] 新用户未达标天数是1

- [x] 修sql

- [x] 查询每日待办

- [x] 查询坚持天数

- [x] 定时任务选题前先计算数据

- [x] 计算做过多少题，多少次题

- [x] 移出时软删除，加入时先判断是否被软删除过

- [x] 首页

- [x] 退出登录删除token

- [x] 数据库初始化sql
  
- [x] 题目跳转的bug

- [x] 怎么配置密码envsubst

- [x] 不能iframe的怎么提示
 
- [x] 怎么持续化部署

- [x] 加入新用户选默认题
 
- [x] 不够？再出几道

- [x] 加入今日任务

- [x] 本地iframe怎么把题目嵌过来

- [x] 点插入题目时不是页面首部

- [x] 做题时全屏，做完退出全屏

- [x] 顶部栏

- [x] 登录注册按钮不对齐

- [x] 用于确认的对话框

- [x] 怎么新增题目

- [x] 怎么修改题目

- [x] 查看答案，怎么用md展示

- [x] 题目广场

- [x] 个人学习计划

- [x] 如果每日的题目不到3题怎么处理

- [x] 查询每日的3题怎么查

- [x] 定时任务选题怎么选

- [x] 插入题目时要不要加入学习计划

- [x] 每日题数支持用户定义 

