<script lang="ts">
import { Ref, Component, Vue } from "vue-facing-decorator"
import { VueFlexWaterfall } from "vue-flex-waterfall"

import { NaApi } from "@/api"
import { ArtworkItem } from "@/api/native/artwork"

import sessionStore from "@/store/session"

import ArtworkUpdate from "./update.vue"

@Component({
    components: {
        ArtworkUpdate, VueFlexWaterfall,
    },
})
export default class ArtworkList extends Vue {
    public session = sessionStore()

    @Ref
    public updateModal!: ArtworkUpdate

    public created() {
        this.getArtworkList()
    }

    // 获取图片列表

    public images: ArtworkItem[] = []

    async getArtworkList() {
        const res = await NaApi.artwork.list({
            UserId: this.session.Level > 1 ? this.session.UserId : 0
        })
        this.images = res.Items
    }

}
</script>

<template>
    <t-space fixed direction="vertical">
        <VueFlexWaterfall align-content="start" col="4" col-spacing="10" :break-by-container="true"
            :break-at="{ 2330: 8, 2070: 7, 1810: 6, 1550: 5, 1290: 4, 1030: 3, 770: 2, 510: 1 }">
            <t-card v-for="item in images" :key="item.Id" theme="poster2" class="item">
                <template #default>
                    <t-image :src="'/upload/' + item.OutputFile" />
                </template>
                <template #footer>
                    <t-comment :author="item.Subject" :content="item.Prompt" />
                </template>
                <template #actions>
                    <t-button variant="text" shape="square" @click="updateModal.open(item)">
                        <t-icon :name="item.Status == 'private' ? 'browse-off' : 'browse'" />
                    </t-button>
                </template>
            </t-card>
        </VueFlexWaterfall>

        <ArtworkUpdate ref="updateModal" @submit="getArtworkList" />
    </t-space>
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