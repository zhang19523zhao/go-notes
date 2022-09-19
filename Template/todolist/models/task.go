package models

import "fmt"

type Task struct {
	Id     int
	Name   string
	Status string
}

var task []*Task = []*Task{&Task{1, "洗衣服", "正在执行"}, &Task{2, "写作业", "已完成"}, &Task{3, "洗衣服", "执行中"}}

func GetTask() []*Task {
	return task
}

// 添加
func Add(name string) {
	task = append(task, &Task{len(task), name, "新增"})
}

// 删除
func Delete(name string) {
	for i, data := range task {
		if data.Name == name {
			fmt.Println(data.Name)
			fmt.Println("删除前", task)
			task = append(task[:i], task[i+1:]...)
			fmt.Println("删除后", task)
		}
	}
}

// 修改
func Modify(name1, name2, values string) {
	fmt.Println(name1, "修改为", name2, "状态: ", values)
	for _, data := range task {
		if data.Name == name1 {
			data.Name = name2
			data.Status = values
		}
	}
}

// 查询
func Query(name string) *Task {
	for _, data := range task {
		if data.Name == name {
			fmt.Println(data)
			return data
		}
	}
	return nil
}
