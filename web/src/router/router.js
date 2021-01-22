import Vue from 'vue'
import Router from 'vue-router'
import store from '../store/store'
import NProgress from 'nprogress' //进度条
import 'nprogress/nprogress.css'
import home from '@/layout/index'

Vue.use(Router)
// 路由懒加载
const getComponent = (name, component) => () => import (`@/views/${name}/${component}.vue`);

const myRouter = new Router({
    routes: [
        {
            path: '/',
            redirect: '/home',
            meta: {
                title: 'home'
            }
        },
        {
            path: '/login',
            component: getComponent('login', 'index')
        },
        {
            path: '/home',
            component: home,
            meta: {
                title: 'home2'
            },
            children: [
                {
                    path: '/home',
                    component: getComponent('home', 'index'),
                    meta: {
                        title: 'home'
                    }
                },
                {
                    path: '/404',
                    component: getComponent('error', '404'),
                    meta: {
                        title: '404'
                    }
                },
                {
                    path: '/403',
                    component: getComponent('error', '403'),
                    meta: {
                        title: '403'
                    }
                },
                {
                    path: '/insertProblem',
                    component: getComponent('problem', 'insert'),
                    meta: {
                        title: 'insertProblem'
                    }
                },
                {
                    path: '/getProblem',
                    component: getComponent('problem', 'daily'),
                    meta: {
                        title: 'getProblem'
                    }
                },
                {
                    path: '/problemSquare',
                    component: getComponent('problem', 'square'),
                    meta: {
                        title: 'problemSquare'
                    }
                },
                {
                    path: '/problemPlan',
                    component: getComponent('problem', 'plan'),
                    meta: {
                        title: 'problemPlan'
                    }
                }
            ]
        },
        {
            path: '*',
            redirect: '/404',
        }
    ]
})

//判断是否存在token
myRouter.beforeEach((to, from, next) => {
    document.title = 'Daily Problem'
    NProgress.start()
    if (to.path !== '/login' && !store.state.token) {
        next('/login')
    } else {
        next();
    }
    NProgress.done() // 结束Progress
})

myRouter.afterEach(() => {
    NProgress.done() // 结束Progress
})
export default myRouter