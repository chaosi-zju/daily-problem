import {get, post} from './request';

// 登陆
export const login = (form) => post('/api/login', form)
// 注册
export const register = (form) => post('/api/register', form)
// 获取每日概况
export const getTodayOverview = (form) => get('/api/user/overview', form)

// 获取每日的题目
export const dailyProblem = (form) => get('/api/user/problem/daily', form)
// 按ID获取题目
export const getProblemByID = (form) => get('/api/user/problem/get', form)
// 完成题目
export const finishProblem = (form) => get('/api/user/problem/finish', form)
// 新增题目
export const insertProblem = (form) => post('/api/user/problem/add', form)
// 更新题目
export const updateProblem = (form) => post('/api/user/problem/update', form)
// 获取所有题目类别
export const getAllTypes = (form) => get('/api/user/problem/types', form)
// 再来几道题
export const moreProblem = (form) => get('/api/user/problem/pickmore', form)

// 获取学习计划
export const problemPlan = (form) => get('/api/user/problem/plan/all', form)
// 加入今日任务
export const addToDaily = (form) => get('/api/user/problem/plan/doitnow', form)
// 移出学习计划
export const removeFromProblemPlan = (form) => get('/api/user/problem/plan/remove', form)

// 获取习题广场
export const problemSquare = (form) => get('/api/user/problem/unplanned', form)
// 加入学习计划
export const addToProblemPlan = (form) => get('/api/user/problem/plan/add', form)