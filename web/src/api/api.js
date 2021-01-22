import {get, post} from './request';
// 登陆
export const login = (form) => post('/api/login', form)
// 注册
export const register = (form) => post('/api/register', form)
// 获取每日的题目
export const dailyProblem = (form) => get('/api/user/problem/daily', form)
// 完成题目
export const finishProblem = (form) => get('/api/user/problem/finish', form)

// 获取学习计划
export const problemPlan = (form) => get('/api/user/problem/plan/all', form)
// 移出学习计划
export const removeFromProblemPlan = (form) => get('/api/user/problem/plan/remove', form)

// 获取习题广场
export const problemSquare = (form) => get('/api/user/problem/unplanned', form)
// 加入学习计划
export const addToProblemPlan = (form) => get('/api/user/problem/plan/add', form)