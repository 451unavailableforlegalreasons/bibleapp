import { createRouter, createWebHistory } from 'vue-router'
import StudyView from '../views/Study.vue'
import ChooseEdition from '../views/ChooseEdition.vue'
import Landing from '../views/Landing.vue'



import Authenticate from '../views/users/Authentication.vue'

const routes = [
    {
        path: "/",
        name: "landing",
        component: Landing
    },
    {
        path: '/study/:edition',
        name: 'study',
        component: StudyView
    },
    {
        path: '/chooseedition',
        name: 'chooseEdition',
        component: ChooseEdition
    },
    {
        path: '/users/auth',
        name: 'loginORregister',
        component: Authenticate
    },
    // {
    //     path: '/chooseedition',
    //     name: 'chooseEdition',
    //     component: ChooseEdition
    // },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
