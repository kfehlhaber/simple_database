package db

import "fmt"

type handler struct {
	key    string
	value  string
	result chan string
	action func(key string, value string)
}

var lock = make(chan int, 1)
var queue = make(chan *handler, 1000)

var dbMap = make(map[string]string)

func Fetch(key string) string {
	result := make(chan string)
	queue <- &handler{key, "", result, func(key string, i string) {
			result <- dbMap[key]
		}}
	return <-result
}

func Persist(key string, value string) (string) {
	result := make(chan string)
	queue <- &handler{key, value, result, func(key string, value string) {
		dbMap[key] = value
		result <- "SUCCESS"
	}}
	return <-result
}

func Serve() {
	go serve(queue)
}

func serve(queue chan *handler) {
	fmt.Println("DB server ready to accept requests")
	for {
		lock <- 1
		req := <-queue
		go handle(req)
	}
}

func handle(h *handler) {
	h.action(h.key, h.value)
	<-lock
}
