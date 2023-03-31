<script lang="ts">
import { Component, Vue } from "vue-facing-decorator"

import layoutStore from "@/store/layout"

@Component
export default class PassportLogin extends Vue {
    public layout = layoutStore()

    public created() {
        this.layout.fetchConfig()
    }
}
</script>

<template>
    <router-view />
</template>
