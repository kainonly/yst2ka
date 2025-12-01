package yst2ka

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type Tq3004Dto struct {
}

func NewTq3004Dto() *Tq3004Dto {
	return &Tq3004Dto{}
}

type Tq3004Result struct {
}

func (x *Yst2Ka) Tm3004(ctx context.Context, dto *Tq3004Dto) (_ *Tq3004Result, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now),
		`/tm/handle`, `3004`, data); err != nil {
		return
	}

	var result Tq3004Result
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
