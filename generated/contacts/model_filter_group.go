/*
Contacts

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package contacts

import (
	"encoding/json"
)

// FilterGroup struct for FilterGroup
type FilterGroup struct {
	Filters []Filter `json:"filters"`
}

// NewFilterGroup instantiates a new FilterGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterGroup(filters []Filter) *FilterGroup {
	this := FilterGroup{}
	this.Filters = filters
	return &this
}

// NewFilterGroupWithDefaults instantiates a new FilterGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterGroupWithDefaults() *FilterGroup {
	this := FilterGroup{}
	return &this
}

// GetFilters returns the Filters field value
func (o *FilterGroup) GetFilters() []Filter {
	if o == nil {
		var ret []Filter
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *FilterGroup) GetFiltersOk() ([]Filter, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *FilterGroup) SetFilters(v []Filter) {
	o.Filters = v
}

func (o FilterGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableFilterGroup struct {
	value *FilterGroup
	isSet bool
}

func (v NullableFilterGroup) Get() *FilterGroup {
	return v.value
}

func (v *NullableFilterGroup) Set(val *FilterGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterGroup(val *FilterGroup) *NullableFilterGroup {
	return &NullableFilterGroup{value: val, isSet: true}
}

func (v NullableFilterGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
