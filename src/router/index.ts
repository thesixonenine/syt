import {createRouter, createWebHistory} from 'vue-router'
// 创建路由器实例
export default createRouter({
    // 路由模式
    history: createWebHistory(),
    // 管理路由
    routes: [
        {
            path: '/home',
            component: () => import('@/pages/home/index.vue'),
        },
        {
            path: '/hospital',
            component: () => import('@/pages/hospital/index.vue'),
            children: [
                {
                    path: 'reg',
                    component: () => import('@/pages/hospital/reg/index.vue'),
                },
                {
                    path: 'detail',
                    component: () => import('@/pages/hospital/detail/index.vue'),
                },
                {
                    path: 'notify',
                    component: () => import('@/pages/hospital/notify/index.vue'),
                },
                {
                    path: 'stop',
                    component: () => import('@/pages/hospital/stop/index.vue'),
                },
                {
                    path: 'cancel',
                    component: () => import('@/pages/hospital/cancel/index.vue'),
                }
            ],
        },
        {
            path: '/sr',
            component: () => import('@/pages/sr/index.vue'),
        },
        {
            path: '/',
            redirect: '/home',
        }
    ],
    // 控制滚动条的位置
    scrollBehavior() {
        return {
            left: 0,
            top: 0
        }
    }
})