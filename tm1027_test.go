package yst2ka_test

import (
	"context"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/kainonly/yst2ka"
	"github.com/stretchr/testify/assert"
)

func TestYst2Ka_Tm1027(t *testing.T) {
	ctx := context.TODO()
	dto := yst2ka.NewTm1027Dto(`bf10006`, "1")

	var r yst2ka.Tm1027Result[yst2ka.PersonInfo]
	err := client.Tm1027(ctx, dto, &r)
	assert.NoError(t, err)

	t.Log(`respCode:`, r.RespCode)
	t.Log(`respMsg:`, r.RespMsg)
	t.Log(`signNum:`, r.SignNum)

	var content string
	content, err = sonic.MarshalString(r.MemberBasicInfo)
	assert.NoError(t, err)
	t.Log(`memberBasicInfo:`, content)
	t.Log(`acctInfo:`, r.AcctInfo)
	t.Log(`agreementArray:`, r.AgreementArray)
	t.Log(`ocrResultJson:`, r.OcrResultJson)
	t.Log(`bindPhoneJson:`, r.BindPhoneJson)
	t.Log(`payAcctOpenJson:`, r.PayAcctOpenJson)
	t.Log(`payAcctAuditJson:`, r.PayAcctAuditJson)
	t.Log(`bankSubAcctInfo:`, r.BankSubAcctInfo)
	t.Log(`settleAcctInfo:`, r.SettleAcctInfo)
	t.Log(`memberControlInfo:`, r.MemberControlInfo)
}

func TestYst2Ka_GetPersonInfo(t *testing.T) {
	ctx := context.TODO()
	info, err := client.GetPersonInfo(ctx, `bf10006`)
	assert.NoError(t, err)

	t.Log(`name:`, info.Name)
	t.Log(`cerType:`, info.CerType)
	t.Log(`cerNum:`, info.CerNum)
	t.Log(`isWithdraw:`, info.IsWithdraw)
	t.Log(`phone:`, info.Phone)
	t.Log(`idValidStartDate:`, info.IdValidStartDate)
	t.Log(`idValidEndDate:`, info.IdValidEndDate)
	t.Log(`registerTime:`, info.RegisterTime)
	t.Log(`isRealNameAuth:`, info.IsRealNameAuth)
	t.Log(`realNameAuthTime:`, info.RealNameAuthTime)
	t.Log(`memberStatus:`, info.MemberStatus)
	t.Log(`memberRole:`, info.MemberRole)
	t.Log(`memberType:`, info.MemberType)
}

func TestYst2Ka_GetEnterpriseInfo(t *testing.T) {
	ctx := context.TODO()
	info, err := client.GetEnterpriseInfo(ctx, `SUP10000`)
	assert.NoError(t, err)

	t.Log(`memberRole:`, info.MemberRole)
	t.Log(`memberType:`, info.MemberType)
	t.Log(`enterpriseName:`, info.EnterpriseName)
	t.Log(`addressCode:`, info.AddressCode)
	t.Log(`enterpriseAdress:`, info.EnterpriseAdress)
	t.Log(`enterpriseNature:`, info.EnterpriseNature)
	t.Log(`unifiedSocialCredit:`, info.UnifiedSocialCredit)
	t.Log(`busLicenseValidDate:`, info.BusLicenseValidDate)
	t.Log(`phone:`, info.Phone)
	t.Log(`legalPersonName:`, info.LegalPersonName)
	t.Log(`legalPersonCerType:`, info.LegalPersonCerType)
	t.Log(`legalPersonCerNum:`, info.LegalPersonCerNum)
	t.Log(`idValidStartDate:`, info.IdValidStartDate)
	t.Log(`idValidEndDate:`, info.IdValidEndDate)
	t.Log(`legalPersonPhone:`, info.LegalPersonPhone)
	t.Log(`memberStatus:`, info.MemberStatus)
	t.Log(`auditTime:`, info.AuditTime)
	t.Log(`isWithdraw:`, info.IsWithdraw)
	t.Log(`respMsg:`, info.RespMsg)
}

func TestYst2Ka_GetAcctInfos(t *testing.T) {
	ctx := context.TODO()

	infos, err := client.GetAcctInfos(ctx, `SUP10000`)
	assert.NoError(t, err)

	for i, info := range infos {
		t.Logf(`========== %d Start ==========`, i)
		t.Log(`bankCardNo:`, info.BankCardNo)
		t.Log(`bankAccountName:`, info.BankAccountName)
		t.Log(`bankName:`, info.BankName)
		t.Log(`bindTime:`, info.BindTime)
		t.Log(`cardType:`, info.CardType)
		t.Log(`bindStatus:`, info.BindStatus)
		t.Log(`bankReservePhone:`, info.BankReservePhone)
		t.Log(`bindType:`, info.BindType)
		t.Log(`acctAttr:`, info.AcctAttr)
		t.Log(`openBankBranchName:`, info.OpenBankBranchName)
		t.Log(`payBankNumber:`, info.PayBankNumber)
		t.Log(`openBankProvince:`, info.OpenBankProvince)
		t.Log(`openBankCity:`, info.OpenBankCity)
		t.Log(`isSpecifyAcct:`, info.IsSpecifyAcct)
		t.Logf(`========== %d End ==========`, i)
	}
}

