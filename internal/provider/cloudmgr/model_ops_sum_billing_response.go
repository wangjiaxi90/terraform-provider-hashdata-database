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

// OpsSumBillingResponse struct for OpsSumBillingResponse
type OpsSumBillingResponse struct {
	SumFee *string `json:"sum_fee,omitempty"`
}

// NewOpsSumBillingResponse instantiates a new OpsSumBillingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsSumBillingResponse() *OpsSumBillingResponse {
	this := OpsSumBillingResponse{}
	return &this
}

// NewOpsSumBillingResponseWithDefaults instantiates a new OpsSumBillingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsSumBillingResponseWithDefaults() *OpsSumBillingResponse {
	this := OpsSumBillingResponse{}
	return &this
}

// GetSumFee returns the SumFee field value if set, zero value otherwise.
func (o *OpsSumBillingResponse) GetSumFee() string {
	if o == nil || o.SumFee == nil {
		var ret string
		return ret
	}
	return *o.SumFee
}

// GetSumFeeOk returns a tuple with the SumFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsSumBillingResponse) GetSumFeeOk() (*string, bool) {
	if o == nil || o.SumFee == nil {
		return nil, false
	}
	return o.SumFee, true
}

// HasSumFee returns a boolean if a field has been set.
func (o *OpsSumBillingResponse) HasSumFee() bool {
	if o != nil && o.SumFee != nil {
		return true
	}

	return false
}

// SetSumFee gets a reference to the given string and assigns it to the SumFee field.
func (o *OpsSumBillingResponse) SetSumFee(v string) {
	o.SumFee = &v
}

func (o OpsSumBillingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SumFee != nil {
		toSerialize["sum_fee"] = o.SumFee
	}
	return json.Marshal(toSerialize)
}

type NullableOpsSumBillingResponse struct {
	value *OpsSumBillingResponse
	isSet bool
}

func (v NullableOpsSumBillingResponse) Get() *OpsSumBillingResponse {
	return v.value
}

func (v *NullableOpsSumBillingResponse) Set(val *OpsSumBillingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsSumBillingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsSumBillingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsSumBillingResponse(val *OpsSumBillingResponse) *NullableOpsSumBillingResponse {
	return &NullableOpsSumBillingResponse{value: val, isSet: true}
}

func (v NullableOpsSumBillingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsSumBillingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


