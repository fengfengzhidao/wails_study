import { createRouter, createWebHashHistory } from "vue-router";

export const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: "/",
            name: "index",
            component: ()=>import("../views/index.vue"),
        },
        {
            path: "/home",
            name: "home",
            component: ()=>import("../views/home.vue"),
        },
        {
            path: "/about",
            name: "about",
            component: ()=>import("../views/about.vue"),
        },
        {
            path: "/settings",
            name: "settings",
            component: ()=>import("../views/settings.vue"),
        }
    ],
});