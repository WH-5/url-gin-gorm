package main

import "github.com/WH-5/url-gin-gorm/configs"

func main() {

	config, err := configs.LoadConfig()
	if err != nil {
		return
	}

}
