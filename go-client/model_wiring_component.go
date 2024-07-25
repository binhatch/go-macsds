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

// checks if the WiringComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WiringComponent{}

// WiringComponent struct for WiringComponent
type WiringComponent struct {
	// Type of the component
	Type *string `json:"type,omitempty"`
	// Id of the component
	Id *string `json:"id,omitempty"`
	Group *Translation `json:"group,omitempty"`
	Text *Translation `json:"text,omitempty"`
}

// NewWiringComponent instantiates a new WiringComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWiringComponent() *WiringComponent {
	this := WiringComponent{}
	return &this
}

// NewWiringComponentWithDefaults instantiates a new WiringComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWiringComponentWithDefaults() *WiringComponent {
	this := WiringComponent{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WiringComponent) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WiringComponent) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WiringComponent) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WiringComponent) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WiringComponent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WiringComponent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WiringComponent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WiringComponent) SetId(v string) {
	o.Id = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *WiringComponent) GetGroup() Translation {
	if o == nil || IsNil(o.Group) {
		var ret Translation
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WiringComponent) GetGroupOk() (*Translation, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *WiringComponent) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given Translation and assigns it to the Group field.
func (o *WiringComponent) SetGroup(v Translation) {
	o.Group = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *WiringComponent) GetText() Translation {
	if o == nil || IsNil(o.Text) {
		var ret Translation
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WiringComponent) GetTextOk() (*Translation, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *WiringComponent) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Translation and assigns it to the Text field.
func (o *WiringComponent) SetText(v Translation) {
	o.Text = &v
}

func (o WiringComponent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WiringComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableWiringComponent struct {
	value *WiringComponent
	isSet bool
}

func (v NullableWiringComponent) Get() *WiringComponent {
	return v.value
}

func (v *NullableWiringComponent) Set(val *WiringComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableWiringComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableWiringComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWiringComponent(val *WiringComponent) *NullableWiringComponent {
	return &NullableWiringComponent{value: val, isSet: true}
}

func (v NullableWiringComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWiringComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

