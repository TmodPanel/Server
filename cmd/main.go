package main

import "TSM-Server/internal/server"

func main() {
	r := server.NewRouter()
	r.Run(":8000")
}
