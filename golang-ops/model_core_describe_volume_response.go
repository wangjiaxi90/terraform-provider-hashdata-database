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

// CoreDescribeVolumeResponse struct for CoreDescribeVolumeResponse
type CoreDescribeVolumeResponse struct {
	CreatedAt *string `json:"created_at,omitempty"`
	DestroyedAt *string `json:"destroyed_at,omitempty"`
	DeviceName *string `json:"device_name,omitempty"`
	Id *string `json:"id,omitempty"`
	Instance *string `json:"instance,omitempty"`
	MountPoint *string `json:"mount_point,omitempty"`
	Name *string `json:"name,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewCoreDescribeVolumeResponse instantiates a new CoreDescribeVolumeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreDescribeVolumeResponse() *CoreDescribeVolumeResponse {
	this := CoreDescribeVolumeResponse{}
	return &this
}

// NewCoreDescribeVolumeResponseWithDefaults instantiates a new CoreDescribeVolumeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreDescribeVolumeResponseWithDefaults() *CoreDescribeVolumeResponse {
	this := CoreDescribeVolumeResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CoreDescribeVolumeResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVolumeResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CoreDescribeVolumeResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CoreDescribeVolumeResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDestroyedAt returns the DestroyedAt field value if set, zero value otherwise.
func (o *CoreDescribeVolumeResponse) GetDestroyedAt() string {
	if o == nil || o.DestroyedAt == nil {
		var ret string
		return ret
	}
	return *o.DestroyedAt
}

// GetDestroyedAtOk returns a tuple with the DestroyedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVolumeResponse) GetDestroyedAtOk() (*string, bool) {
	if o == nil || o.DestroyedAt == nil {
		return nil, false
	}
	return o.DestroyedAt, true
}

// HasDestroyedAt returns a boolean if a field has been set.
func (o *CoreDescribeVolumeResponse) HasDestroyedAt() bool {
	if o != nil && o.DestroyedAt != nil {
		return true
	}

	return false
}

// SetDestroyedAt gets a reference to the given string and assigns it to the DestroyedAt field.
func (o *CoreDescribeVolumeResponse) SetDestroyedAt(v string) {
	o.DestroyedAt = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *CoreDescribeVolumeResponse) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVolumeResponse) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *CoreDescribeVolumeResponse) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *CoreDescribeVolumeResponse) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreDescribeVolumeResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVolumeResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreDescribeVolumeResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CoreDescribeVolumeResponse) SetId(v string) {
	o.Id = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *CoreDescribeVolumeResponse) GetInstance() string {
	if o == nil || o.Instance == nil {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVolumeResponse) GetInstanceOk() (*string, bool) {
	if o == nil || o.Instance == nil {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *CoreDescribeVolumeResponse) HasInstance() bool {
	if o != nil && o.Instance != nil {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *CoreDescribeVolumeResponse) SetInstance(v string) {
	o.Instance = &v
}

// GetMountPoint returns the MountPoint field value if set, zero value otherwise.
func (o *CoreDescribeVolumeResponse) GetMountPoint() string {
	if o == nil || o.MountPoint == nil {
		var ret string
		return ret
	}
	return *o.MountPoint
}

// GetMountPointOk returns a tuple with the MountPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVolumeResponse) GetMountPointOk() (*string, bool) {
	if o == nil || o.MountPoint == nil {
		return nil, false
	}
	return o.MountPoint, true
}

// HasMountPoint returns a boolean if a field has been set.
func (o *CoreDescribeVolumeResponse) HasMountPoint() bool {
	if o != nil && o.MountPoint != nil {
		return true
	}

	return false
}

// SetMountPoint gets a reference to the given string and assigns it to the MountPoint field.
func (o *CoreDescribeVolumeResponse) SetMountPoint(v string) {
	o.MountPoint = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreDescribeVolumeResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVolumeResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreDescribeVolumeResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreDescribeVolumeResponse) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CoreDescribeVolumeResponse) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVolumeResponse) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CoreDescribeVolumeResponse) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *CoreDescribeVolumeResponse) SetSize(v int32) {
	o.Size = &v
}

func (o CoreDescribeVolumeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DestroyedAt != nil {
		toSerialize["destroyed_at"] = o.DestroyedAt
	}
	if o.DeviceName != nil {
		toSerialize["device_name"] = o.DeviceName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Instance != nil {
		toSerialize["instance"] = o.Instance
	}
	if o.MountPoint != nil {
		toSerialize["mount_point"] = o.MountPoint
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableCoreDescribeVolumeResponse struct {
	value *CoreDescribeVolumeResponse
	isSet bool
}

func (v NullableCoreDescribeVolumeResponse) Get() *CoreDescribeVolumeResponse {
	return v.value
}

func (v *NullableCoreDescribeVolumeResponse) Set(val *CoreDescribeVolumeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreDescribeVolumeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreDescribeVolumeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreDescribeVolumeResponse(val *CoreDescribeVolumeResponse) *NullableCoreDescribeVolumeResponse {
	return &NullableCoreDescribeVolumeResponse{value: val, isSet: true}
}

func (v NullableCoreDescribeVolumeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreDescribeVolumeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


