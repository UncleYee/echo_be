package server

import (
	"echo/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()

	port := config.GetString("server.port")
	r.Run(port)
}
