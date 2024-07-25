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

// checks if the RepairTimesTree type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepairTimesTree{}

// RepairTimesTree struct for RepairTimesTree
type RepairTimesTree struct {
	// Categories
	Categories []RepairTimesCategoryResponse `json:"categories,omitempty"`
}

// NewRepairTimesTree instantiates a new RepairTimesTree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepairTimesTree() *RepairTimesTree {
	this := RepairTimesTree{}
	return &this
}

// NewRepairTimesTreeWithDefaults instantiates a new RepairTimesTree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepairTimesTreeWithDefaults() *RepairTimesTree {
	this := RepairTimesTree{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *RepairTimesTree) GetCategories() []RepairTimesCategoryResponse {
	if o == nil || IsNil(o.Categories) {
		var ret []RepairTimesCategoryResponse
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepairTimesTree) GetCategoriesOk() ([]RepairTimesCategoryResponse, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *RepairTimesTree) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []RepairTimesCategoryResponse and assigns it to the Categories field.
func (o *RepairTimesTree) SetCategories(v []RepairTimesCategoryResponse) {
	o.Categories = v
}

func (o RepairTimesTree) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepairTimesTree) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

type NullableRepairTimesTree struct {
	value *RepairTimesTree
	isSet bool
}

func (v NullableRepairTimesTree) Get() *RepairTimesTree {
	return v.value
}

func (v *NullableRepairTimesTree) Set(val *RepairTimesTree) {
	v.value = val
	v.isSet = true
}

func (v NullableRepairTimesTree) IsSet() bool {
	return v.isSet
}

func (v *NullableRepairTimesTree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepairTimesTree(val *RepairTimesTree) *NullableRepairTimesTree {
	return &NullableRepairTimesTree{value: val, isSet: true}
}

func (v NullableRepairTimesTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepairTimesTree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

