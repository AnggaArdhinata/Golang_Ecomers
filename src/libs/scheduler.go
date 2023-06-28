package libs

import (
	"time"

	"github.com/AnggaArdhinata/indochat/src/models"
	"github.com/robfig/cron/v3"
)

func Scheduler() error {

	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	defer scheduler.Stop()

	// Scheduler for send email every 23.50
	// scheduler.AddFunc("50 23 * * *", func() { SendEmail(user) })

	// Schedule every 2 mnute
	// scheduler.AddFunc("*/2 * * * *", func() { SendEmail(user) })

	//Schedule every 22.00
	scheduler.AddFunc("00 22 * * *", func() { SendEmail(models.PendingPayment()) })

	go scheduler.Start()

	return nil
}
