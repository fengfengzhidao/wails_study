import {fileURLToPath, URL} from 'node:url'

import {defineConfig, loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'
import type {ImportMetaEnv} from "./env";

// https://vite.dev/config/
export default defineConfig((data) => {
  const env: Record<keyof ImportMetaEnv, string> = loadEnv(data.mode, process.cwd())
  console.log(env.VITE_BASE_URL)
  const serverUrl = env.VITE_BASE_URL
  const wsUrl = serverUrl.replace("http", "ws")

  return {
    plugins: [
      vue(),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      },
    },
    server: {
      port: 80,
      host: "127.0.0.1",
      proxy: {
        "/api": {
          target: serverUrl,
          changeOrigin: true,
        },
        "/upload": {
          target: serverUrl,
          changeOrigin: true,
        },
        "/ws": {
          target: wsUrl,
          rewrite: (path=>path.replace("/ws", "")),
          changeOrigin: true,
          ws: true,
        }
      }
    }
    // envPrefix: "VITE"  // 读取环境变量的前缀
  }
})