func TestYst2Ka_GetAgreements(t *testing.T) {
	ctx := context.TODO()

	infos, err := client.GetAgreements(ctx, `SUP10000`)
	assert.NoError(t, err)

	for i, info := range infos {
		t.Logf(`========== %d Start ==========`, i)
		t.Log(`signAccount:`, info.SignAccount)
		t.Log(`agreementType:`, info.AgreementType)
		t.Log(`signResult:`, info.SignResult)
		t.Log(`agreeNo:`, info.AgreeNo)
		t.Log(`signTime:`, info.SignTime)
		t.Log(`anotherMemberInfo:`, info.AnotherMemberInfo)
		t.Logf(`========== %d End ==========`, i)
	}
}

func TestYst2Ka_GetOcrResultJson(t *testing.T) {
	ctx := context.TODO()
	info, err := client.GetOcrResultJson(ctx, `SUP10000`)
	assert.NoError(t, err)

	t.Log(`========== OcrResultJson Start ==========`)
	t.Log(`enterpriseCompareResult:`, info.EnterpriseCompareResult)
	t.Log(`legalPersonCompareResult:`, info.LegalPersonCompareResult)
	t.Log(`========== OcrResultJson End ==========`)
}

func TestYst2Ka_GetBindPhoneJson(t *testing.T) {
	ctx := context.TODO()
	info, err := client.GetBindPhoneJson(ctx, `bf10006`)
	assert.NoError(t, err)

	t.Log(`========== BindPhoneJson Start ==========`)
	t.Log(`isBind:`, info.IsBind)
	t.Log(`phone:`, info.Phone)
	t.Log(`========== BindPhoneJson End ==========`)
}

func TestYst2Ka_GetPayAcctOpenJson(t *testing.T) {
	ctx := context.TODO()
	info, err := client.GetPayAcctOpenJson(ctx, `bf10006`)
	assert.NoError(t, err)

	t.Log(`========== PayAcctOpenJson Start ==========`)
	t.Log(`cusId:`, info.CusId)
	t.Log(`payAcctNo:`, info.PayAcctNo)
	t.Log(`payAcctNoStatus:`, info.PayAcctNoStatus)
	t.Log(`openAcctTime:`, info.OpenAcctTime)
	t.Log(`========== PayAcctOpenJson End ==========`)
}

func TestYst2Ka_GetPayAcctAuditJson(t *testing.T) {
	ctx := context.TODO()
	info, err := client.GetPayAcctAuditJson(ctx, `SUP10000`)
	assert.NoError(t, err)

	t.Log(`========== PayAcctAuditJson Start ==========`)
	t.Log(`enterpriseVerifyResult:`, info.EnterpriseVerifyResult)
	t.Log(`legalIdCardVerifyResult:`, info.LegalIdCardVerifyResult)
	t.Log(`bankAcctVerifyResult:`, info.BankAcctVerifyResult)
	t.Log(`unifiedCreditPhotoResult:`, info.UnifiedCreditPhotoResult)
	t.Log(`legalCerPhotoResult:`, info.LegalCerPhotoResult)
	t.Log(`settleAcctPhotoResult:`, info.SettleAcctPhotoResult)
	t.Log(`busOutdoorPhotoResult:`, info.BusOutdoorPhotoResult)
	t.Log(`busInnerPhotoResult:`, info.BusInnerPhotoResult)
	t.Log(`acctManOutdoorPhotoResult:`, info.AcctManOutdoorPhotoResult)
	t.Log(`acctManWithIdPhotoResult:`, info.AcctManWithIdPhotoResult)
	t.Log(`busCoopConfirmResult:`, info.BusCoopConfirmResult)
	t.Log(`nonNatBenfitInfoResult:`, info.NonNatBenfitInfoResult)
	t.Log(`tlPayAcctNoAgreeResult:`, info.TlPayAcctNoAgreeResult)
	t.Log(`========== PayAcctAuditJson End ==========`)
}

func TestYst2Ka_GetbankSubAcctInfo(t *testing.T) {
	ctx := context.TODO()
	info, err := client.GetbankSubAcctInfo(ctx, `SUP10000`)
	assert.NoError(t, err)

	t.Log(`========== BankSubAcctInfo Start ==========`)
	t.Log(`bankSubAcctInfo:`, info)
	t.Log(`========== BankSubAcctInfo End ==========`)
}

func TestYst2Ka_GetSettleAcctInfo(t *testing.T) {
	ctx := context.TODO()
	info, err := client.GetSettleAcctInfo(ctx, `SUP10000`)
	assert.NoError(t, err)

	t.Log(`========== SettleAcctInfo Start ==========`)
	t.Log(`vspCusId:`, info.VspCusId)
	t.Log(`acctNo:`, info.AcctNo)
	t.Log(`status:`, info.Status)
	t.Log(`========== SettleAcctInfo End ==========`)
}

func TestYst2Ka_GetMemberControlInfo(t *testing.T) {
	ctx := context.TODO()
	info, err := client.GetMemberControlInfo(ctx, `bf10006`)
	assert.NoError(t, err)

	t.Log(`========== MemberControlInfo Start ==========`)
	t.Log(`sepOutFlag:`, info.SepOutFlag)
	t.Log(`sepInFlag:`, info.SepInFlag)
	t.Log(`memberWithdrawFlag:`, info.MemberWithdrawFlag)
	t.Log(`========== MemberControlInfo End ==========`)
}
