/**
 * 中后台组件
 * https://tdesign.tencent.com/
 */

import { App } from "vue"

import TDesign from "tdesign-vue-next"
import "tdesign-vue-next/es/style/index.css"

// export default

export default (app: App) => {
    app.use(TDesign)
}
