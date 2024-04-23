package main

import (
	"net/http"

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
		if catFacts, err := catFactsController.GetFacts(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": catFacts})
		}

	})

	// starts the server
	server.Run(":8080")
}
