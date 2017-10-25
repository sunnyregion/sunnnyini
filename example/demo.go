package main

// 主函数
import (
	"fmt"
	"github.com/sunnyregion/sunn"
)

func main() {
	fmt.Println("Hello,this is test for parse ini file.")
	f := readini.NewIniFile()
	f.Readfile("../conf/conf.ini")
	// 获取所有的Section
	section := f.GetSection()
	fmt.Println("There are:", len(section), "section")
	for i, v := range section {
		fmt.Println("The index of", i, "section:", v)
	}
	// 获取某一个section下的键值对
	describ, v := f.GetValue("test")
	fmt.Println("Get the map in section test")
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
	describ, value := f.GetValueInSection("test", "id")
	fmt.Println("Get the value in section:test of key:id")
	if describ == "" {
		fmt.Println("The val is", value)
	} else {
		fmt.Println(describ)
	}
}

 评论 (0)