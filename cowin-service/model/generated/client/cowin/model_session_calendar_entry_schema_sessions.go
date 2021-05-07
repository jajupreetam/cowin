/*
 * Co-WIN Protected APIs
 *
 * Co-WIN APIs to register for vaccination, to book a vaccination appointment from partner applications. These APIs have restricted access. You can test these APIs on the Test Server using API Key -  <code>3sjOr2rmM52GzhpMHjDEE1kpQeRxwFDr4YcBEimi</code> Please contact Ministry of Health and Family Welfare, Government of India to get access to these APIs. [<i>Updated on 03 May 2021</i>]
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cowin

import (
	"encoding/json"
)

// SessionCalendarEntrySchemaSessions struct for SessionCalendarEntrySchemaSessions
type SessionCalendarEntrySchemaSessions struct {
	SessionId string `json:"session_id"`
	Date string `json:"date"`
	AvailableCapacity float32 `json:"available_capacity"`
	MinAgeLimit float32 `json:"min_age_limit"`
	Vaccine string `json:"vaccine"`
	// Array of slot names
	Slots []string `json:"slots"`
}

// NewSessionCalendarEntrySchemaSessions instantiates a new SessionCalendarEntrySchemaSessions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionCalendarEntrySchemaSessions(sessionId string, date string, availableCapacity float32, minAgeLimit float32, vaccine string, slots []string, ) *SessionCalendarEntrySchemaSessions {
	this := SessionCalendarEntrySchemaSessions{}
	this.SessionId = sessionId
	this.Date = date
	this.AvailableCapacity = availableCapacity
	this.MinAgeLimit = minAgeLimit
	this.Vaccine = vaccine
	this.Slots = slots
	return &this
}

// NewSessionCalendarEntrySchemaSessionsWithDefaults instantiates a new SessionCalendarEntrySchemaSessions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionCalendarEntrySchemaSessionsWithDefaults() *SessionCalendarEntrySchemaSessions {
	this := SessionCalendarEntrySchemaSessions{}
	return &this
}

// GetSessionId returns the SessionId field value
func (o *SessionCalendarEntrySchemaSessions) GetSessionId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *SessionCalendarEntrySchemaSessions) GetSessionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *SessionCalendarEntrySchemaSessions) SetSessionId(v string) {
	o.SessionId = v
}

// GetDate returns the Date field value
func (o *SessionCalendarEntrySchemaSessions) GetDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *SessionCalendarEntrySchemaSessions) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *SessionCalendarEntrySchemaSessions) SetDate(v string) {
	o.Date = v
}

// GetAvailableCapacity returns the AvailableCapacity field value
func (o *SessionCalendarEntrySchemaSessions) GetAvailableCapacity() float32 {
	if o == nil  {
		var ret float32
		return ret
	}

	return o.AvailableCapacity
}

// GetAvailableCapacityOk returns a tuple with the AvailableCapacity field value
// and a boolean to check if the value has been set.
func (o *SessionCalendarEntrySchemaSessions) GetAvailableCapacityOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailableCapacity, true
}

// SetAvailableCapacity sets field value
func (o *SessionCalendarEntrySchemaSessions) SetAvailableCapacity(v float32) {
	o.AvailableCapacity = v
}

// GetMinAgeLimit returns the MinAgeLimit field value
func (o *SessionCalendarEntrySchemaSessions) GetMinAgeLimit() float32 {
	if o == nil  {
		var ret float32
		return ret
	}

	return o.MinAgeLimit
}

// GetMinAgeLimitOk returns a tuple with the MinAgeLimit field value
// and a boolean to check if the value has been set.
func (o *SessionCalendarEntrySchemaSessions) GetMinAgeLimitOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinAgeLimit, true
}

// SetMinAgeLimit sets field value
func (o *SessionCalendarEntrySchemaSessions) SetMinAgeLimit(v float32) {
	o.MinAgeLimit = v
}

// GetVaccine returns the Vaccine field value
func (o *SessionCalendarEntrySchemaSessions) GetVaccine() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Vaccine
}

// GetVaccineOk returns a tuple with the Vaccine field value
// and a boolean to check if the value has been set.
func (o *SessionCalendarEntrySchemaSessions) GetVaccineOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vaccine, true
}

// SetVaccine sets field value
func (o *SessionCalendarEntrySchemaSessions) SetVaccine(v string) {
	o.Vaccine = v
}

// GetSlots returns the Slots field value
func (o *SessionCalendarEntrySchemaSessions) GetSlots() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Slots
}

// GetSlotsOk returns a tuple with the Slots field value
// and a boolean to check if the value has been set.
func (o *SessionCalendarEntrySchemaSessions) GetSlotsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Slots, true
}

// SetSlots sets field value
func (o *SessionCalendarEntrySchemaSessions) SetSlots(v []string) {
	o.Slots = v
}

func (o SessionCalendarEntrySchemaSessions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["session_id"] = o.SessionId
	}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["available_capacity"] = o.AvailableCapacity
	}
	if true {
		toSerialize["min_age_limit"] = o.MinAgeLimit
	}
	if true {
		toSerialize["vaccine"] = o.Vaccine
	}
	if true {
		toSerialize["slots"] = o.Slots
	}
	return json.Marshal(toSerialize)
}

type NullableSessionCalendarEntrySchemaSessions struct {
	value *SessionCalendarEntrySchemaSessions
	isSet bool
}

func (v NullableSessionCalendarEntrySchemaSessions) Get() *SessionCalendarEntrySchemaSessions {
	return v.value
}

func (v *NullableSessionCalendarEntrySchemaSessions) Set(val *SessionCalendarEntrySchemaSessions) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionCalendarEntrySchemaSessions) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionCalendarEntrySchemaSessions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionCalendarEntrySchemaSessions(val *SessionCalendarEntrySchemaSessions) *NullableSessionCalendarEntrySchemaSessions {
	return &NullableSessionCalendarEntrySchemaSessions{value: val, isSet: true}
}

func (v NullableSessionCalendarEntrySchemaSessions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionCalendarEntrySchemaSessions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}