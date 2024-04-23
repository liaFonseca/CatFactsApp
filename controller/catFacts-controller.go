package controller

// the controller has the handlers for the API endpoints

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/liaFonseca/CatFactsApp/entity"
	"github.com/liaFonseca/CatFactsApp/service"
)

type CatFactsController interface {
	GetFacts(ctx *gin.Context) ([]entity.Fact, error)
}

type catFactsController struct {
	service service.CatFactsService
}

func New(service service.CatFactsService) CatFactsController {
	return &catFactsController{service: service}
}

type getFactsParams struct {
	Count int    `json:"count" binding:"gte=0,lte=10"`
	Lang  string `json:"lang"`
}

func (c *catFactsController) GetFacts(ctx *gin.Context) ([]entity.Fact, error) {
	details := new(getFactsParams)

	if err := ctx.ShouldBindJSON(details); err != nil {
		fmt.Println("Details 1: ", details)
		return []entity.Fact{}, err
	} else {
		fmt.Println("Details 2: ", details)
		return c.service.GetFacts(details.Count, details.Lang), nil
	}
}
