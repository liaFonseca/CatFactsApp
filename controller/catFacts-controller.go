package controller

// the controller has the handlers for the API endpoints

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/liaFonseca/CatFactsApp/entity"
	"github.com/liaFonseca/CatFactsApp/service"
)

type CatFactsController interface {
	GetFacts(ctx *gin.Context) []entity.Fact
}

type catFactsController struct {
	service service.CatFactsService
}

func New(service service.CatFactsService) CatFactsController {
	return &catFactsController{service: service}
}

func (c *catFactsController) GetFacts(ctx *gin.Context) []entity.Fact {
	var details struct {
		Num  string `json:"count"`
		Lang string `json:"lang"`
	}

	if err := ctx.BindJSON(&details); err != nil {
		fmt.Println("Details 1: ", details)
		return c.service.GetFacts("0", "")
	} else {
		fmt.Println("Details 2: ", details)
		return c.service.GetFacts(details.Num, details.Lang)
	}
}
