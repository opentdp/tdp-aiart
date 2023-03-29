<script lang="ts">
import { Component, Vue } from "vue-facing-decorator"

import sessionStore from "@/store/session"

@Component
export default class DashboardIndex extends Vue {
    public session = sessionStore()

    // 初始化

    public created() {
        this.session.Level = 1
    }
}
</script>

<template>
    <t-space fixed direction="vertical">
        <p>使用腾讯云 <b>AiArt</b> 接口实现的智能绘画平台</p>
    </t-space>
</template>
