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

// checks if the ManualImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualImage{}

// ManualImage struct for ManualImage
type ManualImage struct {
	// Image Name
	ImageName *string `json:"imageName,omitempty"`
	// Content
	Content *string `json:"content,omitempty"`
}

// NewManualImage instantiates a new ManualImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualImage() *ManualImage {
	this := ManualImage{}
	return &this
}

// NewManualImageWithDefaults instantiates a new ManualImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualImageWithDefaults() *ManualImage {
	this := ManualImage{}
	return &this
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *ManualImage) GetImageName() string {
	if o == nil || IsNil(o.ImageName) {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImage) GetImageNameOk() (*string, bool) {
	if o == nil || IsNil(o.ImageName) {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *ManualImage) HasImageName() bool {
	if o != nil && !IsNil(o.ImageName) {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *ManualImage) SetImageName(v string) {
	o.ImageName = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ManualImage) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImage) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ManualImage) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ManualImage) SetContent(v string) {
	o.Content = &v
}

func (o ManualImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageName) {
		toSerialize["imageName"] = o.ImageName
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableManualImage struct {
	value *ManualImage
	isSet bool
}

func (v NullableManualImage) Get() *ManualImage {
	return v.value
}

func (v *NullableManualImage) Set(val *ManualImage) {
	v.value = val
	v.isSet = true
}

func (v NullableManualImage) IsSet() bool {
	return v.isSet
}

func (v *NullableManualImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualImage(val *ManualImage) *NullableManualImage {
	return &NullableManualImage{value: val, isSet: true}
}

func (v NullableManualImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


