<script lang="ts">
import { Component, Vue } from "vue-facing-decorator"

@Component
export default class LayoutContent extends Vue {
}
</script>

<template>
    <router-view v-slot="vNode">
        <transition name="move" mode="out-in">
            <component :is="vNode.Component" />
        </transition>
    </router-view>
</template>

<style lang="scss">
.move-enter-from,
.move-leave-to {
    opacity: 0;
}

.move-enter-active,
.move-leave-active {
    transition: opacity 0.1s ease;
}
</style>
