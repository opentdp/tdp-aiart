<script lang="ts">
import { Component, Vue } from "vue-facing-decorator"

import layoutStore from "@/store/layout"
import { onBeforeRouteUpdate, RouteLocationNormalized } from "vue-router"

@Component
export default class LayoutSidebar extends Vue {
    public layout = layoutStore()

    // 菜单列表

    public items: MenuItem[] = [
        {
            icon: "home",
            index: "/dashboard",
            title: "首页",
        },
        {
            icon: "server",
            index: "/aiart/text",
            title: "文生图",
        },
        {
            icon: "image",
            index: "/aiart/image",
            title: "图生图",
        },
        {
            icon: "setting",
            index: "/config/system",
            title: "参数配置",
        },
    ]

    // 侧栏控制

    public expanded = [] as string[]

    public getExpanded(to: RouteLocationNormalized): string[] {
        if (this.items.findIndex(item => item.index === to.path) >= 0) {
            return []
        } else {
            const idx = this.items.findIndex(item => {
                return item.subs && item.subs.findIndex(sub_item => sub_item.index === to.path) >= 0
            })
            if (idx == -1) {
                return []
            } else {
                return [this.items[idx].index]
            }
        }
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
    subs?: MenuItem[]
}
</script>

<template>
    <t-menu v-model:expanded="expanded" :value="$route.path" :collapsed="layout.Collapse">
        <template #logo>
            <div v-if="layout.Collapse" class="logo">
                <img src="/assets/img/icon.svg">
            </div>
            <div v-else class="logo">
                <img src="/assets/img/logo.svg">
            </div>
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
    img {
        height: 40px;
        margin-top: 8px;
        margin-left: 12px;
    }
}
</style>