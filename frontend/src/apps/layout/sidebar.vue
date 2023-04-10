<script lang="ts">
import { Component, Vue } from "vue-facing-decorator"
import { onBeforeRouteUpdate, RouteLocationNormalized } from "vue-router"

import layoutStore from "@/store/layout"
import sessionStore from "@/store/session"

@Component
export default class LayoutSidebar extends Vue {
    public layout = layoutStore()
    public session = sessionStore()

    // 菜单列表

    public items: MenuItem[] = []

    // 初始化

    public created() {
        this.items = this.itemFilter(menuItems)
    }

    public itemFilter(items: MenuItem[]) {
        const level = this.session.Level
        return items.filter(item => {
            if (item.level && item.level < level) {
                return false //用户组权限不足
            }
            if (item.subs) {
                item.subs = this.itemFilter(item.subs)
            }
            return true
        })
    }

    // 侧栏控制

    public expanded = [] as string[]

    public getExpanded(to: RouteLocationNormalized): string[] {
        if (this.items.findIndex(item => item.index === to.path) >= 0) {
            return []
        }
        const idx = this.items.findIndex(item => {
            return item.subs && item.subs.findIndex(sub => sub.index === to.path) >= 0
        })
        return idx == -1 ? [] : [this.items[idx].index]
    }

    public mounted() {
        this.expanded = this.getExpanded(this.$route)
        onBeforeRouteUpdate((to) => {
            const exp = new Set<string>(this.expanded)
            this.getExpanded(to).forEach(item => exp.add(item))
            this.expanded = Array.from(exp)
        })
    }
}

interface MenuItem {
    icon: string
    index: string
    title: string
    level?: number
    subs?: MenuItem[]
}

const menuItems: MenuItem[] = [
    {
        icon: "home",
        index: "/dashboard",
        title: "首页",
    },
    {
        icon: "gift",
        index: "/artwork/list",
        title: "图库",
    },
    {
        icon: "server",
        index: "/artwork/create",
        title: "文生图",
    },
    {
        icon: "image",
        index: "/artwork/create2",
        title: "图生图",
    },
    {
        icon: "user-avatar",
        index: "/user/list",
        title: "用户管理",
        level: 1,
    },
    {
        icon: "setting",
        index: "/config/list",
        title: "系统参数",
        level: 1,
    },
]
</script>

<template>
    <t-menu v-model:expanded="expanded" :value="$route.path" :collapsed="layout.Collapse">
        <template #logo>
            <a v-if="layout.Collapse" class="logo" href="/">
                <b>Aiart</b>
            </a>
            <a v-else class="logo" href="/">
                <img src="/assets/img/logo.svg">
                <b>Aiart</b>
            </a>
        </template>
        <template v-for="item in items">
            <template v-if="item.subs">
                <t-submenu :key="item.index" :value="item.index">
                    <template #title>
                        <span>{{ item.title }}</span>
                    </template>
                    <template #icon>
                        <t-icon :name="item.icon" />
                    </template>
                    <t-menu-item v-for="item2 in item.subs" :key="item2.index" :value="item2.index"
                        :to="{ path: item2.index }">
                        {{ item2.title }}
                        <template #icon>
                            <t-icon :name="item2.icon" />
                        </template>
                    </t-menu-item>
                </t-submenu>
            </template>
            <template v-else>
                <t-menu-item :key="item.index" :value="item.index" :to="{ path: item.index }">
                    <span>{{ item.title }}</span>
                    <template #icon>
                        <t-icon :name="item.icon" />
                    </template>
                </t-menu-item>
            </template>
        </template>
    </t-menu>
</template>

<style lang="scss" scoped>
.logo {
    display: flex;
    align-items: end;
    justify-content: center;
    width: 100%;

    img {
        height: 40px;
        padding-right: 5px;
    }

    b {
        font-size: 18px;
        padding-bottom: 3px;
    }
}
</style>