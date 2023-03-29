/**
 * 本地缓存中间层
 */

export class StorageMan {
    protected prefix: string
    protected driver: Storage

    constructor(n: string, s?: Storage) {
        this.prefix = 'tdp/' + n + '/'
        this.driver = s || sessionStorage
    }

    // 获取数据
    public get(key: unknown) {
        const skey = this.hashkey(key)
        const svalue = this.driver.getItem(skey)
        if (!svalue || !/^\{.+\}$/.test(svalue)) {
            return svalue
        }
        // 尝试反序列化
        try {
            const data = JSON.parse(svalue)
            if (0 < data.expiry && data.expiry < Date.now()) {
                this.driver.removeItem(skey)
                return null
            }
            return data.value
        } catch (e) {
            return null
        }
    }

    // 存储数据
    public set(key: unknown, value: unknown, expiry = 0) {
        if (value == null) {
            return this.remove(key)
        }
        // 数据序列化
        expiry > 0 && (expiry = Date.now() + expiry * 1000)
        const svalue = JSON.stringify({ key, value, expiry })
        // 存储到后端
        const skey = this.hashkey(key)
        try {
            this.driver.setItem(skey, svalue)
        } catch (e) {
            this.clear() // 防止容量超限
            this.driver.setItem(skey, svalue)
        }
    }

    // 删除缓存
    public remove(key: unknown) {
        const skey = this.hashkey(key)
        this.driver.removeItem(skey)
    }

    // 清理数据
    public clear(all?: boolean) {
        for (const skey in this.driver) {
            if (!skey || skey.indexOf(this.prefix)) {
                continue
            }
            if (all == true) {
                this.driver.removeItem(skey)
            } else {
                this.get(skey.replace(this.prefix, ''))
            }
        }
    }

    // 转换键名
    private hashkey(key: unknown): string {
        // 字符串作为key
        if (typeof key != 'object') {
            return this.prefix + key
        }
        // 将对象作为key
        let hash = 0
        const tmp = JSON.stringify(key) || 'null'
        for (let i = 0, l = tmp.length; i < l; i++) {
            hash = ((hash << 5) - hash) + tmp.charCodeAt(i)
            hash |= 0
        }
        return this.prefix + 'hash-' + (hash >>> 0)
    }

}
