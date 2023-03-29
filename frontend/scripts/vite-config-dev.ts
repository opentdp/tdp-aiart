// https://cn.vitejs.dev/config/

import { UserConfig } from "vite"

export const config: UserConfig = {
    server: {
        port: 7900,
        open: false,
        proxy: {
            "/api": {
                target: "http://127.0.0.1:7800",
            },
        },
    },
}
