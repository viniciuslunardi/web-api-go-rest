package main

import (
	"fmt"

	"github.com/viniciuslunardi/api-go-rest.git/database"
	"github.com/viniciuslunardi/api-go-rest.git/routes"
)

func main() {

	database.ConnectDatabase()
	fmt.Println("Application running on 8000")
	routes.HandleRequest()
}
