package singleton

// 饿汉单例模式
type EagerSingleton struct{}

var EagerInstance = &EagerSingleton{} // 程序启动时就创建好了

func GetEagerInstance() *EagerSingleton {
	return EagerInstance
}
