import { HttpClient } from "@/api/basic/http"

export class CacheModel extends HttpClient {
    public clear() {
        this.storage.clear(true)
    }
}
