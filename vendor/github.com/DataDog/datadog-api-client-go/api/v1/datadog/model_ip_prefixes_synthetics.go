/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// IPPrefixesSynthetics Available prefix information for the Synthetics endpoints.
type IPPrefixesSynthetics struct {
	// List of IPv4 prefixes.
	PrefixesIpv4 *[]string `json:"prefixes_ipv4,omitempty"`
	// List of IPv6 prefixes.
	PrefixesIpv6 *[]string `json:"prefixes_ipv6,omitempty"`
}

// NewIPPrefixesSynthetics instantiates a new IPPrefixesSynthetics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPPrefixesSynthetics() *IPPrefixesSynthetics {
	this := IPPrefixesSynthetics{}
	return &this
}

// NewIPPrefixesSyntheticsWithDefaults instantiates a new IPPrefixesSynthetics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPPrefixesSyntheticsWithDefaults() *IPPrefixesSynthetics {
	this := IPPrefixesSynthetics{}
	return &this
}

// GetPrefixesIpv4 returns the PrefixesIpv4 field value if set, zero value otherwise.
func (o *IPPrefixesSynthetics) GetPrefixesIpv4() []string {
	if o == nil || o.PrefixesIpv4 == nil {
		var ret []string
		return ret
	}
	return *o.PrefixesIpv4
}

// GetPrefixesIpv4Ok returns a tuple with the PrefixesIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPPrefixesSynthetics) GetPrefixesIpv4Ok() (*[]string, bool) {
	if o == nil || o.PrefixesIpv4 == nil {
		return nil, false
	}
	return o.PrefixesIpv4, true
}

// HasPrefixesIpv4 returns a boolean if a field has been set.
func (o *IPPrefixesSynthetics) HasPrefixesIpv4() bool {
	if o != nil && o.PrefixesIpv4 != nil {
		return true
	}

	return false
}

// SetPrefixesIpv4 gets a reference to the given []string and assigns it to the PrefixesIpv4 field.
func (o *IPPrefixesSynthetics) SetPrefixesIpv4(v []string) {
	o.PrefixesIpv4 = &v
}

// GetPrefixesIpv6 returns the PrefixesIpv6 field value if set, zero value otherwise.
func (o *IPPrefixesSynthetics) GetPrefixesIpv6() []string {
	if o == nil || o.PrefixesIpv6 == nil {
		var ret []string
		return ret
	}
	return *o.PrefixesIpv6
}

// GetPrefixesIpv6Ok returns a tuple with the PrefixesIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPPrefixesSynthetics) GetPrefixesIpv6Ok() (*[]string, bool) {
	if o == nil || o.PrefixesIpv6 == nil {
		return nil, false
	}
	return o.PrefixesIpv6, true
}

// HasPrefixesIpv6 returns a boolean if a field has been set.
func (o *IPPrefixesSynthetics) HasPrefixesIpv6() bool {
	if o != nil && o.PrefixesIpv6 != nil {
		return true
	}

	return false
}

// SetPrefixesIpv6 gets a reference to the given []string and assigns it to the PrefixesIpv6 field.
func (o *IPPrefixesSynthetics) SetPrefixesIpv6(v []string) {
	o.PrefixesIpv6 = &v
}

func (o IPPrefixesSynthetics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrefixesIpv4 != nil {
		toSerialize["prefixes_ipv4"] = o.PrefixesIpv4
	}
	if o.PrefixesIpv6 != nil {
		toSerialize["prefixes_ipv6"] = o.PrefixesIpv6
	}
	return json.Marshal(toSerialize)
}

type NullableIPPrefixesSynthetics struct {
	value *IPPrefixesSynthetics
	isSet bool
}

func (v NullableIPPrefixesSynthetics) Get() *IPPrefixesSynthetics {
	return v.value
}

func (v *NullableIPPrefixesSynthetics) Set(val *IPPrefixesSynthetics) {
	v.value = val
	v.isSet = true
}

func (v NullableIPPrefixesSynthetics) IsSet() bool {
	return v.isSet
}

func (v *NullableIPPrefixesSynthetics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPPrefixesSynthetics(val *IPPrefixesSynthetics) *NullableIPPrefixesSynthetics {
	return &NullableIPPrefixesSynthetics{value: val, isSet: true}
}

func (v NullableIPPrefixesSynthetics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPPrefixesSynthetics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}