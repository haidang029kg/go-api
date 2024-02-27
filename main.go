package main

import (
	"goapi/src/api"
)

func main() {
	go func() {
		api.ServerGin()
	}()

	go func() {
		api.ServerGrpc()
	}()
	// Keep the main goroutine running
	select {}
}
