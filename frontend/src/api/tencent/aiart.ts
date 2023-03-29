import { TencentClient } from "./base"
import { Aiart as IAiart } from "./typings"

export class AiartModel extends TencentClient {
    protected Service = "aiart"
    protected Version = "2022-12-29"

    public textToImage(query: IAiart.TextToImageRequest): Promise<IAiart.TextToImageResponse> {
        return this.bus({ Region: "ap-guangzhou", Action: "TextToImage", Payload: query })
    }

    public imageToImage(query: IAiart.ImageToImageRequest): Promise<IAiart.ImageToImageResponse> {
        return this.bus({ Region: "ap-guangzhou", Action: "ImageToImage", Payload: query })
    }
}
