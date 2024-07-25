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

// checks if the ManualGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualGroup{}

// ManualGroup struct for ManualGroup
type ManualGroup struct {
	// Id
	Id *int32 `json:"id,omitempty"`
	// Title
	Title *string `json:"title,omitempty"`
	// List of Manuals
	Manuals []Manual `json:"manuals,omitempty"`
}

// NewManualGroup instantiates a new ManualGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualGroup() *ManualGroup {
	this := ManualGroup{}
	return &this
}

// NewManualGroupWithDefaults instantiates a new ManualGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualGroupWithDefaults() *ManualGroup {
	this := ManualGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManualGroup) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualGroup) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManualGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ManualGroup) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ManualGroup) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualGroup) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ManualGroup) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ManualGroup) SetTitle(v string) {
	o.Title = &v
}

// GetManuals returns the Manuals field value if set, zero value otherwise.
func (o *ManualGroup) GetManuals() []Manual {
	if o == nil || IsNil(o.Manuals) {
		var ret []Manual
		return ret
	}
	return o.Manuals
}

// GetManualsOk returns a tuple with the Manuals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualGroup) GetManualsOk() ([]Manual, bool) {
	if o == nil || IsNil(o.Manuals) {
		return nil, false
	}
	return o.Manuals, true
}

// HasManuals returns a boolean if a field has been set.
func (o *ManualGroup) HasManuals() bool {
	if o != nil && !IsNil(o.Manuals) {
		return true
	}

	return false
}

// SetManuals gets a reference to the given []Manual and assigns it to the Manuals field.
func (o *ManualGroup) SetManuals(v []Manual) {
	o.Manuals = v
}

func (o ManualGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Manuals) {
		toSerialize["manuals"] = o.Manuals
	}
	return toSerialize, nil
}

type NullableManualGroup struct {
	value *ManualGroup
	isSet bool
}

func (v NullableManualGroup) Get() *ManualGroup {
	return v.value
}

func (v *NullableManualGroup) Set(val *ManualGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableManualGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableManualGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualGroup(val *ManualGroup) *NullableManualGroup {
	return &NullableManualGroup{value: val, isSet: true}
}

func (v NullableManualGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


