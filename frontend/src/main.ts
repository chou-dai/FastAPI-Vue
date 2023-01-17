import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from "pinia";
import './style/reset.css';
import 'element-plus/dist/index.css';
import './style/style.css';
import ElementPlus from 'element-plus';

const app = createApp(App);
// 外部Storeライブラリ
app.use(createPinia());
// UIライブラリ
app.use(ElementPlus)
// routerライブラリ
app.use(router);
app.mount('#app');