// https://cn.vitejs.dev/config/

import { defineConfig } from "vite"

import { config as devConfig } from './scripts/vite-config-dev'
import { config as prodConfig } from './scripts/vite-config-prod'

export default defineConfig(({ command }) => {

    if (command === 'serve') {
        return { ...prodConfig, ...devConfig }
    }

    return prodConfig

})
