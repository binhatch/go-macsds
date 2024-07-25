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

// checks if the TecDoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TecDoc{}

// TecDoc struct for TecDoc
type TecDoc struct {
	VehicleType *VehicleType `json:"vehicleType,omitempty"`
}

// NewTecDoc instantiates a new TecDoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTecDoc() *TecDoc {
	this := TecDoc{}
	return &this
}

// NewTecDocWithDefaults instantiates a new TecDoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTecDocWithDefaults() *TecDoc {
	this := TecDoc{}
	return &this
}

// GetVehicleType returns the VehicleType field value if set, zero value otherwise.
func (o *TecDoc) GetVehicleType() VehicleType {
	if o == nil || IsNil(o.VehicleType) {
		var ret VehicleType
		return ret
	}
	return *o.VehicleType
}

// GetVehicleTypeOk returns a tuple with the VehicleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TecDoc) GetVehicleTypeOk() (*VehicleType, bool) {
	if o == nil || IsNil(o.VehicleType) {
		return nil, false
	}
	return o.VehicleType, true
}

// HasVehicleType returns a boolean if a field has been set.
func (o *TecDoc) HasVehicleType() bool {
	if o != nil && !IsNil(o.VehicleType) {
		return true
	}

	return false
}

// SetVehicleType gets a reference to the given VehicleType and assigns it to the VehicleType field.
func (o *TecDoc) SetVehicleType(v VehicleType) {
	o.VehicleType = &v
}

func (o TecDoc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TecDoc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VehicleType) {
		toSerialize["vehicleType"] = o.VehicleType
	}
	return toSerialize, nil
}

type NullableTecDoc struct {
	value *TecDoc
	isSet bool
}

func (v NullableTecDoc) Get() *TecDoc {
	return v.value
}

func (v *NullableTecDoc) Set(val *TecDoc) {
	v.value = val
	v.isSet = true
}

func (v NullableTecDoc) IsSet() bool {
	return v.isSet
}

func (v *NullableTecDoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTecDoc(val *TecDoc) *NullableTecDoc {
	return &NullableTecDoc{value: val, isSet: true}
}

func (v NullableTecDoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTecDoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

