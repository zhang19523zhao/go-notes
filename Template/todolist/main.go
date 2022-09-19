package main

import (
	"html/template"
	"net/http"
	"todolist/models"
)

func main() {
	addr := ":9999"
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.New("tpl").ParseFiles("vieiws/list.html"))
		tpl.ExecuteTemplate(w, "list.html", models.GetTask())
		//tpl.ExecuteTemplate(os.Stdout, "list.html", models.GetTask())
	})

	//添加
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.PostFormValue("name")
			models.Add(name)
			http.Redirect(w, r, "/list", 302)
		}
		tpl := template.Must(template.New("tpl").ParseFiles("vieiws/add.html"))
		tpl.ExecuteTemplate(w, "add.html", nil)
	})

	//删除
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.PostFormValue("name")
			models.Delete(name)
			http.Redirect(w, r, "/list", 302)
		}
		tpl := template.Must(template.New("tpl").ParseFiles("vieiws/delete.html"))
		tpl.ExecuteTemplate(w, "delete.html", nil)
	})

	// 修改
	http.HandleFunc("/modify", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name1 := r.PostFormValue("name1")
			name2 := r.PostFormValue("name2")
			status := r.PostFormValue("status")
			models.Modify(name1, name2, status)
			http.Redirect(w, r, "/list", 302)
		}
		tpl := template.Must(template.New("tpl").ParseFiles("vieiws/modify.html"))
		tpl.ExecuteTemplate(w, "modify.html", nil)
	})

	//查询
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		data := &models.Task{}
		if r.Method == http.MethodPost {
			name := r.PostFormValue("name")
			data = models.Query(name)
			//info := data.Id + data.Name + data.Status
			//w.Write([]byte(info))
		}
		tpl := template.Must(template.New("tpl").ParseFiles("vieiws/query.html"))
		tpl.ExecuteTemplate(w, "query.html", data)
		//tpl.ExecuteTemplate(os.Stdout, "query.html", data)
		//tpl, err := template.New("tpl").ParseFiles("vieiws/query.html")
		//fmt.Println(err)
	})
	http.ListenAndServe(addr, nil)
}
