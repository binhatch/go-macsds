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

// checks if the RepairTimesVehicleReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepairTimesVehicleReference{}

// RepairTimesVehicleReference struct for RepairTimesVehicleReference
type RepairTimesVehicleReference struct {
	// KType
	Ktype *int64 `json:"ktype,omitempty"`
}

// NewRepairTimesVehicleReference instantiates a new RepairTimesVehicleReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepairTimesVehicleReference() *RepairTimesVehicleReference {
	this := RepairTimesVehicleReference{}
	return &this
}

// NewRepairTimesVehicleReferenceWithDefaults instantiates a new RepairTimesVehicleReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepairTimesVehicleReferenceWithDefaults() *RepairTimesVehicleReference {
	this := RepairTimesVehicleReference{}
	return &this
}

// GetKtype returns the Ktype field value if set, zero value otherwise.
func (o *RepairTimesVehicleReference) GetKtype() int64 {
	if o == nil || IsNil(o.Ktype) {
		var ret int64
		return ret
	}
	return *o.Ktype
}

// GetKtypeOk returns a tuple with the Ktype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepairTimesVehicleReference) GetKtypeOk() (*int64, bool) {
	if o == nil || IsNil(o.Ktype) {
		return nil, false
	}
	return o.Ktype, true
}

// HasKtype returns a boolean if a field has been set.
func (o *RepairTimesVehicleReference) HasKtype() bool {
	if o != nil && !IsNil(o.Ktype) {
		return true
	}

	return false
}

// SetKtype gets a reference to the given int64 and assigns it to the Ktype field.
func (o *RepairTimesVehicleReference) SetKtype(v int64) {
	o.Ktype = &v
}

func (o RepairTimesVehicleReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepairTimesVehicleReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ktype) {
		toSerialize["ktype"] = o.Ktype
	}
	return toSerialize, nil
}

type NullableRepairTimesVehicleReference struct {
	value *RepairTimesVehicleReference
	isSet bool
}

func (v NullableRepairTimesVehicleReference) Get() *RepairTimesVehicleReference {
	return v.value
}

func (v *NullableRepairTimesVehicleReference) Set(val *RepairTimesVehicleReference) {
	v.value = val
	v.isSet = true
}

func (v NullableRepairTimesVehicleReference) IsSet() bool {
	return v.isSet
}

func (v *NullableRepairTimesVehicleReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepairTimesVehicleReference(val *RepairTimesVehicleReference) *NullableRepairTimesVehicleReference {
	return &NullableRepairTimesVehicleReference{value: val, isSet: true}
}

func (v NullableRepairTimesVehicleReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepairTimesVehicleReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


