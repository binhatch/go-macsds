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

// checks if the TextContentBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextContentBlock{}

// TextContentBlock struct for TextContentBlock
type TextContentBlock struct {
	// Order
	Order *int32 `json:"order,omitempty"`
	// Suggestion of the text format for the respective content block. Values can be e.g. \"text-h1\", \"text-body-1\", \"text-body-2, \"text-caption\", etc.
	TextFormatType *string `json:"textFormatType,omitempty"`
	// Text Content
	TextContent *string `json:"textContent,omitempty"`
	// Text Type
	Type *string `json:"type,omitempty"`
}

// NewTextContentBlock instantiates a new TextContentBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextContentBlock() *TextContentBlock {
	this := TextContentBlock{}
	return &this
}

// NewTextContentBlockWithDefaults instantiates a new TextContentBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextContentBlockWithDefaults() *TextContentBlock {
	this := TextContentBlock{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *TextContentBlock) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextContentBlock) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *TextContentBlock) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *TextContentBlock) SetOrder(v int32) {
	o.Order = &v
}

// GetTextFormatType returns the TextFormatType field value if set, zero value otherwise.
func (o *TextContentBlock) GetTextFormatType() string {
	if o == nil || IsNil(o.TextFormatType) {
		var ret string
		return ret
	}
	return *o.TextFormatType
}

// GetTextFormatTypeOk returns a tuple with the TextFormatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextContentBlock) GetTextFormatTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TextFormatType) {
		return nil, false
	}
	return o.TextFormatType, true
}

// HasTextFormatType returns a boolean if a field has been set.
func (o *TextContentBlock) HasTextFormatType() bool {
	if o != nil && !IsNil(o.TextFormatType) {
		return true
	}

	return false
}

// SetTextFormatType gets a reference to the given string and assigns it to the TextFormatType field.
func (o *TextContentBlock) SetTextFormatType(v string) {
	o.TextFormatType = &v
}

// GetTextContent returns the TextContent field value if set, zero value otherwise.
func (o *TextContentBlock) GetTextContent() string {
	if o == nil || IsNil(o.TextContent) {
		var ret string
		return ret
	}
	return *o.TextContent
}

// GetTextContentOk returns a tuple with the TextContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextContentBlock) GetTextContentOk() (*string, bool) {
	if o == nil || IsNil(o.TextContent) {
		return nil, false
	}
	return o.TextContent, true
}

// HasTextContent returns a boolean if a field has been set.
func (o *TextContentBlock) HasTextContent() bool {
	if o != nil && !IsNil(o.TextContent) {
		return true
	}

	return false
}

// SetTextContent gets a reference to the given string and assigns it to the TextContent field.
func (o *TextContentBlock) SetTextContent(v string) {
	o.TextContent = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TextContentBlock) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextContentBlock) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TextContentBlock) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TextContentBlock) SetType(v string) {
	o.Type = &v
}

func (o TextContentBlock) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextContentBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.TextFormatType) {
		toSerialize["textFormatType"] = o.TextFormatType
	}
	if !IsNil(o.TextContent) {
		toSerialize["textContent"] = o.TextContent
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTextContentBlock struct {
	value *TextContentBlock
	isSet bool
}

func (v NullableTextContentBlock) Get() *TextContentBlock {
	return v.value
}

func (v *NullableTextContentBlock) Set(val *TextContentBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableTextContentBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableTextContentBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextContentBlock(val *TextContentBlock) *NullableTextContentBlock {
	return &NullableTextContentBlock{value: val, isSet: true}
}

func (v NullableTextContentBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextContentBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


