package main

import (
	"pos/configs"
	"pos/routes"
)

func main() {
	configs.InitDB()
	e := routes.New()
	e.Start(":8000")
}
