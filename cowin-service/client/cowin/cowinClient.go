package cowin

import (
	"context"
	"cowin-service/model"
	"cowin-service/model/generated/client/cowin"
	"cowin-service/utils"
	"github.com/go-openapi/swag"
	"net/url"
)

func ScheduleAppointment(ctx context.Context, requestBody *cowin.InlineObject4) *model.ScheduleAppointmentResponse {
	result := new(model.ScheduleAppointmentResponse)

	httpRequest := HttpRequest{}
	httpRequest.
		Post(utils.GetReader(requestBody)).
		Cowin(ctx).
		Url("/v2/appointment/schedule").
		ContentTypeJson()

	_ = DoRequestAndUnmarshal(&httpRequest, result)

	return result
}

func ReScheduleAppointment(ctx context.Context, requestBody *cowin.InlineObject5) bool {
	result := new(model.ScheduleAppointmentResponse)

	httpRequest := HttpRequest{}
	httpRequest.
		Post(utils.GetReader(requestBody)).
		Cowin(ctx).
		Url("/v2/appointment/reschedule").
		ContentTypeJson()

	_ = DoRequestAndUnmarshal(&httpRequest, result)

	return true
}

func GetSessionCalenderByPin(ctx context.Context, pinCode string, date string, vaccine string) *model.ArrayOfSessionCalendarEntrySchema {
	result := new(model.ArrayOfSessionCalendarEntrySchema)
	params := url.Values{}

	if !swag.IsZero(pinCode) {
		params.Set("pincode", pinCode)
	}
	if !swag.IsZero(date) {
		params.Set("date", date)
	}
	if !swag.IsZero(vaccine) {
		params.Set("vaccine", vaccine)
	}

	httpRequest := HttpRequest{}
	httpRequest.
		Get().
		Cowin(ctx).
		Url("/v2/appointment/sessions/calendarByPin").
		QueryParams(params).
		ContentTypeJson()

	_ = DoRequestAndUnmarshal(&httpRequest, result)

	return result
}
