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

// OpsListSilenceResponse struct for OpsListSilenceResponse
type OpsListSilenceResponse struct {
	Content *[]OpsDescribeSilenceResponse `json:"content,omitempty"`
}

// NewOpsListSilenceResponse instantiates a new OpsListSilenceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsListSilenceResponse() *OpsListSilenceResponse {
	this := OpsListSilenceResponse{}
	return &this
}

// NewOpsListSilenceResponseWithDefaults instantiates a new OpsListSilenceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsListSilenceResponseWithDefaults() *OpsListSilenceResponse {
	this := OpsListSilenceResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *OpsListSilenceResponse) GetContent() []OpsDescribeSilenceResponse {
	if o == nil || o.Content == nil {
		var ret []OpsDescribeSilenceResponse
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsListSilenceResponse) GetContentOk() (*[]OpsDescribeSilenceResponse, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *OpsListSilenceResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []OpsDescribeSilenceResponse and assigns it to the Content field.
func (o *OpsListSilenceResponse) SetContent(v []OpsDescribeSilenceResponse) {
	o.Content = &v
}

func (o OpsListSilenceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableOpsListSilenceResponse struct {
	value *OpsListSilenceResponse
	isSet bool
}

func (v NullableOpsListSilenceResponse) Get() *OpsListSilenceResponse {
	return v.value
}

func (v *NullableOpsListSilenceResponse) Set(val *OpsListSilenceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsListSilenceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsListSilenceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsListSilenceResponse(val *OpsListSilenceResponse) *NullableOpsListSilenceResponse {
	return &NullableOpsListSilenceResponse{value: val, isSet: true}
}

func (v NullableOpsListSilenceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsListSilenceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


