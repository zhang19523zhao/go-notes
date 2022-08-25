package main

import (
	"fmt"
	htmlTemplate "html/template"
	"os"
	"text/template"
)

func main() {
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
