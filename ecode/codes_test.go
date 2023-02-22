package codes

import (
	"testing"

	"github.com/go-kratos/kratos/pkg/ecode"
	"github.com/stretchr/testify/assert"
)

// 测试错误码重复添加
func TestNew(t *testing.T) {
	defer func() {
		errStr := recover()
		if errStr != "ecode: 1 already exist" {
			t.Logf("New duplicate ecode should cause panic")
			t.FailNow()
		}
	}()
	var _ error = New(1, "test_001")
	var _ error = New(2, "test_002")
	var _ error = New(1, "test_001")
}

func TestString(t *testing.T) {
	a := assert.New(t)

	Load()

	a.Equal(OK.Code(), 0)
	a.Equal(OK.Message(), "Succeed")

	eStr := ecode.String("123")
	if eStr.Code() != 123 {
		t.Logf("string parsed error code should be 123")
		t.FailNow()
	}
	if eStr.Error() != "123" {
		t.Logf("string parsed error string should be `123`")
		t.FailNow()
	}
}

// 测试错误引起原因追踪
func TestCause(t *testing.T) {
	e1 := New(4, "test_004")

	Load()

	var err error = e1
	// e2的错误由err引起
	e2 := ecode.Cause(err)
	if e2.Code() != 4 {
		t.Logf("parsed error code should be 4")
		t.FailNow()
	}
	a := assert.New(t)
	a.Equal(e1.Message(), "test_004")
}

func TestInt(t *testing.T) {
	e1 := ecode.Int(1)
	if e1.Code() != 1 {
		t.Logf("int parsed error code should be 1")
		t.FailNow()
	}
	if e1.Error() != "1" {
		t.Logf("int parsed error string should be `1`")
		t.FailNow()
	}
}

func TestLoad(t *testing.T) {
	Load()

	var code ecode.Code = 1309
	a := assert.New(t)
	a.Equal(code.Message(), ErrGetProductInfo.Message())

	bcode := ecode.Cause(ErrGetProductInfo)
	a.Equal(bcode.Message(), ErrGetProductInfo.Message())
}
