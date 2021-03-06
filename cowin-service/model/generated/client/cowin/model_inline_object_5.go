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

// InlineObject5 struct for InlineObject5
type InlineObject5 struct {
	AppointmentId string `json:"appointment_id"`
	SessionId string `json:"session_id"`
	Slot string `json:"slot"`
}

// NewInlineObject5 instantiates a new InlineObject5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject5(appointmentId string, sessionId string, slot string, ) *InlineObject5 {
	this := InlineObject5{}
	this.AppointmentId = appointmentId
	this.SessionId = sessionId
	this.Slot = slot
	return &this
}

// NewInlineObject5WithDefaults instantiates a new InlineObject5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject5WithDefaults() *InlineObject5 {
	this := InlineObject5{}
	return &this
}

// GetAppointmentId returns the AppointmentId field value
func (o *InlineObject5) GetAppointmentId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AppointmentId
}

// GetAppointmentIdOk returns a tuple with the AppointmentId field value
// and a boolean to check if the value has been set.
func (o *InlineObject5) GetAppointmentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AppointmentId, true
}

// SetAppointmentId sets field value
func (o *InlineObject5) SetAppointmentId(v string) {
	o.AppointmentId = v
}

// GetSessionId returns the SessionId field value
func (o *InlineObject5) GetSessionId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *InlineObject5) GetSessionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *InlineObject5) SetSessionId(v string) {
	o.SessionId = v
}

// GetSlot returns the Slot field value
func (o *InlineObject5) GetSlot() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Slot
}

// GetSlotOk returns a tuple with the Slot field value
// and a boolean to check if the value has been set.
func (o *InlineObject5) GetSlotOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Slot, true
}

// SetSlot sets field value
func (o *InlineObject5) SetSlot(v string) {
	o.Slot = v
}

func (o InlineObject5) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appointment_id"] = o.AppointmentId
	}
	if true {
		toSerialize["session_id"] = o.SessionId
	}
	if true {
		toSerialize["slot"] = o.Slot
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject5 struct {
	value *InlineObject5
	isSet bool
}

func (v NullableInlineObject5) Get() *InlineObject5 {
	return v.value
}

func (v *NullableInlineObject5) Set(val *InlineObject5) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject5) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject5(val *InlineObject5) *NullableInlineObject5 {
	return &NullableInlineObject5{value: val, isSet: true}
}

func (v NullableInlineObject5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
