import {
    createRouter,
    createWebHashHistory
} from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Search from '../views/Search.vue'
import Profile from '../views/Profile.vue'
import PageNotFound from '../views/PageNotFound.vue'
import Settings from '../views/Settings.vue'

const router = createRouter({
    history: createWebHashHistory(
        import.meta.env.BASE_URL),
    routes: [{
            path: '/',
            redirect: '/login'
        },
        {
            path: '/login',
            component: Login
        },
        {
            path: '/home',
            component: Home
        },
        {
            path: '/search',
            component: Search
        },
        {
            path: '/users/:id',
            component: Profile

        },
        {
            path: '/users/:id/settings',
            component: Settings

        },
        {
            path: "/:catchAll(.*)",
            component: PageNotFound
        },
    ]
})

export default router