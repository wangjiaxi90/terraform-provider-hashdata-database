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

// CommonResourceConfig struct for CommonResourceConfig
type CommonResourceConfig struct {
	Conf *string `json:"conf,omitempty"`
	PriceConf *CommonPriceConfig `json:"priceConf,omitempty"`
	Region *string `json:"region,omitempty"`
}

// NewCommonResourceConfig instantiates a new CommonResourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResourceConfig() *CommonResourceConfig {
	this := CommonResourceConfig{}
	return &this
}

// NewCommonResourceConfigWithDefaults instantiates a new CommonResourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonResourceConfigWithDefaults() *CommonResourceConfig {
	this := CommonResourceConfig{}
	return &this
}

// GetConf returns the Conf field value if set, zero value otherwise.
func (o *CommonResourceConfig) GetConf() string {
	if o == nil || o.Conf == nil {
		var ret string
		return ret
	}
	return *o.Conf
}

// GetConfOk returns a tuple with the Conf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonResourceConfig) GetConfOk() (*string, bool) {
	if o == nil || o.Conf == nil {
		return nil, false
	}
	return o.Conf, true
}

// HasConf returns a boolean if a field has been set.
func (o *CommonResourceConfig) HasConf() bool {
	if o != nil && o.Conf != nil {
		return true
	}

	return false
}

// SetConf gets a reference to the given string and assigns it to the Conf field.
func (o *CommonResourceConfig) SetConf(v string) {
	o.Conf = &v
}

// GetPriceConf returns the PriceConf field value if set, zero value otherwise.
func (o *CommonResourceConfig) GetPriceConf() CommonPriceConfig {
	if o == nil || o.PriceConf == nil {
		var ret CommonPriceConfig
		return ret
	}
	return *o.PriceConf
}

// GetPriceConfOk returns a tuple with the PriceConf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonResourceConfig) GetPriceConfOk() (*CommonPriceConfig, bool) {
	if o == nil || o.PriceConf == nil {
		return nil, false
	}
	return o.PriceConf, true
}

// HasPriceConf returns a boolean if a field has been set.
func (o *CommonResourceConfig) HasPriceConf() bool {
	if o != nil && o.PriceConf != nil {
		return true
	}

	return false
}

// SetPriceConf gets a reference to the given CommonPriceConfig and assigns it to the PriceConf field.
func (o *CommonResourceConfig) SetPriceConf(v CommonPriceConfig) {
	o.PriceConf = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CommonResourceConfig) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonResourceConfig) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CommonResourceConfig) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CommonResourceConfig) SetRegion(v string) {
	o.Region = &v
}

func (o CommonResourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conf != nil {
		toSerialize["conf"] = o.Conf
	}
	if o.PriceConf != nil {
		toSerialize["priceConf"] = o.PriceConf
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableCommonResourceConfig struct {
	value *CommonResourceConfig
	isSet bool
}

func (v NullableCommonResourceConfig) Get() *CommonResourceConfig {
	return v.value
}

func (v *NullableCommonResourceConfig) Set(val *CommonResourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonResourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonResourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonResourceConfig(val *CommonResourceConfig) *NullableCommonResourceConfig {
	return &NullableCommonResourceConfig{value: val, isSet: true}
}

func (v NullableCommonResourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonResourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

