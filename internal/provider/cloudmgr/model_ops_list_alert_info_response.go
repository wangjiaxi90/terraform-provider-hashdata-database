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

// OpsListAlertInfoResponse struct for OpsListAlertInfoResponse
type OpsListAlertInfoResponse struct {
	Content *[]OpsDescribeAlertInfoResponse `json:"content,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewOpsListAlertInfoResponse instantiates a new OpsListAlertInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsListAlertInfoResponse() *OpsListAlertInfoResponse {
	this := OpsListAlertInfoResponse{}
	return &this
}

// NewOpsListAlertInfoResponseWithDefaults instantiates a new OpsListAlertInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsListAlertInfoResponseWithDefaults() *OpsListAlertInfoResponse {
	this := OpsListAlertInfoResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *OpsListAlertInfoResponse) GetContent() []OpsDescribeAlertInfoResponse {
	if o == nil || o.Content == nil {
		var ret []OpsDescribeAlertInfoResponse
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsListAlertInfoResponse) GetContentOk() (*[]OpsDescribeAlertInfoResponse, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *OpsListAlertInfoResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []OpsDescribeAlertInfoResponse and assigns it to the Content field.
func (o *OpsListAlertInfoResponse) SetContent(v []OpsDescribeAlertInfoResponse) {
	o.Content = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *OpsListAlertInfoResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsListAlertInfoResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *OpsListAlertInfoResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *OpsListAlertInfoResponse) SetCount(v int32) {
	o.Count = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OpsListAlertInfoResponse) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsListAlertInfoResponse) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OpsListAlertInfoResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *OpsListAlertInfoResponse) SetTotal(v int32) {
	o.Total = &v
}

func (o OpsListAlertInfoResponse) MarshalJSON() ([]byte, error) {
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

type NullableOpsListAlertInfoResponse struct {
	value *OpsListAlertInfoResponse
	isSet bool
}

func (v NullableOpsListAlertInfoResponse) Get() *OpsListAlertInfoResponse {
	return v.value
}

func (v *NullableOpsListAlertInfoResponse) Set(val *OpsListAlertInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsListAlertInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsListAlertInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsListAlertInfoResponse(val *OpsListAlertInfoResponse) *NullableOpsListAlertInfoResponse {
	return &NullableOpsListAlertInfoResponse{value: val, isSet: true}
}

func (v NullableOpsListAlertInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsListAlertInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


