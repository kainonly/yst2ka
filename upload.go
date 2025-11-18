package yst2ka

import (
	"context"
)

type FileUploadDto struct {
	AppId     string `json:"appId"`     // 应用号
	SpAppId   string `json:"spAppId"`   // 服务商应用号
	FileType  string `json:"fileType"`  // 文件类型
	Md5       string `json:"md5"`       // 文件的MD5值
	Timestamp string `json:"timestamp"` // 时间戳
	Sign      string `json:"sign"`      // 签名
	File      string `json:"file"`      // 文件流
}

func NewFileUploadDto(orgRespTraceNum string, closeReason string) *FileUploadDto {
	return &FileUploadDto{}
}

type FileUploadResult struct {
	FileId   string `json:"FileId"`   // 文件标识
	RespCode string `json:"RespCode"` // 业务返回码
	RespMsg  string `json:"RespMsg"`  // 返回说明
}

func (x *Yst2Ka) FileUpload(ctx context.Context, dto *FileUploadDto) (_ *FileUploadResult, err error) {
	//now := time.Now()
	//var data string
	//if data, err = sonic.MarshalString(*dto); err != nil {
	//	return
	//}
	//
	//
	//var result FileUploadResult
	//if err = sonic.UnmarshalString(bizData, &result); err != nil {
	//	return
	//}
	//return &result, nil
	return
}
