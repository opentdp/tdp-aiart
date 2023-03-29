import { defineStore } from "pinia"

export default defineStore("cache", {
    state: () => ({
        vendorOnce: false,
    }),
    actions: {
        async initVendorList() {
            if (this.vendorOnce) {
                return // 防抖
            }
            this.vendorOnce = true
        }
    },
    persist: {
        enabled: true,
        strategies: [
            {
                key: "tdp/cache",
                storage: sessionStorage,
            },
        ],
    },
})
