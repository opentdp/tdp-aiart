module tdp-aiart

go 1.18

require (
	// 加解密
	github.com/forgoer/openssl v1.5.0
	// 生成 UUID
	github.com/google/uuid v1.3.0
	// 服务管理
	github.com/kardianos/service v1.2.2
	// Map 转结构体
	github.com/mitchellh/mapstructure v1.5.0
	// 类型转换
	github.com/spf13/cast v1.5.0 // indirect
	// CLI 参数解析
	github.com/spf13/cobra v1.6.1
	// 配置文件读取
	github.com/spf13/viper v1.15.0
	// 日志输出
	go.uber.org/zap v1.24.0
	// 加密扩充库
	golang.org/x/crypto v0.6.0
	// 文本操作
	golang.org/x/text v0.7.0
	// 日志切割
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

// HTTP 服务

require (
	// WEB 框架
	github.com/gin-gonic/gin v1.9.0
	// JWT 认证支持
	github.com/golang-jwt/jwt/v4 v4.4.3
)

// 数据库 ORM

require (
	// SQLite 驱动
	github.com/glebarez/sqlite v1.7.0
	// MySQL 驱动
	gorm.io/driver/mysql v1.4.7
	// ORM 核心
	gorm.io/gorm v1.24.5
)

// 云厂商 SDK

// 腾讯云 SDK
require github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.606

// 域名和证书

// Docker SDK

// 间接依赖

require (
	github.com/bytedance/sonic v1.8.2 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/glebarez/go-sqlite v1.20.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.2 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.7 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/spf13/afero v1.9.4 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.10 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/arch v0.2.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	modernc.org/libc v1.22.3 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.5.0 // indirect
	modernc.org/sqlite v1.21.0 // indirect
)
