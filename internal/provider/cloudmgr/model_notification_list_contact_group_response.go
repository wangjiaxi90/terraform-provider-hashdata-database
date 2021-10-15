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

// NotificationListContactGroupResponse struct for NotificationListContactGroupResponse
type NotificationListContactGroupResponse struct {
	Content *[]NotificationDescribeContactGroupResponse `json:"content,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewNotificationListContactGroupResponse instantiates a new NotificationListContactGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationListContactGroupResponse() *NotificationListContactGroupResponse {
	this := NotificationListContactGroupResponse{}
	return &this
}

// NewNotificationListContactGroupResponseWithDefaults instantiates a new NotificationListContactGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationListContactGroupResponseWithDefaults() *NotificationListContactGroupResponse {
	this := NotificationListContactGroupResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *NotificationListContactGroupResponse) GetContent() []NotificationDescribeContactGroupResponse {
	if o == nil || o.Content == nil {
		var ret []NotificationDescribeContactGroupResponse
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationListContactGroupResponse) GetContentOk() (*[]NotificationDescribeContactGroupResponse, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *NotificationListContactGroupResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []NotificationDescribeContactGroupResponse and assigns it to the Content field.
func (o *NotificationListContactGroupResponse) SetContent(v []NotificationDescribeContactGroupResponse) {
	o.Content = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *NotificationListContactGroupResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationListContactGroupResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *NotificationListContactGroupResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *NotificationListContactGroupResponse) SetCount(v int32) {
	o.Count = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *NotificationListContactGroupResponse) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationListContactGroupResponse) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *NotificationListContactGroupResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *NotificationListContactGroupResponse) SetTotal(v int32) {
	o.Total = &v
}

func (o NotificationListContactGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationListContactGroupResponse struct {
	value *NotificationListContactGroupResponse
	isSet bool
}

func (v NullableNotificationListContactGroupResponse) Get() *NotificationListContactGroupResponse {
	return v.value
}

func (v *NullableNotificationListContactGroupResponse) Set(val *NotificationListContactGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationListContactGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationListContactGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationListContactGroupResponse(val *NotificationListContactGroupResponse) *NullableNotificationListContactGroupResponse {
	return &NullableNotificationListContactGroupResponse{value: val, isSet: true}
}

func (v NullableNotificationListContactGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationListContactGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

