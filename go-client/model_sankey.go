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

// checks if the Sankey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sankey{}

// Sankey struct for Sankey
type Sankey struct {
	// Source
	Source *string `json:"source,omitempty"`
	// Source Type
	SourceType *string `json:"sourceType,omitempty"`
	// Targer
	Target *int32 `json:"target,omitempty"`
	// Amount
	Amount *float64 `json:"amount,omitempty"`
}

// NewSankey instantiates a new Sankey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSankey() *Sankey {
	this := Sankey{}
	return &this
}

// NewSankeyWithDefaults instantiates a new Sankey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSankeyWithDefaults() *Sankey {
	this := Sankey{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Sankey) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sankey) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Sankey) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *Sankey) SetSource(v string) {
	o.Source = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *Sankey) GetSourceType() string {
	if o == nil || IsNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sankey) GetSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *Sankey) HasSourceType() bool {
	if o != nil && !IsNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *Sankey) SetSourceType(v string) {
	o.SourceType = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *Sankey) GetTarget() int32 {
	if o == nil || IsNil(o.Target) {
		var ret int32
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sankey) GetTargetOk() (*int32, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *Sankey) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given int32 and assigns it to the Target field.
func (o *Sankey) SetTarget(v int32) {
	o.Target = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Sankey) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sankey) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Sankey) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *Sankey) SetAmount(v float64) {
	o.Amount = &v
}

func (o Sankey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sankey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.SourceType) {
		toSerialize["sourceType"] = o.SourceType
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableSankey struct {
	value *Sankey
	isSet bool
}

func (v NullableSankey) Get() *Sankey {
	return v.value
}

func (v *NullableSankey) Set(val *Sankey) {
	v.value = val
	v.isSet = true
}

func (v NullableSankey) IsSet() bool {
	return v.isSet
}

func (v *NullableSankey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSankey(val *Sankey) *NullableSankey {
	return &NullableSankey{value: val, isSet: true}
}

func (v NullableSankey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSankey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

