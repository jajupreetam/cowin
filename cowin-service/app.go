package main

import (
	"context"
	"cowin-service/config"
	"cowin-service/service/cowin"
	slog "github.com/go-eden/slf4go"
)

const version = "1.1.9"

func init() {
	slog.Info("Running version ", version)
	config.InitProfiles()
}

func main() {
	cowin.ScheduleAppointment(context.Background(), []string{"682557"}, "08-05-2021",
		"COVAXIN", "79299353493330")

	//cowin.ReScheduleAppointment(context.Background(), "", []string{"411027"}, "07-05-2021", "COVAXIN")
}
