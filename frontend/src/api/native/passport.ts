import { HttpClient } from "@/api/basic/http"

export class PassportModel extends HttpClient {
    public login(rq: UserLogin): Promise<LoginResult> {
        return this.post("/passport/login", rq)
    }

    public register(rq: UserRegister): Promise<{ Id: number }> {
        delete rq.Password2
        return this.post("/passport/register", rq)
    }

    public profile(): Promise<UserItem> {
        return this.post("/passport/profile", {})
    }

    public update(rq: UserUpdate): Promise<null> {
        return this.post("/passport/profile/update", rq)
    }

    public summary(): Promise<UserSummary> {
        return this.post("/passport/summary", {})
    }
}

export interface UserLogin {
    Username: string
    Password: string
}

export interface LoginResult {
    Username: string
    Level: number
    AppId: string
    Email: string
    Token: string
}

export interface UserRegister {
    Username: string
    Password: string
    Password2?: string
    Email: string
}

export interface UserItem {
    Id: number
    Username: string
    AppId: string
    Email: string
    Description: string
}

export interface UserUpdate {
    Username: string
    Password: string
    Password2: string
    Email: string
    Description: string
    OldPassword: string
}

export interface UserSummary {
    Certjob: number
    Domain: number
    Keypair: number
    Machine: number
    Script: number
    Vendor: number
}