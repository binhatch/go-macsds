/*
Hella Gutmann - macsDS (Data Services) - API collection

This document contains all relevant APIs for diagnostics (incl. DTCs), repair & maintenance information (RMI) and vehicle identification.

API version: V20240702-130718
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gomacds

import (
	"encoding/json"
)

// checks if the ComponentNameDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentNameDTO{}

// ComponentNameDTO Component Name
type ComponentNameDTO struct {
	// Id
	Id *int32 `json:"id,omitempty"`
	// Two characters defining the language code; allowed values: de, en, nl, zh, it, hu, hr, fr, fi, es, el, tr, da, sv, sk, cs, ru, ro, pt, pl
	Language *string `json:"language,omitempty"`
	// System text
	Text *string `json:"text,omitempty"`
}

// NewComponentNameDTO instantiates a new ComponentNameDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentNameDTO() *ComponentNameDTO {
	this := ComponentNameDTO{}
	return &this
}

// NewComponentNameDTOWithDefaults instantiates a new ComponentNameDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentNameDTOWithDefaults() *ComponentNameDTO {
	this := ComponentNameDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComponentNameDTO) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentNameDTO) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComponentNameDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ComponentNameDTO) SetId(v int32) {
	o.Id = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ComponentNameDTO) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentNameDTO) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ComponentNameDTO) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ComponentNameDTO) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ComponentNameDTO) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentNameDTO) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ComponentNameDTO) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *ComponentNameDTO) SetText(v string) {
	o.Text = &v
}

func (o ComponentNameDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentNameDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableComponentNameDTO struct {
	value *ComponentNameDTO
	isSet bool
}

func (v NullableComponentNameDTO) Get() *ComponentNameDTO {
	return v.value
}

func (v *NullableComponentNameDTO) Set(val *ComponentNameDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentNameDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentNameDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentNameDTO(val *ComponentNameDTO) *NullableComponentNameDTO {
	return &NullableComponentNameDTO{value: val, isSet: true}
}

func (v NullableComponentNameDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentNameDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


