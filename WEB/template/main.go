package main

import (
	"fmt"
	htmlTemplate "html/template"
	"os"
	"text/template"
)

func textAndhtml() {
	// 显示数据
	tplText := "我叫 {{.}}"
	tpl, err := template.New("tpl").Parse(tplText)
	if err != nil {
		fmt.Println(err)
	}
	tpl.Execute(os.Stdout, `<img src="xxxx" />`)
	fmt.Println()

	htmlTpl, err := htmlTemplate.New("tpl").Parse(tplText)
	if err != nil {
		fmt.Println(err)
	}
	htmlTpl.Execute(os.Stdout, `<img src="xxxx" />`)
}

func tplMust() {
	tplText := "{{ . }}"
	//语法错误直接panic
	tpl := template.Must(template.New("tpl").Parse(tplText))
	tpl.Execute(os.Stdout, "bob")

	//切片
	tpl.Execute(os.Stdout, []int{1, 2, 3})
	//映射
	tpl.Execute(os.Stdout, map[int]int{1: 1, 2: 2})
	//结构体
	tpl.Execute(os.Stdout, struct {
		name string
		age  int
	}{"bob", 20})
}

//切片
func rangeSlice() {
	//没有时为空
	tplText := "{{ .name }} - {{ .age }}"
	//语法错误直接panic
	tpl := template.Must(template.New("tpl").Parse(tplText))
	//切片
	tpl.Execute(os.Stdout, []int{1, 2, 3})
}

//map
func rangeMap() {
	//没有时为空
	tplText := "{{ .name }} - {{ .age }}"
	//语法错误直接panic
	tpl := template.Must(template.New("tpl").Parse(tplText))
	//map
	tpl.Execute(os.Stdout, map[string]string{"name": "bob", "age": "21"})
}

//struct
func rangeStruct() {
	//没有时为空，需要大写
	tplText := "{{ .Name }} - {{ .Age }}"
	//语法错误直接panic
	tpl := template.Must(template.New("tpl").Parse(tplText))
	//结构体
	tpl.Execute(os.Stdout, struct {
		Name string
		Age  int
	}{"bob", 20})
}

func main() {
	//textAndhtml()
	//tplMust()
	//rangeSlice()
	//rangeMap()
	rangeStruct()
}
