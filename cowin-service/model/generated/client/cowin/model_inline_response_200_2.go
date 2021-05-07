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

// InlineResponse2002 struct for InlineResponse2002
type InlineResponse2002 struct {
	Types *[]InlineResponse2002Types `json:"types,omitempty"`
	// Time in hours to cache the data. Client applications can use this data for this many number of hours before calling this API again.
	Ttl *float32 `json:"ttl,omitempty"`
}

// NewInlineResponse2002 instantiates a new InlineResponse2002 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002() *InlineResponse2002 {
	this := InlineResponse2002{}
	return &this
}

// NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002WithDefaults() *InlineResponse2002 {
	this := InlineResponse2002{}
	return &this
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *InlineResponse2002) GetTypes() []InlineResponse2002Types {
	if o == nil || o.Types == nil {
		var ret []InlineResponse2002Types
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002) GetTypesOk() (*[]InlineResponse2002Types, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *InlineResponse2002) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []InlineResponse2002Types and assigns it to the Types field.
func (o *InlineResponse2002) SetTypes(v []InlineResponse2002Types) {
	o.Types = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *InlineResponse2002) GetTtl() float32 {
	if o == nil || o.Ttl == nil {
		var ret float32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002) GetTtlOk() (*float32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *InlineResponse2002) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given float32 and assigns it to the Ttl field.
func (o *InlineResponse2002) SetTtl(v float32) {
	o.Ttl = &v
}

func (o InlineResponse2002) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2002 struct {
	value *InlineResponse2002
	isSet bool
}

func (v NullableInlineResponse2002) Get() *InlineResponse2002 {
	return v.value
}

func (v *NullableInlineResponse2002) Set(val *InlineResponse2002) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002(val *InlineResponse2002) *NullableInlineResponse2002 {
	return &NullableInlineResponse2002{value: val, isSet: true}
}

func (v NullableInlineResponse2002) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
