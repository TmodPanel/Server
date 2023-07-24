package main

import (
	"TSM-Server/internal/Route"
)

func main() {

	r := Route.NewRouter()
	r.Run(":9000")
}
