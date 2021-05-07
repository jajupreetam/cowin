package cowin

import (
	"context"
	cowinClient "cowin-service/client/cowin"
	"cowin-service/model"
	"cowin-service/model/generated/client/cowin"
	"fmt"
	"github.com/thoas/go-funk"
	"sync"
	"time"
)

func ScheduleAppointment(ctx context.Context, pinCodes []string, date string, vaccine string, beneficiary string) {
	fmt.Println("Started Scheduling Appointment")
	var (
		availableSessions *[]cowin.SessionCalendarEntrySchemaSessions
		start             time.Time
	)

	for {
		start = time.Now()
		availableSessions = getAvailableSession(ctx, pinCodes, date, vaccine)
		if availableSessions != nil && len(*availableSessions) > 0 {
			break
		}
	}

	if availableSessions != nil && len(*availableSessions) > 0 {
	loop:
		for _, availableSession := range *availableSessions {
			for _, slot := range availableSession.Slots {
				var dose float32 = 1
				requestBody := &cowin.InlineObject4{
					SessionId:     availableSession.SessionId,
					Slot:          slot,
					Beneficiaries: []string{beneficiary},
					Dose:          &dose,
				}
				response := cowinClient.ScheduleAppointment(ctx, requestBody)
				end := time.Now()
				fmt.Println(fmt.Sprintf("APPOINTMENT BOOKED at SUCCESSFULLY %v in time %v", response.AppointmentId, end.Sub(start)))
				break loop
			}
		}
	}
}

func ReScheduleAppointment(ctx context.Context, appointmentId string, pinCode []string, date string, vaccine string) bool {

	availableSessions := getAvailableSession(ctx, pinCode, date, vaccine)
loop2:
	for _, session := range *availableSessions {
		for _, slot := range session.Slots {
			requestBody := &cowin.InlineObject5{
				SessionId:     session.SessionId,
				Slot:          slot,
				AppointmentId: appointmentId,
			}
			isSuccess := cowinClient.ReScheduleAppointment(ctx, requestBody)
			fmt.Println("APPOINTMENT REBOOKED SUCCESSFULLY", isSuccess)
			break loop2
		}
	}
	return true
}

func getAvailableSession(ctx context.Context, pinCodes []string, date string, vaccine string) *[]cowin.SessionCalendarEntrySchemaSessions {
	var allPinHospitalEntries []model.ArrayOfSessionCalendarEntrySchema
	var availableSessions []cowin.SessionCalendarEntrySchemaSessions
	var wg = &sync.WaitGroup{}

	for _, p := range pinCodes {
		wg.Add(1)
		go func() {
			pinCenters := cowinClient.GetSessionCalenderByPin(ctx, wg, p, date, vaccine)
			if pinCenters != nil {
				allPinHospitalEntries = append(allPinHospitalEntries, *pinCenters)
			}
		}()
	}
	wg.Wait()

	if allPinHospitalEntries != nil && len(allPinHospitalEntries) > 0 {
		for _, pinHospital := range allPinHospitalEntries {
			for _, center := range pinHospital.Centers {
				if len(center.Sessions) > 0 {
					filteredAvailableSessions := filterAvailableSession(&center.Sessions, vaccine, date)
					if filteredAvailableSessions != nil && len(filteredAvailableSessions) > 0 {
						availableSessions = append(availableSessions, filteredAvailableSessions...)
					}
				}
			}
		}
	} else {
		fmt.Println(fmt.Sprintf("No center yet available at pinCode: %v...retrying...", pinCodes))
	}

	return &availableSessions
}

func filterAvailableSession(sessions *[]cowin.SessionCalendarEntrySchemaSessions, vaccine string, date string) []cowin.SessionCalendarEntrySchemaSessions {
	return funk.Filter(*sessions, func(session cowin.SessionCalendarEntrySchemaSessions) bool {
		return session.AvailableCapacity > 0 && vaccine == session.Vaccine && session.MinAgeLimit == 18 && session.Date == date
	}).([]cowin.SessionCalendarEntrySchemaSessions)
}
