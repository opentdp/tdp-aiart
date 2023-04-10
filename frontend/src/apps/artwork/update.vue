<script lang="ts">
import { Ref, Component, Vue } from "vue-facing-decorator"

import { FormInstanceFunctions, FormRules, SubmitContext, Data as TData } from "tdesign-vue-next"

import Api, { NaApi } from "@/api"
import { ArtworkItem, ArtworkOrig } from "@/api/native/artwork"

@Component({
    emits: ["submit"],
    expose: ["open"],
})
export default class ArtworkUpdate extends Vue {

    // 创建表单

    @Ref
    public formRef!: FormInstanceFunctions

    public formModel!: ArtworkItem

    public formRules: FormRules<ArtworkOrig> = {
        Subject: [
            { required: true },
        ],
        Status: [
            { required: true },
        ],
    }

    // 提交表单

    async formSubmit(ctx: SubmitContext<TData>) {
        if (ctx.validateResult !== true) {
            Api.msg.err("请检查表单")
            return false
        }
        const query = {
            Id: this.formModel.Id,
            Subject: this.formModel.Subject,
            Status: this.formModel.Status,
        }
        await NaApi.artwork.update(query)
        this.close()
    }

    // 对话框管理

    public visible = false

    public close() {
        this.visible = false
        this.$emit("submit")
    }

    public open(data: ArtworkItem) {
        this.visible = true
        this.formModel = data
    }
}
</script>

<template>
    <t-dialog v-model:visible="visible" destroy-on-close header="修改信息" :footer="false" width="600px">
        <t-form ref="formRef" :data="formModel" :rules="formRules" label-width="90px" @submit="formSubmit">
            <t-form-item name="UserId" label="用户 Id">
                <t-input v-model="formModel.UserId" disabled />
            </t-form-item>
            <t-form-item name="Subject" label="标题">
                <t-input v-model="formModel.Subject" />
            </t-form-item>
            <t-form-item name="Status" label="状态">
                <t-radio-group v-model="formModel.Status">
                    <t-radio value="public" label="全站用户可见" />
                    <t-radio value="private" label="仅作者可见" />
                </t-radio-group>
            </t-form-item>
            <t-form-item>
                <t-space size="small">
                    <t-button theme="primary" type="submit">
                        提交
                    </t-button>
                    <t-button theme="default" type="reset" @click="visible = false">
                        取消
                    </t-button>
                </t-space>
            </t-form-item>
        </t-form>
    </t-dialog>
</template>