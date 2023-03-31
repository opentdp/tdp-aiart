import { defineStore } from "pinia"

import { NaApi } from "@/api"
import * as vars from "@/helper/const"

export default defineStore("layout", {
    state: () => ({
        // 侧栏折叠
        Collapse: false,
        // 主题模式
        ThemeMode: "",
        // 后端本号
        Version: "",
        // 允许注册
        Registrable: false,
        // 前端配置
        SiteName: "",
        SiteLogo: "",
        SiteIcon: "",
        Analytics: "",
        Copylink: "",
        Copytext: "",
        IcpCode: "",
    }),
    actions: {
        // 侧边栏折叠
        setCollapse(data: boolean) {
            this.Collapse = data
        },
        // 获取前端配置
        async initUIConfig() {
            const res = await NaApi.config.ui()
            Object.keys(res).forEach(k => {
                Object.assign(res, { [k]: res[k].trim() })
            })
            this.Registrable = res.Registrable == "true"
            this.SiteName = res.SiteName || vars.SiteName
            this.Analytics = res.Analytics || vars.Analytics
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
