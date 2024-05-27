import {createRouter,createWebHashHistory} from 'vue-router'

import Main from '../components/Main.vue'
import Login from '../components/Login.vue'
import Register from '../components/Register.vue'
import Chat from '../components/Chat.vue'
import Groups from '../components/Groups.vue'
import Profile from '../components/Profile.vue'
import Group from '../components/Group.vue'

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
        {
            path: '/profile',
            name: 'profile',
            component: Profile
        },
        {
            path: '/group/:id',
            name: 'group',
            component: Group
        },
    ]
})

export default router