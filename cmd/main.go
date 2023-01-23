package main

import (
	internal "forum/internal/server"
)

func main() {
	internal.Serve()
	// srv := new(internal.Server)
	// if err := srv.Run("8080"); err != nil {
	// 	log.Fatalf("error runtime: %s", err.Error())
	// 	return
	// }
}
