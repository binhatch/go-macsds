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

// checks if the VehicleReferenceX type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VehicleReferenceX{}

// VehicleReferenceX struct for VehicleReferenceX
type VehicleReferenceX struct {
	Hgs *Hgs `json:"hgs,omitempty"`
	TecDoc *TecDoc `json:"tecDoc,omitempty"`
	DataTypes *Availability `json:"dataTypes,omitempty"`
}

// NewVehicleReferenceX instantiates a new VehicleReferenceX object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleReferenceX() *VehicleReferenceX {
	this := VehicleReferenceX{}
	return &this
}

// NewVehicleReferenceXWithDefaults instantiates a new VehicleReferenceX object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleReferenceXWithDefaults() *VehicleReferenceX {
	this := VehicleReferenceX{}
	return &this
}

// GetHgs returns the Hgs field value if set, zero value otherwise.
func (o *VehicleReferenceX) GetHgs() Hgs {
	if o == nil || IsNil(o.Hgs) {
		var ret Hgs
		return ret
	}
	return *o.Hgs
}

// GetHgsOk returns a tuple with the Hgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleReferenceX) GetHgsOk() (*Hgs, bool) {
	if o == nil || IsNil(o.Hgs) {
		return nil, false
	}
	return o.Hgs, true
}

// HasHgs returns a boolean if a field has been set.
func (o *VehicleReferenceX) HasHgs() bool {
	if o != nil && !IsNil(o.Hgs) {
		return true
	}

	return false
}

// SetHgs gets a reference to the given Hgs and assigns it to the Hgs field.
func (o *VehicleReferenceX) SetHgs(v Hgs) {
	o.Hgs = &v
}

// GetTecDoc returns the TecDoc field value if set, zero value otherwise.
func (o *VehicleReferenceX) GetTecDoc() TecDoc {
	if o == nil || IsNil(o.TecDoc) {
		var ret TecDoc
		return ret
	}
	return *o.TecDoc
}

// GetTecDocOk returns a tuple with the TecDoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleReferenceX) GetTecDocOk() (*TecDoc, bool) {
	if o == nil || IsNil(o.TecDoc) {
		return nil, false
	}
	return o.TecDoc, true
}

// HasTecDoc returns a boolean if a field has been set.
func (o *VehicleReferenceX) HasTecDoc() bool {
	if o != nil && !IsNil(o.TecDoc) {
		return true
	}

	return false
}

// SetTecDoc gets a reference to the given TecDoc and assigns it to the TecDoc field.
func (o *VehicleReferenceX) SetTecDoc(v TecDoc) {
	o.TecDoc = &v
}

// GetDataTypes returns the DataTypes field value if set, zero value otherwise.
func (o *VehicleReferenceX) GetDataTypes() Availability {
	if o == nil || IsNil(o.DataTypes) {
		var ret Availability
		return ret
	}
	return *o.DataTypes
}

// GetDataTypesOk returns a tuple with the DataTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleReferenceX) GetDataTypesOk() (*Availability, bool) {
	if o == nil || IsNil(o.DataTypes) {
		return nil, false
	}
	return o.DataTypes, true
}

// HasDataTypes returns a boolean if a field has been set.
func (o *VehicleReferenceX) HasDataTypes() bool {
	if o != nil && !IsNil(o.DataTypes) {
		return true
	}

	return false
}

// SetDataTypes gets a reference to the given Availability and assigns it to the DataTypes field.
func (o *VehicleReferenceX) SetDataTypes(v Availability) {
	o.DataTypes = &v
}

func (o VehicleReferenceX) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VehicleReferenceX) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hgs) {
		toSerialize["hgs"] = o.Hgs
	}
	if !IsNil(o.TecDoc) {
		toSerialize["tecDoc"] = o.TecDoc
	}
	if !IsNil(o.DataTypes) {
		toSerialize["dataTypes"] = o.DataTypes
	}
	return toSerialize, nil
}

type NullableVehicleReferenceX struct {
	value *VehicleReferenceX
	isSet bool
}

func (v NullableVehicleReferenceX) Get() *VehicleReferenceX {
	return v.value
}

func (v *NullableVehicleReferenceX) Set(val *VehicleReferenceX) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleReferenceX) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleReferenceX) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleReferenceX(val *VehicleReferenceX) *NullableVehicleReferenceX {
	return &NullableVehicleReferenceX{value: val, isSet: true}
}

func (v NullableVehicleReferenceX) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleReferenceX) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

