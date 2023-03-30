import { CacheModel } from "./cache"

import { AiartModel } from "./aiart"
import { ConfigModel } from "./config"
import { PassportModel } from "./passport"

export default {
    cache: new CacheModel(),

    aiart: new AiartModel(),
    config: new ConfigModel(),
    passport: new PassportModel(),
}
