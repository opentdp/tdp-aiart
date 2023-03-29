// 绘画风格

export const textStyles = [
    { value: '000', label: '不限定风格' },
    { value: '101', label: '水墨画' },
    { value: '102', label: '概念艺术' },
    { value: '103', label: '油画' },
    { value: '104', label: '水彩画' },
    { value: '106', label: '厚涂风格' },
    { value: '107', label: '插图' },
    { value: '108', label: '剪纸风格' },
    { value: '109', label: '印象派' },
    { value: '110', label: '2.5D人像' },
    { value: '111', label: '肖像画' },
    { value: '112', label: '黑白素描画' },
    { value: '113', label: '赛博朋克' },
    { value: '114', label: '科幻风格' },
    { value: '115', label: '黑暗艺术' },
    { value: '201', label: '日系动漫' },
    { value: '202', label: '怪兽风格' },
    { value: '301', label: '游戏卡通手绘' },
]

export const imageStyles = [
    { value: '000', label: '不限定' },
    { value: '106', label: '水彩画' },
    { value: '201', label: '日系动漫' },
    { value: '202', label: '美系动漫' },
    { value: '203', label: '唯美古风' },
]

export const styleDesc = `推荐使用且只使用一种风格，不指定则默认为日系动漫风格`

export const promptDesc = `建议详细描述画面主体、细节、场景等，文本描述越丰富，生成效果越精美，推荐使用中文`

export const negativePromptDesc = `用于减少生成结果中出现描述内容的可能，但不能完全杜绝，推荐使用中文`
