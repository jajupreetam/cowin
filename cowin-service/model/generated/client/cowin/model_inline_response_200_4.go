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

// InlineResponse2004 struct for InlineResponse2004
type InlineResponse2004 struct {
	States *[]InlineResponse2004States `json:"states,omitempty"`
	// Time in hours to cache the data. Client applications can use this data for this many number of hours before calling this API again.
	Ttl *float32 `json:"ttl,omitempty"`
}

// NewInlineResponse2004 instantiates a new InlineResponse2004 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2004() *InlineResponse2004 {
	this := InlineResponse2004{}
	return &this
}

// NewInlineResponse2004WithDefaults instantiates a new InlineResponse2004 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2004WithDefaults() *InlineResponse2004 {
	this := InlineResponse2004{}
	return &this
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *InlineResponse2004) GetStates() []InlineResponse2004States {
	if o == nil || o.States == nil {
		var ret []InlineResponse2004States
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetStatesOk() (*[]InlineResponse2004States, bool) {
	if o == nil || o.States == nil {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *InlineResponse2004) HasStates() bool {
	if o != nil && o.States != nil {
		return true
	}

	return false
}

// SetStates gets a reference to the given []InlineResponse2004States and assigns it to the States field.
func (o *InlineResponse2004) SetStates(v []InlineResponse2004States) {
	o.States = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *InlineResponse2004) GetTtl() float32 {
	if o == nil || o.Ttl == nil {
		var ret float32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetTtlOk() (*float32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *InlineResponse2004) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given float32 and assigns it to the Ttl field.
func (o *InlineResponse2004) SetTtl(v float32) {
	o.Ttl = &v
}

func (o InlineResponse2004) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.States != nil {
		toSerialize["states"] = o.States
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2004 struct {
	value *InlineResponse2004
	isSet bool
}

func (v NullableInlineResponse2004) Get() *InlineResponse2004 {
	return v.value
}

func (v *NullableInlineResponse2004) Set(val *InlineResponse2004) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2004) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2004) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2004(val *InlineResponse2004) *NullableInlineResponse2004 {
	return &NullableInlineResponse2004{value: val, isSet: true}
}

func (v NullableInlineResponse2004) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2004) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}