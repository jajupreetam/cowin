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

// InlineObject3 struct for InlineObject3
type InlineObject3 struct {
	BeneficiaryReferenceId *string `json:"beneficiary_reference_id,omitempty"`
}

// NewInlineObject3 instantiates a new InlineObject3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject3() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// NewInlineObject3WithDefaults instantiates a new InlineObject3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject3WithDefaults() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// GetBeneficiaryReferenceId returns the BeneficiaryReferenceId field value if set, zero value otherwise.
func (o *InlineObject3) GetBeneficiaryReferenceId() string {
	if o == nil || o.BeneficiaryReferenceId == nil {
		var ret string
		return ret
	}
	return *o.BeneficiaryReferenceId
}

// GetBeneficiaryReferenceIdOk returns a tuple with the BeneficiaryReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetBeneficiaryReferenceIdOk() (*string, bool) {
	if o == nil || o.BeneficiaryReferenceId == nil {
		return nil, false
	}
	return o.BeneficiaryReferenceId, true
}

// HasBeneficiaryReferenceId returns a boolean if a field has been set.
func (o *InlineObject3) HasBeneficiaryReferenceId() bool {
	if o != nil && o.BeneficiaryReferenceId != nil {
		return true
	}

	return false
}

// SetBeneficiaryReferenceId gets a reference to the given string and assigns it to the BeneficiaryReferenceId field.
func (o *InlineObject3) SetBeneficiaryReferenceId(v string) {
	o.BeneficiaryReferenceId = &v
}

func (o InlineObject3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BeneficiaryReferenceId != nil {
		toSerialize["beneficiary_reference_id"] = o.BeneficiaryReferenceId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject3 struct {
	value *InlineObject3
	isSet bool
}

func (v NullableInlineObject3) Get() *InlineObject3 {
	return v.value
}

func (v *NullableInlineObject3) Set(val *InlineObject3) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject3) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject3(val *InlineObject3) *NullableInlineObject3 {
	return &NullableInlineObject3{value: val, isSet: true}
}

func (v NullableInlineObject3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}