module tdp-aiart

go 1.19

require (
	// WEB 框架
	github.com/gin-gonic/gin v1.9.1
	// JWT 认证支持
	github.com/golang-jwt/jwt/v4 v4.5.0
	// 生成 UUID
	github.com/google/uuid v1.3.0
	// 服务管理
	github.com/kardianos/service v1.2.2
	// Map 转结构体
	github.com/mitchellh/mapstructure v1.5.0
	// 辅助类库
	github.com/opentdp/go-helper v0.1.3
	// OpenAI 接口
	github.com/sashabaranov/go-openai v1.14.0
	// 类型转换
	github.com/spf13/cast v1.5.1
	// CLI 参数解析
	github.com/spf13/cobra v1.7.0
	// 配置文件读取
	github.com/spf13/viper v1.16.0
	// Go 官方模块
	golang.org/x/crypto v0.11.0
	// ORM 核心
	gorm.io/gorm v1.25.2
)

// 间接依赖包

require (
	github.com/bytedance/sonic v1.9.2 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/forgoer/openssl v1.6.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/glebarez/go-sqlite v1.21.2 // indirect
	github.com/glebarez/sqlite v1.9.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.1 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.700 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.4.0 // indirect
	golang.org/x/exp v0.0.0-20230711153332-06a737ee72cb // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.5.1 // indirect
	modernc.org/libc v1.24.1 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.6.0 // indirect
	modernc.org/sqlite v1.24.0 // indirect
)
