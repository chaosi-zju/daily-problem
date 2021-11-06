# daily-problem

站点直通车：[Daily Problem](http://www.chaosi-zju.com)

#### 简介
相信你平时学习各种知识应该也会有记笔记或者写博客的习惯吧？

但是很多时候，我们记到笔记上之后就很少去看它，等到真正想去看时，太多东西可能让你漫无目的地随便翻翻，啥收获也没有

然而你学一个东西，当时会了，你以为你懂了，其实你可能并没有真正理解，你需要反复温习，温故而知新

这时候你会不会想拥有一个出题的工具，每天针对过去做的笔记随机推送几条作为当天的学习任务，督促你把选出来的题看完

如果没有完成任务就会一直堆积直到你看完，如果你有新的感悟也可以更新其内容

于是我就实现了这么个出题工具，每天给自己出几道题，反复强化学习

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

