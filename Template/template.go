package Template

import "fmt"

/*
@Time    : 2021/1/12 13:25
@Author  : austsxk
@Email   : austsxk@163.com
@File    : template.go
@Software: GoLand
*/

/*
模版方法模式:
将不变的行为抽取出来放到父类中，从而去掉子类中重复的代码。然后子类分别实现自己不同的业务逻辑

1. 模版方法模式属于类行为模式

2. 意图
	定义一个操作的类型的骨架，然后将这些实现延迟到子类中
*/

// 定义一个基类接口，该接口包含了一系列相关的操作
type WorkInterface interface {
	GetUp()
	Working()
	Sleep()
}

// 定义一个工作者的接口，每个工作者都有这些基本操作，然后将具体的实现交给子类
type Worker struct {
	WorkInterface
}

// 定义一个工作者的日常行为
func (w *Worker) Daily() {
	w.GetUp()
	w.Working()
	w.Sleep()
}

// 任何实现了WorkInterface的子类，都可以传入到其中，然后调用其日常方法，就能实现子类的三个具体实现，即多态的实现
func NewWorker(w WorkInterface) *Worker {
	return &Worker{w}
}

// 定义一个清洁员类
type EnvironmentalSanitationWorker struct {
}

func (e *EnvironmentalSanitationWorker) GetUp() {
	fmt.Println("环卫阿姨起床了...")
}

func (e *EnvironmentalSanitationWorker) Working() {
	fmt.Println("环卫阿姨正在打扫卫生...")
}

func (e *EnvironmentalSanitationWorker) Sleep() {
	fmt.Println("环卫阿姨要休息了...")
}

// 定义一个程序员类
type Programmer struct {
}

func (e *Programmer) GetUp() {
	fmt.Println("程序员起床了...")
}

func (e *Programmer) Working() {
	fmt.Println("程序员正在疯狂敲代码...")
}

func (e *Programmer) Sleep() {
	fmt.Println("程序员睡觉了...")
}
