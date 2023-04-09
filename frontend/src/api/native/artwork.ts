import { HttpClient } from "@/api/basic/http"

import { CreateImageRequest } from "./typings/artwork"

export class ArtworkModel extends HttpClient {
    public list(rq: ArtworkListParam): Promise<ArtworkPaged> {
        return this.post("/artwork/list", rq)
    }

    public detail(id: number): Promise<ArtworkDetail> {
        return this.post("/artwork/detail", { Id: id })
    }

    public create(query: CreateImageRequest): Promise<ArtworkCreateResult> {
        return this.post("/artwork/create", query)
    }

    public update(rq: Partial<ArtworkItem>): Promise<null> {
        return this.post("/artwork/update", rq)
    }

    public remove(id: number): Promise<null> {
        return this.post("/artwork/delete", { Id: id })
    }
}

export interface ArtworkOrig {
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

export interface ArtworkItem extends ArtworkOrig {
    Id: number
    CreatedAt: number
    UpdatedAt: number
}

export interface ArtworkListParam {
    UserId: number
}

export interface ArtworkCreateResult {
    Id: number
    OutputFile: string
}

export interface ArtworkDetail {
    Item: ArtworkItem
}

export interface ArtworkPaged {
    Items: ArtworkItem[]
}
