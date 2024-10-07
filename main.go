package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/josepheid-aviva/covid-cases/api"
)

func main() {
	handler := api.NewHandler()
	fmt.Println("Server running on http://localhost:3000/")
	if err := http.ListenAndServe("localhost:3000", handler); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
