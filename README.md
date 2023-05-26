# 土豆片人工智能艺术平台

[![TDP Aiart Builder](https://github.com/open-tdp/tdp-aiart/actions/workflows/release.yml/badge.svg)](https://github.com/open-tdp/tdp-aiart/actions/workflows/release.yml)

开源智能聊天、智能绘画服务平台

- 智能绘画基于腾讯云 `Aiart` 接口实现
- 智能聊天基于OpenAI `gpt-4` `gpt-3.5` 接口实现

## 使用指引

内容较多，请参考文档 <https://docs.opentdp.org>

## 功能预览

- 在线体验开发版功能，请进入 [演示站点](https://aiart.opentdp.org)，自行注册账号后登录

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

在项目目录运行 `build.bat` 或 `./build.sh`。你还可以下载 [稳定版](https://aiart.opentdp.org/files)

### 额外参数设置

如果项目无法运行或编译，请尝试设置系统环境变量或临时环境变量

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## 微信交流群

扫码添加开发者好友（请备注 `Open TDP`，不备注可能无法通过好友申请）

![](http://docs.opentdp.org/static/weixin-qr.jpg)

## 其他

License [GPL-3.0](https://www.gnu.org/licenses/gpl-3.0.txt)

Copyright (c) 2022 - 2023 Open TDP
