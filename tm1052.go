package yst2ka

type Tm1052NotifyResult struct {
	ReqTraceNum              string     `json:"reqTraceNum"`                        // 商户请求流水号
	SignNum                  string     `json:"signNum"`                            // 签约方会员编号
	MemberName               string     `json:"memberName"`                         // 签约会员名称
	PayAcctSignStatus        SignStatus `json:"payAcctSignStatus,omitempty"`        // 支付账户协议签约状态
	PayAcctAgreeNo           string     `json:"payAcctAgreeNo,omitempty"`           // 支付账户协议编号
	PayAcctSignTime          string     `json:"payAcctSignTime,omitempty"`          // 支付账户协议签约时间
	PayeeAgreeStatus         SignStatus `json:"payeeAgreeStatus,omitempty"`         // 收款协议签约状态
	PayeeAgreementNo         string     `json:"payeeAgreementNo,omitempty"`         // 收款协议编号
	PayeeAgreementSignTime   string     `json:"payeeAgreementSignTime,omitempty"`   // 收款协议签约时间
	WithdrawAgreeStatus      SignStatus `json:"withdrawAgreeStatus,omitempty"`      // 企业/个人主体账户提现协议状态
	WithdrawAgreeNo          string     `json:"withdrawAgreeNo,omitempty"`          // 企业/个人主体账户提现协议签约编号
	WithdrawAgreeSignTime    string     `json:"withdrawAgreeSignTime,omitempty"`    // 企业/个人主体账户提现协议签约时间
	MembershipStatus         SignStatus `json:"membershipStatus,omitempty"`         // 会员关系证明函签约状态
	MembershipNo             string     `json:"membershipNo,omitempty"`             // 会员关系证明函编号
	MembershipSignTime       string     `json:"membershipSignTime,omitempty"`       // 会员关系证明函签约时间
	AccreditAgreeNo          string     `json:"accreditAgreeNo,omitempty"`          // 授权委托书协议编号
	AccreditAgreeStatus      SignStatus `json:"accreditAgreeStatus,omitempty"`      // 授权委托书协议状态
	AccreditAgreeSignTime    string     `json:"accreditAgreeSignTime,omitempty"`    // 授权委托协议签约时间
	SepMemBusCode            string     `json:"sepMemBusCode,omitempty"`            // 分账方会员编号
	SepMemName               string     `json:"sepMemName,omitempty"`               // 分账方签约户名
	SepAgreeStatus           SignStatus `json:"sepAgreeStatus,omitempty"`           // 分账协议状态
	SepAgreeSignTime         string     `json:"sepAgreeSignTime,omitempty"`         // 分账协议签约时间
	SepAgreeNo               string     `json:"sepAgreeNo,omitempty"`               // 分账协议编号
	SepAgreePayeeStatus      SignStatus `json:"sepAgreePayeeStatus,omitempty"`      // 分账协议-收款方签约状态
	SepAgreeSepStatus        SignStatus `json:"sepAgreeSepStatus,omitempty"`        // 分账协议-分账方签约状态
	SepWithdrawAgreeStatus   SignStatus `json:"sepWithdrawAgreeStatus,omitempty"`   // 分账方提现协议状态
	SepWithdrawAgreeNo       string     `json:"sepWithdrawAgreeNo,omitempty"`       // 分账方提现协议签约编号
	SepWithdrawSignTime      string     `json:"sepWithdrawSignTime,omitempty"`      // 分账方提现协议签约时间
	SepAccreditAgreeNo       string     `json:"sepAccreditAgreeNo,omitempty"`       // 分账方授权委托书协议编号
	SepAccreditAgreeStatus   SignStatus `json:"sepAccreditAgreeStatus,omitempty"`   // 分账方授权委托书协议状态
	SepAccreditAgreeSignTime string     `json:"sepAccreditAgreeSignTime,omitempty"` // 分账方授权委托协议签约时间
	SepMembershipStatus      SignStatus `json:"sepMembershipStatus,omitempty"`      // 分账方会员关系证明函签约状态
	SepMembershipNo          string     `json:"sepMembershipNo,omitempty"`          // 分账方会员关系证明函编号
	SepMembershipSignTime    string     `json:"SepMembershipSignTime,omitempty"`    // 分账方会员关系证明函签约时间
	CouponAgreeStatus        SignStatus `json:"couponAgreeStatus,omitempty"`        // 平台抽佣协议状态
	CouponAgreementNo        string     `json:"couponAgreementNo,omitempty"`        // 平台抽佣协议编号
	CouponAgreeTime          string     `json:"couponAgreeTime,omitempty"`          // 平台抽佣协议签约时间
	ElecAgreeStatus          SignStatus `json:"elecAgreeStatus,omitempty"`          // 收单协议状态
	ElecAgreementNo          string     `json:"elecAgreementNo,omitempty"`          // 收单协议编号
	ElecAgreeTime            string     `json:"elecAgreeTime,omitempty"`            // 收单协议签约时间
	SignErrorMsg             string     `json:"signErrorMsg,omitempty"`             // 签约失败原因
}
