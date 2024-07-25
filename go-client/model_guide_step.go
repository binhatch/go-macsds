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

// checks if the GuideStep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuideStep{}

// GuideStep struct for GuideStep
type GuideStep struct {
	// Id
	Id *int64 `json:"id,omitempty"`
	// Type
	Type *string `json:"type,omitempty"`
	// Order
	Order *int64 `json:"order,omitempty"`
	Title *TitleText `json:"title,omitempty"`
	Icon *Icon `json:"icon,omitempty"`
	// Lines
	Lines []Line `json:"lines,omitempty"`
	// Images
	Images []OrderedImage `json:"images,omitempty"`
}

// NewGuideStep instantiates a new GuideStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuideStep() *GuideStep {
	this := GuideStep{}
	return &this
}

// NewGuideStepWithDefaults instantiates a new GuideStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuideStepWithDefaults() *GuideStep {
	this := GuideStep{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GuideStep) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuideStep) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GuideStep) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GuideStep) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GuideStep) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuideStep) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GuideStep) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GuideStep) SetType(v string) {
	o.Type = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GuideStep) GetOrder() int64 {
	if o == nil || IsNil(o.Order) {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuideStep) GetOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GuideStep) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *GuideStep) SetOrder(v int64) {
	o.Order = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GuideStep) GetTitle() TitleText {
	if o == nil || IsNil(o.Title) {
		var ret TitleText
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuideStep) GetTitleOk() (*TitleText, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GuideStep) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given TitleText and assigns it to the Title field.
func (o *GuideStep) SetTitle(v TitleText) {
	o.Title = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *GuideStep) GetIcon() Icon {
	if o == nil || IsNil(o.Icon) {
		var ret Icon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuideStep) GetIconOk() (*Icon, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *GuideStep) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given Icon and assigns it to the Icon field.
func (o *GuideStep) SetIcon(v Icon) {
	o.Icon = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *GuideStep) GetLines() []Line {
	if o == nil || IsNil(o.Lines) {
		var ret []Line
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuideStep) GetLinesOk() ([]Line, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *GuideStep) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []Line and assigns it to the Lines field.
func (o *GuideStep) SetLines(v []Line) {
	o.Lines = v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *GuideStep) GetImages() []OrderedImage {
	if o == nil || IsNil(o.Images) {
		var ret []OrderedImage
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuideStep) GetImagesOk() ([]OrderedImage, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *GuideStep) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []OrderedImage and assigns it to the Images field.
func (o *GuideStep) SetImages(v []OrderedImage) {
	o.Images = v
}

func (o GuideStep) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuideStep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	return toSerialize, nil
}

type NullableGuideStep struct {
	value *GuideStep
	isSet bool
}

func (v NullableGuideStep) Get() *GuideStep {
	return v.value
}

func (v *NullableGuideStep) Set(val *GuideStep) {
	v.value = val
	v.isSet = true
}

func (v NullableGuideStep) IsSet() bool {
	return v.isSet
}

func (v *NullableGuideStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuideStep(val *GuideStep) *NullableGuideStep {
	return &NullableGuideStep{value: val, isSet: true}
}

func (v NullableGuideStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuideStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


