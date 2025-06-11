package abstractfactory

import "fmt"

func main() {
	var i interface{} = 3.14
	switch v := i.(type) {
	case int:
		fmt.Println("int: ", v)
	case float64:
		fmt.Println("float64: ", v)
	default:
		fmt.Println("unkonwn type")
	}

}
