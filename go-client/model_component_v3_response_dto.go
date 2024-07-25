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

// checks if the ComponentV3ResponseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentV3ResponseDTO{}

// ComponentV3ResponseDTO struct for ComponentV3ResponseDTO
type ComponentV3ResponseDTO struct {
	// Id
	Id *int32 `json:"id,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// System Identification
	SystemIdent *string `json:"systemIdent,omitempty"`
	// Year From
	YearFrom *string `json:"yearFrom,omitempty"`
	// Year To
	YearTo *string `json:"yearTo,omitempty"`
	// Engine Code
	EngineCode *string `json:"engineCode,omitempty"`
}

// NewComponentV3ResponseDTO instantiates a new ComponentV3ResponseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentV3ResponseDTO() *ComponentV3ResponseDTO {
	this := ComponentV3ResponseDTO{}
	return &this
}

// NewComponentV3ResponseDTOWithDefaults instantiates a new ComponentV3ResponseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentV3ResponseDTOWithDefaults() *ComponentV3ResponseDTO {
	this := ComponentV3ResponseDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComponentV3ResponseDTO) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentV3ResponseDTO) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComponentV3ResponseDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ComponentV3ResponseDTO) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComponentV3ResponseDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentV3ResponseDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComponentV3ResponseDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComponentV3ResponseDTO) SetName(v string) {
	o.Name = &v
}

// GetSystemIdent returns the SystemIdent field value if set, zero value otherwise.
func (o *ComponentV3ResponseDTO) GetSystemIdent() string {
	if o == nil || IsNil(o.SystemIdent) {
		var ret string
		return ret
	}
	return *o.SystemIdent
}

// GetSystemIdentOk returns a tuple with the SystemIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentV3ResponseDTO) GetSystemIdentOk() (*string, bool) {
	if o == nil || IsNil(o.SystemIdent) {
		return nil, false
	}
	return o.SystemIdent, true
}

// HasSystemIdent returns a boolean if a field has been set.
func (o *ComponentV3ResponseDTO) HasSystemIdent() bool {
	if o != nil && !IsNil(o.SystemIdent) {
		return true
	}

	return false
}

// SetSystemIdent gets a reference to the given string and assigns it to the SystemIdent field.
func (o *ComponentV3ResponseDTO) SetSystemIdent(v string) {
	o.SystemIdent = &v
}

// GetYearFrom returns the YearFrom field value if set, zero value otherwise.
func (o *ComponentV3ResponseDTO) GetYearFrom() string {
	if o == nil || IsNil(o.YearFrom) {
		var ret string
		return ret
	}
	return *o.YearFrom
}

// GetYearFromOk returns a tuple with the YearFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentV3ResponseDTO) GetYearFromOk() (*string, bool) {
	if o == nil || IsNil(o.YearFrom) {
		return nil, false
	}
	return o.YearFrom, true
}

// HasYearFrom returns a boolean if a field has been set.
func (o *ComponentV3ResponseDTO) HasYearFrom() bool {
	if o != nil && !IsNil(o.YearFrom) {
		return true
	}

	return false
}

// SetYearFrom gets a reference to the given string and assigns it to the YearFrom field.
func (o *ComponentV3ResponseDTO) SetYearFrom(v string) {
	o.YearFrom = &v
}

// GetYearTo returns the YearTo field value if set, zero value otherwise.
func (o *ComponentV3ResponseDTO) GetYearTo() string {
	if o == nil || IsNil(o.YearTo) {
		var ret string
		return ret
	}
	return *o.YearTo
}

// GetYearToOk returns a tuple with the YearTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentV3ResponseDTO) GetYearToOk() (*string, bool) {
	if o == nil || IsNil(o.YearTo) {
		return nil, false
	}
	return o.YearTo, true
}

// HasYearTo returns a boolean if a field has been set.
func (o *ComponentV3ResponseDTO) HasYearTo() bool {
	if o != nil && !IsNil(o.YearTo) {
		return true
	}

	return false
}

// SetYearTo gets a reference to the given string and assigns it to the YearTo field.
func (o *ComponentV3ResponseDTO) SetYearTo(v string) {
	o.YearTo = &v
}

// GetEngineCode returns the EngineCode field value if set, zero value otherwise.
func (o *ComponentV3ResponseDTO) GetEngineCode() string {
	if o == nil || IsNil(o.EngineCode) {
		var ret string
		return ret
	}
	return *o.EngineCode
}

// GetEngineCodeOk returns a tuple with the EngineCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentV3ResponseDTO) GetEngineCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EngineCode) {
		return nil, false
	}
	return o.EngineCode, true
}

// HasEngineCode returns a boolean if a field has been set.
func (o *ComponentV3ResponseDTO) HasEngineCode() bool {
	if o != nil && !IsNil(o.EngineCode) {
		return true
	}

	return false
}

// SetEngineCode gets a reference to the given string and assigns it to the EngineCode field.
func (o *ComponentV3ResponseDTO) SetEngineCode(v string) {
	o.EngineCode = &v
}

func (o ComponentV3ResponseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentV3ResponseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SystemIdent) {
		toSerialize["systemIdent"] = o.SystemIdent
	}
	if !IsNil(o.YearFrom) {
		toSerialize["yearFrom"] = o.YearFrom
	}
	if !IsNil(o.YearTo) {
		toSerialize["yearTo"] = o.YearTo
	}
	if !IsNil(o.EngineCode) {
		toSerialize["engineCode"] = o.EngineCode
	}
	return toSerialize, nil
}

type NullableComponentV3ResponseDTO struct {
	value *ComponentV3ResponseDTO
	isSet bool
}

func (v NullableComponentV3ResponseDTO) Get() *ComponentV3ResponseDTO {
	return v.value
}

func (v *NullableComponentV3ResponseDTO) Set(val *ComponentV3ResponseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentV3ResponseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentV3ResponseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentV3ResponseDTO(val *ComponentV3ResponseDTO) *NullableComponentV3ResponseDTO {
	return &NullableComponentV3ResponseDTO{value: val, isSet: true}
}

func (v NullableComponentV3ResponseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentV3ResponseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


