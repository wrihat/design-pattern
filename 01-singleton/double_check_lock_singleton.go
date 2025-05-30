package singleton

// 双重检查锁定单例模式
import "sync"

type DCLSingleton struct{}

var (
	DCLInstance *DCLSingleton
	mutex_dcl   sync.Mutex
)

func GetDCLInstance() *DCLSingleton {
	if DCLInstance == nil {
		mutex_dcl.Lock()
		defer mutex_dcl.Unlock()
		if DCLInstance == nil {
			DCLInstance = &DCLSingleton{}
		}
	}
	return DCLInstance
}
