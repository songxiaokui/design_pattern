package AbstractFactory

import "testing"

/*
@Time    : 2021/1/9 13:42
@Author  : austsxk
@Email   : austsxk@163.com
@File    : abstractFactory_test.go
@Software: GoLand
*/

func TestNewMealFactory(t *testing.T) {
	factory := NewMealFactory()
	food := factory.CreatFood()
	food.Cook()

	vegetables := factory.CreatVegetables()
	vegetables.Cook()
}

// go test -v abstractFactory_test.go abstractFactory.go
