<script lang="ts">
import { Ref, Component, Vue } from "vue-facing-decorator"

import { FormInstanceFunctions, FormRules, SubmitContext, Data as TData } from "tdesign-vue-next"

import Api, { NaApi } from "@/api"
import * as IArtwork from "@/api/native/typings/artwork"

import * as Tencent from "./provider/tencent"

@Component
export default class ArtworkCreate extends Vue {
    public meta = Tencent

    public loading = false

    public output = ""

    // 创建表单

    @Ref
    public formRef!: FormInstanceFunctions

    public formModel: IArtwork.CreateImageRequest = {
        Action: "TextToImage",
        Subject: "",
        Prompt: "",
        NegativePrompt: "",
        Styles: ["000"],
        ResultConfig: {
            Resolution: "1024:768",
        },
        LogoAdd: 0,
        Status: "public",
    }

    public formRules: FormRules<IArtwork.CreateImageRequest> = {
        Subject: [{ required: true }],
        Prompt: [{ required: true }],
        Styles: [{ required: true }],
    }

    async formSubmit(ctx: SubmitContext<TData>) {
        if (ctx.validateResult !== true) {
            Api.msg.err("请检查表单")
            return false
        }
        this.loading = true
        const res = await NaApi.artwork.create(this.formModel).finally(() => {
            this.loading = false
        })
        this.output = "/upload/" + res.OutputFile
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
                <t-form-item name="Subject" label="作品标题">
                    <t-input v-model="formModel.Subject" placeholder="请输入标题或备注" />
                </t-form-item>
                <t-form-item name="Status" label="作品状态">
                    <t-radio-group v-model="formModel.Status">
                        <t-radio value="public">
                            全站用户可见
                        </t-radio>
                        <t-radio value="private">
                            仅自己可见
                        </t-radio>
                    </t-radio-group>
                </t-form-item>
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
