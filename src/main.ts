import {createApp} from 'vue'
// 引入清除默认样式
import '@/assets/style/reset.scss'
// 引入根组件App
import App from '@/App.vue'
// 引入顶部和底部的全局组件
import Top from '@/components/top/index.vue';
import Bottom from '@/components/bottom/index.vue';
// 引入 vue-router
import router from '@/router/index.ts';

const app = createApp(App);
app.component('Top', Top);
app.component('Bottom', Bottom);
// 安装 vue-router
app.use(router);
app.mount('#app');
