/**
 * 字符串替换插件
 * @author: rehiy
 * @date: 2023/03/13 06:20
 * @site: https://www.rehiy.com
 *
 * @param {ReplaceOptions} options
 */

import type { PluginOption } from 'vite'

type ReplaceOptions = {
    file: RegExp
    source: RegExp | string
    target: string | number | boolean
}[]

export default function (options: ReplaceOptions): PluginOption {
    return {
        name: 'path-load-file',
        transform(code, id) {
            for (const replaceConfig of options) {
                if (!replaceConfig.file || replaceConfig.file?.test(id)) {
                    code = code.replace(replaceConfig.source, `${replaceConfig.target}`)
                }
            }
            return code
        },
    }
}
