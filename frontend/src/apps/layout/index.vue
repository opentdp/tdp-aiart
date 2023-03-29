<script lang="ts">
import { Component, Vue } from "vue-facing-decorator"

import layoutStore from "@/store/layout"

import Header from "./header.vue"
import Content from "./content.vue"
import Footer from "./footer.vue"
import Sidebar from "./sidebar.vue"

@Component({
    components: { Header, Content, Footer, Sidebar }
})
export default class LayoutIndex extends Vue {
    public layout = layoutStore()
}
</script>

<template>
    <t-layout>
        <t-aside :width="layout.Collapse ? '64px' : '232px'">
            <Sidebar />
        </t-aside>
        <t-layout>
            <t-header>
                <Header />
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
</style>
