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

// CoreDescribeAlertConfigResponse struct for CoreDescribeAlertConfigResponse
type CoreDescribeAlertConfigResponse struct {
	EnableAlert *bool `json:"enableAlert,omitempty"`
	Maintenance *bool `json:"maintenance,omitempty"`
}

// NewCoreDescribeAlertConfigResponse instantiates a new CoreDescribeAlertConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreDescribeAlertConfigResponse() *CoreDescribeAlertConfigResponse {
	this := CoreDescribeAlertConfigResponse{}
	return &this
}

// NewCoreDescribeAlertConfigResponseWithDefaults instantiates a new CoreDescribeAlertConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreDescribeAlertConfigResponseWithDefaults() *CoreDescribeAlertConfigResponse {
	this := CoreDescribeAlertConfigResponse{}
	return &this
}

// GetEnableAlert returns the EnableAlert field value if set, zero value otherwise.
func (o *CoreDescribeAlertConfigResponse) GetEnableAlert() bool {
	if o == nil || o.EnableAlert == nil {
		var ret bool
		return ret
	}
	return *o.EnableAlert
}

// GetEnableAlertOk returns a tuple with the EnableAlert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeAlertConfigResponse) GetEnableAlertOk() (*bool, bool) {
	if o == nil || o.EnableAlert == nil {
		return nil, false
	}
	return o.EnableAlert, true
}

// HasEnableAlert returns a boolean if a field has been set.
func (o *CoreDescribeAlertConfigResponse) HasEnableAlert() bool {
	if o != nil && o.EnableAlert != nil {
		return true
	}

	return false
}

// SetEnableAlert gets a reference to the given bool and assigns it to the EnableAlert field.
func (o *CoreDescribeAlertConfigResponse) SetEnableAlert(v bool) {
	o.EnableAlert = &v
}

// GetMaintenance returns the Maintenance field value if set, zero value otherwise.
func (o *CoreDescribeAlertConfigResponse) GetMaintenance() bool {
	if o == nil || o.Maintenance == nil {
		var ret bool
		return ret
	}
	return *o.Maintenance
}

// GetMaintenanceOk returns a tuple with the Maintenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeAlertConfigResponse) GetMaintenanceOk() (*bool, bool) {
	if o == nil || o.Maintenance == nil {
		return nil, false
	}
	return o.Maintenance, true
}

// HasMaintenance returns a boolean if a field has been set.
func (o *CoreDescribeAlertConfigResponse) HasMaintenance() bool {
	if o != nil && o.Maintenance != nil {
		return true
	}

	return false
}

// SetMaintenance gets a reference to the given bool and assigns it to the Maintenance field.
func (o *CoreDescribeAlertConfigResponse) SetMaintenance(v bool) {
	o.Maintenance = &v
}

func (o CoreDescribeAlertConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnableAlert != nil {
		toSerialize["enableAlert"] = o.EnableAlert
	}
	if o.Maintenance != nil {
		toSerialize["maintenance"] = o.Maintenance
	}
	return json.Marshal(toSerialize)
}

type NullableCoreDescribeAlertConfigResponse struct {
	value *CoreDescribeAlertConfigResponse
	isSet bool
}

func (v NullableCoreDescribeAlertConfigResponse) Get() *CoreDescribeAlertConfigResponse {
	return v.value
}

func (v *NullableCoreDescribeAlertConfigResponse) Set(val *CoreDescribeAlertConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreDescribeAlertConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreDescribeAlertConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreDescribeAlertConfigResponse(val *CoreDescribeAlertConfigResponse) *NullableCoreDescribeAlertConfigResponse {
	return &NullableCoreDescribeAlertConfigResponse{value: val, isSet: true}
}

func (v NullableCoreDescribeAlertConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreDescribeAlertConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


