/*
Hella Gutmann - macsDS (Data Services) - API collection

This document contains all relevant APIs for diagnostics (incl. DTCs), repair & maintenance information (RMI) and vehicle identification.

API version: V20240702-130718
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gomacds

import (
	"encoding/json"
)

// checks if the TechData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TechData{}

// TechData struct for TechData
type TechData struct {
	// A list of Group Details
	Groups []GroupDetails `json:"groups,omitempty"`
}

// NewTechData instantiates a new TechData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechData() *TechData {
	this := TechData{}
	return &this
}

// NewTechDataWithDefaults instantiates a new TechData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechDataWithDefaults() *TechData {
	this := TechData{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *TechData) GetGroups() []GroupDetails {
	if o == nil || IsNil(o.Groups) {
		var ret []GroupDetails
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechData) GetGroupsOk() ([]GroupDetails, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *TechData) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupDetails and assigns it to the Groups field.
func (o *TechData) SetGroups(v []GroupDetails) {
	o.Groups = v
}

func (o TechData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TechData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return toSerialize, nil
}

type NullableTechData struct {
	value *TechData
	isSet bool
}

func (v NullableTechData) Get() *TechData {
	return v.value
}

func (v *NullableTechData) Set(val *TechData) {
	v.value = val
	v.isSet = true
}

func (v NullableTechData) IsSet() bool {
	return v.isSet
}

func (v *NullableTechData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechData(val *TechData) *NullableTechData {
	return &NullableTechData{value: val, isSet: true}
}

func (v NullableTechData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


