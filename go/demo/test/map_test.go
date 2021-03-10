package test

import (
	"fmt"
	"testing"
)

func MapDemo() {
	dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(dict)

	//var colors map[string]string //create nil map,is not can used directly
	colors := map[string]string{}
	colors["Red"] = "#da1337"
	colors["Coral"] = "#ff7F50"

	fmt.Println(colors)

	//Check exist
	if value, exists := colors["Blue"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("not exist Blue")
	}

	//Check exist
	value := colors["Red"]
	if value != "" {
		fmt.Println(value)
	}

	//Print
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	//delete
	delete(colors, "Coral")
}
func TestMap(t *testing.T) {
	yinzhengjie := make(map[int]string)
	letter := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for k, v := range letter {
		yinzhengjie[k] = v
	}
	fmt.Printf("字典中的值为：【%v】\n", yinzhengjie) //注意，字典是无序的哟！
	if v, ok := yinzhengjie[1]; ok {
		fmt.Println("存在key=", v)
	} else {
		fmt.Println("没有找到key=", v)
	}
}
