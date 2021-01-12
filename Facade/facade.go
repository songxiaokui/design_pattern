package Facade

import "fmt"

/*
@Time    : 2021/1/12 12:20
@Author  : austsxk
@Email   : austsxk@163.com
@File    : facade.go
@Software: GoLand
*/
/*

外观模式: 门面模式

1. 是常用的封装模式

2. 定义
	要求一个子系统的外部与内部通讯必须通过一个统一的接口对象进行，门面模式只提供一个高度抽象的接口，
使得子系统更方便使用。外观模式提供一个高抽象的接口，对外提供统一的服务，用来协调子系统的调用，降低了子
系统直接的耦合。此外门面模式不应该处理任何的业务逻辑。
3. 门面角色
	知晓子系统的所有角色和责任。客户端会将client的请求委派到相应的子系统，由子系统进行业务逻辑的处理。
门面类只是一个委托类。
4. 子系统角色
	子系统可以同时拥有一个或者多个子系统。每个子系统不是单独的类，而是相关类的集合。对子系统而言，并
不知道门面的存在。
5. 适用性
	a. 为复杂的子系统提供一个简单的接口;
	b. 客户端和抽象类直接存在很大的依赖关系时;
6. 优点
	a. 减少了系统的相互依赖;
	b. 提高了系统的灵活性;
	c. 提高了系统的安全性;
*/

// 提供一个邮寄信件的功能: 写信、填写地址、检查、装信封、投递
type SampleLetters struct {
	letter string
}

func (l *SampleLetters) WriteLetters(content string) {
	l.letter = fmt.Sprintf("写信内容为:%s", content)
}

func (l *SampleLetters) AddAddress(address string) {
	l.letter = l.letter + "\n" + fmt.Sprintf("邮寄至: %s", address)
}

func (l *SampleLetters) PutInToEnvelope() {
	fmt.Println("信件存放到信封中...")
}

func (l *SampleLetters) SendLetters() {
	// fmt.Println("信件内容:", l.letter)
	fmt.Println("信件已发送...")
}

type Facade interface {
	sendLetter(context, address string)
}

//// 定义一个简单门面结构体
//type SampleFacade struct {
//	SampleLetters
//}
//
//func (f *SampleFacade) sendLetter(context, address string) {
//	// 此处进行业务逻辑的实现
//	// 1. 写信
//	f.WriteLetters(context)
//	// 2. 写地址
//	f.AddAddress(address)
//	// 3. 封装
//	f.PutInToEnvelope()
//	// 4.发送
//	f.SendLetters()
//}

func NewSampleFacade() *SampleFacade {
	return &SampleFacade{}
}

// 如果此时需要新拓展功能，对信的内容进行检查
// 则继续添加SampleFacade的组合类
type PoliceInspection struct {
}

func (p *PoliceInspection) CheckLetterSecurity(content string) bool {
	// 如果检查内通通过，则返回true
	fmt.Printf("被检查内容: %s\n", content)
	fmt.Println("内容检查通过...")
	return true
}

type SampleFacade struct {
	SampleLetters
	PoliceInspection
}

func (f *SampleFacade) sendLetter(context, address string) {
	// 此处进行业务逻辑的实现
	// 1. 写信
	f.WriteLetters(context)
	// 2. 写地址
	f.AddAddress(address)
	// 此处新增检查信件安全性
	ok := f.CheckLetterSecurity(f.letter)
	if ok {
		fmt.Println("this is safe!")
	} else {
		fmt.Println("this is danger!")
		return
	}
	// 3. 封装
	f.PutInToEnvelope()
	// 4.发送
	f.SendLetters()
}

// 由此可见，在不改变用户调用的接口的情况下，已经完成了逻辑的修改，增加的新的功能，这就是门面模式带来的
// 好处
