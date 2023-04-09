<script lang="ts">
import { Component, Vue } from "vue-facing-decorator"

import layoutStore from "@/store/layout"
import sessionStore from "@/store/session"

import Header from "./header.vue"
import Content from "./content.vue"
import Footer from "./footer.vue"
import Sidebar from "./sidebar.vue"

@Component({
    components: { Header, Content, Footer, Sidebar }
})
export default class LayoutBottle extends Vue {
    public layout = layoutStore()
    public session = sessionStore()

    // 切换主题模式
    public themeModeChange() {
        const mode = this.layout.ThemeMode == "dark" ? "light" : "dark"
        this.layout.setThemeMode(mode)
    }
}
</script>

<template>
    <t-layout>
        <t-layout>
            <t-header>
                <t-head-menu>
                    <template #logo>
                        <div class="logo">
                            <img src="/assets/img/logo.svg">
                            <b>Aiart</b>
                        </div>
                    </template>
                    <template #operations>
                        <t-button theme="default" variant="text" @click="themeModeChange">
                            <t-icon name="chart-bubble" size="20" />
                        </t-button>
                        <t-button v-if="session.UserId" v-route="'/dashboard'" theme="default" variant="text">
                            控制台
                        </t-button>
                        <t-button v-else v-route="'/passport/login'" theme="default" variant="text">
                            登录
                        </t-button>
                    </template>
                </t-head-menu>
            </t-header>
            <t-layout class="main">
                <t-content>
                    <Content />
                </t-content>
                <t-footer class="footer">
                    <Footer />
                </t-footer>
            </t-layout>
        </t-layout>
    </t-layout>
</template>

<style lang="scss" scoped>
.main {
    padding: 16px 16px 0 16px;
    height: calc(100vh - 56px);
    overflow-y: auto;

    .footer {
        padding: 16px 0
    }
}

.logo {
    display: flex;
    align-items: end;
    justify-content: center;
    width: 100%;

    img {
        height: 40px;
        padding: 0 5px 0 16px;
    }

    b {
        font-size: 18px;
        padding-bottom: 3px;
    }
}
</style>
