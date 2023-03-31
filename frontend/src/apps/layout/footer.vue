<script lang="ts">
import { Component, Vue } from "vue-facing-decorator"

import layoutStore from "@/store/layout"

@Component
export default class LayoutFooter extends Vue {
    public layout = layoutStore()
}
</script>

<template>
    <div class="copyright">
        <t-space :size="4" separator="•">
            <a v-if="layout.Copytext" :href="layout.Copylink" target="_blank" hover="color">
                &copy; {{ layout.Copytext }}
            </a>
            <a v-else href="https://apps.rehiy.com/open-tdp/" target="_blank" hover="color">
                Powered by Open TDP
            </a>
            <a v-if="layout.IcpCode" href="https://beian.miit.gov.cn/" target="_blank" hover="color">
                {{ layout.IcpCode }}
            </a>
            <a href="https://apps.rehiy.com/tdp-cloud/" target="_blank" hover="color">
                Aiart {{ layout.Version ? "v" + layout.Version : "" }}
            </a>
        </t-space>
    </div>
</template>

<style lang="scss" scoped>
.copyright {
    user-select: none;
    text-align: center;
    font-size: 75%;
}
</style>
