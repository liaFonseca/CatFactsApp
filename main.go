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

	// load static assets:
	//    -> serves the js file
	server.Static("/js", "templates/js")
	//    -> loads html files
	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/api")
	{
		// creates a Get endpoint (/catFacts)
		apiRoutes.GET("/catFacts", func(ctx *gin.Context) {
			if catFacts, err := catFactsController.GetFacts(ctx); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": catFacts})
			}

		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/catFacts", catFactsController.ShowAll)
	}

	// starts the server
	server.Run(":8080")
}
