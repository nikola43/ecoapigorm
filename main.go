package main

import "github.com/nikola43/ecoapigorm/app"

func main() {
	a := app.App{}
	a.Initialize(":3001")
}
