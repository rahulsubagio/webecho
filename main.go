package main

import "github.com/rahulsubagio/webecho/routes"

func main() {
	e := routes.Index()

	e.Logger.Fatal(e.Start(":1234"))
}
