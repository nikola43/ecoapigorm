package main

import "github.com/nikola43/ecoapigorm/app"

func main() {
	app := app.App{}
	app.Initialize(":3000")
}
