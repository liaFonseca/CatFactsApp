package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bndr/gopencils"
	"github.com/liaFonseca/CatFactsApp/entity"
)

// interface which forces the implementation of a GetFacts function
type CatFactsService interface {
	GetNewFacts(int, string) ([]entity.Fact, error)
	GetFacts() []entity.Fact
	GetLanguages() []entity.Lang
}

// struct that will be connected to the interface
type catFactsService struct {
	CatFacts []entity.Fact
}

// connection of the struct with the interface --> returns a pointer to the struct as the interface
func New() CatFactsService {
	return &catFactsService{}
}

func (service *catFactsService) GetNewFacts(count int, lang string) ([]entity.Fact, error) {
	if newFacts, err := GetMeowFacts(count, lang); err != nil {
		return service.CatFacts, err
	} else {
		service.CatFacts = append(newFacts, service.CatFacts...)
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

func (service *catFactsService) GetLanguages() []entity.Lang {
	languages, _ := GetMeowLanguages()
	return languages.Lang
}

func GetMeowLanguages() (entity.Languages, error) {
	// connect to meowfacts api
	api := gopencils.Api("https://meowfacts.herokuapp.com")

	// create a new pointer to the response struct
	languages := new(entity.Languages)

	// make a get request and unmarshal json response into response struct
	_, err := api.Res("options", languages).Get()

	if err != nil {
		return *languages, err
	} else {
		for i, lang := range languages.Lang {
			languages.Lang[i].FullName = strings.ToLower(lang.FullName)
		}
		return *languages, err
	}
}

func (service *catFactsService) GetFacts() []entity.Fact {
	return service.CatFacts
}
