package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tq3001Dto struct {
}

func NewTq3001Dto() *Tq3001Dto {
	return &Tq3001Dto{}
}

type Tq3001Result struct {
}

func (x *Yst2Ka) Tq3001(ctx context.Context, dto *Tq3001Dto) (_ *Tq3001Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tq/handle`, `3001`, data); err != nil {
		return
	}

	var result Tq3001Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
