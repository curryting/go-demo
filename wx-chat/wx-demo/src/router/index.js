import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/login.vue'
import Index from '../views/index.vue'
 
const routerHistory = createWebHistory()
const router = createRouter({
    history: routerHistory,
    routes: [
        {
            path: "",
            redirect: '/index'
        }, {
            path: '/login',
            name: '登录',
            component: Login
        }, {
            path: '/index',
            name: 'index',
            component: Index 
        },
        {
            path: '/message-view',
            name: '消息',
            component: "../views/message-view.vue" 
        },
        {
            path: '/mail-list',
            name: '通讯录',
            component: "../views/mail-list.vue" 
        },
        {
            path: '/my',
            name: '我的',
            component: "../views/my.vue" 
        },
        {
            path: '/chat',
            name: '聊天',
            component: "../views/chat.vue" 
        }
    ]
})

// 设置路由拦截
// router.beforeEach((to, from, next) => {
    // if (to.name === null) return
    // // let name = Cookies.get('name') || store.state.name
    // // 如果cookie没有过期或者store中有name值,则允许访问直接通过。否则就让用户登录
    // if (name) {
    // // store.commit('loginIn', name)
    // next()
    // } else {
    // if (to.path == '/login') {
    //  next()
    // } else {
    //  next({
    //  name: 'Login'
    //  })
    // //  store.commit('loginOut')
    // }
    // }
// })
    
   router.afterEach(() => {})
 
export default router