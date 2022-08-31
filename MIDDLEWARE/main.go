package main

import (
	"fmt"
	"net/http"
	"time"
)

var limitCh = make(chan struct{}, 5)

type middleware func(handler http.Handler) http.Handler

type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

//构造函数
func NewRouter() *Router {
	return &Router{
		middlewareChain: make([]middleware, 0, 10),
		mux:             map[string]http.Handler{},
	}
}

//添加中间件
func (self *Router) AddMiddleware(m middleware) {
	self.middlewareChain = append(self.middlewareChain, m)
}

//添加路由
func (self *Router) ADD(path string, handler http.Handler) {
	//margHandler := handler
	var margHandler = handler
	for i := 0; i < len(self.middlewareChain); i++ {
		margHandler = self.middlewareChain[i](margHandler)
	}
	self.mux[path] = margHandler
}

//计算执行时间中间件
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		//执行handler之前的逻辑
		next.ServeHTTP(w, r)
		fmt.Println("一共耗时: ", time.Now().Sub(start))
	})
}

//限流中间件
func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limitCh <- struct{}{}
		fmt.Printf("当前迸发度为: %d\n", len(limitCh))
		next.ServeHTTP(w, r)
		<-limitCh
	})
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It is OK"))
	time.Sleep(time.Second)
}

func main() {
	router := NewRouter()

	router.AddMiddleware(timeMiddleware)
	router.AddMiddleware(limitMiddleware)
	//fmt.Println("middleware", router.middlewareChain)
	router.ADD("/time", http.HandlerFunc(get))

	for path, h := range router.mux {
		fmt.Println(path, h)
		http.Handle(path, h)
	}

	//http.HandleFunc("/", get)
	http.ListenAndServe("localhost:8080", nil)
}
