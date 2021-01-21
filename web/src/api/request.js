import axios from "axios";
import router from "../router/router";
import {Loading} from "element-ui";
import {messages} from '@/assets/js/common'
import store from '../store/store'

axios.defaults.timeout = process.env.VUE_APP_HTTP_TIMEOUT;
axios.defaults.baseURL = process.env.VUE_APP_LOGOUT_URL;
axios.defaults.headers.post["Content-Type"] = "application/json";
let loading = null;

/*
 *请求前拦截
 *用于处理需要请求前的操作
 */
axios.interceptors.request.use(
    config => {
        loading = Loading.service({
            text: "正在加载中......",
            fullscreen: true
        });
        if (store.state.token) {
            config.headers["Authorization"] = "Bearer " + store.state.token;
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

/*
 *请求响应拦截
 *用于处理数据返回后的操作
 */
axios.interceptors.response.use(
    response => {
        return new Promise((resolve, reject) => {
            //请求成功后关闭加载框
            if (loading) {
                loading.close();
            }
            const res = response.data;
            if (res.code === 200) {
                resolve(res.data)
            } else {
                messages("error", res.message)
                reject(res)
            }
        })
    },
    error => {
        console.log(error)
        // 请求成功后关闭加载框
        if (loading) {
            loading.close();
        }
        console.log(error.response)
        // 断网处理或者请求超时
        if (!error.response) {
            //请求超时
            if (error.message.includes("timeout")) {
                messages("error", "请求超时，请检查互联网连接");
            } else {
                //断网，可以展示断网组件
                messages("error", "请检查网络是否已连接");
            }
            return Promise.reject(error);
        }
        const status = error.response.status;
        switch (status) {
            case 500:
                messages("error", "服务器内部错误");
                break;
            case 404:
                messages("error", "404 page not found");
                break;
            case 401:
                messages("warning", "用户登陆过期，请重新登陆");
                store.state.commit('COMMIT_TOKEN', '')
                setTimeout(() => {
                    router.replace({
                        path: "/login",
                        query: {
                            redirect: router.currentRoute.fullPath
                        }
                    });
                }, 1000);
                break;
            case 400:
                messages("error", "数据异常");
                break;
            default:
                messages("error", error.response.data.message);
        }
        return Promise.reject(error);
    }
);

/*
 *get方法，对应get请求
 *@param {String} url [请求的url地址]
 *@param {Object} params [请求时候携带的参数]
 */
export function get(url, params) {
    return new Promise((resolve, reject) => {
        axios
            .get(url, {
                params
            })
            .then(res => {
                resolve(res);
            })
            .catch(err => {
                reject(err);
            });
    });
}

/*
 *post方法，对应post请求
 *@param {String} url [请求的url地址]
 *@param {Object} params [请求时候携带的参数]
 */
export function post(url, params) {
    return new Promise((resolve, reject) => {
        axios
            .post(url, params)
            .then(res => {
                resolve(res);
            })
            .catch(err => {
                reject(err);
            });
    });
}