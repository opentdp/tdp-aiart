<script lang="ts">
import { Component, Vue } from "vue-facing-decorator"

import { NaApi } from "@/api"
import { UserSummary } from "@/api/native/passport"

import sessionStore from "@/store/session"

@Component
export default class DashboardIndex extends Vue {
    public session = sessionStore()

    // 初始化

    public created() {
        this.getUserSummary()
    }

    // 资源统计

    public summary!: UserSummary

    async getUserSummary() {
        const res = await NaApi.passport.summary()
        this.summary = res
    }
}
</script>

<template>
    <t-card :loading="!summary" hover-shadow header-bordered>
        <t-space fixed direction="vertical">
            <p>欢迎使用智能绘画平台，画图能力基于腾讯云<b>Aiart</b>接口实现</p>
            <t-space v-if="summary" fixed break-line>
                <t-card v-route="'/artwork/list'" class="summary" hover-shadow>
                    <div>图库</div>
                    <b>{{ summary.Artwork }}</b>
                </t-card>
                <t-card class="summary" hover-shadow>
                    <div>配额</div>
                    <b>{{ summary.ArtworkQuota }}</b>
                </t-card>
            </t-space>
        </t-space>
    </t-card>
</template>

<style lang="scss" scoped>
.summary {
    min-width: 150px;
    text-align: center;
    background-color: var(--el-color-info-light-9);
    color: var(--el-color-info-light-3);

    div {
        margin-bottom: 10px;
    }

    b {
        color: var(--el-color-primary-dark-2);
        font-size: 2em;
    }
}
</style>
