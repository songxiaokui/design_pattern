package Singleton

import (
	"sync"
	"testing"
)

/*
@Time    : 2021/1/13 12:26
@Author  : austsxk
@Email   : austsxk@163.com
@File    : singleton_test.go
@Software: GoLand
*/
func TestGetBaseSourceInstance(t *testing.T) {
	var (
		wg    sync.WaitGroup
		count = 1000
	)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			instance := GetBaseSourceInstance()
			instance.IncreaseCount()
		}()
	}
	wg.Wait()
	instance := GetBaseSourceInstance()
	t.Logf("最后的阅读总数:%d", instance.ReaderCount)
	if instance.ReaderCount != count {
		t.Fatal("singleton is error")
	}
}

// go test -v singleton_test.go singleton.go
