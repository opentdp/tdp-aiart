import { okMessage, errMessage } from "./basic/notify"

import NaApi from "./native"

import TcApi from "./tencent"

export { NaApi, TcApi }

export default {
    na: NaApi,
    msg: {
        ok: okMessage,
        err: errMessage,
    }
}
