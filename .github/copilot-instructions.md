# Project Guidelines

## Trade Interface Mapping

- 根据文档截图或参数表生成、修改接口时，`XXXDto` 对应业务请求参数，`XXXResult` 对应业务响应参数。
- 根据网站文档生成接口时，先以文档里的请求 path 和 `transCode` 确定交易前缀与编号。例如 path 为 `/tx/handle`、`transCode` 为 `2085` 时，文件与类型命名应为 `Tx2085`、`Tx2085Dto`、`NewTx2085Dto()`、`Tx2085Result`、`func (x *Yst2Ka) Tx2085(...)`。
- 接口前缀约定：`tm*` 表示会员及账户类接口，调用 `/tm/handle`；`tx*` 表示交易类接口，调用 `/tx/handle`；`tq*` 表示交易结果查询类接口，调用 `/tq/handle`。
- 网站目录页索引固定如下：会员及账户类目录为 https://prodoc.allinpay.com/doc/659/；交易类目录为 https://prodoc.allinpay.com/doc/678/；交易结果查询类目录为 https://prodoc.allinpay.com/doc/688/。
- 基于网站文档生成前，先从对应目录页定位具体接口文档，再进入该接口文档页读取请求 path、`transCode`、业务请求参数、业务响应参数；目录页只负责“找文档”，不直接决定文件前缀。
- 最终文件前缀、请求 path、交易号必须以“具体接口页”中的测试/生产地址与 `transCode` 为准，不能只按目录章节号、交易号段或页面所在目录推断。
- 目录页与最终文件前缀可能存在例外，必须按具体页面处理。例如 674【会员账户明细查询】挂在 659 目录下，但具体页面使用 `/tq/handle`、`transCode=3004`，因此应生成 `Tq3004`，不要因为目录归属生成 `Tm3004`。
- 同一页面如果覆盖多个 `transCode`，必须按页面内各小节拆成独立文件。例如 680 同页覆盖 2089/2090/2091，应分别生成 `Tx2089`、`Tx2090`、`Tx2091`。
- 当前仓库已实现页面与文档地址可按目录记忆：659 目录相关为 Tm1010->663、Tm1011->664、Tm1014->2429、Tm1020->661、Tm1022->662、Tm1023->673、Tm1025->671、Tm1026->675、Tm1027->672、Tm1029->1606、Tm1030->665、Tm1031->667、Tm1032->666、Tm1033->852、Tm1035->1884、Tm1050->668、Tm1051->669、Tm1053->1075、Tq3004->674；678 目录相关为 Tx2085->679、Tx2089/Tx2090/Tx2091->680、Tx2084->681、Tx2290->682、Tx2096->1516、Tx2294->683、Tx2295->684、Tx2094->859、Tx3010->685；688 目录相关为 Tq3001->689、Tq3002->690。
- 一个交易文件必须保持成套命名一致。例如 `Tm1029` 文件内应成套定义 `Tm1029Dto`、`NewTm1029Dto()`、`Tm1029Result` 和 `func (x *Yst2Ka) Tm1029(...)`；`Tq3002`、`Txxxxx` 同理。
- 网站文档中的“业务请求参数”表格对应 `XXXDto`，网站文档中的“业务响应参数”表格对应 `XXXResult`。
- 结构体字段类型必须对应文档表格中的字段类型，字段长度信息忽略，不要因为长度列生成额外约束。
- 字段注释优先使用表格中的“字段名称”，如“说明”列存在稳定约束或补充语义，可追加到同一行注释中。
- 必填字段必须出现在 `NewXXXDto(...)` 的参数列表中，并在构造函数返回的结构体字面量里完成赋值。
- 非必填字段不要放进 `NewXXXDto(...)`；应通过 `SetXXX()` 链式函数赋值。
- 如果 `XXXDto` 或 `XXXResult` 中存在子对象，必须把子对象单独定义为结构体，不要用 `map[string]any` 或匿名结构体代替；子对象内部也遵循必填走 `NewXXX`、非必填走 `SetXXX` 的规则。
- 仓库级共享通用属性和数据字典类型统一维护在 `yst2ka.go`，例如 `PayMode`、`MemberType`、`MemberStatus`、`CerType`、`BindType`、`MerchantType`、`OrderStatus` 等；交易文件内直接复用，不要重复定义。
- `payMode` 是通用属性，参考支付模式字典 https://prodoc.allinpay.com/doc/1132/；当交易接口包含 `payMode` 时，应优先复用 `yst2ka.go` 中的共享类型，例如 `PayMode` / `NewPayMode()`，不要按交易号生成 `Tx2085PayMode` 这类私有类型。
- `payMode` 内部字段随具体支付模式变化；应先根据支付模式字典确定所选模式的字段，再补充对应内容。
- 如果字段类型是 `JSONObject` 或 `JSON`，且文档提供了该字段的详细子表，则该字段应映射为独立结构体；但 `payMode` 这种全局复用属性除外，应优先走共享定义。
- 如果字段类型是 `JSONArray`，且文档提供了该字段的详细子表，则该字段应映射为“结构体切片”，例如 `SepDetail []Tx2085SepDetail`；“分账规则列表JSON（sepDetail）” 这类字段中的子表应定义为独立结构体，而不是字符串或 `[]map[string]any`。
- 子结构体命名应由交易号前缀加字段名组成，例如 `sepDetail` 对应 `Tx2085SepDetail`，`channelParamInfo` 对应 `Tx2085ChannelParamInfo`。
- 如果文档只写 `JSON`/`JSONObject`/`JSONArray` 但未给出任何固定字段说明，优先继续查该字段的“详细”说明页；确认确实没有稳定结构后，才考虑使用 `string` 或 `map[string]any` 作为兜底。
- 生成字段时直接使用文档中的参数名对应 JSON 标签；Go 字段名保持导出并使用驼峰命名。
- 修改现有交易文件时，先检查并修正从其他交易号复制来的残留：类型名、构造函数名、交易码、请求路径、字段、JSON 标签、注释。
- 如果同一交易存在多种响应形态或多个语义化包装函数，例如 `Tm1027`、`Tm1023` 这类查询型接口，可以保留一个基础交易函数，再补充更语义化的辅助方法，但基础交易函数仍必须按交易号命名。

