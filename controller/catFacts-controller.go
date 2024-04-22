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

type getFactsParams struct {
	Count string `json:"count" binding:"gte=1, lte=10"`
	Lang  string `json:"lang"`
}

func (c *catFactsController) GetFacts(ctx *gin.Context) []entity.Fact {
	details := new(getFactsParams)

	if err := ctx.ShouldBindJSON(details); err != nil {
		fmt.Println("Details 1: ", details)
		return c.service.GetFacts(0, "")
	} else {
		fmt.Println("Details 2: ", details)
		return c.service.GetFacts(details.Num, details.Lang)
	}
}
