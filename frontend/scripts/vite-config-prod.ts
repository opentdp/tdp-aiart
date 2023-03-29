// https://cn.vitejs.dev/config/

import { UserConfig } from "vite"

import vue from "@vitejs/plugin-vue"
import replace from "./vite-plugin-replace"

export const config: UserConfig = {
    base: "./",
    build: {
        outDir: "../factory/front",
        emptyOutDir: true,
        manifest: false,
    },
    plugins: [
        vue(),
        replace([
            {
                file: /tdesign-.+\.js/,
                source: /https:\/\/[^"']+\/fonts\/index\.js/g,
                target: 'assets/icon.js',
            },
        ])
    ],
    resolve: {
        alias: {
            "@": "/src",
            "~": "/node_modules",
        },
    }
}
