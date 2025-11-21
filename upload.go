package yst2ka

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/bytedance/sonic"
	"github.com/kainonly/go/help"
	"resty.dev/v3"
)

type FileUploadOption struct {
	Name  string
	Type  string
	Bytes []byte
}

type FileUploadResult struct {
	FileId   string `json:"fileId"`   // 文件标识
	RespCode string `json:"respCode"` // 业务返回码
	RespMsg  string `json:"respMsg"`  // 返回说明
}

func (x *Yst2Ka) FileUpload(ctx context.Context, opt FileUploadOption) (_ *FileUploadResult, err error) {
	hash := md5.Sum(opt.Bytes)
	data := map[string]string{
		`appId`:     x.Option.AppID,
		`fileType`:  opt.Type,
		`md5`:       base64.StdEncoding.EncodeToString(hash[:]),
		`timestamp`: time.Now().Format("2006-01-02 15:04:05"),
	}

	if data["sign"], err = help.Sm2Sign(x.priKey, x.Option.AppID+data["fileType"]+data["md5"]+data["timestamp"]); err != nil {
		return
	}

	var resp *resty.Response
	if resp, err = x.Client.R().
		SetContext(ctx).
		SetMultipartFormData(data).
		SetFileReader(`file`, opt.Name, bytes.NewBuffer(opt.Bytes)).
		Post(`/file/upload`); err != nil {
		return
	}

	if resp.StatusCode() != 200 {
		err = help.E(0, `第三方接口响应失败!`)
		return
	}

	var content M
	if err = sonic.Unmarshal(resp.Bytes(), &content); err != nil {
		return
	}

	if content["code"] != "00000" {
		err = help.E(0, fmt.Sprintf(`第三方请求失败![%s]: %s`, content["code"], content["msg"]))
		return
	}

	sign := content["sign"].(string)
	delete(content, "sign")
	delete(content, "signType")

	var verify bool
	if verify, err = help.Sm2Verify(x.pubKey, help.MapToSignText(content), sign); err != nil {
		return
	}
	if !verify {
		err = help.E(0, `第三方响应内容签名存在不一致!`)
		return
	}

	var result FileUploadResult
	if err = sonic.UnmarshalString(content["bizData"].(string), &result); err != nil {
		return
	}

	return &result, nil
}
