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

// checks if the FeatureImportance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureImportance{}

// FeatureImportance Machine-learning plots and calculation
type FeatureImportance struct {
	// Plot
	Plot *string `json:"plot,omitempty"`
	// Css
	Css *string `json:"css,omitempty"`
	// Sankey diagram showing the relevance of each provided DTC and the component prediction it relates to, encoded as base64
	Sankey []Sankey `json:"sankey,omitempty"`
}

// NewFeatureImportance instantiates a new FeatureImportance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureImportance() *FeatureImportance {
	this := FeatureImportance{}
	return &this
}

// NewFeatureImportanceWithDefaults instantiates a new FeatureImportance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureImportanceWithDefaults() *FeatureImportance {
	this := FeatureImportance{}
	return &this
}

// GetPlot returns the Plot field value if set, zero value otherwise.
func (o *FeatureImportance) GetPlot() string {
	if o == nil || IsNil(o.Plot) {
		var ret string
		return ret
	}
	return *o.Plot
}

// GetPlotOk returns a tuple with the Plot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureImportance) GetPlotOk() (*string, bool) {
	if o == nil || IsNil(o.Plot) {
		return nil, false
	}
	return o.Plot, true
}

// HasPlot returns a boolean if a field has been set.
func (o *FeatureImportance) HasPlot() bool {
	if o != nil && !IsNil(o.Plot) {
		return true
	}

	return false
}

// SetPlot gets a reference to the given string and assigns it to the Plot field.
func (o *FeatureImportance) SetPlot(v string) {
	o.Plot = &v
}

// GetCss returns the Css field value if set, zero value otherwise.
func (o *FeatureImportance) GetCss() string {
	if o == nil || IsNil(o.Css) {
		var ret string
		return ret
	}
	return *o.Css
}

// GetCssOk returns a tuple with the Css field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureImportance) GetCssOk() (*string, bool) {
	if o == nil || IsNil(o.Css) {
		return nil, false
	}
	return o.Css, true
}

// HasCss returns a boolean if a field has been set.
func (o *FeatureImportance) HasCss() bool {
	if o != nil && !IsNil(o.Css) {
		return true
	}

	return false
}

// SetCss gets a reference to the given string and assigns it to the Css field.
func (o *FeatureImportance) SetCss(v string) {
	o.Css = &v
}

// GetSankey returns the Sankey field value if set, zero value otherwise.
func (o *FeatureImportance) GetSankey() []Sankey {
	if o == nil || IsNil(o.Sankey) {
		var ret []Sankey
		return ret
	}
	return o.Sankey
}

// GetSankeyOk returns a tuple with the Sankey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureImportance) GetSankeyOk() ([]Sankey, bool) {
	if o == nil || IsNil(o.Sankey) {
		return nil, false
	}
	return o.Sankey, true
}

// HasSankey returns a boolean if a field has been set.
func (o *FeatureImportance) HasSankey() bool {
	if o != nil && !IsNil(o.Sankey) {
		return true
	}

	return false
}

// SetSankey gets a reference to the given []Sankey and assigns it to the Sankey field.
func (o *FeatureImportance) SetSankey(v []Sankey) {
	o.Sankey = v
}

func (o FeatureImportance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureImportance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plot) {
		toSerialize["plot"] = o.Plot
	}
	if !IsNil(o.Css) {
		toSerialize["css"] = o.Css
	}
	if !IsNil(o.Sankey) {
		toSerialize["sankey"] = o.Sankey
	}
	return toSerialize, nil
}

type NullableFeatureImportance struct {
	value *FeatureImportance
	isSet bool
}

func (v NullableFeatureImportance) Get() *FeatureImportance {
	return v.value
}

func (v *NullableFeatureImportance) Set(val *FeatureImportance) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureImportance) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureImportance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureImportance(val *FeatureImportance) *NullableFeatureImportance {
	return &NullableFeatureImportance{value: val, isSet: true}
}

func (v NullableFeatureImportance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureImportance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

