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

// checks if the LocationsSystemDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationsSystemDTO{}

// LocationsSystemDTO struct for LocationsSystemDTO
type LocationsSystemDTO struct {
	// System Id
	SystemId *string `json:"systemId,omitempty"`
	SystemName *Texts `json:"systemName,omitempty"`
	// Components
	Components []ComponentDTO `json:"components,omitempty"`
}

// NewLocationsSystemDTO instantiates a new LocationsSystemDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationsSystemDTO() *LocationsSystemDTO {
	this := LocationsSystemDTO{}
	return &this
}

// NewLocationsSystemDTOWithDefaults instantiates a new LocationsSystemDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationsSystemDTOWithDefaults() *LocationsSystemDTO {
	this := LocationsSystemDTO{}
	return &this
}

// GetSystemId returns the SystemId field value if set, zero value otherwise.
func (o *LocationsSystemDTO) GetSystemId() string {
	if o == nil || IsNil(o.SystemId) {
		var ret string
		return ret
	}
	return *o.SystemId
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsSystemDTO) GetSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemId) {
		return nil, false
	}
	return o.SystemId, true
}

// HasSystemId returns a boolean if a field has been set.
func (o *LocationsSystemDTO) HasSystemId() bool {
	if o != nil && !IsNil(o.SystemId) {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given string and assigns it to the SystemId field.
func (o *LocationsSystemDTO) SetSystemId(v string) {
	o.SystemId = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *LocationsSystemDTO) GetSystemName() Texts {
	if o == nil || IsNil(o.SystemName) {
		var ret Texts
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsSystemDTO) GetSystemNameOk() (*Texts, bool) {
	if o == nil || IsNil(o.SystemName) {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *LocationsSystemDTO) HasSystemName() bool {
	if o != nil && !IsNil(o.SystemName) {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given Texts and assigns it to the SystemName field.
func (o *LocationsSystemDTO) SetSystemName(v Texts) {
	o.SystemName = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *LocationsSystemDTO) GetComponents() []ComponentDTO {
	if o == nil || IsNil(o.Components) {
		var ret []ComponentDTO
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsSystemDTO) GetComponentsOk() ([]ComponentDTO, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *LocationsSystemDTO) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []ComponentDTO and assigns it to the Components field.
func (o *LocationsSystemDTO) SetComponents(v []ComponentDTO) {
	o.Components = v
}

func (o LocationsSystemDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationsSystemDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SystemId) {
		toSerialize["systemId"] = o.SystemId
	}
	if !IsNil(o.SystemName) {
		toSerialize["systemName"] = o.SystemName
	}
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	return toSerialize, nil
}

type NullableLocationsSystemDTO struct {
	value *LocationsSystemDTO
	isSet bool
}

func (v NullableLocationsSystemDTO) Get() *LocationsSystemDTO {
	return v.value
}

func (v *NullableLocationsSystemDTO) Set(val *LocationsSystemDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationsSystemDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationsSystemDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationsSystemDTO(val *LocationsSystemDTO) *NullableLocationsSystemDTO {
	return &NullableLocationsSystemDTO{value: val, isSet: true}
}

func (v NullableLocationsSystemDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationsSystemDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

