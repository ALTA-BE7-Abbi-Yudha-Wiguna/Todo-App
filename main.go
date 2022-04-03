package main

import (
	"todoListApp/config"
	"todoListApp/utils"
)

func main() {
	configs := config.GetConfig()
	utils.InitDB(configs)
}
