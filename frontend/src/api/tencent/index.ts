import { TencentVendor } from "./base"

import { AiartModel } from "./aiart"
import { CamModel } from "./cam"

export default {
    vendor: TencentVendor,

    aiart: new AiartModel(),
    cam: new CamModel(),
}
