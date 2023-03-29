export { }

function errorHandler() {

    window.addEventListener('error', event => {

        event.preventDefault()
        console.warn(event.type, event.message)

    })

    window.addEventListener("rejectionhandled", event => {

        console.warn(event.reason)

    }, false)

    window.addEventListener("unhandledrejection", event => {

        event.preventDefault()
        console.warn(event.reason)

    })

}

import.meta.env.PROD && errorHandler()
