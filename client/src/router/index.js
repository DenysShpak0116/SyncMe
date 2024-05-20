import {createRouter,createWebHashHistory} from 'vue-router'

import Main from '../components/Main.vue'
import Login from '../components/Login.vue'
import Register from '../components/Register.vue'
import Chat from '../components/Chat.vue'
import Groups from '../components/Groups.vue'

const router = createRouter({
    history:createWebHashHistory(),
    routes:[
        {
            path: '/',
            name: 'main',
            component: Main
        },
        {
            path: '/login',
            name: 'login',
            component: Login
        },
        {
            path: '/register',
            name: 'register',
            component: Register
        },
        {
            path: '/chat',
            name: 'chat',
            component: Chat
        },
        {
            path: '/groups',
            name: 'groups',
            component: Groups
        },
    ]
})

export default router