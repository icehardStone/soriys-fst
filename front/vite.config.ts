import { fileURLToPath, URL } from 'url'

import { defineConfig } from 'vite'
import { createProxyMiddleware } from 'http-proxy-middleware';
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), vueJsx()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    host: "0.0.0.0",
    proxy: {
      '/api': {
        "target": "http://43.134.115.113:44300",
        // "target": "http://localhost:8080",
        "changeOrigin": false,
      },
      '/.well-known': {
        target: 'https://43.134.115.113:5000',
        changeOrigin: true,
        secure: false,
        configure: (proxy, options) => {
          // proxy will be an instance of 'http-proxy'
          // proxy..headers['Access-Control-Allow-Origin'] = '*';
          // options.headers['Access-Control-Allow-Origin'] = '*';
        },
        agent: new (require('https').Agent as any)({
          rejectUnauthorized: false, // 忽略证书验证错误
          // secureProtocol: 'TLSv1_2_method', // 指定TLS版本
        }),
      },
    }
  },
  define: {
    'process.env.ESLINT_CONFIG_PRETTIER_NO_DEPRECATED': JSON.stringify(process.env.ESLINT_CONFIG_PRETTIER_NO_DEPRECATED || 'false')
  },
  build: {
    rollupOptions: {
      input: {
        main: resolve(__dirname, 'index.html'),
        nested: resolve(__dirname, 'signin-oidc.html')
      }
    }
  }
})
