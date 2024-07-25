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

// checks if the CoordinatesDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoordinatesDTO{}

// CoordinatesDTO struct for CoordinatesDTO
type CoordinatesDTO struct {
	// X1
	X1 *string `json:"x1,omitempty"`
	// Y1
	Y1 *string `json:"y1,omitempty"`
	// X2
	X2 *string `json:"x2,omitempty"`
	// Y2
	Y2 *string `json:"y2,omitempty"`
}

// NewCoordinatesDTO instantiates a new CoordinatesDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoordinatesDTO() *CoordinatesDTO {
	this := CoordinatesDTO{}
	return &this
}

// NewCoordinatesDTOWithDefaults instantiates a new CoordinatesDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoordinatesDTOWithDefaults() *CoordinatesDTO {
	this := CoordinatesDTO{}
	return &this
}

// GetX1 returns the X1 field value if set, zero value otherwise.
func (o *CoordinatesDTO) GetX1() string {
	if o == nil || IsNil(o.X1) {
		var ret string
		return ret
	}
	return *o.X1
}

// GetX1Ok returns a tuple with the X1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoordinatesDTO) GetX1Ok() (*string, bool) {
	if o == nil || IsNil(o.X1) {
		return nil, false
	}
	return o.X1, true
}

// HasX1 returns a boolean if a field has been set.
func (o *CoordinatesDTO) HasX1() bool {
	if o != nil && !IsNil(o.X1) {
		return true
	}

	return false
}

// SetX1 gets a reference to the given string and assigns it to the X1 field.
func (o *CoordinatesDTO) SetX1(v string) {
	o.X1 = &v
}

// GetY1 returns the Y1 field value if set, zero value otherwise.
func (o *CoordinatesDTO) GetY1() string {
	if o == nil || IsNil(o.Y1) {
		var ret string
		return ret
	}
	return *o.Y1
}

// GetY1Ok returns a tuple with the Y1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoordinatesDTO) GetY1Ok() (*string, bool) {
	if o == nil || IsNil(o.Y1) {
		return nil, false
	}
	return o.Y1, true
}

// HasY1 returns a boolean if a field has been set.
func (o *CoordinatesDTO) HasY1() bool {
	if o != nil && !IsNil(o.Y1) {
		return true
	}

	return false
}

// SetY1 gets a reference to the given string and assigns it to the Y1 field.
func (o *CoordinatesDTO) SetY1(v string) {
	o.Y1 = &v
}

// GetX2 returns the X2 field value if set, zero value otherwise.
func (o *CoordinatesDTO) GetX2() string {
	if o == nil || IsNil(o.X2) {
		var ret string
		return ret
	}
	return *o.X2
}

// GetX2Ok returns a tuple with the X2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoordinatesDTO) GetX2Ok() (*string, bool) {
	if o == nil || IsNil(o.X2) {
		return nil, false
	}
	return o.X2, true
}

// HasX2 returns a boolean if a field has been set.
func (o *CoordinatesDTO) HasX2() bool {
	if o != nil && !IsNil(o.X2) {
		return true
	}

	return false
}

// SetX2 gets a reference to the given string and assigns it to the X2 field.
func (o *CoordinatesDTO) SetX2(v string) {
	o.X2 = &v
}

// GetY2 returns the Y2 field value if set, zero value otherwise.
func (o *CoordinatesDTO) GetY2() string {
	if o == nil || IsNil(o.Y2) {
		var ret string
		return ret
	}
	return *o.Y2
}

// GetY2Ok returns a tuple with the Y2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoordinatesDTO) GetY2Ok() (*string, bool) {
	if o == nil || IsNil(o.Y2) {
		return nil, false
	}
	return o.Y2, true
}

// HasY2 returns a boolean if a field has been set.
func (o *CoordinatesDTO) HasY2() bool {
	if o != nil && !IsNil(o.Y2) {
		return true
	}

	return false
}

// SetY2 gets a reference to the given string and assigns it to the Y2 field.
func (o *CoordinatesDTO) SetY2(v string) {
	o.Y2 = &v
}

func (o CoordinatesDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoordinatesDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.X1) {
		toSerialize["x1"] = o.X1
	}
	if !IsNil(o.Y1) {
		toSerialize["y1"] = o.Y1
	}
	if !IsNil(o.X2) {
		toSerialize["x2"] = o.X2
	}
	if !IsNil(o.Y2) {
		toSerialize["y2"] = o.Y2
	}
	return toSerialize, nil
}

type NullableCoordinatesDTO struct {
	value *CoordinatesDTO
	isSet bool
}

func (v NullableCoordinatesDTO) Get() *CoordinatesDTO {
	return v.value
}

func (v *NullableCoordinatesDTO) Set(val *CoordinatesDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCoordinatesDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCoordinatesDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoordinatesDTO(val *CoordinatesDTO) *NullableCoordinatesDTO {
	return &NullableCoordinatesDTO{value: val, isSet: true}
}

func (v NullableCoordinatesDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoordinatesDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


