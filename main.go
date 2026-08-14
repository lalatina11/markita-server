package main

import (
	"fmt"
	"log"

	"github.com/lalatina11/markita.git/src/app"
	"github.com/lalatina11/markita.git/src/config"
)

func main() {
	config := config.NewAppConfig()
	app := app.App()
	url := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	log.Fatal(app.Listen(url))
}
