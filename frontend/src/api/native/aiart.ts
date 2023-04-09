import { HttpClient } from "@/api/basic/http"

import { CreateImageRequest } from "./typings/aiart"

export class AiartModel extends HttpClient {
    public list(rq: AiartListParam): Promise<AiartPaged> {
        return this.post("/aiart/list", rq)
    }

    public detail(id: number): Promise<AiartDetail> {
        return this.post("/aiart/detail", { Id: id })
    }

    public create(query: CreateImageRequest): Promise<{ ResultImage: string }> {
        const action = query.InputImage ? "ImageToImage" : "TextToImage"
        return this.post("/aiart/create", { Action: action, Payload: query })
    }

    public update(rq: Partial<AiartItem>): Promise<null> {
        return this.post("/aiart/update", rq)
    }

    public remove(id: number): Promise<null> {
        return this.post("/aiart/delete", { Id: id })
    }
}

export interface AiartOrig {
    UserId: string
    Subject: string
    Prompt: string
    NegativePrompt: string
    Styles: string
    Strength: number
    InputFile: string
    OutputFile: string
    Status: string
}

export interface AiartItem extends AiartOrig {
    Id: number
    CreatedAt: number
    UpdatedAt: number
}

export interface AiartListParam {
    UserId: number
}

export interface AiartDetail {
    Item: AiartItem
}

export interface AiartPaged {
    Items: AiartItem[]
}
