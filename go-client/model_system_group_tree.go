/*
Hella Gutmann - macsDS (Data Services) - API collection

This document contains all relevant APIs for diagnostics (incl. DTCs), repair & maintenance information (RMI) and vehicle identification.

API version: V20240702-130718
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/binhatch/go-macds

import (
	"encoding/json"
)

// checks if the SystemGroupTree type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemGroupTree{}

// SystemGroupTree struct for SystemGroupTree
type SystemGroupTree struct {
	// Recursive element representing the categories of wiring diagrams. Most recurring categories are engine, ABS, Airbag, climate. All other categories are considered 'Additional' and include (sub)categories such as 'Comfort'.
	SystemGroups []SystemGroup `json:"systemGroups,omitempty"`
}

// NewSystemGroupTree instantiates a new SystemGroupTree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemGroupTree() *SystemGroupTree {
	this := SystemGroupTree{}
	return &this
}

// NewSystemGroupTreeWithDefaults instantiates a new SystemGroupTree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemGroupTreeWithDefaults() *SystemGroupTree {
	this := SystemGroupTree{}
	return &this
}

// GetSystemGroups returns the SystemGroups field value if set, zero value otherwise.
func (o *SystemGroupTree) GetSystemGroups() []SystemGroup {
	if o == nil || IsNil(o.SystemGroups) {
		var ret []SystemGroup
		return ret
	}
	return o.SystemGroups
}

// GetSystemGroupsOk returns a tuple with the SystemGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemGroupTree) GetSystemGroupsOk() ([]SystemGroup, bool) {
	if o == nil || IsNil(o.SystemGroups) {
		return nil, false
	}
	return o.SystemGroups, true
}

// HasSystemGroups returns a boolean if a field has been set.
func (o *SystemGroupTree) HasSystemGroups() bool {
	if o != nil && !IsNil(o.SystemGroups) {
		return true
	}

	return false
}

// SetSystemGroups gets a reference to the given []SystemGroup and assigns it to the SystemGroups field.
func (o *SystemGroupTree) SetSystemGroups(v []SystemGroup) {
	o.SystemGroups = v
}

func (o SystemGroupTree) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemGroupTree) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SystemGroups) {
		toSerialize["systemGroups"] = o.SystemGroups
	}
	return toSerialize, nil
}

type NullableSystemGroupTree struct {
	value *SystemGroupTree
	isSet bool
}

func (v NullableSystemGroupTree) Get() *SystemGroupTree {
	return v.value
}

func (v *NullableSystemGroupTree) Set(val *SystemGroupTree) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemGroupTree) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemGroupTree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemGroupTree(val *SystemGroupTree) *NullableSystemGroupTree {
	return &NullableSystemGroupTree{value: val, isSet: true}
}

func (v NullableSystemGroupTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemGroupTree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


