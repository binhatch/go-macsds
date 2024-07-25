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

// checks if the Line type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Line{}

// Line struct for Line
type Line struct {
	// Id
	Id *int64 `json:"id,omitempty"`
	// Order
	Order *int64 `json:"order,omitempty"`
	TextReference *IdTranslation `json:"textReference,omitempty"`
	// Is List Item
	IsListItem *bool `json:"isListItem,omitempty"`
}

// NewLine instantiates a new Line object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLine() *Line {
	this := Line{}
	return &this
}

// NewLineWithDefaults instantiates a new Line object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineWithDefaults() *Line {
	this := Line{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Line) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Line) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Line) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Line) SetId(v int64) {
	o.Id = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Line) GetOrder() int64 {
	if o == nil || IsNil(o.Order) {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Line) GetOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Line) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *Line) SetOrder(v int64) {
	o.Order = &v
}

// GetTextReference returns the TextReference field value if set, zero value otherwise.
func (o *Line) GetTextReference() IdTranslation {
	if o == nil || IsNil(o.TextReference) {
		var ret IdTranslation
		return ret
	}
	return *o.TextReference
}

// GetTextReferenceOk returns a tuple with the TextReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Line) GetTextReferenceOk() (*IdTranslation, bool) {
	if o == nil || IsNil(o.TextReference) {
		return nil, false
	}
	return o.TextReference, true
}

// HasTextReference returns a boolean if a field has been set.
func (o *Line) HasTextReference() bool {
	if o != nil && !IsNil(o.TextReference) {
		return true
	}

	return false
}

// SetTextReference gets a reference to the given IdTranslation and assigns it to the TextReference field.
func (o *Line) SetTextReference(v IdTranslation) {
	o.TextReference = &v
}

// GetIsListItem returns the IsListItem field value if set, zero value otherwise.
func (o *Line) GetIsListItem() bool {
	if o == nil || IsNil(o.IsListItem) {
		var ret bool
		return ret
	}
	return *o.IsListItem
}

// GetIsListItemOk returns a tuple with the IsListItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Line) GetIsListItemOk() (*bool, bool) {
	if o == nil || IsNil(o.IsListItem) {
		return nil, false
	}
	return o.IsListItem, true
}

// HasIsListItem returns a boolean if a field has been set.
func (o *Line) HasIsListItem() bool {
	if o != nil && !IsNil(o.IsListItem) {
		return true
	}

	return false
}

// SetIsListItem gets a reference to the given bool and assigns it to the IsListItem field.
func (o *Line) SetIsListItem(v bool) {
	o.IsListItem = &v
}

func (o Line) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Line) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.TextReference) {
		toSerialize["textReference"] = o.TextReference
	}
	if !IsNil(o.IsListItem) {
		toSerialize["isListItem"] = o.IsListItem
	}
	return toSerialize, nil
}

type NullableLine struct {
	value *Line
	isSet bool
}

func (v NullableLine) Get() *Line {
	return v.value
}

func (v *NullableLine) Set(val *Line) {
	v.value = val
	v.isSet = true
}

func (v NullableLine) IsSet() bool {
	return v.isSet
}

func (v *NullableLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLine(val *Line) *NullableLine {
	return &NullableLine{value: val, isSet: true}
}

func (v NullableLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

