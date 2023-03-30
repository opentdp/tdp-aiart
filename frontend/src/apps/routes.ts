import { RouteRecordRaw } from "vue-router"

import layout from "./layout/index.vue"

export const routes: RouteRecordRaw[] = [
    {
        path: "/",
        redirect: "/dashboard",
    },
    //////
    {
        path: "/",
        name: "home",
        component: layout,
        children: [
            {
                path: "/dashboard",
                name: "dashboard",
                meta: {
                    title: "系统首页",
                    login: true,
                },
                component: () => import("./dashboard/index.vue"),
            },
            {
                path: "/config/system",
                name: "config-system",
                meta: {
                    title: "系统参数",
                    admin: true,
                    login: true,
                },
                component: () => import("./config/system.vue"),
            },
            //////
            {
                path: "/aiart/text",
                name: "aiart-text",
                meta: {
                    title: "文生图",
                    login: true,
                },
                component: () => import("./aiart/text.vue"),
            },
            {
                path: "/aiart/image",
                name: "aiart-image",
                meta: {
                    title: "图生图",
                    login: true,
                },
                component: () => import("./aiart/image.vue"),
            },
            //////
            {
                path: "/passport/profile",
                name: "passport-profile",
                meta: {
                    title: "个人资料",
                    login: true,
                },
                component: () => import("./passport/profile.vue"),
            },
            //////
            {
                path: "/error/403",
                name: "error-403",
                meta: {
                    title: "没有权限",
                    login: true,
                },
                component: () => import("./error/403.vue"),
            },
            {
                path: "/error/404",
                name: "error-404",
                meta: {
                    title: "找不到页面",
                    login: true,
                },
                component: () => import("./error/404.vue"),
            },
        ],
    },
    //////
    {
        path: "/passport/login",
        name: "passport-login",
        meta: {
            title: "登录",
        },
        component: () => import("./passport/login.vue"),
    },
    {
        path: "/passport/register",
        name: "passport-register",
        meta: {
            title: "注册",
        },
        component: () => import("./passport/register.vue"),
    },
    //////
    {
        path: "/:catchAll(.*)",
        name: "NotFound",
        meta: {
            title: "找不到页面",
            login: true,
        },
        component: () => import("./error/404.vue"),
    },
]
