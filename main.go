package main

import (
	"fmt"
	"webserver/src/app/routes"
	dep "webserver/src/dependencies"
)

func main() {
	fmt.Println("Starting Application")
	dependencies, err := dep.BuildDependencies()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Dependencies Created With Succes")
	routes.Router(dependencies)
}
