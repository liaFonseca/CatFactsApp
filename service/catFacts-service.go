package service

import (
	"fmt"
	"strconv"

	"github.com/bndr/gopencils"
	"github.com/liaFonseca/CatFactsApp/entity"
)

// interface which forces the implementation of a GetFacts function
type CatFactsService interface {
	GetFacts(int, string) ([]entity.Fact, error)
}

// struct that will be connected to the interface
type catFactsService struct {
	CatFacts []entity.Fact
}

// connection of the struct with the interface --> returns a pointer to the struct as the interface
func New() CatFactsService {
	return &catFactsService{}
}

func (service *catFactsService) GetFacts(count int, lang string) ([]entity.Fact, error) {
	if newFacts, err := GetMeowFacts(count, lang); err != nil {
		return service.CatFacts, err
	} else {
		service.CatFacts = append(service.CatFacts, newFacts...)
		return service.CatFacts, nil
	}
}

type MeowFacts struct {
	Facts []entity.Fact `json:"data"`
}

func GetMeowFacts(count int, lang string) ([]entity.Fact, error) {
	// connect to meowfacts api
	api := gopencils.Api("https://meowfacts.herokuapp.com")

	// create a new pointer to the response struct
	newFacts := new(MeowFacts)

	if count == 0 {
		count = 1
	}

	// make a get request and unmarshal json response into response struct
	_, err := api.Res("", newFacts).Get(map[string]string{"count": strconv.Itoa(count), "lang": lang})

	if err != nil {
		return []entity.Fact{}, err
	} else {
		for _, fact := range newFacts.Facts {
			fmt.Println("fact:", fact)
		}
		return newFacts.Facts, err
	}
}
