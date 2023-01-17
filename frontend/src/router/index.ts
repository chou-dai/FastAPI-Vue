import MyPage from '@/views/MyPage.vue';
import Home from '@/views/Home.vue';
import { createRouter,createWebHistory } from 'vue-router';

const routes = [
    {
        path: "/",
        name: "timeline",
        component: Home
    },
    {
        path: "/mypage",
        name: "mypage",
        component: MyPage
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router;