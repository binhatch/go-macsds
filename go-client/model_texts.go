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

// checks if the Texts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Texts{}

// Texts struct for Texts
type Texts struct {
	// Translations
	Translations *map[string]string `json:"translations,omitempty"`
}

// NewTexts instantiates a new Texts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTexts() *Texts {
	this := Texts{}
	return &this
}

// NewTextsWithDefaults instantiates a new Texts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextsWithDefaults() *Texts {
	this := Texts{}
	return &this
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *Texts) GetTranslations() map[string]string {
	if o == nil || IsNil(o.Translations) {
		var ret map[string]string
		return ret
	}
	return *o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Texts) GetTranslationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Translations) {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *Texts) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given map[string]string and assigns it to the Translations field.
func (o *Texts) SetTranslations(v map[string]string) {
	o.Translations = &v
}

func (o Texts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Texts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Translations) {
		toSerialize["translations"] = o.Translations
	}
	return toSerialize, nil
}

type NullableTexts struct {
	value *Texts
	isSet bool
}

func (v NullableTexts) Get() *Texts {
	return v.value
}

func (v *NullableTexts) Set(val *Texts) {
	v.value = val
	v.isSet = true
}

func (v NullableTexts) IsSet() bool {
	return v.isSet
}

func (v *NullableTexts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTexts(val *Texts) *NullableTexts {
	return &NullableTexts{value: val, isSet: true}
}

func (v NullableTexts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTexts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

