package model

import "cowin-service/model/generated/client/cowin"

type ArrayOfSessionCalendarEntrySchema struct {
	Centers []cowin.SessionCalendarEntrySchema `json:"centers"`
}
