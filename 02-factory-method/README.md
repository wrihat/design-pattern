# 工厂模式
## 1. 简单工厂模式
简单工厂模式的概念就是建立一个工厂类，对实现接口的一些类进行实例的创建。简单工厂类的实质是由一个工厂类根据传入的参数，动态决定应该创建哪一个产品类的实例。(这些产品类都继承自一个父类或接口)
* 注意
`go`语言没有构造函数一说，所以一般会定义`NewXX`函数来初始化相关类。`NewXX`函数返回接口时就是简单工厂模式，也就是`go`的一般推荐做法就是简单工厂模式。
## 2. 工厂方法模式
简单工厂只有一个工厂类，负责创建所有产品，如果添加新的产品，通常需要添加新的产品，通常需要修改工厂类的代码。而工厂类方法模式引入了抽象工厂和具体工厂的概念。每个具体工厂只负责创建一个具体产品，添加新的产品只需要添加新的工厂类而无需修改原来的代码，这样就使得产品的生产更加更加灵活，支持扩展，符合开闭原则。
工厂方法模式分为几个角色：
* 抽象工厂：一个接口，包含一个抽象的工厂方法（用于创建产品对象）
* 抽象产品：定义产品的接口
* 具体工厂：实现抽象工厂接口，创建具体产品
* 具体产品：实现抽象产品接口，是工厂创建的对象

工厂方法模式使用子类的方式延迟生成对象到子类中实现。
Go中不存在继承，所以使用匿名组合来实现

工厂方法模式使得每个工厂类的职责单一，每个工厂只负责创建一种产品，当创建对象涉及到一系列复杂的初始化逻辑，而这些逻辑在不同的子类中可能有所不同时，可以使用工厂方法模式将这些初始化逻辑封装在子类的工厂中。

## 3. 抽象工厂模式



## 4. 三种模式的区别




