package main

import (
	"fmt"
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
	tpl.Execute(os.Stdout, "BOB")
}
