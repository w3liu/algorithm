package main

import (
	"errors"
	"fmt"
)

func main() {
	f := NewParamFilter()
	f.Use(
		newFirst("11"),
		newSecond("22"),
	)
	resp := f.Valid()
	fmt.Println(resp.Code, resp.Msg, resp.Err)
}

type CheckResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Err  error  `json:"err"`
}

type IParam interface {
	Check() CheckResp
}

func NewCheckRespOk() CheckResp {
	return CheckResp{
		Code: "0000",
		Msg:  "success",
	}
}

func NewCheckRespErr() CheckResp {
	return CheckResp{
		Err: errors.New("param error"),
	}
}

type ParamFilter struct {
	checkChain []IParam
}

func NewParamFilter(fns ...IParam) ParamFilter {
	chain := make([]IParam, 0, len(fns))
	chain = append(chain, fns...)
	return ParamFilter{
		checkChain: chain,
	}
}

func (f *ParamFilter) Use(obj ...IParam) {
	f.checkChain = append(f.checkChain, obj...)
}

func (f *ParamFilter) Valid() CheckResp {
	for i, _ := range f.checkChain {
		resp := f.checkChain[i].Check()
		if resp.Err != nil {
			return resp
		}
		if resp.Code != "0000" {
			return resp
		}
	}
	return NewCheckRespOk()
}

type firstValidator struct {
	phone string
}

func newFirst(p string) firstValidator {
	return firstValidator{
		phone: p,
	}
}

func (v firstValidator) Check() CheckResp {
	if v.phone == "86-1111" {
		return NewCheckRespOk()
	}
	return NewCheckRespErr()
}

type secondValidator struct {
	phone string
}

func newSecond(p string) secondValidator {
	return secondValidator{
		phone: p,
	}
}

func (v secondValidator) Check() CheckResp {
	if v.phone == "86-1111" {
		return NewCheckRespOk()
	}
	return NewCheckRespErr()
}
