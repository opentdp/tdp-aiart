import { defineStore } from "pinia"

import { LoginResult } from "@/api/native/passport"

export default defineStore("session", {
    state: () => ({
        Username: "",
        Level: 0,
        AppId: "",
        Email: "",
        Token: "",
    }),
    actions: {
        update(res: LoginResult) {
            Object.assign(this, res)
        }
    },
    persist: {
        enabled: true,
        strategies: [
            {
                key: "tdp/session",
                storage: localStorage,
            },
        ],
    },
})
