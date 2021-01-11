package Factory

import "testing"

/*
@Time    : 2021/1/9 13:30
@Author  : austsxk
@Email   : austsxk@163.com
@File    : factory_test.go
@Software: GoLand
*/

func TestNewCars(t *testing.T) {
	brands1 := NewCars("WLHG").GetCarBrandName()
	if brands1 != "五菱宏光" {
		t.Fatalf("v 的值错误，期望值：五菱宏光")
	}
	brands2 := NewCars("SHDZ").GetCarBrandName()
	if brands2 != "上海大众" {
		t.Fatalf("v 的值错误，期望值：上海大众")
	}

}

//  go test -v factory_test.go factory.go
