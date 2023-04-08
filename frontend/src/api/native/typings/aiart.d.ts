/**
 * ImageToImage请求参数结构体
 */
export interface ImageToImageRequest extends TextToImageRequest {
  /**
    * 输入图 Base64 数据。
      算法将根据输入的图片，结合文本描述智能生成与之相关的图像。
      Base64 和 Url 必须提供一个，如果都提供以 Base64 为准。
      图片限制：单边分辨率小于2000，转成 Base64 字符串后小于 5MB。
    */
  InputImage?: string;
  /**
    * 生成自由度。
      Strength 值越小，生成图和原图越接近。取值范围0~1，不传默认为0.6。
    */
  Strength?: number;
}

/**
 * TextToImage请求参数结构体
 */
export interface TextToImageRequest {
  /**
    * 文本描述。
      算法将根据输入的文本智能生成与之相关的图像。建议详细描述画面主体、细节、场景等，文本描述越丰富，生成效果越精美。
      不能为空，推荐使用中文。最多可传512个 utf-8 字符。
    */
  Prompt: string;
  /**
    * 反向文本描述。
    */
  NegativePrompt?: string;
  /**
    * 绘画风格。
      推荐使用且只使用一种风格。不传默认使用201（日系动漫风格）。
      如果想要探索风格列表之外的风格，也可以尝试在 Prompt 中输入其他的风格描述。
    */
  Styles?: Array<string>;
  /**
    * 生成图结果的配置，包括输出图片分辨率和尺寸等。
    */
  ResultConfig?: ResultConfig;
  /**
    * 为生成结果图添加标识的开关，默认为1。
      1：添加标识。
      0：不添加标识。
      其他数值：默认按1处理。
      建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
    */
  LogoAdd?: number;
  /**
    * 标识内容设置。
      默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
    */
  LogoParam?: LogoParam;
}

/**
 * 返回结果配置
 */
export interface ResultConfig {
  /**
    * 生成图分辨率
      支持生成以下不同分辨率的图片，对应1:1方图、3:4竖图、4:3横图三种尺寸规格，不传默认为"768:768"
      取值：
      ● 768:768
      ● 768:1024
      ● 1024:768
    * 注意：此字段可能返回 null，表示取不到有效值。
    */
  Resolution?: string;
}

/**
 * logo参数
 */
export interface LogoParam {
  /**
    * 水印url
    * 注意：此字段可能返回 null，表示取不到有效值。
    */
  LogoUrl?: string;
  /**
    * 水印base64，url和base64二选一传入
    * 注意：此字段可能返回 null，表示取不到有效值。
    */
  LogoImage?: string;
  /**
    * 水印图片位于融合结果图中的坐标，将按照坐标对标识图片进行位置和大小的拉伸匹配
    * 注意：此字段可能返回 null，表示取不到有效值。
    */
  LogoRect?: LogoRect;
}

/**
 * logo坐标
 */
export interface LogoRect {
  /**
    * 左上角X坐标
    * 注意：此字段可能返回 null，表示取不到有效值。
    */
  X?: number;
  /**
    * 左上角Y坐标
    * 注意：此字段可能返回 null，表示取不到有效值。
    */
  Y?: number;
  /**
    * 方框宽度
    * 注意：此字段可能返回 null，表示取不到有效值。
    */
  Width?: number;
  /**
    * 方框高度
    * 注意：此字段可能返回 null，表示取不到有效值。
    */
  Height?: number;
}
