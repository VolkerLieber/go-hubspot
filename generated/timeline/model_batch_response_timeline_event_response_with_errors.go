/*
Timeline events

This feature allows an app to create and configure custom events that can show up in the timelines of certain CRM objects like contacts, companies, tickets, or deals. You'll find multiple use cases for this API in the sections below.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package timeline

import (
	"encoding/json"
	"time"
)

// BatchResponseTimelineEventResponseWithErrors struct for BatchResponseTimelineEventResponseWithErrors
type BatchResponseTimelineEventResponseWithErrors struct {
	Status      string                  `json:"status"`
	Results     []TimelineEventResponse `json:"results"`
	NumErrors   *int32                  `json:"numErrors,omitempty"`
	Errors      []StandardError         `json:"errors,omitempty"`
	RequestedAt *time.Time              `json:"requestedAt,omitempty"`
	StartedAt   time.Time               `json:"startedAt"`
	CompletedAt time.Time               `json:"completedAt"`
	Links       *map[string]string      `json:"links,omitempty"`
}

// NewBatchResponseTimelineEventResponseWithErrors instantiates a new BatchResponseTimelineEventResponseWithErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchResponseTimelineEventResponseWithErrors(status string, results []TimelineEventResponse, startedAt time.Time, completedAt time.Time) *BatchResponseTimelineEventResponseWithErrors {
	this := BatchResponseTimelineEventResponseWithErrors{}
	this.Status = status
	this.Results = results
	this.StartedAt = startedAt
	this.CompletedAt = completedAt
	return &this
}

// NewBatchResponseTimelineEventResponseWithErrorsWithDefaults instantiates a new BatchResponseTimelineEventResponseWithErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchResponseTimelineEventResponseWithErrorsWithDefaults() *BatchResponseTimelineEventResponseWithErrors {
	this := BatchResponseTimelineEventResponseWithErrors{}
	return &this
}

// GetStatus returns the Status field value
func (o *BatchResponseTimelineEventResponseWithErrors) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BatchResponseTimelineEventResponseWithErrors) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BatchResponseTimelineEventResponseWithErrors) SetStatus(v string) {
	o.Status = v
}

// GetResults returns the Results field value
func (o *BatchResponseTimelineEventResponseWithErrors) GetResults() []TimelineEventResponse {
	if o == nil {
		var ret []TimelineEventResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BatchResponseTimelineEventResponseWithErrors) GetResultsOk() ([]TimelineEventResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BatchResponseTimelineEventResponseWithErrors) SetResults(v []TimelineEventResponse) {
	o.Results = v
}

// GetNumErrors returns the NumErrors field value if set, zero value otherwise.
func (o *BatchResponseTimelineEventResponseWithErrors) GetNumErrors() int32 {
	if o == nil || o.NumErrors == nil {
		var ret int32
		return ret
	}
	return *o.NumErrors
}

// GetNumErrorsOk returns a tuple with the NumErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseTimelineEventResponseWithErrors) GetNumErrorsOk() (*int32, bool) {
	if o == nil || o.NumErrors == nil {
		return nil, false
	}
	return o.NumErrors, true
}

// HasNumErrors returns a boolean if a field has been set.
func (o *BatchResponseTimelineEventResponseWithErrors) HasNumErrors() bool {
	if o != nil && o.NumErrors != nil {
		return true
	}

	return false
}

// SetNumErrors gets a reference to the given int32 and assigns it to the NumErrors field.
func (o *BatchResponseTimelineEventResponseWithErrors) SetNumErrors(v int32) {
	o.NumErrors = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BatchResponseTimelineEventResponseWithErrors) GetErrors() []StandardError {
	if o == nil || o.Errors == nil {
		var ret []StandardError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseTimelineEventResponseWithErrors) GetErrorsOk() ([]StandardError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BatchResponseTimelineEventResponseWithErrors) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []StandardError and assigns it to the Errors field.
func (o *BatchResponseTimelineEventResponseWithErrors) SetErrors(v []StandardError) {
	o.Errors = v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *BatchResponseTimelineEventResponseWithErrors) GetRequestedAt() time.Time {
	if o == nil || o.RequestedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseTimelineEventResponseWithErrors) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || o.RequestedAt == nil {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *BatchResponseTimelineEventResponseWithErrors) HasRequestedAt() bool {
	if o != nil && o.RequestedAt != nil {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *BatchResponseTimelineEventResponseWithErrors) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetStartedAt returns the StartedAt field value
func (o *BatchResponseTimelineEventResponseWithErrors) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseTimelineEventResponseWithErrors) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *BatchResponseTimelineEventResponseWithErrors) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetCompletedAt returns the CompletedAt field value
func (o *BatchResponseTimelineEventResponseWithErrors) GetCompletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseTimelineEventResponseWithErrors) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedAt, true
}

// SetCompletedAt sets field value
func (o *BatchResponseTimelineEventResponseWithErrors) SetCompletedAt(v time.Time) {
	o.CompletedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchResponseTimelineEventResponseWithErrors) GetLinks() map[string]string {
	if o == nil || o.Links == nil {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseTimelineEventResponseWithErrors) GetLinksOk() (*map[string]string, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchResponseTimelineEventResponseWithErrors) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *BatchResponseTimelineEventResponseWithErrors) SetLinks(v map[string]string) {
	o.Links = &v
}

func (o BatchResponseTimelineEventResponseWithErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["results"] = o.Results
	}
	if o.NumErrors != nil {
		toSerialize["numErrors"] = o.NumErrors
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.RequestedAt != nil {
		toSerialize["requestedAt"] = o.RequestedAt
	}
	if true {
		toSerialize["startedAt"] = o.StartedAt
	}
	if true {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableBatchResponseTimelineEventResponseWithErrors struct {
	value *BatchResponseTimelineEventResponseWithErrors
	isSet bool
}

func (v NullableBatchResponseTimelineEventResponseWithErrors) Get() *BatchResponseTimelineEventResponseWithErrors {
	return v.value
}

func (v *NullableBatchResponseTimelineEventResponseWithErrors) Set(val *BatchResponseTimelineEventResponseWithErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchResponseTimelineEventResponseWithErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchResponseTimelineEventResponseWithErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchResponseTimelineEventResponseWithErrors(val *BatchResponseTimelineEventResponseWithErrors) *NullableBatchResponseTimelineEventResponseWithErrors {
	return &NullableBatchResponseTimelineEventResponseWithErrors{value: val, isSet: true}
}

func (v NullableBatchResponseTimelineEventResponseWithErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchResponseTimelineEventResponseWithErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
