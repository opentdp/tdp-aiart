<script lang="ts">
import { Component, Vue } from "vue-facing-decorator"

import layoutStore from "@/store/layout"

@Component
export default class PassportLogin extends Vue {
    public layout = layoutStore()

    public created() {
        this.layout.initUIConfig().then(() => {
            this.runScript(this.layout.Analytics)
        })
    }

    public runScript(code: string) {
        const script = document.createElement('script')
        script.innerHTML = code
        document.body.appendChild(script)
        document.body.removeChild(script)
    }
}
</script>

<template>
    <router-view />
</template>
