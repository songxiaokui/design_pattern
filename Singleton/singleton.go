package Singleton

import (
	"sync"
	"time"
)

/*
@Time    : 2021/1/13 12:26
@Author  : austsxk
@Email   : austsxk@163.com
@File    : singleton.go
@Software: GoLand
*/

/*
单例模式:

1. 建造者模式

2. 确保了一个类对象只有一个实例，而且实例化的对象会向整个系统提供该实例

3. 优点
	减少了内存开支（频繁创建与销毁对象）
	减少了系统性能开销（配置的读取与对象的依赖）
	避免对资源的多重占用（例如一个写文件的动作）
	设置系统全局资源访问点，优化和共享资源的访问

4. 缺点
	没有接口、不方便拓展
	对测试不利，不存在接口mock一个实例
	与单一职责原则有冲突

5. 使用场景
	在一个系统中，要求类实例的对象仅仅只有一个；
	比如数据库的连接对象等
	提供唯一序列号的环境
*/

// 系统一运行，就为系统提供一个可调用的系统对象
var (
	BaseData *BaseSource
	once     sync.Once
)

func init() {
	// 为了保证线程安全，所以只进行执行一次
	once.Do(func() {
		if BaseData == nil {
			BaseData = &BaseSource{}
		}
	})
}

type BaseSource struct {
	RealTime    time.Time
	ReaderCount int
	m           sync.Mutex
}

// 提供一个获取系统实例的方法，只要通过该方法就可以获得系统实例
func GetBaseSourceInstance() *BaseSource {
	return BaseData
}

func (b *BaseSource) GetNowTime() time.Time {
	return time.Now()
}

func (b *BaseSource) IncreaseCount() {
	// 调用后自增阅读量
	b.m.Lock()
	defer b.m.Unlock()
	b.ReaderCount++
}

func (b *BaseSource) GetUserReaderCount() int {
	return b.ReaderCount
}
