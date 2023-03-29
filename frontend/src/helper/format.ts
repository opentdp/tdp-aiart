/**
 * 工具仓库
 */

export function codeTrim(s: string) {
    return s.replace(/\n {8}/g, "\n").trim()
}

export function bytesToSize(a: number, d = 2) {
    if (a == 0) {
        return "0 Bytes"
    }
    const c = 1024
    const e = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"]
    const f = Math.floor(Math.log(a) / Math.log(c))
    return parseFloat((a / Math.pow(c, f)).toFixed(d)) + " " + e[f]
}

export function dateFormat(t: string | number, fmt: string) {
    const d = new Date(t)
    const o = {
        "M+": d.getMonth() + 1,
        "d+": d.getDate(),
        "h+": d.getHours(),
        "m+": d.getMinutes(),
        "s+": d.getSeconds(),
        "q+": Math.floor((d.getMonth() + 3) / 3),
        S: d.getMilliseconds(),
    }

    const ry = fmt.match(/y+/)
    if (ry) {
        fmt = fmt.replace(ry[0], (d.getFullYear() + "").substring(4 - ry[0].length))
    }

    let k: keyof typeof o
    for (k in o) {
        const rk = fmt.match(new RegExp(k))
        if (rk) {
            const ok = o[k] + ""
            fmt = fmt.replace(rk[0], rk[0].length == 1 ? ok : ("00" + ok).substring(ok.length))
        }
    }

    return fmt
}
