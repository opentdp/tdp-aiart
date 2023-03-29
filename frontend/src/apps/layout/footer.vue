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
        <a href="https://apps.rehiy.com/tdp-aiart/" target="_blank" hover="color">
            Powered by TDP Aiart
        </a>
        {{ layout.Version ? "v" + layout.Version : "" }}
    </div>
</template>

<style lang="scss" scoped>
.copyright {
    user-select: none;
    text-align: center;
    font-size: 75%;
}
</style>
