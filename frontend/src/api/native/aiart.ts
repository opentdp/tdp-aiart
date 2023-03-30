import { HttpClient } from "@/api/basic/http"

import * as IAiart from "./typings/aiart"

export class AiartModel extends HttpClient {
    public byText(query: IAiart.TextToImageRequest): Promise<IAiart.TextToImageResponse> {
        return this.post("/aiart/create", { Action: "TextToImage", Payload: query })
    }

    public byImage(query: IAiart.ImageToImageRequest): Promise<IAiart.ImageToImageResponse> {
        return this.post("/aiart/create", { Action: "ImageToImage", Payload: query })
    }
}
