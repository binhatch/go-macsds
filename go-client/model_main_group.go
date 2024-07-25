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

// checks if the MainGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MainGroup{}

// MainGroup struct for MainGroup
type MainGroup struct {
	// Id
	Id *int32 `json:"id,omitempty"`
	// Title
	Title *string `json:"title,omitempty"`
	// List of Sub Groups
	SubGroups []SubGroup `json:"subGroups,omitempty"`
}

// NewMainGroup instantiates a new MainGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMainGroup() *MainGroup {
	this := MainGroup{}
	return &this
}

// NewMainGroupWithDefaults instantiates a new MainGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMainGroupWithDefaults() *MainGroup {
	this := MainGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MainGroup) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainGroup) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MainGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MainGroup) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MainGroup) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainGroup) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MainGroup) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MainGroup) SetTitle(v string) {
	o.Title = &v
}

// GetSubGroups returns the SubGroups field value if set, zero value otherwise.
func (o *MainGroup) GetSubGroups() []SubGroup {
	if o == nil || IsNil(o.SubGroups) {
		var ret []SubGroup
		return ret
	}
	return o.SubGroups
}

// GetSubGroupsOk returns a tuple with the SubGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainGroup) GetSubGroupsOk() ([]SubGroup, bool) {
	if o == nil || IsNil(o.SubGroups) {
		return nil, false
	}
	return o.SubGroups, true
}

// HasSubGroups returns a boolean if a field has been set.
func (o *MainGroup) HasSubGroups() bool {
	if o != nil && !IsNil(o.SubGroups) {
		return true
	}

	return false
}

// SetSubGroups gets a reference to the given []SubGroup and assigns it to the SubGroups field.
func (o *MainGroup) SetSubGroups(v []SubGroup) {
	o.SubGroups = v
}

func (o MainGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MainGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.SubGroups) {
		toSerialize["subGroups"] = o.SubGroups
	}
	return toSerialize, nil
}

type NullableMainGroup struct {
	value *MainGroup
	isSet bool
}

func (v NullableMainGroup) Get() *MainGroup {
	return v.value
}

func (v *NullableMainGroup) Set(val *MainGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableMainGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableMainGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMainGroup(val *MainGroup) *NullableMainGroup {
	return &NullableMainGroup{value: val, isSet: true}
}

func (v NullableMainGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMainGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


