package main

import (
	"github.com/kejne/simple_database/db"
	"github.com/kejne/simple_database/web"
)

func main() {
	db.Serve()
	web.Serve()
}

