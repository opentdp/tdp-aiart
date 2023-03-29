import { defineStore } from "pinia"

import { NaApi } from "@/api"

export default defineStore("layout", {
    state: () => ({
        // 版本号
        Version: "",
        // 允许注册
        Registrable: false,
        // 侧栏折叠
        Collapse: false,
        // 主题模式
        ThemeMode: "",
    }),
    actions: {
        // 侧边栏折叠
        setCollapse(data: boolean) {
            this.Collapse = data
        },
        // 获取前端配置
        initUIConfig() {
            NaApi.config.ui().then(res => {
                this.Version = res.Version
                this.Registrable = res.Registrable == "true"
            })
        },
        // 设置主题模式
        setThemeMode(mode: "dark" | "light") {
            this.ThemeMode = mode
            this.syncThemeMode()
        },
        syncThemeMode() {
            if (this.ThemeMode == "dark") {
                document.documentElement.setAttribute('theme-mode', 'dark')
            } else {
                document.documentElement.removeAttribute('theme-mode')
            }
        }
    },
    persist: {
        enabled: true,
        strategies: [
            {
                key: "tdp/layout",
                storage: sessionStorage,
            },
        ],
    },
})
