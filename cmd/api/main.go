package main

import (
	"fmt"
	"log"

	"github.com/cie-upb/resource-management/config"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("initializing... resource-management")

	app, err := config.App()
	if err != nil {
		log.Fatal(err)
	}
	defer app.CloseDBConnection()

	//timeout := time.Duration(config.NewEnv().ContextTimeout) * time.Second
	gin := gin.Default()

	//route.Setup(config.Env, timeout, app, gin)

	gin.Run(":" + config.NewEnv().ServerPort)
}
