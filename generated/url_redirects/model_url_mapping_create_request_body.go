/*
URL redirects

URL redirect operations

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package url_redirects

import (
	"encoding/json"
)

// UrlMappingCreateRequestBody struct for UrlMappingCreateRequestBody
type UrlMappingCreateRequestBody struct {
	Precedence              *int32 `json:"precedence,omitempty"`
	IsOnlyAfterNotFound     *bool  `json:"isOnlyAfterNotFound,omitempty"`
	IsMatchFullUrl          *bool  `json:"isMatchFullUrl,omitempty"`
	IsMatchQueryString      *bool  `json:"isMatchQueryString,omitempty"`
	IsPattern               *bool  `json:"isPattern,omitempty"`
	IsTrailingSlashOptional *bool  `json:"isTrailingSlashOptional,omitempty"`
	IsProtocolAgnostic      *bool  `json:"isProtocolAgnostic,omitempty"`
	RoutePrefix             string `json:"routePrefix"`
	Destination             string `json:"destination"`
	RedirectStyle           int32  `json:"redirectStyle"`
}

// NewUrlMappingCreateRequestBody instantiates a new UrlMappingCreateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrlMappingCreateRequestBody(routePrefix string, destination string, redirectStyle int32) *UrlMappingCreateRequestBody {
	this := UrlMappingCreateRequestBody{}
	this.RoutePrefix = routePrefix
	this.Destination = destination
	this.RedirectStyle = redirectStyle
	return &this
}

// NewUrlMappingCreateRequestBodyWithDefaults instantiates a new UrlMappingCreateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlMappingCreateRequestBodyWithDefaults() *UrlMappingCreateRequestBody {
	this := UrlMappingCreateRequestBody{}
	return &this
}

// GetPrecedence returns the Precedence field value if set, zero value otherwise.
func (o *UrlMappingCreateRequestBody) GetPrecedence() int32 {
	if o == nil || o.Precedence == nil {
		var ret int32
		return ret
	}
	return *o.Precedence
}

// GetPrecedenceOk returns a tuple with the Precedence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlMappingCreateRequestBody) GetPrecedenceOk() (*int32, bool) {
	if o == nil || o.Precedence == nil {
		return nil, false
	}
	return o.Precedence, true
}

// HasPrecedence returns a boolean if a field has been set.
func (o *UrlMappingCreateRequestBody) HasPrecedence() bool {
	if o != nil && o.Precedence != nil {
		return true
	}

	return false
}

// SetPrecedence gets a reference to the given int32 and assigns it to the Precedence field.
func (o *UrlMappingCreateRequestBody) SetPrecedence(v int32) {
	o.Precedence = &v
}

// GetIsOnlyAfterNotFound returns the IsOnlyAfterNotFound field value if set, zero value otherwise.
func (o *UrlMappingCreateRequestBody) GetIsOnlyAfterNotFound() bool {
	if o == nil || o.IsOnlyAfterNotFound == nil {
		var ret bool
		return ret
	}
	return *o.IsOnlyAfterNotFound
}

// GetIsOnlyAfterNotFoundOk returns a tuple with the IsOnlyAfterNotFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlMappingCreateRequestBody) GetIsOnlyAfterNotFoundOk() (*bool, bool) {
	if o == nil || o.IsOnlyAfterNotFound == nil {
		return nil, false
	}
	return o.IsOnlyAfterNotFound, true
}

// HasIsOnlyAfterNotFound returns a boolean if a field has been set.
func (o *UrlMappingCreateRequestBody) HasIsOnlyAfterNotFound() bool {
	if o != nil && o.IsOnlyAfterNotFound != nil {
		return true
	}

	return false
}

// SetIsOnlyAfterNotFound gets a reference to the given bool and assigns it to the IsOnlyAfterNotFound field.
func (o *UrlMappingCreateRequestBody) SetIsOnlyAfterNotFound(v bool) {
	o.IsOnlyAfterNotFound = &v
}

// GetIsMatchFullUrl returns the IsMatchFullUrl field value if set, zero value otherwise.
func (o *UrlMappingCreateRequestBody) GetIsMatchFullUrl() bool {
	if o == nil || o.IsMatchFullUrl == nil {
		var ret bool
		return ret
	}
	return *o.IsMatchFullUrl
}

// GetIsMatchFullUrlOk returns a tuple with the IsMatchFullUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlMappingCreateRequestBody) GetIsMatchFullUrlOk() (*bool, bool) {
	if o == nil || o.IsMatchFullUrl == nil {
		return nil, false
	}
	return o.IsMatchFullUrl, true
}

// HasIsMatchFullUrl returns a boolean if a field has been set.
func (o *UrlMappingCreateRequestBody) HasIsMatchFullUrl() bool {
	if o != nil && o.IsMatchFullUrl != nil {
		return true
	}

	return false
}

// SetIsMatchFullUrl gets a reference to the given bool and assigns it to the IsMatchFullUrl field.
func (o *UrlMappingCreateRequestBody) SetIsMatchFullUrl(v bool) {
	o.IsMatchFullUrl = &v
}

// GetIsMatchQueryString returns the IsMatchQueryString field value if set, zero value otherwise.
func (o *UrlMappingCreateRequestBody) GetIsMatchQueryString() bool {
	if o == nil || o.IsMatchQueryString == nil {
		var ret bool
		return ret
	}
	return *o.IsMatchQueryString
}

// GetIsMatchQueryStringOk returns a tuple with the IsMatchQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlMappingCreateRequestBody) GetIsMatchQueryStringOk() (*bool, bool) {
	if o == nil || o.IsMatchQueryString == nil {
		return nil, false
	}
	return o.IsMatchQueryString, true
}

// HasIsMatchQueryString returns a boolean if a field has been set.
func (o *UrlMappingCreateRequestBody) HasIsMatchQueryString() bool {
	if o != nil && o.IsMatchQueryString != nil {
		return true
	}

	return false
}

// SetIsMatchQueryString gets a reference to the given bool and assigns it to the IsMatchQueryString field.
func (o *UrlMappingCreateRequestBody) SetIsMatchQueryString(v bool) {
	o.IsMatchQueryString = &v
}

// GetIsPattern returns the IsPattern field value if set, zero value otherwise.
func (o *UrlMappingCreateRequestBody) GetIsPattern() bool {
	if o == nil || o.IsPattern == nil {
		var ret bool
		return ret
	}
	return *o.IsPattern
}

// GetIsPatternOk returns a tuple with the IsPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlMappingCreateRequestBody) GetIsPatternOk() (*bool, bool) {
	if o == nil || o.IsPattern == nil {
		return nil, false
	}
	return o.IsPattern, true
}

// HasIsPattern returns a boolean if a field has been set.
func (o *UrlMappingCreateRequestBody) HasIsPattern() bool {
	if o != nil && o.IsPattern != nil {
		return true
	}

	return false
}

// SetIsPattern gets a reference to the given bool and assigns it to the IsPattern field.
func (o *UrlMappingCreateRequestBody) SetIsPattern(v bool) {
	o.IsPattern = &v
}

// GetIsTrailingSlashOptional returns the IsTrailingSlashOptional field value if set, zero value otherwise.
func (o *UrlMappingCreateRequestBody) GetIsTrailingSlashOptional() bool {
	if o == nil || o.IsTrailingSlashOptional == nil {
		var ret bool
		return ret
	}
	return *o.IsTrailingSlashOptional
}

// GetIsTrailingSlashOptionalOk returns a tuple with the IsTrailingSlashOptional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlMappingCreateRequestBody) GetIsTrailingSlashOptionalOk() (*bool, bool) {
	if o == nil || o.IsTrailingSlashOptional == nil {
		return nil, false
	}
	return o.IsTrailingSlashOptional, true
}

// HasIsTrailingSlashOptional returns a boolean if a field has been set.
func (o *UrlMappingCreateRequestBody) HasIsTrailingSlashOptional() bool {
	if o != nil && o.IsTrailingSlashOptional != nil {
		return true
	}

	return false
}

// SetIsTrailingSlashOptional gets a reference to the given bool and assigns it to the IsTrailingSlashOptional field.
func (o *UrlMappingCreateRequestBody) SetIsTrailingSlashOptional(v bool) {
	o.IsTrailingSlashOptional = &v
}

// GetIsProtocolAgnostic returns the IsProtocolAgnostic field value if set, zero value otherwise.
func (o *UrlMappingCreateRequestBody) GetIsProtocolAgnostic() bool {
	if o == nil || o.IsProtocolAgnostic == nil {
		var ret bool
		return ret
	}
	return *o.IsProtocolAgnostic
}

// GetIsProtocolAgnosticOk returns a tuple with the IsProtocolAgnostic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlMappingCreateRequestBody) GetIsProtocolAgnosticOk() (*bool, bool) {
	if o == nil || o.IsProtocolAgnostic == nil {
		return nil, false
	}
	return o.IsProtocolAgnostic, true
}

// HasIsProtocolAgnostic returns a boolean if a field has been set.
func (o *UrlMappingCreateRequestBody) HasIsProtocolAgnostic() bool {
	if o != nil && o.IsProtocolAgnostic != nil {
		return true
	}

	return false
}

// SetIsProtocolAgnostic gets a reference to the given bool and assigns it to the IsProtocolAgnostic field.
func (o *UrlMappingCreateRequestBody) SetIsProtocolAgnostic(v bool) {
	o.IsProtocolAgnostic = &v
}

// GetRoutePrefix returns the RoutePrefix field value
func (o *UrlMappingCreateRequestBody) GetRoutePrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutePrefix
}

// GetRoutePrefixOk returns a tuple with the RoutePrefix field value
// and a boolean to check if the value has been set.
func (o *UrlMappingCreateRequestBody) GetRoutePrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutePrefix, true
}

// SetRoutePrefix sets field value
func (o *UrlMappingCreateRequestBody) SetRoutePrefix(v string) {
	o.RoutePrefix = v
}

// GetDestination returns the Destination field value
func (o *UrlMappingCreateRequestBody) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *UrlMappingCreateRequestBody) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *UrlMappingCreateRequestBody) SetDestination(v string) {
	o.Destination = v
}

// GetRedirectStyle returns the RedirectStyle field value
func (o *UrlMappingCreateRequestBody) GetRedirectStyle() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RedirectStyle
}

// GetRedirectStyleOk returns a tuple with the RedirectStyle field value
// and a boolean to check if the value has been set.
func (o *UrlMappingCreateRequestBody) GetRedirectStyleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectStyle, true
}

// SetRedirectStyle sets field value
func (o *UrlMappingCreateRequestBody) SetRedirectStyle(v int32) {
	o.RedirectStyle = v
}

func (o UrlMappingCreateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Precedence != nil {
		toSerialize["precedence"] = o.Precedence
	}
	if o.IsOnlyAfterNotFound != nil {
		toSerialize["isOnlyAfterNotFound"] = o.IsOnlyAfterNotFound
	}
	if o.IsMatchFullUrl != nil {
		toSerialize["isMatchFullUrl"] = o.IsMatchFullUrl
	}
	if o.IsMatchQueryString != nil {
		toSerialize["isMatchQueryString"] = o.IsMatchQueryString
	}
	if o.IsPattern != nil {
		toSerialize["isPattern"] = o.IsPattern
	}
	if o.IsTrailingSlashOptional != nil {
		toSerialize["isTrailingSlashOptional"] = o.IsTrailingSlashOptional
	}
	if o.IsProtocolAgnostic != nil {
		toSerialize["isProtocolAgnostic"] = o.IsProtocolAgnostic
	}
	if true {
		toSerialize["routePrefix"] = o.RoutePrefix
	}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if true {
		toSerialize["redirectStyle"] = o.RedirectStyle
	}
	return json.Marshal(toSerialize)
}

type NullableUrlMappingCreateRequestBody struct {
	value *UrlMappingCreateRequestBody
	isSet bool
}

func (v NullableUrlMappingCreateRequestBody) Get() *UrlMappingCreateRequestBody {
	return v.value
}

func (v *NullableUrlMappingCreateRequestBody) Set(val *UrlMappingCreateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUrlMappingCreateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUrlMappingCreateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrlMappingCreateRequestBody(val *UrlMappingCreateRequestBody) *NullableUrlMappingCreateRequestBody {
	return &NullableUrlMappingCreateRequestBody{value: val, isSet: true}
}

func (v NullableUrlMappingCreateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrlMappingCreateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
