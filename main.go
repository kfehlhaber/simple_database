package main

import (
	"github.com/einride/pair-programming-kaj-fehlhaber/db"
	"github.com/einride/pair-programming-kaj-fehlhaber/web"
)

func main() {
	db.Serve()
	web.Serve()
}

