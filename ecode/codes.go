package ecode

import (
	"fmt"

	"github.com/go-kratos/kratos/pkg/ecode"
)

// All common ecode
var _codes map[int]string

// 主程序需要调用这个，注册全部已经定定义好的错误码
func Load() {
	ecode.Register(_codes)
}

// New new a Codes by int value.
// NOTE: ecode must unique in global, the New will check repeat and then panic.
func New(e int, msg string) ecode.Code {
	if _codes == nil {
		_codes = make(map[int]string, 128)
	}
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = msg
	return ecode.Int(e)
}
