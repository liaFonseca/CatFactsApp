package main

import (
	"fmt"

	"github.com/bndr/gopencils"
	"github.com/gin-gonic/gin"
)

type catFacts struct {
	Facts []string `json:"data"`
}

func main() {
	// creates a gin default server
	server := gin.Default()

	// creates a Get endpoint (/test)
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "OK!"})
	})

	// starts the server
	server.Run(":8080")

	// connect to meowfacts api
	api := gopencils.Api("https://meowfacts.herokuapp.com")

	// create a new pointer to the response struct
	res := new(catFacts)

	// make a get request and unmarshal json response into response struct
	_, err := api.Res("", res).Get(map[string]string{"count": "3", "lang": "por-br"})

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, fact := range res.Facts {
			fmt.Println("fact:", fact)
		}
	}
}
