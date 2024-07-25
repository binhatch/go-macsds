/*
Hella Gutmann - macsDS (Data Services) - API collection

This document contains all relevant APIs for diagnostics (incl. DTCs), repair & maintenance information (RMI) and vehicle identification.

API version: V20240702-130718
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package go-macds

import (
	"encoding/json"
)

// checks if the MakeV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MakeV2{}

// MakeV2 struct for MakeV2
type MakeV2 struct {
	// Id of the make
	Id *int32 `json:"id,omitempty"`
	// Name of the make
	Name *string `json:"name,omitempty"`
}

// NewMakeV2 instantiates a new MakeV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMakeV2() *MakeV2 {
	this := MakeV2{}
	return &this
}

// NewMakeV2WithDefaults instantiates a new MakeV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMakeV2WithDefaults() *MakeV2 {
	this := MakeV2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MakeV2) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MakeV2) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MakeV2) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MakeV2) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MakeV2) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MakeV2) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MakeV2) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MakeV2) SetName(v string) {
	o.Name = &v
}

func (o MakeV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MakeV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableMakeV2 struct {
	value *MakeV2
	isSet bool
}

func (v NullableMakeV2) Get() *MakeV2 {
	return v.value
}

func (v *NullableMakeV2) Set(val *MakeV2) {
	v.value = val
	v.isSet = true
}

func (v NullableMakeV2) IsSet() bool {
	return v.isSet
}

func (v *NullableMakeV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMakeV2(val *MakeV2) *NullableMakeV2 {
	return &NullableMakeV2{value: val, isSet: true}
}

func (v NullableMakeV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMakeV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


