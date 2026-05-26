package yst2ka

type Tx2080NotifyResult struct {
	OpenBankNo      string `json:"openBankNo,omitempty"`   // 开户银行编码
	ChnlTradeCode   string `json:"chnlTradeCode"`          // 渠道流水号
	ReqTraceNum     string `json:"reqTraceNum,omitempty"`  // 商户订单号
	RespTraceNum    string `json:"respTraceNum,omitempty"` // 通联订单号
	PayAcctNo       string `json:"payAcctNo"`              // 来款账号
	PayAcctName     string `json:"payAcctName,omitempty"`  // 来款账户名
	ReceiveAcctType string `json:"receiveAcctType"`        // 入账账户类型
	SignNum         string `json:"signNum,omitempty"`      // 入账商户会员编号
	InAcctNo        string `json:"inAcctNo,omitempty"`     // 入账账号
	InAcctName      string `json:"inAcctName,omitempty"`   // 入账会员名称
	OrderAmount     int64  `json:"orderAmount"`            // 入账金额，单位分
	TransDateTime   string `json:"transDateTime"`          // 入账时间
	Result          string `json:"result"`                 // 入账状态
	Message         string `json:"message,omitempty"`      // 入账失败说明
	Summary         string `json:"summary,omitempty"`      // 交易附言
	Fee             string `json:"fee,omitempty"`          // 手续费
	FeeCycle        string `json:"feeCycle,omitempty"`     // 收费周期
}
