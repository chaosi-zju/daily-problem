import {get, post} from './request';
// 登陆
export const login = (form) => post('/api/login', form)
// 注册
export const register = (form) => post('/api/register', form)
// 获取每日的题目
export const dailyProblem = (form) => get('/api/user/problem/daily', form)
// 完成题目
export const finishProblem = (form) => get('/api/user/problem/finish', form)