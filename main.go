package main

import "github.com/WH-5/url-gin-gorm/cmd/app"

func main() {
	err := app.Run("configs/config.yaml")
	if err != nil {
		return
	}
}
