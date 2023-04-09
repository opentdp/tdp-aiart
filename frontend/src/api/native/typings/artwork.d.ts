/**
 * 请求参数结构体
 */
export interface CreateImageRequest {
    // 作品类型
    Action?: "ImageToImage" | "TextToImage"
    // 作品标题
    Subject?: string
    // 文本描述，文本描述越丰富，生成效果越精美
    Prompt?: string
    // 反向文本描述
    NegativePrompt?: string
    // 输入图 Base64 数据。与 InputUrl 不能同时为空
    InputImage?: string
    // 生成自由度。值越小，生成图和原图越接近。取值范围0~1
    Strength?: number
    // 绘画风格
    Styles?: Array<string>
    // 生成图结果的配置，包括输出图片分辨率和尺寸等
    ResultConfig?: ResultConfig
    // 为生成结果图添加标识的开关，1为添加，0为不添加
    LogoAdd?: 0 | 1
    // 标识内容设置
    LogoParam?: LogoParam
    // 作品状态
    Status?: "public" | "private"
}

/**
 * 返回结果配置
 */
export interface ResultConfig {
    // 生成图分辨率  ● 768:768  ● 768:1024 ● 1024:768
    Resolution?: string
}

/**
 * logo参数
 */
export interface LogoParam {
    // 水印url
    LogoUrl?: string
    // 水印base64，url和base64二选一传入
    LogoImage?: string
    // 水印图片位于融合结果图中的坐标，将按照坐标对标识图片进行位置和大小的拉伸匹配
    LogoRect?: LogoRect
}

/**
 * logo坐标
 */
export interface LogoRect {
    // 左上角X坐标
    X?: number
    // 左上角Y坐标
    Y?: number
    // 方框宽度
    Width?: number
    // 方框高度
    Height?: number
}
