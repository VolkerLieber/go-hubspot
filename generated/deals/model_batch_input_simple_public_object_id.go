/*
Deals

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deals

import (
	"encoding/json"
)

// BatchInputSimplePublicObjectId struct for BatchInputSimplePublicObjectId
type BatchInputSimplePublicObjectId struct {
	Inputs []SimplePublicObjectId `json:"inputs"`
}

// NewBatchInputSimplePublicObjectId instantiates a new BatchInputSimplePublicObjectId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputSimplePublicObjectId(inputs []SimplePublicObjectId) *BatchInputSimplePublicObjectId {
	this := BatchInputSimplePublicObjectId{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputSimplePublicObjectIdWithDefaults instantiates a new BatchInputSimplePublicObjectId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputSimplePublicObjectIdWithDefaults() *BatchInputSimplePublicObjectId {
	this := BatchInputSimplePublicObjectId{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputSimplePublicObjectId) GetInputs() []SimplePublicObjectId {
	if o == nil {
		var ret []SimplePublicObjectId
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputSimplePublicObjectId) GetInputsOk() ([]SimplePublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputSimplePublicObjectId) SetInputs(v []SimplePublicObjectId) {
	o.Inputs = v
}

func (o BatchInputSimplePublicObjectId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inputs"] = o.Inputs
	}
	return json.Marshal(toSerialize)
}

type NullableBatchInputSimplePublicObjectId struct {
	value *BatchInputSimplePublicObjectId
	isSet bool
}

func (v NullableBatchInputSimplePublicObjectId) Get() *BatchInputSimplePublicObjectId {
	return v.value
}

func (v *NullableBatchInputSimplePublicObjectId) Set(val *BatchInputSimplePublicObjectId) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputSimplePublicObjectId) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputSimplePublicObjectId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputSimplePublicObjectId(val *BatchInputSimplePublicObjectId) *NullableBatchInputSimplePublicObjectId {
	return &NullableBatchInputSimplePublicObjectId{value: val, isSet: true}
}

func (v NullableBatchInputSimplePublicObjectId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputSimplePublicObjectId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
