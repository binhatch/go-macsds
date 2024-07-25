/*
Hella Gutmann - macsDS (Data Services) - API collection

This document contains all relevant APIs for diagnostics (incl. DTCs), repair & maintenance information (RMI) and vehicle identification.

API version: V20240702-130718
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/binhatch/go-macds

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the LocationsRequestDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationsRequestDTO{}

// LocationsRequestDTO struct for LocationsRequestDTO
type LocationsRequestDTO struct {
	// KType
	Ktype int32 `json:"ktype"`
	// Two characters defining the language code; allowed values: de, en, nl, zh, it, hu, hr, fr, fi, es, el, tr, da, sv, sk, cs, ru, ro, pt, pl
	Language string `json:"language"`
	// Year From
	YearFrom *string `json:"yearFrom,omitempty"`
	// Year To
	YearTo *string `json:"yearTo,omitempty"`
	// Engine Code
	EngineCode *string `json:"engineCode,omitempty"`
	// Component Id
	ComponentId *int32 `json:"componentId,omitempty"`
	// Image Id
	ImageId *int32 `json:"imageId,omitempty"`
}

type _LocationsRequestDTO LocationsRequestDTO

// NewLocationsRequestDTO instantiates a new LocationsRequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationsRequestDTO(ktype int32, language string) *LocationsRequestDTO {
	this := LocationsRequestDTO{}
	this.Ktype = ktype
	this.Language = language
	return &this
}

// NewLocationsRequestDTOWithDefaults instantiates a new LocationsRequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationsRequestDTOWithDefaults() *LocationsRequestDTO {
	this := LocationsRequestDTO{}
	return &this
}

// GetKtype returns the Ktype field value
func (o *LocationsRequestDTO) GetKtype() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ktype
}

// GetKtypeOk returns a tuple with the Ktype field value
// and a boolean to check if the value has been set.
func (o *LocationsRequestDTO) GetKtypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ktype, true
}

// SetKtype sets field value
func (o *LocationsRequestDTO) SetKtype(v int32) {
	o.Ktype = v
}

// GetLanguage returns the Language field value
func (o *LocationsRequestDTO) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *LocationsRequestDTO) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *LocationsRequestDTO) SetLanguage(v string) {
	o.Language = v
}

// GetYearFrom returns the YearFrom field value if set, zero value otherwise.
func (o *LocationsRequestDTO) GetYearFrom() string {
	if o == nil || IsNil(o.YearFrom) {
		var ret string
		return ret
	}
	return *o.YearFrom
}

// GetYearFromOk returns a tuple with the YearFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsRequestDTO) GetYearFromOk() (*string, bool) {
	if o == nil || IsNil(o.YearFrom) {
		return nil, false
	}
	return o.YearFrom, true
}

// HasYearFrom returns a boolean if a field has been set.
func (o *LocationsRequestDTO) HasYearFrom() bool {
	if o != nil && !IsNil(o.YearFrom) {
		return true
	}

	return false
}

// SetYearFrom gets a reference to the given string and assigns it to the YearFrom field.
func (o *LocationsRequestDTO) SetYearFrom(v string) {
	o.YearFrom = &v
}

// GetYearTo returns the YearTo field value if set, zero value otherwise.
func (o *LocationsRequestDTO) GetYearTo() string {
	if o == nil || IsNil(o.YearTo) {
		var ret string
		return ret
	}
	return *o.YearTo
}

// GetYearToOk returns a tuple with the YearTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsRequestDTO) GetYearToOk() (*string, bool) {
	if o == nil || IsNil(o.YearTo) {
		return nil, false
	}
	return o.YearTo, true
}

// HasYearTo returns a boolean if a field has been set.
func (o *LocationsRequestDTO) HasYearTo() bool {
	if o != nil && !IsNil(o.YearTo) {
		return true
	}

	return false
}

// SetYearTo gets a reference to the given string and assigns it to the YearTo field.
func (o *LocationsRequestDTO) SetYearTo(v string) {
	o.YearTo = &v
}

// GetEngineCode returns the EngineCode field value if set, zero value otherwise.
func (o *LocationsRequestDTO) GetEngineCode() string {
	if o == nil || IsNil(o.EngineCode) {
		var ret string
		return ret
	}
	return *o.EngineCode
}

// GetEngineCodeOk returns a tuple with the EngineCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsRequestDTO) GetEngineCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EngineCode) {
		return nil, false
	}
	return o.EngineCode, true
}

// HasEngineCode returns a boolean if a field has been set.
func (o *LocationsRequestDTO) HasEngineCode() bool {
	if o != nil && !IsNil(o.EngineCode) {
		return true
	}

	return false
}

// SetEngineCode gets a reference to the given string and assigns it to the EngineCode field.
func (o *LocationsRequestDTO) SetEngineCode(v string) {
	o.EngineCode = &v
}

// GetComponentId returns the ComponentId field value if set, zero value otherwise.
func (o *LocationsRequestDTO) GetComponentId() int32 {
	if o == nil || IsNil(o.ComponentId) {
		var ret int32
		return ret
	}
	return *o.ComponentId
}

// GetComponentIdOk returns a tuple with the ComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsRequestDTO) GetComponentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ComponentId) {
		return nil, false
	}
	return o.ComponentId, true
}

// HasComponentId returns a boolean if a field has been set.
func (o *LocationsRequestDTO) HasComponentId() bool {
	if o != nil && !IsNil(o.ComponentId) {
		return true
	}

	return false
}

// SetComponentId gets a reference to the given int32 and assigns it to the ComponentId field.
func (o *LocationsRequestDTO) SetComponentId(v int32) {
	o.ComponentId = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *LocationsRequestDTO) GetImageId() int32 {
	if o == nil || IsNil(o.ImageId) {
		var ret int32
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsRequestDTO) GetImageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *LocationsRequestDTO) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given int32 and assigns it to the ImageId field.
func (o *LocationsRequestDTO) SetImageId(v int32) {
	o.ImageId = &v
}

func (o LocationsRequestDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationsRequestDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ktype"] = o.Ktype
	toSerialize["language"] = o.Language
	if !IsNil(o.YearFrom) {
		toSerialize["yearFrom"] = o.YearFrom
	}
	if !IsNil(o.YearTo) {
		toSerialize["yearTo"] = o.YearTo
	}
	if !IsNil(o.EngineCode) {
		toSerialize["engineCode"] = o.EngineCode
	}
	if !IsNil(o.ComponentId) {
		toSerialize["componentId"] = o.ComponentId
	}
	if !IsNil(o.ImageId) {
		toSerialize["imageId"] = o.ImageId
	}
	return toSerialize, nil
}

func (o *LocationsRequestDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ktype",
		"language",
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

	varLocationsRequestDTO := _LocationsRequestDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLocationsRequestDTO)

	if err != nil {
		return err
	}

	*o = LocationsRequestDTO(varLocationsRequestDTO)

	return err
}

type NullableLocationsRequestDTO struct {
	value *LocationsRequestDTO
	isSet bool
}

func (v NullableLocationsRequestDTO) Get() *LocationsRequestDTO {
	return v.value
}

func (v *NullableLocationsRequestDTO) Set(val *LocationsRequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationsRequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationsRequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationsRequestDTO(val *LocationsRequestDTO) *NullableLocationsRequestDTO {
	return &NullableLocationsRequestDTO{value: val, isSet: true}
}

func (v NullableLocationsRequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationsRequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


