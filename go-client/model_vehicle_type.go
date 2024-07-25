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

// checks if the VehicleType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VehicleType{}

// VehicleType struct for VehicleType
type VehicleType struct {
	// Additional KType Numbers
	AdditionalKTypeNos []int32 `json:"additionalKTypeNos,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	Lifespan *Lifespan `json:"lifespan,omitempty"`
	// KType Number
	KTypeNo *int32 `json:"kTypeNo,omitempty"`
}

// NewVehicleType instantiates a new VehicleType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleType() *VehicleType {
	this := VehicleType{}
	return &this
}

// NewVehicleTypeWithDefaults instantiates a new VehicleType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleTypeWithDefaults() *VehicleType {
	this := VehicleType{}
	return &this
}

// GetAdditionalKTypeNos returns the AdditionalKTypeNos field value if set, zero value otherwise.
func (o *VehicleType) GetAdditionalKTypeNos() []int32 {
	if o == nil || IsNil(o.AdditionalKTypeNos) {
		var ret []int32
		return ret
	}
	return o.AdditionalKTypeNos
}

// GetAdditionalKTypeNosOk returns a tuple with the AdditionalKTypeNos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleType) GetAdditionalKTypeNosOk() ([]int32, bool) {
	if o == nil || IsNil(o.AdditionalKTypeNos) {
		return nil, false
	}
	return o.AdditionalKTypeNos, true
}

// HasAdditionalKTypeNos returns a boolean if a field has been set.
func (o *VehicleType) HasAdditionalKTypeNos() bool {
	if o != nil && !IsNil(o.AdditionalKTypeNos) {
		return true
	}

	return false
}

// SetAdditionalKTypeNos gets a reference to the given []int32 and assigns it to the AdditionalKTypeNos field.
func (o *VehicleType) SetAdditionalKTypeNos(v []int32) {
	o.AdditionalKTypeNos = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VehicleType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VehicleType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VehicleType) SetName(v string) {
	o.Name = &v
}

// GetLifespan returns the Lifespan field value if set, zero value otherwise.
func (o *VehicleType) GetLifespan() Lifespan {
	if o == nil || IsNil(o.Lifespan) {
		var ret Lifespan
		return ret
	}
	return *o.Lifespan
}

// GetLifespanOk returns a tuple with the Lifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleType) GetLifespanOk() (*Lifespan, bool) {
	if o == nil || IsNil(o.Lifespan) {
		return nil, false
	}
	return o.Lifespan, true
}

// HasLifespan returns a boolean if a field has been set.
func (o *VehicleType) HasLifespan() bool {
	if o != nil && !IsNil(o.Lifespan) {
		return true
	}

	return false
}

// SetLifespan gets a reference to the given Lifespan and assigns it to the Lifespan field.
func (o *VehicleType) SetLifespan(v Lifespan) {
	o.Lifespan = &v
}

// GetKTypeNo returns the KTypeNo field value if set, zero value otherwise.
func (o *VehicleType) GetKTypeNo() int32 {
	if o == nil || IsNil(o.KTypeNo) {
		var ret int32
		return ret
	}
	return *o.KTypeNo
}

// GetKTypeNoOk returns a tuple with the KTypeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleType) GetKTypeNoOk() (*int32, bool) {
	if o == nil || IsNil(o.KTypeNo) {
		return nil, false
	}
	return o.KTypeNo, true
}

// HasKTypeNo returns a boolean if a field has been set.
func (o *VehicleType) HasKTypeNo() bool {
	if o != nil && !IsNil(o.KTypeNo) {
		return true
	}

	return false
}

// SetKTypeNo gets a reference to the given int32 and assigns it to the KTypeNo field.
func (o *VehicleType) SetKTypeNo(v int32) {
	o.KTypeNo = &v
}

func (o VehicleType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VehicleType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalKTypeNos) {
		toSerialize["additionalKTypeNos"] = o.AdditionalKTypeNos
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Lifespan) {
		toSerialize["lifespan"] = o.Lifespan
	}
	if !IsNil(o.KTypeNo) {
		toSerialize["kTypeNo"] = o.KTypeNo
	}
	return toSerialize, nil
}

type NullableVehicleType struct {
	value *VehicleType
	isSet bool
}

func (v NullableVehicleType) Get() *VehicleType {
	return v.value
}

func (v *NullableVehicleType) Set(val *VehicleType) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleType) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleType(val *VehicleType) *NullableVehicleType {
	return &NullableVehicleType{value: val, isSet: true}
}

func (v NullableVehicleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

