package main

import (
	db "github.com/AnggaArdhinata/indochat/src/configs"
	// "github.com/AnggaArdhinata/indochat/src/models"

	"github.com/AnggaArdhinata/indochat/src/libs"

	"github.com/AnggaArdhinata/indochat/src/routers"
)

func main() {

	db.Init()

	// fmt.Println(models.GenerateCsv())

	libs.Scheduler()

	e := routers.Init()

	e.Logger.Fatal(e.Start(":6625"))

}
