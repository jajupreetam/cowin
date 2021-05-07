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

// InlineResponse2005 struct for InlineResponse2005
type InlineResponse2005 struct {
	Districts *[]InlineResponse2005Districts `json:"districts,omitempty"`
	// Time in hours to cache the data. Client applications can use this data for this many number of hours before calling this API again.
	Ttl *float32 `json:"ttl,omitempty"`
}

// NewInlineResponse2005 instantiates a new InlineResponse2005 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2005() *InlineResponse2005 {
	this := InlineResponse2005{}
	return &this
}

// NewInlineResponse2005WithDefaults instantiates a new InlineResponse2005 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2005WithDefaults() *InlineResponse2005 {
	this := InlineResponse2005{}
	return &this
}

// GetDistricts returns the Districts field value if set, zero value otherwise.
func (o *InlineResponse2005) GetDistricts() []InlineResponse2005Districts {
	if o == nil || o.Districts == nil {
		var ret []InlineResponse2005Districts
		return ret
	}
	return *o.Districts
}

// GetDistrictsOk returns a tuple with the Districts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005) GetDistrictsOk() (*[]InlineResponse2005Districts, bool) {
	if o == nil || o.Districts == nil {
		return nil, false
	}
	return o.Districts, true
}

// HasDistricts returns a boolean if a field has been set.
func (o *InlineResponse2005) HasDistricts() bool {
	if o != nil && o.Districts != nil {
		return true
	}

	return false
}

// SetDistricts gets a reference to the given []InlineResponse2005Districts and assigns it to the Districts field.
func (o *InlineResponse2005) SetDistricts(v []InlineResponse2005Districts) {
	o.Districts = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *InlineResponse2005) GetTtl() float32 {
	if o == nil || o.Ttl == nil {
		var ret float32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005) GetTtlOk() (*float32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *InlineResponse2005) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given float32 and assigns it to the Ttl field.
func (o *InlineResponse2005) SetTtl(v float32) {
	o.Ttl = &v
}

func (o InlineResponse2005) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Districts != nil {
		toSerialize["districts"] = o.Districts
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2005 struct {
	value *InlineResponse2005
	isSet bool
}

func (v NullableInlineResponse2005) Get() *InlineResponse2005 {
	return v.value
}

func (v *NullableInlineResponse2005) Set(val *InlineResponse2005) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2005) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2005) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2005(val *InlineResponse2005) *NullableInlineResponse2005 {
	return &NullableInlineResponse2005{value: val, isSet: true}
}

func (v NullableInlineResponse2005) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2005) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
