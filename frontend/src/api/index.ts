import { okMessage, errMessage } from "./basic/notify"

import NaApi from "./native"

export { NaApi }

export default {
    na: NaApi,
    msg: {
        ok: okMessage,
        err: errMessage,
    }
}
