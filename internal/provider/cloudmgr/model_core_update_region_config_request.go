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

// CoreUpdateRegionConfigRequest struct for CoreUpdateRegionConfigRequest
type CoreUpdateRegionConfigRequest struct {
	AccessKeyId *string `json:"access_key_id,omitempty"`
	AccessKeySecret *string `json:"access_key_secret,omitempty"`
	Project *string `json:"project,omitempty"`
	RecyclingSecurityGroup *string `json:"recycling_security_group,omitempty"`
	ReverseConnect *bool `json:"reverse_connect,omitempty"`
	Vpc *string `json:"vpc,omitempty"`
}

// NewCoreUpdateRegionConfigRequest instantiates a new CoreUpdateRegionConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUpdateRegionConfigRequest() *CoreUpdateRegionConfigRequest {
	this := CoreUpdateRegionConfigRequest{}
	return &this
}

// NewCoreUpdateRegionConfigRequestWithDefaults instantiates a new CoreUpdateRegionConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUpdateRegionConfigRequestWithDefaults() *CoreUpdateRegionConfigRequest {
	this := CoreUpdateRegionConfigRequest{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *CoreUpdateRegionConfigRequest) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateRegionConfigRequest) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *CoreUpdateRegionConfigRequest) HasAccessKeyId() bool {
	if o != nil && o.AccessKeyId != nil {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *CoreUpdateRegionConfigRequest) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetAccessKeySecret returns the AccessKeySecret field value if set, zero value otherwise.
func (o *CoreUpdateRegionConfigRequest) GetAccessKeySecret() string {
	if o == nil || o.AccessKeySecret == nil {
		var ret string
		return ret
	}
	return *o.AccessKeySecret
}

// GetAccessKeySecretOk returns a tuple with the AccessKeySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateRegionConfigRequest) GetAccessKeySecretOk() (*string, bool) {
	if o == nil || o.AccessKeySecret == nil {
		return nil, false
	}
	return o.AccessKeySecret, true
}

// HasAccessKeySecret returns a boolean if a field has been set.
func (o *CoreUpdateRegionConfigRequest) HasAccessKeySecret() bool {
	if o != nil && o.AccessKeySecret != nil {
		return true
	}

	return false
}

// SetAccessKeySecret gets a reference to the given string and assigns it to the AccessKeySecret field.
func (o *CoreUpdateRegionConfigRequest) SetAccessKeySecret(v string) {
	o.AccessKeySecret = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *CoreUpdateRegionConfigRequest) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateRegionConfigRequest) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *CoreUpdateRegionConfigRequest) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *CoreUpdateRegionConfigRequest) SetProject(v string) {
	o.Project = &v
}

// GetRecyclingSecurityGroup returns the RecyclingSecurityGroup field value if set, zero value otherwise.
func (o *CoreUpdateRegionConfigRequest) GetRecyclingSecurityGroup() string {
	if o == nil || o.RecyclingSecurityGroup == nil {
		var ret string
		return ret
	}
	return *o.RecyclingSecurityGroup
}

// GetRecyclingSecurityGroupOk returns a tuple with the RecyclingSecurityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateRegionConfigRequest) GetRecyclingSecurityGroupOk() (*string, bool) {
	if o == nil || o.RecyclingSecurityGroup == nil {
		return nil, false
	}
	return o.RecyclingSecurityGroup, true
}

// HasRecyclingSecurityGroup returns a boolean if a field has been set.
func (o *CoreUpdateRegionConfigRequest) HasRecyclingSecurityGroup() bool {
	if o != nil && o.RecyclingSecurityGroup != nil {
		return true
	}

	return false
}

// SetRecyclingSecurityGroup gets a reference to the given string and assigns it to the RecyclingSecurityGroup field.
func (o *CoreUpdateRegionConfigRequest) SetRecyclingSecurityGroup(v string) {
	o.RecyclingSecurityGroup = &v
}

// GetReverseConnect returns the ReverseConnect field value if set, zero value otherwise.
func (o *CoreUpdateRegionConfigRequest) GetReverseConnect() bool {
	if o == nil || o.ReverseConnect == nil {
		var ret bool
		return ret
	}
	return *o.ReverseConnect
}

// GetReverseConnectOk returns a tuple with the ReverseConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateRegionConfigRequest) GetReverseConnectOk() (*bool, bool) {
	if o == nil || o.ReverseConnect == nil {
		return nil, false
	}
	return o.ReverseConnect, true
}

// HasReverseConnect returns a boolean if a field has been set.
func (o *CoreUpdateRegionConfigRequest) HasReverseConnect() bool {
	if o != nil && o.ReverseConnect != nil {
		return true
	}

	return false
}

// SetReverseConnect gets a reference to the given bool and assigns it to the ReverseConnect field.
func (o *CoreUpdateRegionConfigRequest) SetReverseConnect(v bool) {
	o.ReverseConnect = &v
}

// GetVpc returns the Vpc field value if set, zero value otherwise.
func (o *CoreUpdateRegionConfigRequest) GetVpc() string {
	if o == nil || o.Vpc == nil {
		var ret string
		return ret
	}
	return *o.Vpc
}

// GetVpcOk returns a tuple with the Vpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateRegionConfigRequest) GetVpcOk() (*string, bool) {
	if o == nil || o.Vpc == nil {
		return nil, false
	}
	return o.Vpc, true
}

// HasVpc returns a boolean if a field has been set.
func (o *CoreUpdateRegionConfigRequest) HasVpc() bool {
	if o != nil && o.Vpc != nil {
		return true
	}

	return false
}

// SetVpc gets a reference to the given string and assigns it to the Vpc field.
func (o *CoreUpdateRegionConfigRequest) SetVpc(v string) {
	o.Vpc = &v
}

func (o CoreUpdateRegionConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKeyId != nil {
		toSerialize["access_key_id"] = o.AccessKeyId
	}
	if o.AccessKeySecret != nil {
		toSerialize["access_key_secret"] = o.AccessKeySecret
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.RecyclingSecurityGroup != nil {
		toSerialize["recycling_security_group"] = o.RecyclingSecurityGroup
	}
	if o.ReverseConnect != nil {
		toSerialize["reverse_connect"] = o.ReverseConnect
	}
	if o.Vpc != nil {
		toSerialize["vpc"] = o.Vpc
	}
	return json.Marshal(toSerialize)
}

type NullableCoreUpdateRegionConfigRequest struct {
	value *CoreUpdateRegionConfigRequest
	isSet bool
}

func (v NullableCoreUpdateRegionConfigRequest) Get() *CoreUpdateRegionConfigRequest {
	return v.value
}

func (v *NullableCoreUpdateRegionConfigRequest) Set(val *CoreUpdateRegionConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUpdateRegionConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUpdateRegionConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUpdateRegionConfigRequest(val *CoreUpdateRegionConfigRequest) *NullableCoreUpdateRegionConfigRequest {
	return &NullableCoreUpdateRegionConfigRequest{value: val, isSet: true}
}

func (v NullableCoreUpdateRegionConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUpdateRegionConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

