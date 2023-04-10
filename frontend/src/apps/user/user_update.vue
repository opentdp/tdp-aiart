<script lang="ts">
import { Ref, Component, Vue } from "vue-facing-decorator"

import { FormInstanceFunctions, FormRules, SubmitContext, Data as TData } from "tdesign-vue-next"

import Api, { NaApi } from "@/api"
import { UserOrig, UserItem } from "@/api/native/user"

@Component({
    emits: ["submit"],
    expose: ["open"],
})
export default class UserUpdate extends Vue {

    // 创建表单

    @Ref
    public formRef!: FormInstanceFunctions

    public formModel!: UserOrig

    public formRules: FormRules<UserOrig> = {
        Email: [
            { required: true },
        ],
        Level: [
            { required: true },
            { max: 9 },
        ],
        ArtworkQuota: [
            { required: true },
            { min: 1 },
            { max: 9999 },
        ],
    }

    // 提交表单

    async formSubmit(ctx: SubmitContext<TData>) {
        if (ctx.validateResult !== true) {
            Api.msg.err("请检查表单")
            return false
        }
        await NaApi.user.update(this.formModel)
        this.close()
    }

    // 对话框管理

    public visible = false

    public close() {
        this.visible = false
        this.$emit("submit")
    }

    public open(data: UserItem) {
        this.visible = true
        this.formModel = data
    }
}
</script>

<template>
    <t-dialog v-model:visible="visible" destroy-on-close header="修改信息" :footer="false" width="600px">
        <t-form ref="formRef" :data="formModel" :rules="formRules" label-width="90px" @submit="formSubmit">
            <t-form-item name="Username" label="用户名">
                <t-input v-model="formModel.Username" readonly />
            </t-form-item>
            <t-form-item name="Email" label="邮箱">
                <t-input v-model="formModel.Email" />
            </t-form-item>
            <t-form-item name="Level" label="用户组">
                <t-select v-model="formModel.Level">
                    <t-option :value="1" label="创始人" />
                    <t-option :value="5" label="注册用户" />
                </t-select>
            </t-form-item>
            <t-form-item name="ArtworkQuota" label="绘图配额">
                <t-input-number v-model="formModel.ArtworkQuota" />
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