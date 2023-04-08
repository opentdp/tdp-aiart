<script lang="ts">
import { Component, Vue } from "vue-facing-decorator"

import { NaApi } from "@/api"
import { AiartItem } from "@/api/native/aiart"

@Component
export default class DashboardIndex extends Vue {

    public created() {
        this.getImages()
    }

    // 获取图片列表

    public images: AiartItem[] = []

    async getImages() {
        const res = await NaApi.aiart.list({})
        this.images = res.Items
    }

}
</script>

<template>
    <t-space fixed direction="vertical">
        <p>使用腾讯云 <b>AiArt</b> 接口实现的智能绘画平台</p>
    </t-space>
</template>
