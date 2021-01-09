package Factory

/*
@Time    : 2021/1/9 13:10
@Author  : austsxk
@Email   : austsxk@163.com
@File    : factory.go
@Software: GoLand
*/

// 工厂模式
/*
创建型模式

提供创建对象的最佳方式方法 通过工厂批量产生实例化对象，在创建工厂时，不会对客户端暴露创建逻辑，
通过一个共同的接口指向新创建的对象

定义一个创建对象的接口，然后通过子类去实例具体的对象，可以通过参数进行区分。

使用条件: 可以明确的计划在不同条件下创建不同的实例
*/

// 创建一个车接口，包含一个方法 获取各种实例类品牌的名字
type Cars interface {
	GetCarBrandName() string
}

// 五菱宏光
type wuLingHongGuang struct {
}

// 上海大众
type shangHaiDaZhong struct {
}

func (c *wuLingHongGuang) GetCarBrandName() string {
	return "五菱宏光"
}

func (c *shangHaiDaZhong) GetCarBrandName() string {
	return "上海大众"
}

func NewCars(brand string) Cars {
	switch brand {
	case "WLHG":
		return &wuLingHongGuang{}
	case "SHDZ":
		return &shangHaiDaZhong{}
	default:
		return nil
	}
}
