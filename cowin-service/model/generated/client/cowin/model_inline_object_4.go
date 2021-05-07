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

// InlineObject4 struct for InlineObject4
type InlineObject4 struct {
	Dose *float32 `json:"dose,omitempty"`
	SessionId string `json:"session_id"`
	Slot string `json:"slot"`
	// Array of beneficiary reference ids
	Beneficiaries []string `json:"beneficiaries"`
}

// NewInlineObject4 instantiates a new InlineObject4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject4(sessionId string, slot string, beneficiaries []string, ) *InlineObject4 {
	this := InlineObject4{}
	this.SessionId = sessionId
	this.Slot = slot
	this.Beneficiaries = beneficiaries
	return &this
}

// NewInlineObject4WithDefaults instantiates a new InlineObject4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject4WithDefaults() *InlineObject4 {
	this := InlineObject4{}
	return &this
}

// GetDose returns the Dose field value if set, zero value otherwise.
func (o *InlineObject4) GetDose() float32 {
	if o == nil || o.Dose == nil {
		var ret float32
		return ret
	}
	return *o.Dose
}

// GetDoseOk returns a tuple with the Dose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetDoseOk() (*float32, bool) {
	if o == nil || o.Dose == nil {
		return nil, false
	}
	return o.Dose, true
}

// HasDose returns a boolean if a field has been set.
func (o *InlineObject4) HasDose() bool {
	if o != nil && o.Dose != nil {
		return true
	}

	return false
}

// SetDose gets a reference to the given float32 and assigns it to the Dose field.
func (o *InlineObject4) SetDose(v float32) {
	o.Dose = &v
}

// GetSessionId returns the SessionId field value
func (o *InlineObject4) GetSessionId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetSessionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *InlineObject4) SetSessionId(v string) {
	o.SessionId = v
}

// GetSlot returns the Slot field value
func (o *InlineObject4) GetSlot() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Slot
}

// GetSlotOk returns a tuple with the Slot field value
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetSlotOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Slot, true
}

// SetSlot sets field value
func (o *InlineObject4) SetSlot(v string) {
	o.Slot = v
}

// GetBeneficiaries returns the Beneficiaries field value
func (o *InlineObject4) GetBeneficiaries() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Beneficiaries
}

// GetBeneficiariesOk returns a tuple with the Beneficiaries field value
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetBeneficiariesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Beneficiaries, true
}

// SetBeneficiaries sets field value
func (o *InlineObject4) SetBeneficiaries(v []string) {
	o.Beneficiaries = v
}

func (o InlineObject4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dose != nil {
		toSerialize["dose"] = o.Dose
	}
	if true {
		toSerialize["session_id"] = o.SessionId
	}
	if true {
		toSerialize["slot"] = o.Slot
	}
	if true {
		toSerialize["beneficiaries"] = o.Beneficiaries
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject4 struct {
	value *InlineObject4
	isSet bool
}

func (v NullableInlineObject4) Get() *InlineObject4 {
	return v.value
}

func (v *NullableInlineObject4) Set(val *InlineObject4) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject4) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject4(val *InlineObject4) *NullableInlineObject4 {
	return &NullableInlineObject4{value: val, isSet: true}
}

func (v NullableInlineObject4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
