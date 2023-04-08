import { HttpClient } from "@/api/basic/http"

import { ImageToImageRequest, TextToImageRequest } from "./typings/aiart"

export class AiartModel extends HttpClient {
    public list(rq: { Module?: string }): Promise<AiartPaged> {
        return this.post("/Aiart/list", rq)
    }

    public detail(name: string): Promise<AiartDetail> {
        return this.post("/Aiart/detail", { name: name })
    }

    public createByText(query: TextToImageRequest): Promise<{ ResultImage: string }> {
        return this.post("/aiart/create", { Action: "TextToImage", Payload: query })
    }

    public createByImage(query: ImageToImageRequest): Promise<{ ResultImage: string }> {
        return this.post("/aiart/create", { Action: "ImageToImage", Payload: query })
    }

    public update(rq: Partial<AiartItem>): Promise<null> {
        return this.post("/Aiart/update", rq)
    }

    public remove(id: number): Promise<null> {
        return this.post("/Aiart/delete", { Id: id })
    }
}

export interface AiartOrig {
    UserId: string
    Payload: TextToImageRequest | ImageToImageRequest
    OutputFile: string
}

export interface AiartItem extends AiartOrig {
    Id: number
    CreatedAt: number
    UpdatedAt: number
}

export interface AiartDetail {
    Item: AiartItem
}

export interface AiartPaged {
    Items: AiartItem[]
}
