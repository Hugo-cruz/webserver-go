package main

import (
	"fmt"
	"webserver/src/app/routes"
	dep "webserver/src/dependencies"
)

func main() {
	fmt.Println("Starting")
	dependencies, err := dep.BuildDependencies()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("dependencies ok")
	routes.Router(dependencies)
}
