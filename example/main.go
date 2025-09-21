package main

import (
	"fmt"

	bigObj "github.com/J-Siu/go-basestruct/example/bigobj"
)

func main() {

	var (
		str  string
		obj1 bigObj.BigObject
	)

	fmt.Println("-- obj1: Using [bigObj.BigObject] with no initialization\n")
	str = obj1.SetName("Cube").SetHeight(10).SetLength(10).SetWidth(10).String()
	if obj1.Err != nil {
		fmt.Println(obj1.Err.Error()) // Normal to reach this line
	} else {
		fmt.Println(str)
	}

	fmt.Println("\n-- obj2: Name not set\n")
	obj2 := bigObj.New()
	str = obj2.SetHeight(10).SetLength(10).SetWidth(10).String()
	if obj2.Err != nil {
		fmt.Println(obj2.Err.Error()) // Normal to reach this line
	} else {
		fmt.Println(str)
	}

	fmt.Println("\n--obj3: No error\n")
	obj3 := bigObj.New()
	str = obj3.SetName("Cube").SetHeight(10).SetLength(10).SetWidth(10).String()
	if obj3.Err != nil {
		fmt.Println(obj3.Err.Error()) // Should not reach this line
	} else {
		fmt.Println(str)
	}

}
