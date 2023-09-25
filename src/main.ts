import { createApp } from 'vue'
// 引入清除默认样式
import '@/reset.scss'
// 引入根组件App
import App from '@/App.vue'

const app = createApp(App);
app.mount('#app')
