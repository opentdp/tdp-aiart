import { HttpClient } from "@/api/basic/http"

export class ConfigModel extends HttpClient {
    public list(rq: { Module?: string }): Promise<ConfigPaged> {
        return this.post("/config/list", rq)
    }

    public detail(name: string): Promise<ConfigDetail> {
        return this.post("/config/detail", { name: name })
    }

    public create(rq: ConfigOrig): Promise<{ Id: number }> {
        return this.post("/config/create", rq)
    }

    public update(rq: Partial<ConfigItem>): Promise<null> {
        return this.post("/config/update", rq)
    }

    public remove(id: number): Promise<null> {
        return this.post("/config/delete", { Id: id })
    }

    public ui(): Promise<Record<string, string>> {
        return this.post("/config/ui", {})
    }
}

export interface ConfigOrig {
    Name: string
    Value: string
    Type: "bool" | "string" | "text" | "upload"
    Module: string
    Description: string
}

export interface ConfigItem extends ConfigOrig {
    Id: number
    CreatedAt: number
    UpdatedAt: number
}

export interface ConfigDetail {
    Item: ConfigItem
}

export interface ConfigPaged {
    Items: ConfigItem[]
}
