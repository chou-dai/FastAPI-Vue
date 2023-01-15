import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './style/reset.css';
import './style/style.css';

const app = createApp(App);
app.use(router);
app.mount('#app');