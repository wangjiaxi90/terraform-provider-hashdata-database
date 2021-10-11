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

// CoreListResourcePoolResponse struct for CoreListResourcePoolResponse
type CoreListResourcePoolResponse struct {
	Content *[]CoreDescribeResourcePoolResponse `json:"content,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewCoreListResourcePoolResponse instantiates a new CoreListResourcePoolResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreListResourcePoolResponse() *CoreListResourcePoolResponse {
	this := CoreListResourcePoolResponse{}
	return &this
}

// NewCoreListResourcePoolResponseWithDefaults instantiates a new CoreListResourcePoolResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreListResourcePoolResponseWithDefaults() *CoreListResourcePoolResponse {
	this := CoreListResourcePoolResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *CoreListResourcePoolResponse) GetContent() []CoreDescribeResourcePoolResponse {
	if o == nil || o.Content == nil {
		var ret []CoreDescribeResourcePoolResponse
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListResourcePoolResponse) GetContentOk() (*[]CoreDescribeResourcePoolResponse, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *CoreListResourcePoolResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []CoreDescribeResourcePoolResponse and assigns it to the Content field.
func (o *CoreListResourcePoolResponse) SetContent(v []CoreDescribeResourcePoolResponse) {
	o.Content = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CoreListResourcePoolResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListResourcePoolResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CoreListResourcePoolResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CoreListResourcePoolResponse) SetCount(v int32) {
	o.Count = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CoreListResourcePoolResponse) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListResourcePoolResponse) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *CoreListResourcePoolResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *CoreListResourcePoolResponse) SetTotal(v int32) {
	o.Total = &v
}

func (o CoreListResourcePoolResponse) MarshalJSON() ([]byte, error) {
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

type NullableCoreListResourcePoolResponse struct {
	value *CoreListResourcePoolResponse
	isSet bool
}

func (v NullableCoreListResourcePoolResponse) Get() *CoreListResourcePoolResponse {
	return v.value
}

func (v *NullableCoreListResourcePoolResponse) Set(val *CoreListResourcePoolResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreListResourcePoolResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreListResourcePoolResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreListResourcePoolResponse(val *CoreListResourcePoolResponse) *NullableCoreListResourcePoolResponse {
	return &NullableCoreListResourcePoolResponse{value: val, isSet: true}
}

func (v NullableCoreListResourcePoolResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreListResourcePoolResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


