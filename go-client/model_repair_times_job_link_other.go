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

// checks if the RepairTimesJobLinkOther type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepairTimesJobLinkOther{}

// RepairTimesJobLinkOther struct for RepairTimesJobLinkOther
type RepairTimesJobLinkOther struct {
	// Order
	Order *int32 `json:"order,omitempty"`
	Link *RepairTimesJobLinkDetails `json:"link,omitempty"`
}

// NewRepairTimesJobLinkOther instantiates a new RepairTimesJobLinkOther object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepairTimesJobLinkOther() *RepairTimesJobLinkOther {
	this := RepairTimesJobLinkOther{}
	return &this
}

// NewRepairTimesJobLinkOtherWithDefaults instantiates a new RepairTimesJobLinkOther object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepairTimesJobLinkOtherWithDefaults() *RepairTimesJobLinkOther {
	this := RepairTimesJobLinkOther{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RepairTimesJobLinkOther) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepairTimesJobLinkOther) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RepairTimesJobLinkOther) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *RepairTimesJobLinkOther) SetOrder(v int32) {
	o.Order = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *RepairTimesJobLinkOther) GetLink() RepairTimesJobLinkDetails {
	if o == nil || IsNil(o.Link) {
		var ret RepairTimesJobLinkDetails
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepairTimesJobLinkOther) GetLinkOk() (*RepairTimesJobLinkDetails, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *RepairTimesJobLinkOther) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given RepairTimesJobLinkDetails and assigns it to the Link field.
func (o *RepairTimesJobLinkOther) SetLink(v RepairTimesJobLinkDetails) {
	o.Link = &v
}

func (o RepairTimesJobLinkOther) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepairTimesJobLinkOther) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	return toSerialize, nil
}

type NullableRepairTimesJobLinkOther struct {
	value *RepairTimesJobLinkOther
	isSet bool
}

func (v NullableRepairTimesJobLinkOther) Get() *RepairTimesJobLinkOther {
	return v.value
}

func (v *NullableRepairTimesJobLinkOther) Set(val *RepairTimesJobLinkOther) {
	v.value = val
	v.isSet = true
}

func (v NullableRepairTimesJobLinkOther) IsSet() bool {
	return v.isSet
}

func (v *NullableRepairTimesJobLinkOther) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepairTimesJobLinkOther(val *RepairTimesJobLinkOther) *NullableRepairTimesJobLinkOther {
	return &NullableRepairTimesJobLinkOther{value: val, isSet: true}
}

func (v NullableRepairTimesJobLinkOther) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepairTimesJobLinkOther) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


