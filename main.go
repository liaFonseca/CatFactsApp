package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liaFonseca/CatFactsApp/controller"
	"github.com/liaFonseca/CatFactsApp/service"
)

var (
	catFactsService    service.CatFactsService       = service.New()
	catFactsController controller.CatFactsController = controller.New(catFactsService)
)

func main() {
	// creates a gin default server
	server := gin.Default()

	// creates a Get endpoint (/test)
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "OK!"})
	})

	// creates a Get endpoint (/catFacts)
	server.GET("/catFacts", func(ctx *gin.Context) {
		ctx.JSON(200, catFactsController.GetFacts(ctx))
	})

	// starts the server
	server.Run(":8080")
}
