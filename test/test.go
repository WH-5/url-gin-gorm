package main

import (
	"fmt"
	"github.com/WH-5/url-gin-gorm/configs"
)

func main() {
	config, err := configs.LoadConfig("configs/config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)
}
