package Facade

import "testing"

/*
@Time    : 2021/1/12 12:57
@Author  : austsxk
@Email   : austsxk@163.com
@File    : facade_test.go
@Software: GoLand
*/

func TestNewSampleFacade(t *testing.T) {
	facade := NewSampleFacade()
	facade.sendLetter(
		"hello Song, this is facade pattern testing...",
		"austsxk@vip.qq.com")
}

// go test -v facade_test.go facade.go
