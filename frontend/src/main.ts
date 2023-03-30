import "@/helper/globals"

import { createApp } from "vue"

import App from "@/app.vue"
import "@/theme/index.scss"

import setupPinia from "@/setup/pinia"

import setupRoute from "@/setup/route"
import setupRouter from "@/setup/router"

import setupTDesign from "@/setup/tdesign"

import setupCropper from "@/setup/cropper"

const app = createApp(App)

setupPinia(app)

setupRoute(app)
setupRouter(app)

setupTDesign(app)

setupCropper(app)

app.mount("#app")
