import MyMemory from '@/pages/MyMemory.vue';
import Timeline from '@/pages/Timeline.vue';
import { createRouter,createWebHistory } from 'vue-router';

const routes = [
    {
        path: "/",
        name: "myMemory",
        component: MyMemory
    },
    {
        path: "/timeline",
        name: "timeline",
        component: Timeline
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router;