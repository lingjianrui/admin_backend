package main

import (
	"backend/api"
)

func main() {
	server := &api.Server{}
	server.Initialize("mysql", "root", "xiaohei", "3306", "127.0.0.1", "backend")
	server.Run("0.0.0.0:8009")
}
