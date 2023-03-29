import { MessagePlugin } from "tdesign-vue-next"

const Message: Record<string, string> = {
    UKN: "未知错误",
    EOF: "参数错误",
}

export function okMessage(data: string | { Message: string }) {
    let msg = (typeof data != "string" ? data.Message : data) || "UKN"
    msg = Message[msg] || msg
    MessagePlugin.success(msg)
    return msg
}

export function errMessage(data: string | { Message: string }) {
    let msg = (typeof data != "string" ? data.Message : data) || "UKN"
    msg = Message[msg] || msg
    MessagePlugin.error(msg)
    return msg
}
