<script lang="ts">
import { Ref, Component, Vue } from "vue-facing-decorator"

import { FormInstanceFunctions, FormRules, SubmitContext, Data as TData } from "tdesign-vue-next"

import Api, { NaApi } from "@/api"
import * as IAiart from "@/api/native/typings/aiart"

import * as Metadata from "./metadata"

@Component
export default class AiartText extends Vue {
    public meta = Metadata

    public loading = false

    public output = ""

    // 创建表单

    @Ref
    public formRef!: FormInstanceFunctions

    public formModel: IAiart.CreateImageRequest = {
        Prompt: "",
        NegativePrompt: "",
        Styles: ["000"],
        ResultConfig: {
            Resolution: "1024:768",
        },
        LogoAdd: 0,
    }

    public formRules: FormRules<IAiart.CreateImageRequest> = {
        Prompt: [{ required: true }],
        Styles: [{ required: true }],
    }

    async formSubmit(ctx: SubmitContext<TData>) {
        if (ctx.validateResult !== true) {
            Api.msg.err("请检查表单")
            return false
        }
        this.loading = true
        const res = await NaApi.aiart.create(this.formModel).finally(() => {
            this.loading = false
        })
        this.output = 'data:image/jpeg;base64,' + res.ResultImage
    }
}
</script>

<template>
    <t-space fixed direction="vertical">
        <t-breadcrumb>
            <t-breadcrumb-item to="/">
                首页
            </t-breadcrumb-item>
            <t-breadcrumb-item>
                文生图
            </t-breadcrumb-item>
        </t-breadcrumb>

        <t-card title="绘图参数" hover-shadow header-bordered>
            <t-form ref="formRef" :data="formModel" :rules="formRules" label-width="90px" @submit="formSubmit">
                <t-form-item name="Styles" label="绘画风格">
                    <t-select v-model="formModel.Styles" :placeholder="meta.styleDesc" :max="3" multiple>
                        <t-option v-for="item in meta.textStyles" :key="item.value" :value="item.value"
                            :label="item.label" />
                    </t-select>
                </t-form-item>
                <t-form-item name="Prompt" label="文本描述">
                    <t-textarea v-model="formModel.Prompt" :autosize="{ minRows: 3, maxRows: 15 }" :maxlength="512"
                        :placeholder="meta.promptDesc" />
                </t-form-item>
                <t-form-item name="NegativePrompt" label="反向描述">
                    <t-textarea v-model="formModel.NegativePrompt" :autosize="{ minRows: 3, maxRows: 15 }" :maxlength="512"
                        :placeholder="meta.negativePromptDesc" />
                </t-form-item>
                <t-form-item>
                    <t-button theme="primary" type="submit" :loading="loading">
                        提交
                    </t-button>
                </t-form-item>
            </t-form>
        </t-card>

        <t-card v-if="output" title="生成结果" hover-shadow header-bordered>
            <t-image :src="output" />
        </t-card>
    </t-space>
</template>
