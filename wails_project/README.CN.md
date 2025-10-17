
# Wails 模板 ArcoDesign (TS)

[English](./README.md)

## 使用方法

```
wails init -n "Your Project Name" -t https://github.com/fengfengzhidao/wails-template-arcodesign-ts
```

## 关于

- 这个项目是 [Wails](https://wails.io/) 的模板。
- 它使用了 [ArcoDesign](https://arco.design/vue/docs/start) 框架。
- 以下选项用于生成前端：
    - Vite 5
    - Vue 3
    - TypeScript
    - Vue Router
    - Pinia 用于状态管理
    - ArcoDesign UI 组件
    - Axios 用于 HTTP 请求

## 实时开发

要以实时开发模式运行，首先在项目目录中运行 `wails dev`。然后，在另一个终端中，导航到 `frontend` 目录并运行 `npm run dev`。前端开发服务器将在 http://localhost:80 上运行。您可以在浏览器中连接到此地址以访问您的应用程序。

## 构建

要构建可分发的生产模式包，请使用 `wails build`。
