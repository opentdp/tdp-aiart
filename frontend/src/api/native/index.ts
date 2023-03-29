import { CacheModel } from "./cache"

import { ConfigModel } from "./config"
import { PassportModel } from "./passport"

export default {
    cache: new CacheModel(),

    config: new ConfigModel(),
    passport: new PassportModel(),
}
