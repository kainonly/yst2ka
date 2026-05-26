# yst2ka

yst2ka 是一个面向通联云商通 KA 客户版接口的 Go SDK，封装了会员及账户类、交易类、交易结果查询类接口，以及文件上传、请求签名验签和常用数据字典。

## 特性

- 基于 resty v3 发起 HTTP 请求，统一封装请求发送流程。
- 内置 SM2 签名与验签，业务调用侧只需要关心 DTO/Result。
- 交易接口统一采用 NewXXXDto() + SetXXX() 的建模方式。
- 在 yst2ka.go 中集中维护通用属性和数据字典类型，例如 PayMode、MemberType、CerType、OrderStatus 等。
- 项目内置真实请求测试，以及银行卡、身份证、营业执照等测试数据生成工具。

## 环境要求

- Go 1.25+
- 可用的通联测试或生产环境应用参数
- 对应环境的商户私钥、通联公钥、AppID、SecretKey

## 安装

```bash
go get github.com/kainonly/yst2ka
```

## 快速开始

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kainonly/yst2ka"
)

func main() {
	client, err := yst2ka.NewYst2Ka(yst2ka.Option{
		BaseURL:           "https://ibstest.allinpay.com/yst/yst-service-api",
		PrivateKey:        "<your_private_key_base64>",
		AllinpayPublicKey: "<allinpay_public_key_base64>",
		AppID:             "<your_app_id>",
		SecretKey:         "<your_secret_key>",
	})
	if err != nil {
		log.Fatal(err)
	}

	// 如需排查请求问题，可打开 Trace。
	client.Debug()

	dto := yst2ka.NewTm1029Dto(
		"XPS1001-202605261200000010",
		"PS1001",
		"https://example.com/tm1029/callback",
	).SetJumpUrl("https://example.com/tm1029/return")

	resp, err := client.Tm1029(context.TODO(), dto)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.RespCode, resp.RespMsg, resp.OpenAcctUrl)
}
```

## 配置项

运行时初始化使用 yst2ka.Option：

| 字段              | 说明                                                                            |
| ----------------- | ------------------------------------------------------------------------------- |
| BaseURL           | 通联接口基础地址，例如测试环境 https://ibstest.allinpay.com/yst/yst-service-api |
| PrivateKey        | 商户 SM2 私钥，Base64 编码                                                      |
| AllinpayPublicKey | 通联 SM2 公钥，Base64 编码                                                      |
| AppID             | 应用 ID                                                                         |
| SpAppID           | 服务商应用 ID，当前仓库大多数接口未直接使用                                     |
| SecretKey         | 对称加密密钥，测试中常用于敏感字段加密                                          |

测试基座默认从 config/values.yml 读取配置，示例结构如下：

```yaml
context: dev
configs:
  dev:
    base_url: https://ibstest.allinpay.com/yst/yst-service-api
    notify_url: http://localhost:3000
    private_key_str: <your_private_key_base64>
    allinpay_public_key_str: <allinpay_public_key_base64>
    app_id: <your_app_id>
    secret_key: <your_secret_key>
    person_code: <your_person_member_code>
    enterprise_code: <your_enterprise_member_code>
    phone: <your_bound_phone>
```

可参考 [config/values.example.yml](config/values.example.yml) 组织自己的测试配置。

## 已实现能力

### 公共能力

- 文件上传：[upload.go](upload.go)
- 客户端初始化、请求签名验签、通用字典类型：[yst2ka.go](yst2ka.go)

### 会员及账户类接口

当前已实现的 Tm 接口代码：

- 1010
- 1011
- 1014
- 1020
- 1022
- 1023
- 1025
- 1026
- 1027
- 1029
- 1030
- 1031
- 1032
- 1033
- 1035
- 1050
- 1051
- 1053

其中部分查询类接口还提供了语义化辅助方法，例如：

- Tm1023：GetPlatformBalanceDetail、GetMemberBalanceDetails
- Tm1027：GetPersonInfo、GetEnterpriseInfo、GetAcctInfos、GetAgreements 等

### 交易类接口

当前已实现的 Tx 接口代码：

- 2085
- 2089
- 2090
- 2091

### 交易结果查询类接口

当前已实现的 Tq 接口代码：

- 3001
- 3002
- 3004

接口文件命名与前缀约定如下：

- tm\*：会员及账户类接口，调用 /tm/handle
- tx\*：交易类接口，调用 /tx/handle
- tq\*：交易结果查询类接口，调用 /tq/handle

## 通用字典类型

以下共享类型统一维护在 [yst2ka.go](yst2ka.go)，交易文件内直接复用：

- PayMode
- MemberType
- MemberStatus
- CerType
- BindStatus
- CardType
- SignStatus
- BindType
- MerchantType
- AccountType
- PayAccountStatus
- OrderStatus
- ChannelTradeType

其中 PayMode 参考通联支付模式字典，按具体支付模式补充所需字段；不要在单个交易文件中重复定义私有 PayMode 类型。

## 项目结构

- [yst2ka.go](yst2ka.go)：客户端、请求发送、签名验签、共享字典类型
- [upload.go](upload.go)：文件上传接口
- tm\*.go：会员及账户类接口实现
- tx\*.go：交易类接口实现
- tq\*.go：交易结果查询类接口实现
- \*\_test.go：真实请求测试
- [util](util)：测试数据生成工具，目前包含银行卡、营业执照、身份证号生成器
- [config](config)：测试配置示例

## 测试

项目测试默认以真实请求为主，不使用本地 mock 替代第三方接口。

常用命令：

```bash
# 仅做编译校验
go test ./... -run '^$'

# 执行单个真实请求测试
go test ./... -run '^TestYst2Ka_Tm1029$' -v

# 执行文件上传测试
go test ./... -run '^TestYst2Ka_FileUpload$' -v
```

注意事项：

- 大部分接口测试会直接请求通联测试环境。
- 部分测试依赖短时效 fileId、短信验证码、支付模式参数、收款会员号等人工前置条件，这类场景会在对应测试文件内保留清晰占位或 Skip 说明。
- 共享基座配置（如 person_code、enterprise_code、phone）由 config/values.yml 提供，真实请求测试默认直接依赖该配置执行。

## 维护约定

新增或修改交易接口时，建议遵循以下约定：

- XXXDto 对应业务请求参数，XXXResult 对应业务响应参数。
- 必填字段进入 NewXXXDto() 参数列表，并在构造函数内完成赋值。
- 非必填字段通过 SetXXX() 链式方法赋值。
- 如果字段类型为 JSON/JSONObject/JSONArray，且文档给出固定结构，应拆成独立结构体或结构体切片。
- 仓库级共享属性和数据字典统一放在 [yst2ka.go](yst2ka.go)，不要在交易文件里重复定义。
- 新增或修改接口后，优先补对应真实请求测试；若暂时无法联调，至少做窄范围编译校验。

更细的生成与维护规则见 [.github/copilot-instructions.md](.github/copilot-instructions.md)。

## 许可证

见 [LICENSE](LICENSE)。
