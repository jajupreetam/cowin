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

// InlineResponse2004States struct for InlineResponse2004States
type InlineResponse2004States struct {
	StateId float32 `json:"state_id"`
	StateName string `json:"state_name"`
	// State name in preferred language as specified in Accept-Language header parameter.
	StateNameL *string `json:"state_name_l,omitempty"`
}

// NewInlineResponse2004States instantiates a new InlineResponse2004States object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2004States(stateId float32, stateName string, ) *InlineResponse2004States {
	this := InlineResponse2004States{}
	this.StateId = stateId
	this.StateName = stateName
	return &this
}

// NewInlineResponse2004StatesWithDefaults instantiates a new InlineResponse2004States object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2004StatesWithDefaults() *InlineResponse2004States {
	this := InlineResponse2004States{}
	return &this
}

// GetStateId returns the StateId field value
func (o *InlineResponse2004States) GetStateId() float32 {
	if o == nil  {
		var ret float32
		return ret
	}

	return o.StateId
}

// GetStateIdOk returns a tuple with the StateId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004States) GetStateIdOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StateId, true
}

// SetStateId sets field value
func (o *InlineResponse2004States) SetStateId(v float32) {
	o.StateId = v
}

// GetStateName returns the StateName field value
func (o *InlineResponse2004States) GetStateName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004States) GetStateNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StateName, true
}

// SetStateName sets field value
func (o *InlineResponse2004States) SetStateName(v string) {
	o.StateName = v
}

// GetStateNameL returns the StateNameL field value if set, zero value otherwise.
func (o *InlineResponse2004States) GetStateNameL() string {
	if o == nil || o.StateNameL == nil {
		var ret string
		return ret
	}
	return *o.StateNameL
}

// GetStateNameLOk returns a tuple with the StateNameL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004States) GetStateNameLOk() (*string, bool) {
	if o == nil || o.StateNameL == nil {
		return nil, false
	}
	return o.StateNameL, true
}

// HasStateNameL returns a boolean if a field has been set.
func (o *InlineResponse2004States) HasStateNameL() bool {
	if o != nil && o.StateNameL != nil {
		return true
	}

	return false
}

// SetStateNameL gets a reference to the given string and assigns it to the StateNameL field.
func (o *InlineResponse2004States) SetStateNameL(v string) {
	o.StateNameL = &v
}

func (o InlineResponse2004States) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state_id"] = o.StateId
	}
	if true {
		toSerialize["state_name"] = o.StateName
	}
	if o.StateNameL != nil {
		toSerialize["state_name_l"] = o.StateNameL
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2004States struct {
	value *InlineResponse2004States
	isSet bool
}

func (v NullableInlineResponse2004States) Get() *InlineResponse2004States {
	return v.value
}

func (v *NullableInlineResponse2004States) Set(val *InlineResponse2004States) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2004States) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2004States) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2004States(val *InlineResponse2004States) *NullableInlineResponse2004States {
	return &NullableInlineResponse2004States{value: val, isSet: true}
}

func (v NullableInlineResponse2004States) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2004States) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}