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

// checks if the Manual type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Manual{}

// Manual struct for Manual
type Manual struct {
	// Id
	Id *int32 `json:"id,omitempty"`
	Qualifier *Qualifier `json:"qualifier,omitempty"`
}

// NewManual instantiates a new Manual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManual() *Manual {
	this := Manual{}
	return &this
}

// NewManualWithDefaults instantiates a new Manual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualWithDefaults() *Manual {
	this := Manual{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Manual) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Manual) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Manual) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Manual) SetId(v int32) {
	o.Id = &v
}

// GetQualifier returns the Qualifier field value if set, zero value otherwise.
func (o *Manual) GetQualifier() Qualifier {
	if o == nil || IsNil(o.Qualifier) {
		var ret Qualifier
		return ret
	}
	return *o.Qualifier
}

// GetQualifierOk returns a tuple with the Qualifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Manual) GetQualifierOk() (*Qualifier, bool) {
	if o == nil || IsNil(o.Qualifier) {
		return nil, false
	}
	return o.Qualifier, true
}

// HasQualifier returns a boolean if a field has been set.
func (o *Manual) HasQualifier() bool {
	if o != nil && !IsNil(o.Qualifier) {
		return true
	}

	return false
}

// SetQualifier gets a reference to the given Qualifier and assigns it to the Qualifier field.
func (o *Manual) SetQualifier(v Qualifier) {
	o.Qualifier = &v
}

func (o Manual) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Manual) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Qualifier) {
		toSerialize["qualifier"] = o.Qualifier
	}
	return toSerialize, nil
}

type NullableManual struct {
	value *Manual
	isSet bool
}

func (v NullableManual) Get() *Manual {
	return v.value
}

func (v *NullableManual) Set(val *Manual) {
	v.value = val
	v.isSet = true
}

func (v NullableManual) IsSet() bool {
	return v.isSet
}

func (v *NullableManual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManual(val *Manual) *NullableManual {
	return &NullableManual{value: val, isSet: true}
}

func (v NullableManual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


