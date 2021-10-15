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

// CoreDescribeVpcDomainMappingResponse struct for CoreDescribeVpcDomainMappingResponse
type CoreDescribeVpcDomainMappingResponse struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Id *string `json:"id,omitempty"`
	Ipaddress *string `json:"ipaddress,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
	Vpc *string `json:"vpc,omitempty"`
}

// NewCoreDescribeVpcDomainMappingResponse instantiates a new CoreDescribeVpcDomainMappingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreDescribeVpcDomainMappingResponse() *CoreDescribeVpcDomainMappingResponse {
	this := CoreDescribeVpcDomainMappingResponse{}
	return &this
}

// NewCoreDescribeVpcDomainMappingResponseWithDefaults instantiates a new CoreDescribeVpcDomainMappingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreDescribeVpcDomainMappingResponseWithDefaults() *CoreDescribeVpcDomainMappingResponse {
	this := CoreDescribeVpcDomainMappingResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CoreDescribeVpcDomainMappingResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVpcDomainMappingResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CoreDescribeVpcDomainMappingResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CoreDescribeVpcDomainMappingResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *CoreDescribeVpcDomainMappingResponse) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVpcDomainMappingResponse) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *CoreDescribeVpcDomainMappingResponse) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *CoreDescribeVpcDomainMappingResponse) SetDomain(v string) {
	o.Domain = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreDescribeVpcDomainMappingResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVpcDomainMappingResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreDescribeVpcDomainMappingResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CoreDescribeVpcDomainMappingResponse) SetId(v string) {
	o.Id = &v
}

// GetIpaddress returns the Ipaddress field value if set, zero value otherwise.
func (o *CoreDescribeVpcDomainMappingResponse) GetIpaddress() string {
	if o == nil || o.Ipaddress == nil {
		var ret string
		return ret
	}
	return *o.Ipaddress
}

// GetIpaddressOk returns a tuple with the Ipaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVpcDomainMappingResponse) GetIpaddressOk() (*string, bool) {
	if o == nil || o.Ipaddress == nil {
		return nil, false
	}
	return o.Ipaddress, true
}

// HasIpaddress returns a boolean if a field has been set.
func (o *CoreDescribeVpcDomainMappingResponse) HasIpaddress() bool {
	if o != nil && o.Ipaddress != nil {
		return true
	}

	return false
}

// SetIpaddress gets a reference to the given string and assigns it to the Ipaddress field.
func (o *CoreDescribeVpcDomainMappingResponse) SetIpaddress(v string) {
	o.Ipaddress = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *CoreDescribeVpcDomainMappingResponse) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVpcDomainMappingResponse) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *CoreDescribeVpcDomainMappingResponse) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *CoreDescribeVpcDomainMappingResponse) SetVendor(v string) {
	o.Vendor = &v
}

// GetVpc returns the Vpc field value if set, zero value otherwise.
func (o *CoreDescribeVpcDomainMappingResponse) GetVpc() string {
	if o == nil || o.Vpc == nil {
		var ret string
		return ret
	}
	return *o.Vpc
}

// GetVpcOk returns a tuple with the Vpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeVpcDomainMappingResponse) GetVpcOk() (*string, bool) {
	if o == nil || o.Vpc == nil {
		return nil, false
	}
	return o.Vpc, true
}

// HasVpc returns a boolean if a field has been set.
func (o *CoreDescribeVpcDomainMappingResponse) HasVpc() bool {
	if o != nil && o.Vpc != nil {
		return true
	}

	return false
}

// SetVpc gets a reference to the given string and assigns it to the Vpc field.
func (o *CoreDescribeVpcDomainMappingResponse) SetVpc(v string) {
	o.Vpc = &v
}

func (o CoreDescribeVpcDomainMappingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Ipaddress != nil {
		toSerialize["ipaddress"] = o.Ipaddress
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	if o.Vpc != nil {
		toSerialize["vpc"] = o.Vpc
	}
	return json.Marshal(toSerialize)
}

type NullableCoreDescribeVpcDomainMappingResponse struct {
	value *CoreDescribeVpcDomainMappingResponse
	isSet bool
}

func (v NullableCoreDescribeVpcDomainMappingResponse) Get() *CoreDescribeVpcDomainMappingResponse {
	return v.value
}

func (v *NullableCoreDescribeVpcDomainMappingResponse) Set(val *CoreDescribeVpcDomainMappingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreDescribeVpcDomainMappingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreDescribeVpcDomainMappingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreDescribeVpcDomainMappingResponse(val *CoreDescribeVpcDomainMappingResponse) *NullableCoreDescribeVpcDomainMappingResponse {
	return &NullableCoreDescribeVpcDomainMappingResponse{value: val, isSet: true}
}

func (v NullableCoreDescribeVpcDomainMappingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreDescribeVpcDomainMappingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

