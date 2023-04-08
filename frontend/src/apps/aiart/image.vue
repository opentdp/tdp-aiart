<script lang="ts">
import { Ref, Component, Vue } from "vue-facing-decorator"

import { FormInstanceFunctions, FormRules, SubmitContext, Data as TData, UploadFile, RequestMethodResponse } from "tdesign-vue-next"
import { VueCropper } from "vue-cropper"

import Api, { NaApi } from "@/api"
import * as IAiart from "@/api/native/typings/aiart"

import * as Metadata from "./metadata"

@Component
export default class AiartImage extends Vue {
    public meta = Metadata

    public loading = false

    public output = ""

    // 创建表单

    @Ref
    public formRef!: FormInstanceFunctions

    public formModel: IAiart.CreateImageRequest = {
        Prompt: "",
        NegativePrompt: "",
        Styles: [],
        InputImage: "",
        Strength: 65,
        ResultConfig: {
            Resolution: "1024:768",
        },
        LogoAdd: 0,
    }

    public formRules: FormRules<IAiart.CreateImageRequest> = {
        InputImage: [{ required: true }],
    }

    async formSubmit(ctx: SubmitContext<TData>) {
        if (ctx.validateResult !== true) {
            Api.msg.err("请检查表单")
            return false
        }
        this.output = ""
        this.loading = true
        const query = { ...this.formModel }
        if (query.Strength) {
            query.Strength = query.Strength / 100
        }
        const res = await NaApi.aiart.create(this.formModel).finally(() => {
            this.loading = false
        })
        this.output = 'data:image/jpeg;base64,' + res.ResultImage
    }

    // 图片选择

    public imageOrigin = ""

    public imageSelect(file: UploadFile) {
        const reader = new FileReader()
        file.raw && reader.readAsDataURL(file.raw)
        reader.onload = () => this.imageOrigin = String(reader.result)
        const data: RequestMethodResponse = {
            status: 'success', response: {
                url: 'assets/img/avatar.jpg'
            }
        }
        return Promise.resolve(data)
    }

    public imageClear() {
        this.imageOrigin = ''
        this.formModel.InputImage = ''
    }

    // 图片裁剪

    @Ref
    public cropper: VueCropper

    public fixedNumber = [4, 3]
    public limitMinSize = [200, 150]

    public cropperPreview() {
        this.cropper.getCropData((data: string) => {
            this.formModel.InputImage = data
        })
    }

    public onResolutionChange() {
        const img = this.imageOrigin; this.imageOrigin = ''
        switch (this.formModel.ResultConfig?.Resolution) {
            case "768:768":
                this.fixedNumber = [1, 1]
                this.limitMinSize = [200, 200]
                break
            case "768:1024":
                this.fixedNumber = [3, 4]
                this.limitMinSize = [150, 200]
                break
            case "1024:768":
                this.fixedNumber = [4, 3]
                this.limitMinSize = [200, 150]
                break
        }
        setTimeout(() => this.imageOrigin = img)
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
                图生图
            </t-breadcrumb-item>
        </t-breadcrumb>

        <t-card title="绘图参数" hover-shadow header-bordered>
            <t-form ref="formRef" :data="formModel" :rules="formRules" label-width="120px" @submit="formSubmit">
                <t-form-item name="Styles" label="绘画风格">
                    <t-select v-model="formModel.Styles" :placeholder="meta.styleDesc" :max="2" multiple>
                        <t-option v-for="item in meta.imageStyles" :key="item.value" :value="item.value"
                            :label="item.label" />
                    </t-select>
                </t-form-item>
                <t-form-item name="Styles" label="输出尺寸">
                    <t-select v-model="formModel.ResultConfig!.Resolution" @change="onResolutionChange">
                        <t-option value="768:768" label="768x768" />
                        <t-option value="768:1024" label="768x1024" />
                        <t-option value="1024:768" label="1024x768" />
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
                <t-form-item name="Strength" label="生成自由度" help="值越小，生成图和原图越接近">
                    <t-slider v-model="formModel.Strength" :show-tooltip="true" :max="100" :min="0" />
                </t-form-item>
                <br>
                <t-form-item name="InputImage" label="输入图片">
                    <t-space v-if="imageOrigin">
                        <vueCropper ref="cropper" class="cropper" mode="cover" output-type="png" :img="imageOrigin"
                            :auto-crop="true" :info-true="true" :fixed="true" :fixed-number="fixedNumber"
                            :limit-min-size="limitMinSize" @real-time="cropperPreview" />
                        <img v-if="formModel.InputImage" :src="formModel.InputImage" width="200">
                    </t-space>
                    <t-upload v-else theme="custom" draggable :request-method="imageSelect">
                        <template #dragContent="params">
                            {{ params && params.dragActive ? "释放鼠标" : "点击或拖拽上传" }}
                        </template>
                    </t-upload>
                </t-form-item>
                <t-form-item>
                    <t-space>
                        <t-button theme="primary" type="submit" :loading="loading">
                            提交
                        </t-button>
                        <t-button theme="warning" @click="imageClear">
                            清除图片
                        </t-button>
                    </t-space>
                </t-form-item>
            </t-form>
        </t-card>

        <t-card v-if="output" title="生成结果" hover-shadow header-bordered>
            <t-image :src="output" />
        </t-card>
    </t-space>
</template>

<style lang="scss" scoped>
.cropper {
    width: 480px;
    height: 320px;
}
</style>