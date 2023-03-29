/**
 * 代码高亮组件
 * https://prismjs.com
 */

import type { App, Plugin } from "vue"

import Prism from "prismjs"
import "prismjs/themes/prism-okaidia.css"

import "prismjs/components/prism-bash"
import "prismjs/components/prism-batch"
import "prismjs/components/prism-json"
import "prismjs/components/prism-powershell"

import 'prismjs/plugins/toolbar/prism-toolbar'
import "prismjs/plugins/toolbar/prism-toolbar.css"
import "prismjs/plugins/show-language/prism-show-language"
import "prismjs/plugins/copy-to-clipboard/prism-copy-to-clipboard"

import 'prismjs/plugins/line-numbers/prism-line-numbers'
import "prismjs/plugins/line-numbers/prism-line-numbers.css"

import "prismjs/plugins/normalize-whitespace/prism-normalize-whitespace"

const plugin: Plugin = {
    install: (app: App) => {
        app.directive('highlight', pre => {
            pre.querySelectorAll('code').forEach((code: Element) => {
                Prism.highlightElement(code)
            })
        })
    },
}

// export default

export default (app: App) => {
    app.use(plugin)
}
