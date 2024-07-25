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

// checks if the RichTextDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RichTextDTO{}

// RichTextDTO struct for RichTextDTO
type RichTextDTO struct {
	// id
	Id *string `json:"id,omitempty"`
	// Defines the type of text, e.g. \"third-order-headline\", \"table\", \"tbody\", \"trow\", \"tcell\"
	Type *string `json:"type,omitempty"`
	// Text
	Text *string `json:"text,omitempty"`
	// List of Texts
	Texts []string `json:"texts,omitempty"`
	// Text content value
	ContentValue *string `json:"contentValue,omitempty"`
	// List of text content values
	ContentValues []string `json:"contentValues,omitempty"`
	Content *RichTextDTO `json:"content,omitempty"`
	Attributes *Attributes `json:"attributes,omitempty"`
}

// NewRichTextDTO instantiates a new RichTextDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRichTextDTO() *RichTextDTO {
	this := RichTextDTO{}
	return &this
}

// NewRichTextDTOWithDefaults instantiates a new RichTextDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRichTextDTOWithDefaults() *RichTextDTO {
	this := RichTextDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RichTextDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichTextDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RichTextDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RichTextDTO) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RichTextDTO) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichTextDTO) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RichTextDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RichTextDTO) SetType(v string) {
	o.Type = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *RichTextDTO) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichTextDTO) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *RichTextDTO) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *RichTextDTO) SetText(v string) {
	o.Text = &v
}

// GetTexts returns the Texts field value if set, zero value otherwise.
func (o *RichTextDTO) GetTexts() []string {
	if o == nil || IsNil(o.Texts) {
		var ret []string
		return ret
	}
	return o.Texts
}

// GetTextsOk returns a tuple with the Texts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichTextDTO) GetTextsOk() ([]string, bool) {
	if o == nil || IsNil(o.Texts) {
		return nil, false
	}
	return o.Texts, true
}

// HasTexts returns a boolean if a field has been set.
func (o *RichTextDTO) HasTexts() bool {
	if o != nil && !IsNil(o.Texts) {
		return true
	}

	return false
}

// SetTexts gets a reference to the given []string and assigns it to the Texts field.
func (o *RichTextDTO) SetTexts(v []string) {
	o.Texts = v
}

// GetContentValue returns the ContentValue field value if set, zero value otherwise.
func (o *RichTextDTO) GetContentValue() string {
	if o == nil || IsNil(o.ContentValue) {
		var ret string
		return ret
	}
	return *o.ContentValue
}

// GetContentValueOk returns a tuple with the ContentValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichTextDTO) GetContentValueOk() (*string, bool) {
	if o == nil || IsNil(o.ContentValue) {
		return nil, false
	}
	return o.ContentValue, true
}

// HasContentValue returns a boolean if a field has been set.
func (o *RichTextDTO) HasContentValue() bool {
	if o != nil && !IsNil(o.ContentValue) {
		return true
	}

	return false
}

// SetContentValue gets a reference to the given string and assigns it to the ContentValue field.
func (o *RichTextDTO) SetContentValue(v string) {
	o.ContentValue = &v
}

// GetContentValues returns the ContentValues field value if set, zero value otherwise.
func (o *RichTextDTO) GetContentValues() []string {
	if o == nil || IsNil(o.ContentValues) {
		var ret []string
		return ret
	}
	return o.ContentValues
}

// GetContentValuesOk returns a tuple with the ContentValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichTextDTO) GetContentValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentValues) {
		return nil, false
	}
	return o.ContentValues, true
}

// HasContentValues returns a boolean if a field has been set.
func (o *RichTextDTO) HasContentValues() bool {
	if o != nil && !IsNil(o.ContentValues) {
		return true
	}

	return false
}

// SetContentValues gets a reference to the given []string and assigns it to the ContentValues field.
func (o *RichTextDTO) SetContentValues(v []string) {
	o.ContentValues = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *RichTextDTO) GetContent() RichTextDTO {
	if o == nil || IsNil(o.Content) {
		var ret RichTextDTO
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichTextDTO) GetContentOk() (*RichTextDTO, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *RichTextDTO) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given RichTextDTO and assigns it to the Content field.
func (o *RichTextDTO) SetContent(v RichTextDTO) {
	o.Content = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RichTextDTO) GetAttributes() Attributes {
	if o == nil || IsNil(o.Attributes) {
		var ret Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RichTextDTO) GetAttributesOk() (*Attributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RichTextDTO) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Attributes and assigns it to the Attributes field.
func (o *RichTextDTO) SetAttributes(v Attributes) {
	o.Attributes = &v
}

func (o RichTextDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RichTextDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Texts) {
		toSerialize["texts"] = o.Texts
	}
	if !IsNil(o.ContentValue) {
		toSerialize["contentValue"] = o.ContentValue
	}
	if !IsNil(o.ContentValues) {
		toSerialize["contentValues"] = o.ContentValues
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableRichTextDTO struct {
	value *RichTextDTO
	isSet bool
}

func (v NullableRichTextDTO) Get() *RichTextDTO {
	return v.value
}

func (v *NullableRichTextDTO) Set(val *RichTextDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableRichTextDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableRichTextDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRichTextDTO(val *RichTextDTO) *NullableRichTextDTO {
	return &NullableRichTextDTO{value: val, isSet: true}
}

func (v NullableRichTextDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRichTextDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


