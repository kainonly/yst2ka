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