## Trade Tests

- 每个公开交易函数都需要对应真实请求测试，默认放在同名的 `_test.go` 文件中，测试函数命名为 `TestYst2Ka_<TradeCode>`。
- 真实请求测试默认使用 `yst2ka_test` 包、`context.TODO()`、共享的 `client`、`cfg`、`v`、`Num(...)` 基座，不要用本地 mock 替代第三方挡板接口。
- 不要为 `cfg.PersonCode`、`cfg.EnterpriseCode`、`cfg.Phone` 这类共享基座配置额外生成 `if ... == "" { t.Skip(...) }` 检测逻辑；测试应直接依赖现有基座配置执行。
- 敏感字段在测试中发送前先按现有基座加密，例如用 `v.Encrypt(...)` 处理证件号、银行卡号等。
- 测试至少记录关键返回字段，通常包括 `respCode`、`respMsg`、`respTraceNum`，以及该接口最关键的业务字段，例如 `signNum`、`signAgreementUrl`、`openAcctStatus`。
- 如果真实请求依赖短时效文件 ID、短信验证码、支付模式参数、收款会员号或其他人工前置条件，测试中显式 `t.Skip(...)` 或保留清晰占位说明，不要伪造数据绕过真实链路；这类 `Skip` 仅用于真实前置条件，不用于共享配置判空。

## Existing Patterns

- 优先沿用仓库现有风格：构造函数返回指针；请求方法内部使用 `sonic.MarshalString` / `sonic.UnmarshalString`；通过 `x.Request(x.SetNow(ctx, now), path, code, data)` 发起调用。
- 新增或修改交易接口后，先做该接口的定向测试；如果当前不能直接联调，至少做窄范围编译校验。
