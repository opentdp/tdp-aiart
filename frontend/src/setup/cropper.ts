/**
 * 图片裁切插件
 * https://github.com/xyxiao001/vue-cropper
 */

import { App } from "vue"

import VueCropper from "vue-cropper"
import "vue-cropper/dist/index.css"

// export default

export default (app: App) => {
    app.use(VueCropper)
}
