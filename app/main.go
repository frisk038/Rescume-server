package main

import (
	"fmt"
	"rescueme-server/adapters/database"
	"rescueme-server/app/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	repo, err := database.NewPGStore()
	if err != nil {
		fmt.Println(err)
	}

	r := gin.Default()
	r.POST("/user/:userid/coordinates", handler.POSTCoordinates(repo))
	r.Run()
}
