import { HttpClient } from "@/api/basic/http"

export class UserModel extends HttpClient {
    public list(rq: { Module?: string }): Promise<UserPaged> {
        return this.post("/user/list", rq)
    }

    public detail(name: string): Promise<UserDetail> {
        return this.post("/user/detail", { name: name })
    }

    public create(rq: UserOrig): Promise<{ Id: number }> {
        return this.post("/user/create", rq)
    }

    public update(rq: Partial<UserItem>): Promise<null> {
        return this.post("/user/update", rq)
    }

    public remove(id: number): Promise<null> {
        return this.post("/user/delete", { Id: id })
    }

    public ui(): Promise<Record<string, string>> {
        return this.post("/user/ui", {})
    }
}

export interface UserOrig {
    Username: string
    Password: string
    Level: number
    AppId: string
    AppKey: string
    Email: string
    Description: string
    ArtworkQuota: number
}

export interface UserItem extends UserOrig {
    Id: number
    CreatedAt: number
    UpdatedAt: number
}

export interface UserDetail {
    Item: UserItem
}

export interface UserPaged {
    Items: UserItem[]
}
