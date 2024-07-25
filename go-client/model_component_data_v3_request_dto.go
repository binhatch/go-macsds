/*
Hella Gutmann - macsDS (Data Services) - API collection

This document contains all relevant APIs for diagnostics (incl. DTCs), repair & maintenance information (RMI) and vehicle identification.

API version: V20240702-130718
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package go-macds

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ComponentDataV3RequestDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentDataV3RequestDTO{}

// ComponentDataV3RequestDTO struct for ComponentDataV3RequestDTO
type ComponentDataV3RequestDTO struct {
	// KType
	Ktype int32 `json:"ktype"`
	// Two characters defining the language code; allowed values: de, en, nl, zh, it, hu, hr, fr, fi, es, el, tr, da, sv, sk, cs, ru, ro, pt, pl
	Language string `json:"language"`
	// System identification
	SystemIdent string `json:"systemIdent"`
	// Component Id
	ComponentId int32 `json:"componentId"`
	// Year From
	YearFrom *string `json:"yearFrom,omitempty"`
	// Year To
	YearTo *string `json:"yearTo,omitempty"`
	// Engine Code
	EngineCode *string `json:"engineCode,omitempty"`
}

type _ComponentDataV3RequestDTO ComponentDataV3RequestDTO

// NewComponentDataV3RequestDTO instantiates a new ComponentDataV3RequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentDataV3RequestDTO(ktype int32, language string, systemIdent string, componentId int32) *ComponentDataV3RequestDTO {
	this := ComponentDataV3RequestDTO{}
	this.Ktype = ktype
	this.Language = language
	this.SystemIdent = systemIdent
	this.ComponentId = componentId
	return &this
}

// NewComponentDataV3RequestDTOWithDefaults instantiates a new ComponentDataV3RequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentDataV3RequestDTOWithDefaults() *ComponentDataV3RequestDTO {
	this := ComponentDataV3RequestDTO{}
	return &this
}

// GetKtype returns the Ktype field value
func (o *ComponentDataV3RequestDTO) GetKtype() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ktype
}

// GetKtypeOk returns a tuple with the Ktype field value
// and a boolean to check if the value has been set.
func (o *ComponentDataV3RequestDTO) GetKtypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ktype, true
}

// SetKtype sets field value
func (o *ComponentDataV3RequestDTO) SetKtype(v int32) {
	o.Ktype = v
}

// GetLanguage returns the Language field value
func (o *ComponentDataV3RequestDTO) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *ComponentDataV3RequestDTO) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *ComponentDataV3RequestDTO) SetLanguage(v string) {
	o.Language = v
}

// GetSystemIdent returns the SystemIdent field value
func (o *ComponentDataV3RequestDTO) GetSystemIdent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemIdent
}

// GetSystemIdentOk returns a tuple with the SystemIdent field value
// and a boolean to check if the value has been set.
func (o *ComponentDataV3RequestDTO) GetSystemIdentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemIdent, true
}

// SetSystemIdent sets field value
func (o *ComponentDataV3RequestDTO) SetSystemIdent(v string) {
	o.SystemIdent = v
}

// GetComponentId returns the ComponentId field value
func (o *ComponentDataV3RequestDTO) GetComponentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ComponentId
}

// GetComponentIdOk returns a tuple with the ComponentId field value
// and a boolean to check if the value has been set.
func (o *ComponentDataV3RequestDTO) GetComponentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentId, true
}

// SetComponentId sets field value
func (o *ComponentDataV3RequestDTO) SetComponentId(v int32) {
	o.ComponentId = v
}

// GetYearFrom returns the YearFrom field value if set, zero value otherwise.
func (o *ComponentDataV3RequestDTO) GetYearFrom() string {
	if o == nil || IsNil(o.YearFrom) {
		var ret string
		return ret
	}
	return *o.YearFrom
}

// GetYearFromOk returns a tuple with the YearFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentDataV3RequestDTO) GetYearFromOk() (*string, bool) {
	if o == nil || IsNil(o.YearFrom) {
		return nil, false
	}
	return o.YearFrom, true
}

// HasYearFrom returns a boolean if a field has been set.
func (o *ComponentDataV3RequestDTO) HasYearFrom() bool {
	if o != nil && !IsNil(o.YearFrom) {
		return true
	}

	return false
}

// SetYearFrom gets a reference to the given string and assigns it to the YearFrom field.
func (o *ComponentDataV3RequestDTO) SetYearFrom(v string) {
	o.YearFrom = &v
}

// GetYearTo returns the YearTo field value if set, zero value otherwise.
func (o *ComponentDataV3RequestDTO) GetYearTo() string {
	if o == nil || IsNil(o.YearTo) {
		var ret string
		return ret
	}
	return *o.YearTo
}

// GetYearToOk returns a tuple with the YearTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentDataV3RequestDTO) GetYearToOk() (*string, bool) {
	if o == nil || IsNil(o.YearTo) {
		return nil, false
	}
	return o.YearTo, true
}

// HasYearTo returns a boolean if a field has been set.
func (o *ComponentDataV3RequestDTO) HasYearTo() bool {
	if o != nil && !IsNil(o.YearTo) {
		return true
	}

	return false
}

// SetYearTo gets a reference to the given string and assigns it to the YearTo field.
func (o *ComponentDataV3RequestDTO) SetYearTo(v string) {
	o.YearTo = &v
}

// GetEngineCode returns the EngineCode field value if set, zero value otherwise.
func (o *ComponentDataV3RequestDTO) GetEngineCode() string {
	if o == nil || IsNil(o.EngineCode) {
		var ret string
		return ret
	}
	return *o.EngineCode
}

// GetEngineCodeOk returns a tuple with the EngineCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentDataV3RequestDTO) GetEngineCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EngineCode) {
		return nil, false
	}
	return o.EngineCode, true
}

// HasEngineCode returns a boolean if a field has been set.
func (o *ComponentDataV3RequestDTO) HasEngineCode() bool {
	if o != nil && !IsNil(o.EngineCode) {
		return true
	}

	return false
}

// SetEngineCode gets a reference to the given string and assigns it to the EngineCode field.
func (o *ComponentDataV3RequestDTO) SetEngineCode(v string) {
	o.EngineCode = &v
}

func (o ComponentDataV3RequestDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentDataV3RequestDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ktype"] = o.Ktype
	toSerialize["language"] = o.Language
	toSerialize["systemIdent"] = o.SystemIdent
	toSerialize["componentId"] = o.ComponentId
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

func (o *ComponentDataV3RequestDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ktype",
		"language",
		"systemIdent",
		"componentId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varComponentDataV3RequestDTO := _ComponentDataV3RequestDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varComponentDataV3RequestDTO)

	if err != nil {
		return err
	}

	*o = ComponentDataV3RequestDTO(varComponentDataV3RequestDTO)

	return err
}

type NullableComponentDataV3RequestDTO struct {
	value *ComponentDataV3RequestDTO
	isSet bool
}

func (v NullableComponentDataV3RequestDTO) Get() *ComponentDataV3RequestDTO {
	return v.value
}

func (v *NullableComponentDataV3RequestDTO) Set(val *ComponentDataV3RequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentDataV3RequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentDataV3RequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentDataV3RequestDTO(val *ComponentDataV3RequestDTO) *NullableComponentDataV3RequestDTO {
	return &NullableComponentDataV3RequestDTO{value: val, isSet: true}
}

func (v NullableComponentDataV3RequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentDataV3RequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


