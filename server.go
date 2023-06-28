package main

import (
	_ "github.com/joho/godotenv/autoload"

	db "github.com/AnggaArdhinata/indochat/src/configs"
	"github.com/AnggaArdhinata/indochat/src/libs"

	"github.com/AnggaArdhinata/indochat/src/routers"
)

func main() {

	db.Init()
	
	e := routers.Init()

	libs.Scheduler()

	e.Logger.Fatal(e.Start(":6625"))

}
