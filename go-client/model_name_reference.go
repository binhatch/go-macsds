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

// checks if the NameReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NameReference{}

// NameReference struct for NameReference
type NameReference struct {
	TextReference *TextReference `json:"textReference,omitempty"`
	// Order
	Order *int32 `json:"order,omitempty"`
}

// NewNameReference instantiates a new NameReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameReference() *NameReference {
	this := NameReference{}
	return &this
}

// NewNameReferenceWithDefaults instantiates a new NameReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameReferenceWithDefaults() *NameReference {
	this := NameReference{}
	return &this
}

// GetTextReference returns the TextReference field value if set, zero value otherwise.
func (o *NameReference) GetTextReference() TextReference {
	if o == nil || IsNil(o.TextReference) {
		var ret TextReference
		return ret
	}
	return *o.TextReference
}

// GetTextReferenceOk returns a tuple with the TextReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameReference) GetTextReferenceOk() (*TextReference, bool) {
	if o == nil || IsNil(o.TextReference) {
		return nil, false
	}
	return o.TextReference, true
}

// HasTextReference returns a boolean if a field has been set.
func (o *NameReference) HasTextReference() bool {
	if o != nil && !IsNil(o.TextReference) {
		return true
	}

	return false
}

// SetTextReference gets a reference to the given TextReference and assigns it to the TextReference field.
func (o *NameReference) SetTextReference(v TextReference) {
	o.TextReference = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *NameReference) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameReference) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *NameReference) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *NameReference) SetOrder(v int32) {
	o.Order = &v
}

func (o NameReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NameReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TextReference) {
		toSerialize["textReference"] = o.TextReference
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullableNameReference struct {
	value *NameReference
	isSet bool
}

func (v NullableNameReference) Get() *NameReference {
	return v.value
}

func (v *NullableNameReference) Set(val *NameReference) {
	v.value = val
	v.isSet = true
}

func (v NullableNameReference) IsSet() bool {
	return v.isSet
}

func (v *NullableNameReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameReference(val *NameReference) *NullableNameReference {
	return &NullableNameReference{value: val, isSet: true}
}

func (v NullableNameReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


