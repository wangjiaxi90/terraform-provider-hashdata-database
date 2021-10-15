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

// CoreListWorkloadDetailResponse struct for CoreListWorkloadDetailResponse
type CoreListWorkloadDetailResponse struct {
	Content *[]CoreDescribeWorkloadDetailResponse `json:"content,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewCoreListWorkloadDetailResponse instantiates a new CoreListWorkloadDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreListWorkloadDetailResponse() *CoreListWorkloadDetailResponse {
	this := CoreListWorkloadDetailResponse{}
	return &this
}

// NewCoreListWorkloadDetailResponseWithDefaults instantiates a new CoreListWorkloadDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreListWorkloadDetailResponseWithDefaults() *CoreListWorkloadDetailResponse {
	this := CoreListWorkloadDetailResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *CoreListWorkloadDetailResponse) GetContent() []CoreDescribeWorkloadDetailResponse {
	if o == nil || o.Content == nil {
		var ret []CoreDescribeWorkloadDetailResponse
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListWorkloadDetailResponse) GetContentOk() (*[]CoreDescribeWorkloadDetailResponse, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *CoreListWorkloadDetailResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []CoreDescribeWorkloadDetailResponse and assigns it to the Content field.
func (o *CoreListWorkloadDetailResponse) SetContent(v []CoreDescribeWorkloadDetailResponse) {
	o.Content = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CoreListWorkloadDetailResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListWorkloadDetailResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CoreListWorkloadDetailResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CoreListWorkloadDetailResponse) SetCount(v int32) {
	o.Count = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CoreListWorkloadDetailResponse) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListWorkloadDetailResponse) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *CoreListWorkloadDetailResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *CoreListWorkloadDetailResponse) SetTotal(v int32) {
	o.Total = &v
}

func (o CoreListWorkloadDetailResponse) MarshalJSON() ([]byte, error) {
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

type NullableCoreListWorkloadDetailResponse struct {
	value *CoreListWorkloadDetailResponse
	isSet bool
}

func (v NullableCoreListWorkloadDetailResponse) Get() *CoreListWorkloadDetailResponse {
	return v.value
}

func (v *NullableCoreListWorkloadDetailResponse) Set(val *CoreListWorkloadDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreListWorkloadDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreListWorkloadDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreListWorkloadDetailResponse(val *CoreListWorkloadDetailResponse) *NullableCoreListWorkloadDetailResponse {
	return &NullableCoreListWorkloadDetailResponse{value: val, isSet: true}
}

func (v NullableCoreListWorkloadDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreListWorkloadDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

