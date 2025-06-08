package factory

import "fmt"

// 定义一个接口，所有产品类都需要实现它
type API interface {
	Say(name string) string
}

// NewAPI函数用来通过传入的参数返回一个API的具体实例
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{} // &hiAPI{}表示的是一个*hiAPI类型的值，也就是hiAPI类型的指针
	} else if t == 2 { // 而在haiAPI的实现中，接收者是*hiAPI
		return &helloAPI{}
	}
	return nil
}

// hiAPI是一个API接口实现的产品类
type hiAPI struct{}

// Say方法的实现
func (hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// helloAPI 是另一个API接口实现的产品类
type helloAPI struct{}

// Say方法的实现
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
