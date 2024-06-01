package main

import (
	"fmt"
	"server/internal/server"
	"server/internal/auth"
	"server/internal/utils"
)

func main() {
	utils.InitPhotos()
	auth.NewAuth()
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
