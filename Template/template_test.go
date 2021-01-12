package Template

import "testing"

/*
@Time    : 2021/1/12 18:14
@Author  : austsxk
@Email   : austsxk@163.com
@File    : template_test.go
@Software: GoLand
*/

func TestNewWorker(t *testing.T) {
	ewer := NewWorker(&EnvironmentalSanitationWorker{})
	ewer.Daily()
	programmer := NewWorker(&Programmer{})
	programmer.Daily()
}

// go test -v template.go template_test.go

/*
test result:
=== RUN   TestNewWorker
环卫阿姨起床了...
环卫阿姨正在打扫卫生...
环卫阿姨要休息了...
程序员起床了...
程序员正在疯狂敲代码...
程序员睡觉了...
--- PASS: TestNewWorker (0.00s)
PASS
ok      command-line-arguments  0.647s
*/
