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

// KsyunDescribeResourceUsageResponse struct for KsyunDescribeResourceUsageResponse
type KsyunDescribeResourceUsageResponse struct {
	Message *string `json:"message,omitempty"`
	Result *[]CommonIaasResource `json:"result,omitempty"`
	Status *int32 `json:"status,omitempty"`
}

// NewKsyunDescribeResourceUsageResponse instantiates a new KsyunDescribeResourceUsageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKsyunDescribeResourceUsageResponse() *KsyunDescribeResourceUsageResponse {
	this := KsyunDescribeResourceUsageResponse{}
	return &this
}

// NewKsyunDescribeResourceUsageResponseWithDefaults instantiates a new KsyunDescribeResourceUsageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKsyunDescribeResourceUsageResponseWithDefaults() *KsyunDescribeResourceUsageResponse {
	this := KsyunDescribeResourceUsageResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *KsyunDescribeResourceUsageResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsyunDescribeResourceUsageResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *KsyunDescribeResourceUsageResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *KsyunDescribeResourceUsageResponse) SetMessage(v string) {
	o.Message = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *KsyunDescribeResourceUsageResponse) GetResult() []CommonIaasResource {
	if o == nil || o.Result == nil {
		var ret []CommonIaasResource
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsyunDescribeResourceUsageResponse) GetResultOk() (*[]CommonIaasResource, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *KsyunDescribeResourceUsageResponse) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given []CommonIaasResource and assigns it to the Result field.
func (o *KsyunDescribeResourceUsageResponse) SetResult(v []CommonIaasResource) {
	o.Result = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KsyunDescribeResourceUsageResponse) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsyunDescribeResourceUsageResponse) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KsyunDescribeResourceUsageResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KsyunDescribeResourceUsageResponse) SetStatus(v int32) {
	o.Status = &v
}

func (o KsyunDescribeResourceUsageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableKsyunDescribeResourceUsageResponse struct {
	value *KsyunDescribeResourceUsageResponse
	isSet bool
}

func (v NullableKsyunDescribeResourceUsageResponse) Get() *KsyunDescribeResourceUsageResponse {
	return v.value
}

func (v *NullableKsyunDescribeResourceUsageResponse) Set(val *KsyunDescribeResourceUsageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKsyunDescribeResourceUsageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKsyunDescribeResourceUsageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKsyunDescribeResourceUsageResponse(val *KsyunDescribeResourceUsageResponse) *NullableKsyunDescribeResourceUsageResponse {
	return &NullableKsyunDescribeResourceUsageResponse{value: val, isSet: true}
}

func (v NullableKsyunDescribeResourceUsageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKsyunDescribeResourceUsageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


