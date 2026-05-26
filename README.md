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

文档目录索引：会员及账户类目录为 [659](https://prodoc.allinpay.com/doc/659/)，交易类目录为 [678](https://prodoc.allinpay.com/doc/678/)，交易结果查询类目录为 [688](https://prodoc.allinpay.com/doc/688/)；生成或核对接口时，应先从目录页定位具体文档，再以具体页面里的请求 path 和 transCode 为准。

### 公共能力

- 文件上传：[upload.go](upload.go)
- 客户端初始化、请求签名验签、通用字典类型：[yst2ka.go](yst2ka.go)

### 会员及账户类接口

Tm 接口实现情况：

| 接口代码 | 接口标题                       | 文档地址                                      | 是否实现 |
| -------- | ------------------------------ | --------------------------------------------- | -------- |
| Tm1010   | 个人会员实名及绑卡（申请）     | [663](https://prodoc.allinpay.com/doc/663/)   | ✓        |
| Tm1011   | 个人会员实名及绑卡（确认）     | [664](https://prodoc.allinpay.com/doc/664/)   | ✓        |
| Tm1014   | 会员解绑银行卡                 | [2429](https://prodoc.allinpay.com/doc/2429/) | ✓        |
| Tm1020   | 企业会员实名开户               | [661](https://prodoc.allinpay.com/doc/661/)   | ✓        |
| Tm1022   | 会员资料补录                   | [662](https://prodoc.allinpay.com/doc/662/)   | ✓        |
| Tm1023   | 账户余额查询                   | [673](https://prodoc.allinpay.com/doc/673/)   | ✓        |
| Tm1025   | 企业会员支付账户开户           | [671](https://prodoc.allinpay.com/doc/671/)   | ✓        |
| Tm1026   | 平台资金查询                   | [675](https://prodoc.allinpay.com/doc/675/)   | ✓        |
| Tm1027   | 查询会员信息                   | [672](https://prodoc.allinpay.com/doc/672/)   | ✓        |
| Tm1029   | 个人支付账户开户H5             | [1606](https://prodoc.allinpay.com/doc/1606/) | ✓        |
| Tm1030   | 会员绑定手机号申请             | [665](https://prodoc.allinpay.com/doc/665/)   | ✓        |
| Tm1031   | 会员解绑手机号（原手机号）申请 | [667](https://prodoc.allinpay.com/doc/667/)   | ✓        |
| Tm1032   | 确认绑定/解绑手机号            | [666](https://prodoc.allinpay.com/doc/666/)   | ✓        |
| Tm1033   | 企业会员新增绑定对公户         | [852](https://prodoc.allinpay.com/doc/852/)   | ✓        |
| Tm1035   | 企业会员信息修改               | [1884](https://prodoc.allinpay.com/doc/1884/) | ✓        |
| Tm1050   | 会员线上协议签约申请           | [668](https://prodoc.allinpay.com/doc/668/)   | ✓        |
| Tm1051   | 线下协议文件上传               | [669](https://prodoc.allinpay.com/doc/669/)   | ✓        |
| Tm1053   | 签约协议查看H5                 | [1075](https://prodoc.allinpay.com/doc/1075/) | ✓        |

其中部分查询类接口还提供了语义化辅助方法，例如：

- Tm1023：GetPlatformBalanceDetail、GetMemberBalanceDetails
- Tm1027：GetPersonInfo、GetEnterpriseInfo、GetAcctInfos、GetAgreements 等

### 交易类接口

Tx 接口实现情况：

| 接口代码 | 接口标题                        | 文档地址                                      | 是否实现 |
| -------- | ------------------------------- | --------------------------------------------- | -------- |
| Tx2084   | 转账申请                        | [681](https://prodoc.allinpay.com/doc/681/)   | ✓        |
| Tx2085   | 消费申请                        | [679](https://prodoc.allinpay.com/doc/679/)   | ✓        |
| Tx2089   | 担保消费申请                    | [680](https://prodoc.allinpay.com/doc/680/)   | ✓        |
| Tx2090   | 单订单担保确认                  | [680](https://prodoc.allinpay.com/doc/680/)   | ✓        |
| Tx2091   | 单会员担保确认                  | [680](https://prodoc.allinpay.com/doc/680/)   | ✓        |
| Tx2094   | 储值卡订单核销                  | [859](https://prodoc.allinpay.com/doc/859/)   | ✓        |
| Tx2096   | 充值申请                        | [1516](https://prodoc.allinpay.com/doc/1516/) | ✓        |
| Tx2290   | 提现申请                        | [682](https://prodoc.allinpay.com/doc/682/)   | ✓        |
| Tx2294   | 退款申请                        | [683](https://prodoc.allinpay.com/doc/683/)   | ✓        |
| Tx2295   | 订单关闭                        | [684](https://prodoc.allinpay.com/doc/684/)   | ✓        |
| Tx3010   | 确认支付（后台+短信验证码确认） | [685](https://prodoc.allinpay.com/doc/685/)   | ✓        |

### 交易结果查询类接口

Tq 接口实现情况：

| 接口代码 | 接口标题         | 文档地址                                    | 是否实现 |
| -------- | ---------------- | ------------------------------------------- | -------- |
| Tq3001   | 订单状态查询     | [689](https://prodoc.allinpay.com/doc/689/) | ✓        |
| Tq3002   | 订单详情查询     | [690](https://prodoc.allinpay.com/doc/690/) | ✓        |
| Tq3004   | 会员账户明细查询 | [674](https://prodoc.allinpay.com/doc/674/) | ✓        |

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
