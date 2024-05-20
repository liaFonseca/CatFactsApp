package controller

// the controller has the handlers for the API endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liaFonseca/CatFactsApp/entity"
	"github.com/liaFonseca/CatFactsApp/service"
)

type CatFactsController interface {
	GetFacts(ctx *gin.Context) ([]entity.Fact, error)
	ShowAll(ctx *gin.Context)
}

type catFactsController struct {
	service service.CatFactsService
}

func New(service service.CatFactsService) CatFactsController {
	return &catFactsController{service: service}
}

type getFactsParams struct {
	Count int    `json:"count" form:"count" binding:"gte=0,lte=10"`
	Lang  string `json:"lang" form:"lang"`
}

func (c *catFactsController) GetFacts(ctx *gin.Context) ([]entity.Fact, error) {
	details := new(getFactsParams)

	if err := ctx.ShouldBindQuery(details); err != nil {
		return []entity.Fact{}, err
	} else {
		return c.service.GetNewFacts(details.Count, details.Lang)
	}
}

func (c *catFactsController) ShowAll(ctx *gin.Context) {
	facts := c.service.GetFacts()
	lang := c.service.GetLanguages()

	data := gin.H{
		"Title": "Cat Facts",
		"Facts": facts,
		"Lang":  lang,
	}

	// send current Cat Facts to the html page
	ctx.HTML(http.StatusOK, "index.html", data)
}
