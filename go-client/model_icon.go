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

// checks if the Icon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Icon{}

// Icon struct for Icon
type Icon struct {
	// Order
	Order *int64 `json:"order,omitempty"`
	ObjectReference *ObjectReference `json:"objectReference,omitempty"`
	// File Extension
	FileExtension *string `json:"fileExtension,omitempty"`
}

// NewIcon instantiates a new Icon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIcon() *Icon {
	this := Icon{}
	return &this
}

// NewIconWithDefaults instantiates a new Icon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIconWithDefaults() *Icon {
	this := Icon{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Icon) GetOrder() int64 {
	if o == nil || IsNil(o.Order) {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Icon) GetOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Icon) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *Icon) SetOrder(v int64) {
	o.Order = &v
}

// GetObjectReference returns the ObjectReference field value if set, zero value otherwise.
func (o *Icon) GetObjectReference() ObjectReference {
	if o == nil || IsNil(o.ObjectReference) {
		var ret ObjectReference
		return ret
	}
	return *o.ObjectReference
}

// GetObjectReferenceOk returns a tuple with the ObjectReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Icon) GetObjectReferenceOk() (*ObjectReference, bool) {
	if o == nil || IsNil(o.ObjectReference) {
		return nil, false
	}
	return o.ObjectReference, true
}

// HasObjectReference returns a boolean if a field has been set.
func (o *Icon) HasObjectReference() bool {
	if o != nil && !IsNil(o.ObjectReference) {
		return true
	}

	return false
}

// SetObjectReference gets a reference to the given ObjectReference and assigns it to the ObjectReference field.
func (o *Icon) SetObjectReference(v ObjectReference) {
	o.ObjectReference = &v
}

// GetFileExtension returns the FileExtension field value if set, zero value otherwise.
func (o *Icon) GetFileExtension() string {
	if o == nil || IsNil(o.FileExtension) {
		var ret string
		return ret
	}
	return *o.FileExtension
}

// GetFileExtensionOk returns a tuple with the FileExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Icon) GetFileExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.FileExtension) {
		return nil, false
	}
	return o.FileExtension, true
}

// HasFileExtension returns a boolean if a field has been set.
func (o *Icon) HasFileExtension() bool {
	if o != nil && !IsNil(o.FileExtension) {
		return true
	}

	return false
}

// SetFileExtension gets a reference to the given string and assigns it to the FileExtension field.
func (o *Icon) SetFileExtension(v string) {
	o.FileExtension = &v
}

func (o Icon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Icon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.ObjectReference) {
		toSerialize["objectReference"] = o.ObjectReference
	}
	if !IsNil(o.FileExtension) {
		toSerialize["fileExtension"] = o.FileExtension
	}
	return toSerialize, nil
}

type NullableIcon struct {
	value *Icon
	isSet bool
}

func (v NullableIcon) Get() *Icon {
	return v.value
}

func (v *NullableIcon) Set(val *Icon) {
	v.value = val
	v.isSet = true
}

func (v NullableIcon) IsSet() bool {
	return v.isSet
}

func (v *NullableIcon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIcon(val *Icon) *NullableIcon {
	return &NullableIcon{value: val, isSet: true}
}

func (v NullableIcon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIcon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


