/**
 * 路由指令
 */

import { App, Directive } from "vue"

const route: Directive = {

    created(el, binding) {
        el.style.cursor = 'pointer'
        el.addEventListener('click', (evt: Event) => {
            const { value, instance } = binding
            if (typeof value == 'string' && instance) {
                instance.$router.push(value)
            }
            evt.stopPropagation()
        })
    }

}

// export default

export default (app: App) => {
    app.directive('route', route)
}
