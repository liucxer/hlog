package hlog_test

import (
	"errors"
	"github.com/liucxer/hlog"
	"testing"
)

func TestLog(t *testing.T) {
	a := 10
	err := errors.New("test")
	hlog.Error("a:%d, err:%v", a, err)
	hlog.Warning("a:%d, err:%v", a, err)
	hlog.Info("a:%d, err:%v", a, err)
}
