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

// CoreUpdateQuotaRequest struct for CoreUpdateQuotaRequest
type CoreUpdateQuotaRequest struct {
	Quota *int32 `json:"quota,omitempty"`
	Tenant *string `json:"tenant,omitempty"`
	Type *string `json:"type,omitempty"`
	Zone *string `json:"zone,omitempty"`
}

// NewCoreUpdateQuotaRequest instantiates a new CoreUpdateQuotaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUpdateQuotaRequest() *CoreUpdateQuotaRequest {
	this := CoreUpdateQuotaRequest{}
	return &this
}

// NewCoreUpdateQuotaRequestWithDefaults instantiates a new CoreUpdateQuotaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUpdateQuotaRequestWithDefaults() *CoreUpdateQuotaRequest {
	this := CoreUpdateQuotaRequest{}
	return &this
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *CoreUpdateQuotaRequest) GetQuota() int32 {
	if o == nil || o.Quota == nil {
		var ret int32
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateQuotaRequest) GetQuotaOk() (*int32, bool) {
	if o == nil || o.Quota == nil {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *CoreUpdateQuotaRequest) HasQuota() bool {
	if o != nil && o.Quota != nil {
		return true
	}

	return false
}

// SetQuota gets a reference to the given int32 and assigns it to the Quota field.
func (o *CoreUpdateQuotaRequest) SetQuota(v int32) {
	o.Quota = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *CoreUpdateQuotaRequest) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateQuotaRequest) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *CoreUpdateQuotaRequest) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *CoreUpdateQuotaRequest) SetTenant(v string) {
	o.Tenant = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CoreUpdateQuotaRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateQuotaRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CoreUpdateQuotaRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CoreUpdateQuotaRequest) SetType(v string) {
	o.Type = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *CoreUpdateQuotaRequest) GetZone() string {
	if o == nil || o.Zone == nil {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateQuotaRequest) GetZoneOk() (*string, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *CoreUpdateQuotaRequest) HasZone() bool {
	if o != nil && o.Zone != nil {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *CoreUpdateQuotaRequest) SetZone(v string) {
	o.Zone = &v
}

func (o CoreUpdateQuotaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Quota != nil {
		toSerialize["quota"] = o.Quota
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}
	return json.Marshal(toSerialize)
}

type NullableCoreUpdateQuotaRequest struct {
	value *CoreUpdateQuotaRequest
	isSet bool
}

func (v NullableCoreUpdateQuotaRequest) Get() *CoreUpdateQuotaRequest {
	return v.value
}

func (v *NullableCoreUpdateQuotaRequest) Set(val *CoreUpdateQuotaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUpdateQuotaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUpdateQuotaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUpdateQuotaRequest(val *CoreUpdateQuotaRequest) *NullableCoreUpdateQuotaRequest {
	return &NullableCoreUpdateQuotaRequest{value: val, isSet: true}
}

func (v NullableCoreUpdateQuotaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUpdateQuotaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

