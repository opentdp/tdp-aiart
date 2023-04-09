<script lang="ts">
import { Component, Vue } from "vue-facing-decorator"
import { VueFlexWaterfall } from "vue-flex-waterfall"

import { NaApi } from "@/api"
import { ArtworkItem } from "@/api/native/artwork"

import sessionStore from "@/store/session"

@Component({
    components: {
        VueFlexWaterfall,
    },
})
export default class HomeIndex extends Vue {
    public session = sessionStore()

    public created() {
        this.getImages()
    }

    // 获取图片列表

    public images: ArtworkItem[] = []

    async getImages() {
        const res = await NaApi.artwork.list({
            UserId: 0
        })
        this.images = res.Items
    }

}
</script>

<template>
    <VueFlexWaterfall align-content="start" col="4" col-spacing="10" :break-by-container="true"
        :break-at="{ 2330: 8, 2070: 7, 1810: 6, 1550: 5, 1290: 4, 1030: 3, 770: 2, 510: 1 }">
        <t-card v-for="item in images" :key="item.Id" class="item">
            <t-image :src="'/upload/' + item.OutputFile" />
            <template #footer>
                <div v-if="item.InputFile">
                    <b>图生图</b>
                </div>
                <div v-if="item.Prompt">
                    <b>提示词：</b>{{ item.Prompt }}
                </div>
                <div v-if="item.NegativePrompt">
                    <b>反向词：</b>{{ item.NegativePrompt }}
                </div>
            </template>
        </t-card>
    </VueFlexWaterfall>
</template>

<style lang="scss" scoped>
.item {
    width: 250px;
    margin-bottom: 10px;

    ::v-deep .t-card__body {
        line-height: 0;
        padding: 0;
    }

    ::v-deep .t-image__wrapper {
        min-height: 100px;
    }
}
</style>