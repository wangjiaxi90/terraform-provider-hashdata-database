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

// CoreCreateServiceFeatureRequest struct for CoreCreateServiceFeatureRequest
type CoreCreateServiceFeatureRequest struct {
	LocalStorage *bool `json:"local_storage,omitempty"`
	MirrorStandby *bool `json:"mirror_standby,omitempty"`
}

// NewCoreCreateServiceFeatureRequest instantiates a new CoreCreateServiceFeatureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCreateServiceFeatureRequest() *CoreCreateServiceFeatureRequest {
	this := CoreCreateServiceFeatureRequest{}
	return &this
}

// NewCoreCreateServiceFeatureRequestWithDefaults instantiates a new CoreCreateServiceFeatureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCreateServiceFeatureRequestWithDefaults() *CoreCreateServiceFeatureRequest {
	this := CoreCreateServiceFeatureRequest{}
	return &this
}

// GetLocalStorage returns the LocalStorage field value if set, zero value otherwise.
func (o *CoreCreateServiceFeatureRequest) GetLocalStorage() bool {
	if o == nil || o.LocalStorage == nil {
		var ret bool
		return ret
	}
	return *o.LocalStorage
}

// GetLocalStorageOk returns a tuple with the LocalStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateServiceFeatureRequest) GetLocalStorageOk() (*bool, bool) {
	if o == nil || o.LocalStorage == nil {
		return nil, false
	}
	return o.LocalStorage, true
}

// HasLocalStorage returns a boolean if a field has been set.
func (o *CoreCreateServiceFeatureRequest) HasLocalStorage() bool {
	if o != nil && o.LocalStorage != nil {
		return true
	}

	return false
}

// SetLocalStorage gets a reference to the given bool and assigns it to the LocalStorage field.
func (o *CoreCreateServiceFeatureRequest) SetLocalStorage(v bool) {
	o.LocalStorage = &v
}

// GetMirrorStandby returns the MirrorStandby field value if set, zero value otherwise.
func (o *CoreCreateServiceFeatureRequest) GetMirrorStandby() bool {
	if o == nil || o.MirrorStandby == nil {
		var ret bool
		return ret
	}
	return *o.MirrorStandby
}

// GetMirrorStandbyOk returns a tuple with the MirrorStandby field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateServiceFeatureRequest) GetMirrorStandbyOk() (*bool, bool) {
	if o == nil || o.MirrorStandby == nil {
		return nil, false
	}
	return o.MirrorStandby, true
}

// HasMirrorStandby returns a boolean if a field has been set.
func (o *CoreCreateServiceFeatureRequest) HasMirrorStandby() bool {
	if o != nil && o.MirrorStandby != nil {
		return true
	}

	return false
}

// SetMirrorStandby gets a reference to the given bool and assigns it to the MirrorStandby field.
func (o *CoreCreateServiceFeatureRequest) SetMirrorStandby(v bool) {
	o.MirrorStandby = &v
}

func (o CoreCreateServiceFeatureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LocalStorage != nil {
		toSerialize["local_storage"] = o.LocalStorage
	}
	if o.MirrorStandby != nil {
		toSerialize["mirror_standby"] = o.MirrorStandby
	}
	return json.Marshal(toSerialize)
}

type NullableCoreCreateServiceFeatureRequest struct {
	value *CoreCreateServiceFeatureRequest
	isSet bool
}

func (v NullableCoreCreateServiceFeatureRequest) Get() *CoreCreateServiceFeatureRequest {
	return v.value
}

func (v *NullableCoreCreateServiceFeatureRequest) Set(val *CoreCreateServiceFeatureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCreateServiceFeatureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCreateServiceFeatureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCreateServiceFeatureRequest(val *CoreCreateServiceFeatureRequest) *NullableCoreCreateServiceFeatureRequest {
	return &NullableCoreCreateServiceFeatureRequest{value: val, isSet: true}
}

func (v NullableCoreCreateServiceFeatureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCreateServiceFeatureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

