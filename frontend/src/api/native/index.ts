import { CacheModel } from "./cache"

import { ArtworkModel } from "./artwork"
import { ConfigModel } from "./config"
import { PassportModel } from "./passport"
import { UserModel } from "./user"

export default {
    cache: new CacheModel(),

    artwork: new ArtworkModel(),
    config: new ConfigModel(),
    passport: new PassportModel(),
    user: new UserModel(),
}
