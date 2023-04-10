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
        this.getArtworkList()
    }

    // 获取图片列表

    public images: ArtworkItem[] = []

    async getArtworkList() {
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
        <t-card v-for="item in images" :key="item.Id" theme="poster2" class="item">
            <template #default>
                <t-image-viewer :images="['/upload/' + item.OutputFile]">
                    <template #trigger="{ open }">
                        <t-image :src="'/upload/' + item.OutputFile" @click="open" />
                    </template>
                </t-image-viewer>
            </template>
            <template #footer>
                <t-comment :author="item.Subject" :content="item.Prompt" />
            </template>
        </t-card>
    </VueFlexWaterfall>
</template>

<style lang="scss" scoped>
.item {
    width: 250px;
    margin-bottom: 10px;

    :deep(.t-card__body) {
        line-height: 0;
        padding: 0;
    }

    :deep(.t-image__wrapper) {
        min-height: 100px;
    }
}
</style>