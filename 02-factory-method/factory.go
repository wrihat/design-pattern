package factory

// Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase是Operator接口的实现类，封装公用方法
type OperatorBase struct {
	a, b int
}

// OperatorBase的公用方法具体实现
func (o *OperatorBase) SetA(a int) {
	o.a = a
}
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// ----首先实现具体的产品类，也就是不通的运算方法类----
// 因为OperatorBase已经实现了两个通用方法，所以后续的具体产品类实现Result()方法则就是实现了Operator接口
// PlusOperator Operator的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

// MinusOperator Operator的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}

// ----接下来实现不同具体产品的具体工厂----
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}
