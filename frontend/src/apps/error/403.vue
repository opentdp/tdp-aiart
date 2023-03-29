<script lang="ts">
import { Component, Vue } from "vue-facing-decorator"

@Component
export default class Error403 extends Vue {
    public goBack() {
        this.$router.go(-1)
    }
}
</script>

<template>
    <div class="error-page">
        <div class="error-code">
            <span>4</span>
            <span>0</span>
            <span>3</span>
        </div>
        <div class="error-desc">
            啊哦~ 你没有权限访问该页面哦
        </div>
        <div class="error-handle">
            <t-button v-route="'/'" theme="primary" size="large">
                返回首页
            </t-button>
            <t-button class="error-btn" theme="primary" size="large" @click="goBack">
                返回上一页
            </t-button>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.error-page {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    width: 100%;
    height: 100%;
    background: var(--td-bg-color-component);
}

.error-code {
    line-height: 1;
    font-size: 250px;
    font-weight: bolder;
    span:nth-child(1) {
        color: var(--td-brand-color-6);
    }

    span:nth-child(2) {
        color: var(--td-success-color-5);
    }

    span:nth-child(3) {
        color: var(--td-error-color-6);
    }
}

.error-desc {
    font-size: 30px;
    color: var(--td-text-color-placeholder);
}

.error-handle {
    margin-top: 30px;
    padding-bottom: 200px;

    .error-btn {
        margin-left: 100px;
    }
}
</style>
