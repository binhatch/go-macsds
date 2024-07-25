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

// checks if the Unit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Unit{}

// Unit struct for Unit
type Unit struct {
	// Unit Type
	UnitType *string `json:"unitType,omitempty"`
	// Minimum Value
	MinValue *int32 `json:"minValue,omitempty"`
	// Maximum Value
	MaxValue *int32 `json:"maxValue,omitempty"`
}

// NewUnit instantiates a new Unit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnit() *Unit {
	this := Unit{}
	return &this
}

// NewUnitWithDefaults instantiates a new Unit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnitWithDefaults() *Unit {
	this := Unit{}
	return &this
}

// GetUnitType returns the UnitType field value if set, zero value otherwise.
func (o *Unit) GetUnitType() string {
	if o == nil || IsNil(o.UnitType) {
		var ret string
		return ret
	}
	return *o.UnitType
}

// GetUnitTypeOk returns a tuple with the UnitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unit) GetUnitTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UnitType) {
		return nil, false
	}
	return o.UnitType, true
}

// HasUnitType returns a boolean if a field has been set.
func (o *Unit) HasUnitType() bool {
	if o != nil && !IsNil(o.UnitType) {
		return true
	}

	return false
}

// SetUnitType gets a reference to the given string and assigns it to the UnitType field.
func (o *Unit) SetUnitType(v string) {
	o.UnitType = &v
}

// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *Unit) GetMinValue() int32 {
	if o == nil || IsNil(o.MinValue) {
		var ret int32
		return ret
	}
	return *o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unit) GetMinValueOk() (*int32, bool) {
	if o == nil || IsNil(o.MinValue) {
		return nil, false
	}
	return o.MinValue, true
}

// HasMinValue returns a boolean if a field has been set.
func (o *Unit) HasMinValue() bool {
	if o != nil && !IsNil(o.MinValue) {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given int32 and assigns it to the MinValue field.
func (o *Unit) SetMinValue(v int32) {
	o.MinValue = &v
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *Unit) GetMaxValue() int32 {
	if o == nil || IsNil(o.MaxValue) {
		var ret int32
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unit) GetMaxValueOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxValue) {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *Unit) HasMaxValue() bool {
	if o != nil && !IsNil(o.MaxValue) {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given int32 and assigns it to the MaxValue field.
func (o *Unit) SetMaxValue(v int32) {
	o.MaxValue = &v
}

func (o Unit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Unit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UnitType) {
		toSerialize["unitType"] = o.UnitType
	}
	if !IsNil(o.MinValue) {
		toSerialize["minValue"] = o.MinValue
	}
	if !IsNil(o.MaxValue) {
		toSerialize["maxValue"] = o.MaxValue
	}
	return toSerialize, nil
}

type NullableUnit struct {
	value *Unit
	isSet bool
}

func (v NullableUnit) Get() *Unit {
	return v.value
}

func (v *NullableUnit) Set(val *Unit) {
	v.value = val
	v.isSet = true
}

func (v NullableUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnit(val *Unit) *NullableUnit {
	return &NullableUnit{value: val, isSet: true}
}

func (v NullableUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


