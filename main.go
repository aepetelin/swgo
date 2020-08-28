package main

import "log"

// @title swgoABC API
// @version 1.0
// @description This is a sample server
// @termsOfService http://swagger.io/terms/
// @host localhost:8801
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-JWT

func main() {
	log.Println("started ...")
	_ = Route()
	log.Println("exit")
}
