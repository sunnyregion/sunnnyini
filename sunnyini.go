package sunnyini

// Copyright 2017 sunnyini authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Source code and project home:
//
// https://github.com/sunnyregion/sunnyini
//
// Installation:
//
// go get  github.com/sunnyregion/sunnyini
//
// Example:
//
//   f := sunnyini.NewIniFile()
//   f.Readfile("conf/demo.ini")
//   section := f.GetSection()
//   …………
//   describ, v := f.GetValue("example")
//

import (
	"bufio"
	"io"
	"os"
	"strings"
)

//Element section下面的键值对
type Element map[string]string

//IniFile ini文件结构(对象)
// section为key
type IniFile struct {
	section string
	Object  map[string][]Element
}

// NewIniFile IniFile's construct fun
func NewIniFile() *IniFile {
	o := make(map[string][]Element)
	f := &IniFile{Object: o}
	return f
}

// Readfile read file读取文件
func (i *IniFile) Readfile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// 转换成一个bufio
	r := bufio.NewReader(f)
	for {
		// 以'\n'结束符读入一行
		line, err := r.ReadString('\n')
		if err != nil || io.EOF == err {
			if len(strings.TrimSpace(line)) < 3 {
				break
			}
		}
		// 删除行两端的空白字符
		line = strings.TrimSpace(line)
		// 解析一行中的内容
		// 1.[]格式
		// 2."" = ""格式
		i.Parse(line)
	}
}

// Parse 解析一行中的内容
func (i *IniFile) Parse(str string) {
	length := len(str)
	// fmt.Println("The string is:", str, "and the length is:", length)
	// 匹配字符串
	if str[0] == '[' && str[length-1] == ']' { // section
		root := str[1 : length-1]
		// 如果map中没有这个key，添加进来
		if _, ok := i.Object[root]; !ok {
			// 一个空的[]Element{},这里需要注意
			i.Object[root] = []Element{}
			i.section = root
		}
	} else if str == "" { // 空行

	} else if str[0] == '#' || str[0] == ';' {

	} else { // 键值对数据
		// 分割字符串
		s := strings.Split(str, "=")
		key := strings.TrimSpace(s[0])
		value := strings.TrimSpace(s[1])
		iValue := strings.Index(value, `//`)
		if iValue > 0 {
			value = value[:iValue]
			value = strings.TrimSpace(value)
		}

		element := make(Element)
		element[key] = value
		// 判断是否是在本key下面的
		if _, ok := i.Object[i.section]; ok {
			i.Object[i.section] = append(i.Object[i.section], element)
		}
	}
}

// GetSection 获取所有的Section
func (i *IniFile) GetSection() []string {
	key := []string{}
	for k := range i.Object {
		key = append(key, k)
	}
	return key
}

// GetValue 获取section下的键值对数据
func (i *IniFile) GetValue(section string) (string, []Element) {
	if v, ok := i.Object[section]; ok {
		return "", v
	}
	return "There is no data in " + section, nil
}

// GetValueInSection 获取某个section下的某个键的值
func (i *IniFile) GetValueInSection(section, key string) (string, string) {
	if v, ok := i.Object[section]; ok {
		for _, ele := range v { // 循环数组
			if m, ok := ele[key]; ok {
				return "", m
			}
			return "There is no key:" + key + " in section:" + section, ""
		}
	}
	return "There is no section:" + section, ""
}
