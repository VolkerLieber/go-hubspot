/*
Companies

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package companies

import (
	"encoding/json"
)

// BatchInputSimplePublicObjectBatchInput struct for BatchInputSimplePublicObjectBatchInput
type BatchInputSimplePublicObjectBatchInput struct {
	Inputs []SimplePublicObjectBatchInput `json:"inputs"`
}

// NewBatchInputSimplePublicObjectBatchInput instantiates a new BatchInputSimplePublicObjectBatchInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputSimplePublicObjectBatchInput(inputs []SimplePublicObjectBatchInput) *BatchInputSimplePublicObjectBatchInput {
	this := BatchInputSimplePublicObjectBatchInput{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputSimplePublicObjectBatchInputWithDefaults instantiates a new BatchInputSimplePublicObjectBatchInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputSimplePublicObjectBatchInputWithDefaults() *BatchInputSimplePublicObjectBatchInput {
	this := BatchInputSimplePublicObjectBatchInput{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputSimplePublicObjectBatchInput) GetInputs() []SimplePublicObjectBatchInput {
	if o == nil {
		var ret []SimplePublicObjectBatchInput
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputSimplePublicObjectBatchInput) GetInputsOk() ([]SimplePublicObjectBatchInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputSimplePublicObjectBatchInput) SetInputs(v []SimplePublicObjectBatchInput) {
	o.Inputs = v
}

func (o BatchInputSimplePublicObjectBatchInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inputs"] = o.Inputs
	}
	return json.Marshal(toSerialize)
}

type NullableBatchInputSimplePublicObjectBatchInput struct {
	value *BatchInputSimplePublicObjectBatchInput
	isSet bool
}

func (v NullableBatchInputSimplePublicObjectBatchInput) Get() *BatchInputSimplePublicObjectBatchInput {
	return v.value
}

func (v *NullableBatchInputSimplePublicObjectBatchInput) Set(val *BatchInputSimplePublicObjectBatchInput) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputSimplePublicObjectBatchInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputSimplePublicObjectBatchInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputSimplePublicObjectBatchInput(val *BatchInputSimplePublicObjectBatchInput) *NullableBatchInputSimplePublicObjectBatchInput {
	return &NullableBatchInputSimplePublicObjectBatchInput{value: val, isSet: true}
}

func (v NullableBatchInputSimplePublicObjectBatchInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputSimplePublicObjectBatchInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
