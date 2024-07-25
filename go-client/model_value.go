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

// checks if the Value type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Value{}

// Value struct for Value
type Value struct {
	// Id
	Id *int32 `json:"id,omitempty"`
	TextReference *TranslationContainer `json:"textReference,omitempty"`
}

// NewValue instantiates a new Value object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValue() *Value {
	this := Value{}
	return &this
}

// NewValueWithDefaults instantiates a new Value object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueWithDefaults() *Value {
	this := Value{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Value) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Value) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Value) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Value) SetId(v int32) {
	o.Id = &v
}

// GetTextReference returns the TextReference field value if set, zero value otherwise.
func (o *Value) GetTextReference() TranslationContainer {
	if o == nil || IsNil(o.TextReference) {
		var ret TranslationContainer
		return ret
	}
	return *o.TextReference
}

// GetTextReferenceOk returns a tuple with the TextReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Value) GetTextReferenceOk() (*TranslationContainer, bool) {
	if o == nil || IsNil(o.TextReference) {
		return nil, false
	}
	return o.TextReference, true
}

// HasTextReference returns a boolean if a field has been set.
func (o *Value) HasTextReference() bool {
	if o != nil && !IsNil(o.TextReference) {
		return true
	}

	return false
}

// SetTextReference gets a reference to the given TranslationContainer and assigns it to the TextReference field.
func (o *Value) SetTextReference(v TranslationContainer) {
	o.TextReference = &v
}

func (o Value) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Value) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TextReference) {
		toSerialize["textReference"] = o.TextReference
	}
	return toSerialize, nil
}

type NullableValue struct {
	value *Value
	isSet bool
}

func (v NullableValue) Get() *Value {
	return v.value
}

func (v *NullableValue) Set(val *Value) {
	v.value = val
	v.isSet = true
}

func (v NullableValue) IsSet() bool {
	return v.isSet
}

func (v *NullableValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValue(val *Value) *NullableValue {
	return &NullableValue{value: val, isSet: true}
}

func (v NullableValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


