/*
 * Cloud Manager API
 *
 * Cloud Manager Restful API Documentation.
 *
 * API version: v2.0
 * Contact: wang@hashdata.cn
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudmgr

import (
	"encoding/json"
)

// NotificationDescribeMessageResponse struct for NotificationDescribeMessageResponse
type NotificationDescribeMessageResponse struct {
	Category *string `json:"category,omitempty"`
	Contacts *[]NotificationDescribeContactResponse `json:"contacts,omitempty"`
	Content *string `json:"content,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	Id *string `json:"id,omitempty"`
	SendAt *string `json:"send_at,omitempty"`
	Status *string `json:"status,omitempty"`
	Subject *string `json:"subject,omitempty"`
}

// NewNotificationDescribeMessageResponse instantiates a new NotificationDescribeMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationDescribeMessageResponse() *NotificationDescribeMessageResponse {
	this := NotificationDescribeMessageResponse{}
	return &this
}

// NewNotificationDescribeMessageResponseWithDefaults instantiates a new NotificationDescribeMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationDescribeMessageResponseWithDefaults() *NotificationDescribeMessageResponse {
	this := NotificationDescribeMessageResponse{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *NotificationDescribeMessageResponse) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDescribeMessageResponse) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *NotificationDescribeMessageResponse) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *NotificationDescribeMessageResponse) SetCategory(v string) {
	o.Category = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *NotificationDescribeMessageResponse) GetContacts() []NotificationDescribeContactResponse {
	if o == nil || o.Contacts == nil {
		var ret []NotificationDescribeContactResponse
		return ret
	}
	return *o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDescribeMessageResponse) GetContactsOk() (*[]NotificationDescribeContactResponse, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *NotificationDescribeMessageResponse) HasContacts() bool {
	if o != nil && o.Contacts != nil {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []NotificationDescribeContactResponse and assigns it to the Contacts field.
func (o *NotificationDescribeMessageResponse) SetContacts(v []NotificationDescribeContactResponse) {
	o.Contacts = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *NotificationDescribeMessageResponse) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDescribeMessageResponse) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *NotificationDescribeMessageResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *NotificationDescribeMessageResponse) SetContent(v string) {
	o.Content = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationDescribeMessageResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDescribeMessageResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationDescribeMessageResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *NotificationDescribeMessageResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *NotificationDescribeMessageResponse) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDescribeMessageResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *NotificationDescribeMessageResponse) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *NotificationDescribeMessageResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationDescribeMessageResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDescribeMessageResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationDescribeMessageResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NotificationDescribeMessageResponse) SetId(v string) {
	o.Id = &v
}

// GetSendAt returns the SendAt field value if set, zero value otherwise.
func (o *NotificationDescribeMessageResponse) GetSendAt() string {
	if o == nil || o.SendAt == nil {
		var ret string
		return ret
	}
	return *o.SendAt
}

// GetSendAtOk returns a tuple with the SendAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDescribeMessageResponse) GetSendAtOk() (*string, bool) {
	if o == nil || o.SendAt == nil {
		return nil, false
	}
	return o.SendAt, true
}

// HasSendAt returns a boolean if a field has been set.
func (o *NotificationDescribeMessageResponse) HasSendAt() bool {
	if o != nil && o.SendAt != nil {
		return true
	}

	return false
}

// SetSendAt gets a reference to the given string and assigns it to the SendAt field.
func (o *NotificationDescribeMessageResponse) SetSendAt(v string) {
	o.SendAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NotificationDescribeMessageResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDescribeMessageResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NotificationDescribeMessageResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NotificationDescribeMessageResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *NotificationDescribeMessageResponse) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDescribeMessageResponse) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *NotificationDescribeMessageResponse) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *NotificationDescribeMessageResponse) SetSubject(v string) {
	o.Subject = &v
}

func (o NotificationDescribeMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SendAt != nil {
		toSerialize["send_at"] = o.SendAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationDescribeMessageResponse struct {
	value *NotificationDescribeMessageResponse
	isSet bool
}

func (v NullableNotificationDescribeMessageResponse) Get() *NotificationDescribeMessageResponse {
	return v.value
}

func (v *NullableNotificationDescribeMessageResponse) Set(val *NotificationDescribeMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationDescribeMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationDescribeMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationDescribeMessageResponse(val *NotificationDescribeMessageResponse) *NullableNotificationDescribeMessageResponse {
	return &NullableNotificationDescribeMessageResponse{value: val, isSet: true}
}

func (v NullableNotificationDescribeMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationDescribeMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


