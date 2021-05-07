package cowin

import (
	"context"
	cowinClient "cowin-service/client/cowin"
	"cowin-service/model"
	"cowin-service/model/generated/client/cowin"
	"fmt"
	"github.com/thoas/go-funk"
)

func ScheduleAppointment(ctx context.Context, pinCode []string, date string, vaccine string, beneficiary string) {
	availableSessions := getAvailableSession(ctx, pinCode, date, vaccine)

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
				fmt.Println(fmt.Sprintf("APPOINTMENT BOOKED at SUCCESSFULLY %v", response.AppointmentId))
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

func getAvailableSession(ctx context.Context, pinCode []string, date string, vaccine string) *[]cowin.SessionCalendarEntrySchemaSessions {
	var centerEntries *model.ArrayOfSessionCalendarEntrySchema
	var availableSessions []cowin.SessionCalendarEntrySchemaSessions
	for _, p := range pinCode {
		centerEntries = cowinClient.GetSessionCalenderByPin(ctx, p, date, vaccine)
	}
	if centerEntries != nil && len(centerEntries.Centers) > 0 {

		for _, center := range centerEntries.Centers {
			if len(center.Sessions) > 0 {
				filteredAvailableSessions := filterAvailableSession(&center.Sessions, vaccine, date)
				if filteredAvailableSessions != nil && len(filteredAvailableSessions) > 0 {
					availableSessions = append(availableSessions, filteredAvailableSessions...)
				}
			}
		}
	} else {
		fmt.Println(fmt.Sprintf("No center yet available at pinCode: %v...retrying...", pinCode))
	}

	return &availableSessions
}

func filterAvailableSession(sessions *[]cowin.SessionCalendarEntrySchemaSessions, vaccine string, date string) []cowin.SessionCalendarEntrySchemaSessions {
	return funk.Filter(*sessions, func(session cowin.SessionCalendarEntrySchemaSessions) bool {
		return session.AvailableCapacity > 0 && vaccine == session.Vaccine && session.MinAgeLimit == 45 && session.Date == date
	}).([]cowin.SessionCalendarEntrySchemaSessions)
}
