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

// checks if the Criteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Criteria{}

// Criteria struct for Criteria
type Criteria struct {
	// Id
	Id *string `json:"id,omitempty"`
	Name *Name `json:"name,omitempty"`
}

// NewCriteria instantiates a new Criteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCriteria() *Criteria {
	this := Criteria{}
	return &this
}

// NewCriteriaWithDefaults instantiates a new Criteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCriteriaWithDefaults() *Criteria {
	this := Criteria{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Criteria) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Criteria) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Criteria) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Criteria) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Criteria) GetName() Name {
	if o == nil || IsNil(o.Name) {
		var ret Name
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Criteria) GetNameOk() (*Name, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Criteria) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given Name and assigns it to the Name field.
func (o *Criteria) SetName(v Name) {
	o.Name = &v
}

func (o Criteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Criteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCriteria struct {
	value *Criteria
	isSet bool
}

func (v NullableCriteria) Get() *Criteria {
	return v.value
}

func (v *NullableCriteria) Set(val *Criteria) {
	v.value = val
	v.isSet = true
}

func (v NullableCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCriteria(val *Criteria) *NullableCriteria {
	return &NullableCriteria{value: val, isSet: true}
}

func (v NullableCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


