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

// checks if the Attributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Attributes{}

// Attributes struct for Attributes
type Attributes struct {
	// id
	Id *string `json:"id,omitempty"`
	// type
	Type *string `json:"type,omitempty"`
	// returns an alt-term for the image
	Alt *string `json:"alt,omitempty"`
	// fullscreen
	Fullscreen *string `json:"fullscreen,omitempty"`
	// indicates the mime type of the file
	Mimetype *string `json:"mimetype,omitempty"`
}

// NewAttributes instantiates a new Attributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributes() *Attributes {
	this := Attributes{}
	return &this
}

// NewAttributesWithDefaults instantiates a new Attributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributesWithDefaults() *Attributes {
	this := Attributes{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Attributes) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attributes) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Attributes) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Attributes) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Attributes) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attributes) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Attributes) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Attributes) SetType(v string) {
	o.Type = &v
}

// GetAlt returns the Alt field value if set, zero value otherwise.
func (o *Attributes) GetAlt() string {
	if o == nil || IsNil(o.Alt) {
		var ret string
		return ret
	}
	return *o.Alt
}

// GetAltOk returns a tuple with the Alt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attributes) GetAltOk() (*string, bool) {
	if o == nil || IsNil(o.Alt) {
		return nil, false
	}
	return o.Alt, true
}

// HasAlt returns a boolean if a field has been set.
func (o *Attributes) HasAlt() bool {
	if o != nil && !IsNil(o.Alt) {
		return true
	}

	return false
}

// SetAlt gets a reference to the given string and assigns it to the Alt field.
func (o *Attributes) SetAlt(v string) {
	o.Alt = &v
}

// GetFullscreen returns the Fullscreen field value if set, zero value otherwise.
func (o *Attributes) GetFullscreen() string {
	if o == nil || IsNil(o.Fullscreen) {
		var ret string
		return ret
	}
	return *o.Fullscreen
}

// GetFullscreenOk returns a tuple with the Fullscreen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attributes) GetFullscreenOk() (*string, bool) {
	if o == nil || IsNil(o.Fullscreen) {
		return nil, false
	}
	return o.Fullscreen, true
}

// HasFullscreen returns a boolean if a field has been set.
func (o *Attributes) HasFullscreen() bool {
	if o != nil && !IsNil(o.Fullscreen) {
		return true
	}

	return false
}

// SetFullscreen gets a reference to the given string and assigns it to the Fullscreen field.
func (o *Attributes) SetFullscreen(v string) {
	o.Fullscreen = &v
}

// GetMimetype returns the Mimetype field value if set, zero value otherwise.
func (o *Attributes) GetMimetype() string {
	if o == nil || IsNil(o.Mimetype) {
		var ret string
		return ret
	}
	return *o.Mimetype
}

// GetMimetypeOk returns a tuple with the Mimetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attributes) GetMimetypeOk() (*string, bool) {
	if o == nil || IsNil(o.Mimetype) {
		return nil, false
	}
	return o.Mimetype, true
}

// HasMimetype returns a boolean if a field has been set.
func (o *Attributes) HasMimetype() bool {
	if o != nil && !IsNil(o.Mimetype) {
		return true
	}

	return false
}

// SetMimetype gets a reference to the given string and assigns it to the Mimetype field.
func (o *Attributes) SetMimetype(v string) {
	o.Mimetype = &v
}

func (o Attributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Attributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Alt) {
		toSerialize["alt"] = o.Alt
	}
	if !IsNil(o.Fullscreen) {
		toSerialize["fullscreen"] = o.Fullscreen
	}
	if !IsNil(o.Mimetype) {
		toSerialize["mimetype"] = o.Mimetype
	}
	return toSerialize, nil
}

type NullableAttributes struct {
	value *Attributes
	isSet bool
}

func (v NullableAttributes) Get() *Attributes {
	return v.value
}

func (v *NullableAttributes) Set(val *Attributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributes(val *Attributes) *NullableAttributes {
	return &NullableAttributes{value: val, isSet: true}
}

func (v NullableAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


