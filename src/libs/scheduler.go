package libs

import (
	"time"

	"github.com/AnggaArdhinata/indochat/src/models"
	"github.com/robfig/cron/v3"
)

func Scheduler() error {

	user, err := models.PendingPayment()
	if err != nil {
		return err
	}

	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	defer scheduler.Stop()

	// Scheduler for send email every 23.50
	// scheduler.AddFunc("50 23 * * *", func() { SendEmail(email) })

	scheduler.AddFunc("*/2 * * * *", func() { SendEmail(user) })

	go scheduler.Start()

	return nil
}
