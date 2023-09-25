import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
// 引入node提供的path,获取绝对路径
import path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  base: '/dashboard',
  plugins: [vue()],
  resolve: {
    alias:{
      "@": path.resolve(__dirname, 'src')
    }
  }
})
