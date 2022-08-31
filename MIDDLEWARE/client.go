package main

import (
	"fmt"
	"net/http"
	"time"
)

func Get() {
	client := http.Client{}
	requst, err := http.NewRequest("GET", "http://localhost:8080/time", nil)
	if err != nil {
		fmt.Println("err: ", err)
	}
	client.Do(requst)
}

func main() {

	for i := 0; i <= 15; i++ {
		go Get()
	}
	time.Sleep(time.Minute)
}
