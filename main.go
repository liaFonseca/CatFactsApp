package main

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liaFonseca/CatFactsApp/controller"
	"github.com/liaFonseca/CatFactsApp/service"
	"github.com/liaFonseca/CatFactsApp/templates"
)

var (
	catFactsService    service.CatFactsService       = service.New()
	catFactsController controller.CatFactsController = controller.New(catFactsService)
)

func main() {
	// creates a gin default server
	server := gin.Default()
	server.SetTrustedProxies([]string{"localhost"})

	// load static assets:
	//    -> serves the js file
	jsDir, _ := fs.Sub(templates.TempDir, "js")
	fileSystem := http.FS(jsDir)
	server.StaticFS("/js", fileSystem)

	//    -> loads html filess
	templHTML := template.Must(template.ParseFS(templates.TempDir, "*.html"))
	server.SetHTMLTemplate(templHTML)

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

	server.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/view/catFacts"
		server.HandleContext(c)
	})

	// starts the server
	server.Run(":80")
}
