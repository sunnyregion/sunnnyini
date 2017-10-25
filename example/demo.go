package main

// 主函数
import (
	"fmt"

	"github.com/sunnyregion/sunnyini"
)

func main() {
	fmt.Println("Hello,this is test for parse ini file.")
	f := sunnyini.NewIniFile()
	f.Readfile("conf/demo.ini")
	// 获取所有的Section
	section := f.GetSection()
	fmt.Println("There are:", len(section), "section")
	for i, v := range section {
		fmt.Println("The index of", i, "section:", v)
	}
	// 获取某一个section下的键值对
	describ, v := f.GetValue("example")
	fmt.Println("Get the map in section example1")
	if describ == "" { // 有数据
		for _, value := range v {
			for k, v := range value {
				fmt.Println("The key is:", k, "and the value is:", v)
			}
		}
	} else {
		fmt.Println(describ)
	}
	// 获取某个section下的某个key的值
	describ, value := f.GetValueInSection("example", "id")
	fmt.Println("Get the value in section:example of key:id")
	if describ == "" {
		fmt.Println("The val is", value)
	} else {
		fmt.Println(describ)
	}
}
