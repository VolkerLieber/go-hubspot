/*
Transactional Email

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transactional

import (
	"encoding/json"
	"time"
)

// SmtpApiTokenView A SMTP API token provides both an ID and password that can be used to send email through the HubSpot SMTP API.
type SmtpApiTokenView struct {
	// User name to log into the HubSpot SMTP server.
	Id string `json:"id"`
	// Email address of the user that sent the token creation request.
	CreatedBy string `json:"createdBy"`
	// Password used to log into the HubSpot SMTP server.
	Password *string `json:"password,omitempty"`
	// Identifier assigned to the campaign provided in the token creation request.
	EmailCampaignId string `json:"emailCampaignId"`
	// Timestamp generated when a token is created.
	CreatedAt time.Time `json:"createdAt"`
	// Indicates whether a contact should be created for email recipients.
	CreateContact bool `json:"createContact"`
	// A name for the campaign tied to the token.
	CampaignName string `json:"campaignName"`
}

// NewSmtpApiTokenView instantiates a new SmtpApiTokenView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmtpApiTokenView(id string, createdBy string, emailCampaignId string, createdAt time.Time, createContact bool, campaignName string) *SmtpApiTokenView {
	this := SmtpApiTokenView{}
	this.Id = id
	this.CreatedBy = createdBy
	this.EmailCampaignId = emailCampaignId
	this.CreatedAt = createdAt
	this.CreateContact = createContact
	this.CampaignName = campaignName
	return &this
}

// NewSmtpApiTokenViewWithDefaults instantiates a new SmtpApiTokenView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmtpApiTokenViewWithDefaults() *SmtpApiTokenView {
	this := SmtpApiTokenView{}
	return &this
}

// GetId returns the Id field value
func (o *SmtpApiTokenView) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SmtpApiTokenView) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SmtpApiTokenView) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *SmtpApiTokenView) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SmtpApiTokenView) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *SmtpApiTokenView) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SmtpApiTokenView) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpApiTokenView) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SmtpApiTokenView) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SmtpApiTokenView) SetPassword(v string) {
	o.Password = &v
}

// GetEmailCampaignId returns the EmailCampaignId field value
func (o *SmtpApiTokenView) GetEmailCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailCampaignId
}

// GetEmailCampaignIdOk returns a tuple with the EmailCampaignId field value
// and a boolean to check if the value has been set.
func (o *SmtpApiTokenView) GetEmailCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailCampaignId, true
}

// SetEmailCampaignId sets field value
func (o *SmtpApiTokenView) SetEmailCampaignId(v string) {
	o.EmailCampaignId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SmtpApiTokenView) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SmtpApiTokenView) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SmtpApiTokenView) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreateContact returns the CreateContact field value
func (o *SmtpApiTokenView) GetCreateContact() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CreateContact
}

// GetCreateContactOk returns a tuple with the CreateContact field value
// and a boolean to check if the value has been set.
func (o *SmtpApiTokenView) GetCreateContactOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateContact, true
}

// SetCreateContact sets field value
func (o *SmtpApiTokenView) SetCreateContact(v bool) {
	o.CreateContact = v
}

// GetCampaignName returns the CampaignName field value
func (o *SmtpApiTokenView) GetCampaignName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignName
}

// GetCampaignNameOk returns a tuple with the CampaignName field value
// and a boolean to check if the value has been set.
func (o *SmtpApiTokenView) GetCampaignNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignName, true
}

// SetCampaignName sets field value
func (o *SmtpApiTokenView) SetCampaignName(v string) {
	o.CampaignName = v
}

func (o SmtpApiTokenView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["emailCampaignId"] = o.EmailCampaignId
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["createContact"] = o.CreateContact
	}
	if true {
		toSerialize["campaignName"] = o.CampaignName
	}
	return json.Marshal(toSerialize)
}

type NullableSmtpApiTokenView struct {
	value *SmtpApiTokenView
	isSet bool
}

func (v NullableSmtpApiTokenView) Get() *SmtpApiTokenView {
	return v.value
}

func (v *NullableSmtpApiTokenView) Set(val *SmtpApiTokenView) {
	v.value = val
	v.isSet = true
}

func (v NullableSmtpApiTokenView) IsSet() bool {
	return v.isSet
}

func (v *NullableSmtpApiTokenView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmtpApiTokenView(val *SmtpApiTokenView) *NullableSmtpApiTokenView {
	return &NullableSmtpApiTokenView{value: val, isSet: true}
}

func (v NullableSmtpApiTokenView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmtpApiTokenView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
