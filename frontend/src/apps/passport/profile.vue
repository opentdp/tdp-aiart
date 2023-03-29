<script lang="ts">
import { Component, Ref, Vue } from "vue-facing-decorator"

import { FormInstanceFunctions, FormRules, SubmitContext, Data as TData } from "tdesign-vue-next"

import Api, { NaApi } from "@/api"
import { UserUpdate } from "@/api/native/passport"

import sessionStore from "@/store/session"

@Component
export default class PassportProfile extends Vue {
    public session = sessionStore()
    public loading = true

    // 初始化

    public created() {
        this.getProfile()
    }

    // 获取用户信息

    async getProfile() {
        const res = await NaApi.passport.profile()
        Object.assign(this.formModel, res)
        this.loading = false
    }

    // 创建表单

    @Ref
    public formRef!: FormInstanceFunctions

    public formModel: UserUpdate = {
        Username: "",
        Password: "",
        Password2: "",
        Email: "",
        Description: "",
        OldPassword: "",
    }

    public formRules: FormRules<UserUpdate> = {
        Username: [{ required: true }],
        Password2: [{ validator: val => val == this.formModel.Password, message: '两次密码不一致' }],
        Email: [{ required: true }],
        Description: [{ required: true }],
        OldPassword: [{ required: true }],
    }

    public formSubmit(ctx: SubmitContext<TData>) {
        if (ctx.validateResult !== true) {
            Api.msg.err("请检查表单")
            return false
        }
        NaApi.passport.update(this.formModel)
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
                个人资料
            </t-breadcrumb-item>
        </t-breadcrumb>

        <t-row :gutter="16">
            <t-col :span="5">
                <t-card title="基础信息" hover-shadow header-bordered>
                    <div class="info">
                        <t-avatar size="96px" image="assets/img/avatar.jpg" />
                        <div class="info-name">
                            {{ session.Username }}
                        </div>
                        <div class="info-desc">
                            {{ formModel.Description }}
                        </div>
                    </div>
                </t-card>
            </t-col>
            <t-col :span="7">
                <t-card title="账户编辑" hover-shadow header-bordered>
                    <t-form ref="formRef" :data="formModel" :rules="formRules" label-width="70px" @submit="formSubmit">
                        <t-form-item label="AppId">
                            {{ session.AppId }}
                        </t-form-item>
                        <t-form-item name="Username" label="用户名">
                            <t-input v-model="formModel.Username" />
                        </t-form-item>
                        <t-form-item name="OldPassword" label="原密码">
                            <t-input v-model="formModel.OldPassword" type="password" />
                        </t-form-item>
                        <t-form-item name="Password" label="新密码">
                            <t-input v-model="formModel.Password" type="password" />
                        </t-form-item>
                        <t-form-item name="Password2" label="确认密码">
                            <t-input v-model="formModel.Password2" type="password" />
                        </t-form-item>
                        <t-form-item name="Email" label="邮箱">
                            <t-input v-model="formModel.Email" />
                        </t-form-item>
                        <t-form-item name="Description" label="简介">
                            <t-textarea v-model="formModel.Description" :autosize="{ minRows: 3, maxRows: 15 }" />
                        </t-form-item>
                        <t-form-item>
                            <t-button theme="primary" type="submit">
                                保存
                            </t-button>
                        </t-form-item>
                    </t-form>
                </t-card>
            </t-col>
        </t-row>
    </t-space>
</template>

<style lang="scss" scoped>
.info {
    text-align: center;
    padding: 35px 0;

    .info-name {
        margin: 15px 0 10px;
        font-size: 150%;
        font-weight: 600;
    }
}
</style>
