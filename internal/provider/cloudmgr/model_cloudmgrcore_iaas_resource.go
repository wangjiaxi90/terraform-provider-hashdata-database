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

// CloudmgrcoreIaasResource struct for CloudmgrcoreIaasResource
type CloudmgrcoreIaasResource struct {
	Count *int32 `json:"count,omitempty"`
	ElasticNic *CoreCreateElasticNicRequest `json:"elastic_nic,omitempty"`
	Image *string `json:"image,omitempty"`
	InstanceType *string `json:"instance_type,omitempty"`
	InternetAccess *CoreCreateInternetAccessRequest `json:"internet_access,omitempty"`
	VolumeSize *int32 `json:"volume_size,omitempty"`
	VolumeType *string `json:"volume_type,omitempty"`
	Zone *string `json:"zone,omitempty"`
}

// NewCloudmgrcoreIaasResource instantiates a new CloudmgrcoreIaasResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudmgrcoreIaasResource() *CloudmgrcoreIaasResource {
	this := CloudmgrcoreIaasResource{}
	return &this
}

// NewCloudmgrcoreIaasResourceWithDefaults instantiates a new CloudmgrcoreIaasResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudmgrcoreIaasResourceWithDefaults() *CloudmgrcoreIaasResource {
	this := CloudmgrcoreIaasResource{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CloudmgrcoreIaasResource) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcoreIaasResource) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CloudmgrcoreIaasResource) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CloudmgrcoreIaasResource) SetCount(v int32) {
	o.Count = &v
}

// GetElasticNic returns the ElasticNic field value if set, zero value otherwise.
func (o *CloudmgrcoreIaasResource) GetElasticNic() CoreCreateElasticNicRequest {
	if o == nil || o.ElasticNic == nil {
		var ret CoreCreateElasticNicRequest
		return ret
	}
	return *o.ElasticNic
}

// GetElasticNicOk returns a tuple with the ElasticNic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcoreIaasResource) GetElasticNicOk() (*CoreCreateElasticNicRequest, bool) {
	if o == nil || o.ElasticNic == nil {
		return nil, false
	}
	return o.ElasticNic, true
}

// HasElasticNic returns a boolean if a field has been set.
func (o *CloudmgrcoreIaasResource) HasElasticNic() bool {
	if o != nil && o.ElasticNic != nil {
		return true
	}

	return false
}

// SetElasticNic gets a reference to the given CoreCreateElasticNicRequest and assigns it to the ElasticNic field.
func (o *CloudmgrcoreIaasResource) SetElasticNic(v CoreCreateElasticNicRequest) {
	o.ElasticNic = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CloudmgrcoreIaasResource) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcoreIaasResource) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CloudmgrcoreIaasResource) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *CloudmgrcoreIaasResource) SetImage(v string) {
	o.Image = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *CloudmgrcoreIaasResource) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcoreIaasResource) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *CloudmgrcoreIaasResource) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *CloudmgrcoreIaasResource) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetInternetAccess returns the InternetAccess field value if set, zero value otherwise.
func (o *CloudmgrcoreIaasResource) GetInternetAccess() CoreCreateInternetAccessRequest {
	if o == nil || o.InternetAccess == nil {
		var ret CoreCreateInternetAccessRequest
		return ret
	}
	return *o.InternetAccess
}

// GetInternetAccessOk returns a tuple with the InternetAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcoreIaasResource) GetInternetAccessOk() (*CoreCreateInternetAccessRequest, bool) {
	if o == nil || o.InternetAccess == nil {
		return nil, false
	}
	return o.InternetAccess, true
}

// HasInternetAccess returns a boolean if a field has been set.
func (o *CloudmgrcoreIaasResource) HasInternetAccess() bool {
	if o != nil && o.InternetAccess != nil {
		return true
	}

	return false
}

// SetInternetAccess gets a reference to the given CoreCreateInternetAccessRequest and assigns it to the InternetAccess field.
func (o *CloudmgrcoreIaasResource) SetInternetAccess(v CoreCreateInternetAccessRequest) {
	o.InternetAccess = &v
}

// GetVolumeSize returns the VolumeSize field value if set, zero value otherwise.
func (o *CloudmgrcoreIaasResource) GetVolumeSize() int32 {
	if o == nil || o.VolumeSize == nil {
		var ret int32
		return ret
	}
	return *o.VolumeSize
}

// GetVolumeSizeOk returns a tuple with the VolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcoreIaasResource) GetVolumeSizeOk() (*int32, bool) {
	if o == nil || o.VolumeSize == nil {
		return nil, false
	}
	return o.VolumeSize, true
}

// HasVolumeSize returns a boolean if a field has been set.
func (o *CloudmgrcoreIaasResource) HasVolumeSize() bool {
	if o != nil && o.VolumeSize != nil {
		return true
	}

	return false
}

// SetVolumeSize gets a reference to the given int32 and assigns it to the VolumeSize field.
func (o *CloudmgrcoreIaasResource) SetVolumeSize(v int32) {
	o.VolumeSize = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *CloudmgrcoreIaasResource) GetVolumeType() string {
	if o == nil || o.VolumeType == nil {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcoreIaasResource) GetVolumeTypeOk() (*string, bool) {
	if o == nil || o.VolumeType == nil {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *CloudmgrcoreIaasResource) HasVolumeType() bool {
	if o != nil && o.VolumeType != nil {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *CloudmgrcoreIaasResource) SetVolumeType(v string) {
	o.VolumeType = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *CloudmgrcoreIaasResource) GetZone() string {
	if o == nil || o.Zone == nil {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcoreIaasResource) GetZoneOk() (*string, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *CloudmgrcoreIaasResource) HasZone() bool {
	if o != nil && o.Zone != nil {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *CloudmgrcoreIaasResource) SetZone(v string) {
	o.Zone = &v
}

func (o CloudmgrcoreIaasResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.ElasticNic != nil {
		toSerialize["elastic_nic"] = o.ElasticNic
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.InstanceType != nil {
		toSerialize["instance_type"] = o.InstanceType
	}
	if o.InternetAccess != nil {
		toSerialize["internet_access"] = o.InternetAccess
	}
	if o.VolumeSize != nil {
		toSerialize["volume_size"] = o.VolumeSize
	}
	if o.VolumeType != nil {
		toSerialize["volume_type"] = o.VolumeType
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}
	return json.Marshal(toSerialize)
}

type NullableCloudmgrcoreIaasResource struct {
	value *CloudmgrcoreIaasResource
	isSet bool
}

func (v NullableCloudmgrcoreIaasResource) Get() *CloudmgrcoreIaasResource {
	return v.value
}

func (v *NullableCloudmgrcoreIaasResource) Set(val *CloudmgrcoreIaasResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudmgrcoreIaasResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudmgrcoreIaasResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudmgrcoreIaasResource(val *CloudmgrcoreIaasResource) *NullableCloudmgrcoreIaasResource {
	return &NullableCloudmgrcoreIaasResource{value: val, isSet: true}
}

func (v NullableCloudmgrcoreIaasResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudmgrcoreIaasResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

