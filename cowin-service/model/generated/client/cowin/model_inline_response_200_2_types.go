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

// InlineResponse2002Types struct for InlineResponse2002Types
type InlineResponse2002Types struct {
	Type string `json:"type"`
	// Type in preferred language as specified in Accept-Language header parameter.
	TypeL *string `json:"type_l,omitempty"`
	Id float32 `json:"id"`
}

// NewInlineResponse2002Types instantiates a new InlineResponse2002Types object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002Types(type_ string, id float32, ) *InlineResponse2002Types {
	this := InlineResponse2002Types{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewInlineResponse2002TypesWithDefaults instantiates a new InlineResponse2002Types object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002TypesWithDefaults() *InlineResponse2002Types {
	this := InlineResponse2002Types{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse2002Types) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Types) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse2002Types) SetType(v string) {
	o.Type = v
}

// GetTypeL returns the TypeL field value if set, zero value otherwise.
func (o *InlineResponse2002Types) GetTypeL() string {
	if o == nil || o.TypeL == nil {
		var ret string
		return ret
	}
	return *o.TypeL
}

// GetTypeLOk returns a tuple with the TypeL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Types) GetTypeLOk() (*string, bool) {
	if o == nil || o.TypeL == nil {
		return nil, false
	}
	return o.TypeL, true
}

// HasTypeL returns a boolean if a field has been set.
func (o *InlineResponse2002Types) HasTypeL() bool {
	if o != nil && o.TypeL != nil {
		return true
	}

	return false
}

// SetTypeL gets a reference to the given string and assigns it to the TypeL field.
func (o *InlineResponse2002Types) SetTypeL(v string) {
	o.TypeL = &v
}

// GetId returns the Id field value
func (o *InlineResponse2002Types) GetId() float32 {
	if o == nil  {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Types) GetIdOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse2002Types) SetId(v float32) {
	o.Id = v
}

func (o InlineResponse2002Types) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.TypeL != nil {
		toSerialize["type_l"] = o.TypeL
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2002Types struct {
	value *InlineResponse2002Types
	isSet bool
}

func (v NullableInlineResponse2002Types) Get() *InlineResponse2002Types {
	return v.value
}

func (v *NullableInlineResponse2002Types) Set(val *InlineResponse2002Types) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002Types) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002Types) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002Types(val *InlineResponse2002Types) *NullableInlineResponse2002Types {
	return &NullableInlineResponse2002Types{value: val, isSet: true}
}

func (v NullableInlineResponse2002Types) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002Types) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}