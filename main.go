package main

import (
	"github.com/monikanaico-hub/goapi/app"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app.CreateUrlMappings()
	// Listen and server on 0.0.0.0:8080
	app.Router.Run(":8080")
}
