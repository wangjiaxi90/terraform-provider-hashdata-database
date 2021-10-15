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

// CommonDescribeJobResponse struct for CommonDescribeJobResponse
type CommonDescribeJobResponse struct {
	Action *string `json:"action,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	ErrorCode *int32 `json:"error_code,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	Extras *string `json:"extras,omitempty"`
	Id *string `json:"id,omitempty"`
	RecoveryId *string `json:"recovery_id,omitempty"`
	ResourceIds *[]string `json:"resource_ids,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	User *string `json:"user,omitempty"`
}

// NewCommonDescribeJobResponse instantiates a new CommonDescribeJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonDescribeJobResponse() *CommonDescribeJobResponse {
	this := CommonDescribeJobResponse{}
	return &this
}

// NewCommonDescribeJobResponseWithDefaults instantiates a new CommonDescribeJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonDescribeJobResponseWithDefaults() *CommonDescribeJobResponse {
	this := CommonDescribeJobResponse{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *CommonDescribeJobResponse) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDescribeJobResponse) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *CommonDescribeJobResponse) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *CommonDescribeJobResponse) SetAction(v string) {
	o.Action = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CommonDescribeJobResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDescribeJobResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CommonDescribeJobResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CommonDescribeJobResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CommonDescribeJobResponse) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDescribeJobResponse) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CommonDescribeJobResponse) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *CommonDescribeJobResponse) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *CommonDescribeJobResponse) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDescribeJobResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CommonDescribeJobResponse) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *CommonDescribeJobResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *CommonDescribeJobResponse) GetExtras() string {
	if o == nil || o.Extras == nil {
		var ret string
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDescribeJobResponse) GetExtrasOk() (*string, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *CommonDescribeJobResponse) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given string and assigns it to the Extras field.
func (o *CommonDescribeJobResponse) SetExtras(v string) {
	o.Extras = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CommonDescribeJobResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDescribeJobResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CommonDescribeJobResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CommonDescribeJobResponse) SetId(v string) {
	o.Id = &v
}

// GetRecoveryId returns the RecoveryId field value if set, zero value otherwise.
func (o *CommonDescribeJobResponse) GetRecoveryId() string {
	if o == nil || o.RecoveryId == nil {
		var ret string
		return ret
	}
	return *o.RecoveryId
}

// GetRecoveryIdOk returns a tuple with the RecoveryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDescribeJobResponse) GetRecoveryIdOk() (*string, bool) {
	if o == nil || o.RecoveryId == nil {
		return nil, false
	}
	return o.RecoveryId, true
}

// HasRecoveryId returns a boolean if a field has been set.
func (o *CommonDescribeJobResponse) HasRecoveryId() bool {
	if o != nil && o.RecoveryId != nil {
		return true
	}

	return false
}

// SetRecoveryId gets a reference to the given string and assigns it to the RecoveryId field.
func (o *CommonDescribeJobResponse) SetRecoveryId(v string) {
	o.RecoveryId = &v
}

// GetResourceIds returns the ResourceIds field value if set, zero value otherwise.
func (o *CommonDescribeJobResponse) GetResourceIds() []string {
	if o == nil || o.ResourceIds == nil {
		var ret []string
		return ret
	}
	return *o.ResourceIds
}

// GetResourceIdsOk returns a tuple with the ResourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDescribeJobResponse) GetResourceIdsOk() (*[]string, bool) {
	if o == nil || o.ResourceIds == nil {
		return nil, false
	}
	return o.ResourceIds, true
}

// HasResourceIds returns a boolean if a field has been set.
func (o *CommonDescribeJobResponse) HasResourceIds() bool {
	if o != nil && o.ResourceIds != nil {
		return true
	}

	return false
}

// SetResourceIds gets a reference to the given []string and assigns it to the ResourceIds field.
func (o *CommonDescribeJobResponse) SetResourceIds(v []string) {
	o.ResourceIds = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CommonDescribeJobResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDescribeJobResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CommonDescribeJobResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CommonDescribeJobResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CommonDescribeJobResponse) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDescribeJobResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CommonDescribeJobResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *CommonDescribeJobResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CommonDescribeJobResponse) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDescribeJobResponse) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CommonDescribeJobResponse) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *CommonDescribeJobResponse) SetUser(v string) {
	o.User = &v
}

func (o CommonDescribeJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RecoveryId != nil {
		toSerialize["recovery_id"] = o.RecoveryId
	}
	if o.ResourceIds != nil {
		toSerialize["resource_ids"] = o.ResourceIds
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableCommonDescribeJobResponse struct {
	value *CommonDescribeJobResponse
	isSet bool
}

func (v NullableCommonDescribeJobResponse) Get() *CommonDescribeJobResponse {
	return v.value
}

func (v *NullableCommonDescribeJobResponse) Set(val *CommonDescribeJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonDescribeJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonDescribeJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonDescribeJobResponse(val *CommonDescribeJobResponse) *NullableCommonDescribeJobResponse {
	return &NullableCommonDescribeJobResponse{value: val, isSet: true}
}

func (v NullableCommonDescribeJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonDescribeJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

