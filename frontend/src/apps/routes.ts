import { RouteRecordRaw } from "vue-router"

import LayoutBottle from "./layout/bottle.vue"
import LayoutConsole from "./layout/console.vue"

import HomeIndex from "./home/index.vue"

export const routes: RouteRecordRaw[] = [
    {
        path: "/",
        name: "bottle",
        component: LayoutBottle,
        children: [
            {
                path: "/",
                name: "home-index",
                meta: {
                    title: "首页",
                },
                component: HomeIndex,
            },
        ],
    },
    //////
    {
        path: "/",
        name: "console",
        component: LayoutConsole,
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
                path: "/artwork/list",
                name: "artwork-list",
                meta: {
                    title: "图库",
                    login: true,
                },
                component: () => import("./artwork/list.vue"),
            },
            {
                path: "/artwork/create",
                name: "artwork-create",
                meta: {
                    title: "文生图",
                    login: true,
                },
                component: () => import("./artwork/create.vue"),
            },
            {
                path: "/artwork/create2",
                name: "artwork-create2",
                meta: {
                    title: "图生图",
                    login: true,
                },
                component: () => import("./artwork/create2.vue"),
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
