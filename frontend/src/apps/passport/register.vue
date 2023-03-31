<script lang="ts">
import { Ref, Component, Vue } from "vue-facing-decorator"

import { FormInstanceFunctions, FormRules, SubmitContext, Data as TData } from "tdesign-vue-next"

import Api, { NaApi } from "@/api"
import { UserRegister } from "@/api/native/passport"

import layoutStore from "@/store/layout"
import sessionStore from "@/store/session"

@Component
export default class PassportRegister extends Vue {
    public layout = layoutStore()
    public session = sessionStore()

    // 创建表单

    @Ref
    public formRef!: FormInstanceFunctions

    public formModel: UserRegister = {
        Username: "",
        Password: "",
        Password2: "",
        Email: "",
    }

    public formRules: FormRules<UserRegister> = {
        Username: [{ required: true }],
        Password: [{ required: true }],
        Password2: [
            { required: true },
            { validator: val => val == this.formModel.Password, message: '两次密码不一致' }
        ],
        Email: [{ required: true }],
    }

    async formSubmit(ctx: SubmitContext<TData>) {
        if (ctx.validateResult !== true) {
            Api.msg.err("请检查表单")
            return false
        }
        await NaApi.passport.register(this.formModel)
        // 切换到登陆页面
        this.$router.push("/passport/login")
    }
}
</script>

<template>
    <t-layout class="layout">
        <t-content class="content">
            <t-space fixed direction="vertical" class="magic-box">
                <div class="magic-title">
                    登录到<br>{{ layout.SiteName }} 控制面板
                </div>
                <t-form ref="formRef" :data="formModel" :rules="formRules" label-width="0px" class="magic-body"
                    @submit="formSubmit">
                    <t-form-item name="Username" :required-mark="false">
                        <t-input v-model="formModel.Username" size="large" placeholder="用户名">
                            <template #prefix-icon>
                                <t-icon name="user" />
                            </template>
                        </t-input>
                    </t-form-item>
                    <t-form-item name="Email" :required-mark="false">
                        <t-input v-model="formModel.Email" size="large" placeholder="邮箱">
                            <template #prefix-icon>
                                <t-icon name="mail" />
                            </template>
                        </t-input>
                    </t-form-item>
                    <t-form-item name="Password" :required-mark="false">
                        <t-input v-model="formModel.Password" type="password" size="large" placeholder="密码">
                            <template #prefix-icon>
                                <t-icon name="lock-on" />
                            </template>
                        </t-input>
                    </t-form-item>
                    <t-form-item name="Password2" :required-mark="false">
                        <t-input v-model="formModel.Password2" type="password" size="large" placeholder="确认密码">
                            <template #prefix-icon>
                                <t-icon name="lock-on" />
                            </template>
                        </t-input>
                    </t-form-item>
                    <t-form-item>
                        <t-button theme="primary" size="large" type="submit" block>
                            注册
                        </t-button>
                    </t-form-item>
                </t-form>
                <div class="magic-guide">
                    已有账号？
                    <t-link v-route="'/passport/login'" theme="default">
                        登录
                    </t-link>
                </div>
            </t-space>
        </t-content>
        <t-footer class="copyright">
            <t-space :size="4" separator="•">
                <a v-if="layout.Copytext" :href="layout.Copylink" target="_blank" hover="color">
                    &copy; {{ layout.Copytext }}
                </a>
                <a v-else href="https://apps.rehiy.com/open-tdp/" target="_blank" hover="color">
                    Powered by Open TDP
                </a>
                <a v-if="layout.IcpCode" href="https://beian.miit.gov.cn/" target="_blank" hover="color">
                    {{ layout.IcpCode }}
                </a>
                <a href="https://apps.rehiy.com/tdp-cloud/" target="_blank" hover="color">
                    Aiart {{ layout.Version ? "v" + layout.Version : "" }}
                </a>
            </t-space>
        </t-footer>
    </t-layout>
</template>

<style lang="scss" scoped>
.layout {
    width: 100%;
    height: 100%;
    background-image: url(/assets/img/bg01.jpg);
    background-position: center center;
    background-size: cover;

    .content {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100%;
    }
}

.magic-box {
    width: 360px;
    overflow: hidden;

    .magic-title {
        font-size: 200%;
        line-height: 1.5;
        margin-bottom: 20px;
    }

    .magic-guide {
        text-align: right;
    }
}

.copyright {
    user-select: none;
    text-align: center;
    font-size: 75%;
    color: var(--td-gray-color-1);
}
</style>
