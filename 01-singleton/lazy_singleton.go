package singleton

// 懒汉单例模式
import "sync"

type LazySingleton struct{}

var (
	LazyInstance *LazySingleton // 全局实例，但是并未初始化，近声明未定义
	onceLazy     sync.Once
)

func GetLazyInstance() *LazySingleton {
	// if LazyInstance == nil {   // 非线程安全版的懒汉模式
	// 	LazyInstance = &LazySingleton{}
	// }
	onceLazy.Do(func() { // 线程安全版推荐懒汉模式实现
		LazyInstance = &LazySingleton{}
	})
	return LazyInstance
}
