import {get, post} from './request';
//登陆
export const login = (loginForm) => post('/api/login1', loginForm)
//注册
export const register = (registerForm) => post('/api/register', registerForm)
//上传
export const upload = (upload) => get('/api/get/upload', upload)