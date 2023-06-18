package main

import (

	"fmt"

	db "github.com/AnggaArdhinata/indochat/src/configs"
	"github.com/AnggaArdhinata/indochat/src/libs"
	"github.com/AnggaArdhinata/indochat/src/models"

	"github.com/AnggaArdhinata/indochat/src/routers"
)

func main() {

	db.Init()

	str := models.PendingPayment()
	fmt.Println(str)

	libs.Scheduler()

	e := routers.Init()
	e.Logger.Fatal(e.Start(":8080"))

	
}
