package main

import (
	"fmt"

	"github.com/bndr/gopencils"
)

type catFacts struct {
	Facts []string `json:"data"`
}

func main() {
	// connect to meowfacts api
	api := gopencils.Api("https://meowfacts.herokuapp.com")

	// create a new pointer to the response struct
	res := new(catFacts)

	// make a get request and unmarshal json response into response struct
	_, err := api.Res("", res).Get()

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, fact := range res.Facts {
			fmt.Println("fact:", fact)
		}
	}
}
