import { HttpClient, HttpRequest } from "@/api/basic/http"

export class TencentClient extends HttpClient {
    static VendorId = ""

    protected Service = ""
    protected Version = ""

    protected bus(query: TencentRequest, expiry = 0) {
        const req: HttpRequest = {
            method: "POST",
            url: "/tencent/" + TencentClient.VendorId,
            query: Object.assign({
                Service: this.Service,
                Version: this.Version,
            }, query)
        }

        if (expiry > 0) {
            return this.rcache(req, expiry)
        }

        return this.request(req)
    }
}

export function TencentVendor(id: number | string) {
    TencentClient.VendorId = id + ""
}

export interface TencentRequest {
    Service?: string
    Version?: string
    Action: string
    Payload?: unknown
    Region?: string
    Endpoint?: string
}
