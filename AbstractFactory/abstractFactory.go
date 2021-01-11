package AbstractFactory

import "fmt"

/*
@Time    : 2021/1/9 13:41
@Author  : austsxk
@Email   : austsxk@163.com
@File    : abstrcutFactory.go
@Software: GoLand
*/
// 工厂模式生成的是不同实例的对象，但是具有相同的方法，并实现不同的业务逻辑

// 抽象工厂，则是围绕一个超级工厂生成其他工厂

// 创建型类型

// 在抽象工厂模式中，接口负责生产一个相关对象的工厂。每个工厂又能够按照工厂模式提供对象

// 定义一个饭接口
type Meal interface {
	Cook()
}

// 定义两种实现饭的结构体
type StapleFood struct{}

func (s *StapleFood) Cook() {
	fmt.Println("It is Rice.")
}

type Vegetables struct{}

func (v *Vegetables) Cook() {
	fmt.Println("It is Potato.")
}

// 定义一个工厂，用来生产饭（主食、菜...）
type MealFactory interface {
	CreatFood() Meal
	CreatVegetables() Meal
}

// 定义一个简单的饭工厂，然后实现饭的全部实现
type SampleMealFactory struct{}

func (s *SampleMealFactory) CreatFood() Meal {
	return &StapleFood{}
}

func (s *SampleMealFactory) CreatVegetables() Meal {
	return &Vegetables{}
}

func NewMealFactory() MealFactory {
	return &SampleMealFactory{}
}
