package service

import (
	"fmt"

	"github.com/bndr/gopencils"
	"github.com/liaFonseca/CatFactsApp/entity"
)

// interface which forces the implementation of a GetFacts function
type CatFactsService interface {
	GetFacts(string, string) []entity.Fact
}

// struct that will be connected to the interface
type catFactsService struct {
	CatFacts []entity.Fact `json:"data"`
}

// connection of the struct with the interface --> returns a pointer to the struct as the interface
func New() CatFactsService {
	return &catFactsService{}
}

func (service *catFactsService) GetFacts(count string, lang string) []entity.Fact {
	newFacts := GetMeowFacts(count, lang)
	service.CatFacts = append(service.CatFacts, newFacts...)
	return service.CatFacts
}

func GetMeowFacts(count string, lang string) []entity.Fact {
	// connect to meowfacts api
	api := gopencils.Api("https://meowfacts.herokuapp.com")

	// create a new pointer to the response struct
	newFacts := []entity.Fact{}

	// make a get request and unmarshal json response into response struct
	_, err := api.Res("", newFacts).Get(map[string]string{"count": count, "lang": lang})

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, fact := range newFacts {
			fmt.Println("fact:", fact)
		}
	}

	return newFacts
}
