package main

import "gogle-class/backend/internal/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
