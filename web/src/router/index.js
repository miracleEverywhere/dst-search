import {createRouter, createWebHashHistory} from "vue-router";
import Home from "../views/home.vue"
import Result from "../views/result.vue"

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        { path: '/', component: Home },
        { path: '/result', component: Result },
    ]
})

export default router