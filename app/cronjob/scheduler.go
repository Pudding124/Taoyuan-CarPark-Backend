package cronjob

import (
	"github.com/go-co-op/gocron"
	"taoyuan_carpark/park"
	"time"
)

func EnableSchedule() {
	// update postgres
	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(30).Seconds().Do(park.UpdateCarPark)
	s1.Every(1).Hour().Do(park.UpdateHistoryCarPark)
	s1.StartAsync()
}
