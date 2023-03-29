# 土豆片智能绘画

[![TDP Aiart Builder](https://github.com/open-tdp/tdp-aiart/actions/workflows/release.yml/badge.svg)](https://github.com/open-tdp/tdp-aiart/actions/workflows/release.yml)

使用腾讯云 `Ai Art` 接口实现的智能绘画平台

## 开发说明

### 启动开发服务

在项目目录运行  `serve.bat` 或 `./serve.sh`

### 提交代码时请使用下面标识

- `feat` 新功能（feature）
- `fix` 错误修复
- `docs` 文档更改（documentation）
- `style` 格式（不影响代码含义的更改，空格、格式、缺少分号等）
- `refactor` 重构（即不是新功能，也不是修补bug的代码变动）
- `perf` 优化（提高性能的代码更改）
- `test` 测试（添加缺失的测试或更正现有测试）
- `chore` 构建过程或辅助工具的变动
- `revert` 还原以前的提交

### 编译为二进制

在项目目录运行 `build.bat` 或 `./build.sh`

### 额外参数设置

如果项目无法运行或编译，请尝试设置系统环境变量或临时环境变量

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## 开发工具配置

项目编辑器推荐使用 `vscode`，并设置代码格式化插件如下

```json
{
    "[css]": {
        "editor.defaultFormatter": "vscode.css-language-features"
    },
    "[scss]": {
        "editor.defaultFormatter": "vscode.css-language-features"
    },
    "[html]": {
        "editor.defaultFormatter": "vscode.html-language-features"
    },
    "[json]": {
        "editor.defaultFormatter": "vscode.json-language-features"
    },
    "[jsonc]": {
        "editor.defaultFormatter": "vscode.json-language-features"
    },
    "[javascript]": {
        "editor.defaultFormatter": "vscode.typescript-language-features"
    },
    "[typescript]": {
        "editor.defaultFormatter": "vscode.typescript-language-features"
    },
    "[vue]": {
        "editor.defaultFormatter": "Vue.volar"
    }
}
```

## 微信交流群

扫码添加开发者好友（请备注 `Open TDP`，不备注可能无法通过好友申请）

![](./docs/weixin-qr.jpg)

## 其他

License [GPL-3.0](https://www.gnu.org/licenses/gpl-3.0.txt)

Copyright (c) 2022 - 2023 Open TDP
