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

// OpsDescribeAlertInfoResponse struct for OpsDescribeAlertInfoResponse
type OpsDescribeAlertInfoResponse struct {
	AlertStatus *string `json:"alert_status,omitempty"`
	AlertType *string `json:"alert_type,omitempty"`
	ConfirmAt *string `json:"confirm_at,omitempty"`
	CreateAt *string `json:"create_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Info *string `json:"info,omitempty"`
	ProcessAt *string `json:"process_at,omitempty"`
	Processor *string `json:"processor,omitempty"`
}

// NewOpsDescribeAlertInfoResponse instantiates a new OpsDescribeAlertInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsDescribeAlertInfoResponse() *OpsDescribeAlertInfoResponse {
	this := OpsDescribeAlertInfoResponse{}
	return &this
}

// NewOpsDescribeAlertInfoResponseWithDefaults instantiates a new OpsDescribeAlertInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsDescribeAlertInfoResponseWithDefaults() *OpsDescribeAlertInfoResponse {
	this := OpsDescribeAlertInfoResponse{}
	return &this
}

// GetAlertStatus returns the AlertStatus field value if set, zero value otherwise.
func (o *OpsDescribeAlertInfoResponse) GetAlertStatus() string {
	if o == nil || o.AlertStatus == nil {
		var ret string
		return ret
	}
	return *o.AlertStatus
}

// GetAlertStatusOk returns a tuple with the AlertStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeAlertInfoResponse) GetAlertStatusOk() (*string, bool) {
	if o == nil || o.AlertStatus == nil {
		return nil, false
	}
	return o.AlertStatus, true
}

// HasAlertStatus returns a boolean if a field has been set.
func (o *OpsDescribeAlertInfoResponse) HasAlertStatus() bool {
	if o != nil && o.AlertStatus != nil {
		return true
	}

	return false
}

// SetAlertStatus gets a reference to the given string and assigns it to the AlertStatus field.
func (o *OpsDescribeAlertInfoResponse) SetAlertStatus(v string) {
	o.AlertStatus = &v
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *OpsDescribeAlertInfoResponse) GetAlertType() string {
	if o == nil || o.AlertType == nil {
		var ret string
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeAlertInfoResponse) GetAlertTypeOk() (*string, bool) {
	if o == nil || o.AlertType == nil {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *OpsDescribeAlertInfoResponse) HasAlertType() bool {
	if o != nil && o.AlertType != nil {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given string and assigns it to the AlertType field.
func (o *OpsDescribeAlertInfoResponse) SetAlertType(v string) {
	o.AlertType = &v
}

// GetConfirmAt returns the ConfirmAt field value if set, zero value otherwise.
func (o *OpsDescribeAlertInfoResponse) GetConfirmAt() string {
	if o == nil || o.ConfirmAt == nil {
		var ret string
		return ret
	}
	return *o.ConfirmAt
}

// GetConfirmAtOk returns a tuple with the ConfirmAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeAlertInfoResponse) GetConfirmAtOk() (*string, bool) {
	if o == nil || o.ConfirmAt == nil {
		return nil, false
	}
	return o.ConfirmAt, true
}

// HasConfirmAt returns a boolean if a field has been set.
func (o *OpsDescribeAlertInfoResponse) HasConfirmAt() bool {
	if o != nil && o.ConfirmAt != nil {
		return true
	}

	return false
}

// SetConfirmAt gets a reference to the given string and assigns it to the ConfirmAt field.
func (o *OpsDescribeAlertInfoResponse) SetConfirmAt(v string) {
	o.ConfirmAt = &v
}

// GetCreateAt returns the CreateAt field value if set, zero value otherwise.
func (o *OpsDescribeAlertInfoResponse) GetCreateAt() string {
	if o == nil || o.CreateAt == nil {
		var ret string
		return ret
	}
	return *o.CreateAt
}

// GetCreateAtOk returns a tuple with the CreateAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeAlertInfoResponse) GetCreateAtOk() (*string, bool) {
	if o == nil || o.CreateAt == nil {
		return nil, false
	}
	return o.CreateAt, true
}

// HasCreateAt returns a boolean if a field has been set.
func (o *OpsDescribeAlertInfoResponse) HasCreateAt() bool {
	if o != nil && o.CreateAt != nil {
		return true
	}

	return false
}

// SetCreateAt gets a reference to the given string and assigns it to the CreateAt field.
func (o *OpsDescribeAlertInfoResponse) SetCreateAt(v string) {
	o.CreateAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OpsDescribeAlertInfoResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeAlertInfoResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OpsDescribeAlertInfoResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OpsDescribeAlertInfoResponse) SetId(v string) {
	o.Id = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *OpsDescribeAlertInfoResponse) GetInfo() string {
	if o == nil || o.Info == nil {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeAlertInfoResponse) GetInfoOk() (*string, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *OpsDescribeAlertInfoResponse) HasInfo() bool {
	if o != nil && o.Info != nil {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *OpsDescribeAlertInfoResponse) SetInfo(v string) {
	o.Info = &v
}

// GetProcessAt returns the ProcessAt field value if set, zero value otherwise.
func (o *OpsDescribeAlertInfoResponse) GetProcessAt() string {
	if o == nil || o.ProcessAt == nil {
		var ret string
		return ret
	}
	return *o.ProcessAt
}

// GetProcessAtOk returns a tuple with the ProcessAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeAlertInfoResponse) GetProcessAtOk() (*string, bool) {
	if o == nil || o.ProcessAt == nil {
		return nil, false
	}
	return o.ProcessAt, true
}

// HasProcessAt returns a boolean if a field has been set.
func (o *OpsDescribeAlertInfoResponse) HasProcessAt() bool {
	if o != nil && o.ProcessAt != nil {
		return true
	}

	return false
}

// SetProcessAt gets a reference to the given string and assigns it to the ProcessAt field.
func (o *OpsDescribeAlertInfoResponse) SetProcessAt(v string) {
	o.ProcessAt = &v
}

// GetProcessor returns the Processor field value if set, zero value otherwise.
func (o *OpsDescribeAlertInfoResponse) GetProcessor() string {
	if o == nil || o.Processor == nil {
		var ret string
		return ret
	}
	return *o.Processor
}

// GetProcessorOk returns a tuple with the Processor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeAlertInfoResponse) GetProcessorOk() (*string, bool) {
	if o == nil || o.Processor == nil {
		return nil, false
	}
	return o.Processor, true
}

// HasProcessor returns a boolean if a field has been set.
func (o *OpsDescribeAlertInfoResponse) HasProcessor() bool {
	if o != nil && o.Processor != nil {
		return true
	}

	return false
}

// SetProcessor gets a reference to the given string and assigns it to the Processor field.
func (o *OpsDescribeAlertInfoResponse) SetProcessor(v string) {
	o.Processor = &v
}

func (o OpsDescribeAlertInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AlertStatus != nil {
		toSerialize["alert_status"] = o.AlertStatus
	}
	if o.AlertType != nil {
		toSerialize["alert_type"] = o.AlertType
	}
	if o.ConfirmAt != nil {
		toSerialize["confirm_at"] = o.ConfirmAt
	}
	if o.CreateAt != nil {
		toSerialize["create_at"] = o.CreateAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.ProcessAt != nil {
		toSerialize["process_at"] = o.ProcessAt
	}
	if o.Processor != nil {
		toSerialize["processor"] = o.Processor
	}
	return json.Marshal(toSerialize)
}

type NullableOpsDescribeAlertInfoResponse struct {
	value *OpsDescribeAlertInfoResponse
	isSet bool
}

func (v NullableOpsDescribeAlertInfoResponse) Get() *OpsDescribeAlertInfoResponse {
	return v.value
}

func (v *NullableOpsDescribeAlertInfoResponse) Set(val *OpsDescribeAlertInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsDescribeAlertInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsDescribeAlertInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsDescribeAlertInfoResponse(val *OpsDescribeAlertInfoResponse) *NullableOpsDescribeAlertInfoResponse {
	return &NullableOpsDescribeAlertInfoResponse{value: val, isSet: true}
}

func (v NullableOpsDescribeAlertInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsDescribeAlertInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


