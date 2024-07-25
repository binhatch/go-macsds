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

// checks if the RepairTimesTreeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepairTimesTreeResponse{}

// RepairTimesTreeResponse struct for RepairTimesTreeResponse
type RepairTimesTreeResponse struct {
	VehicleReference *RepairTimesVehicleReference `json:"vehicleReference,omitempty"`
	RepairTimesTree *RepairTimesTree `json:"repairTimesTree,omitempty"`
}

// NewRepairTimesTreeResponse instantiates a new RepairTimesTreeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepairTimesTreeResponse() *RepairTimesTreeResponse {
	this := RepairTimesTreeResponse{}
	return &this
}

// NewRepairTimesTreeResponseWithDefaults instantiates a new RepairTimesTreeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepairTimesTreeResponseWithDefaults() *RepairTimesTreeResponse {
	this := RepairTimesTreeResponse{}
	return &this
}

// GetVehicleReference returns the VehicleReference field value if set, zero value otherwise.
func (o *RepairTimesTreeResponse) GetVehicleReference() RepairTimesVehicleReference {
	if o == nil || IsNil(o.VehicleReference) {
		var ret RepairTimesVehicleReference
		return ret
	}
	return *o.VehicleReference
}

// GetVehicleReferenceOk returns a tuple with the VehicleReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepairTimesTreeResponse) GetVehicleReferenceOk() (*RepairTimesVehicleReference, bool) {
	if o == nil || IsNil(o.VehicleReference) {
		return nil, false
	}
	return o.VehicleReference, true
}

// HasVehicleReference returns a boolean if a field has been set.
func (o *RepairTimesTreeResponse) HasVehicleReference() bool {
	if o != nil && !IsNil(o.VehicleReference) {
		return true
	}

	return false
}

// SetVehicleReference gets a reference to the given RepairTimesVehicleReference and assigns it to the VehicleReference field.
func (o *RepairTimesTreeResponse) SetVehicleReference(v RepairTimesVehicleReference) {
	o.VehicleReference = &v
}

// GetRepairTimesTree returns the RepairTimesTree field value if set, zero value otherwise.
func (o *RepairTimesTreeResponse) GetRepairTimesTree() RepairTimesTree {
	if o == nil || IsNil(o.RepairTimesTree) {
		var ret RepairTimesTree
		return ret
	}
	return *o.RepairTimesTree
}

// GetRepairTimesTreeOk returns a tuple with the RepairTimesTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepairTimesTreeResponse) GetRepairTimesTreeOk() (*RepairTimesTree, bool) {
	if o == nil || IsNil(o.RepairTimesTree) {
		return nil, false
	}
	return o.RepairTimesTree, true
}

// HasRepairTimesTree returns a boolean if a field has been set.
func (o *RepairTimesTreeResponse) HasRepairTimesTree() bool {
	if o != nil && !IsNil(o.RepairTimesTree) {
		return true
	}

	return false
}

// SetRepairTimesTree gets a reference to the given RepairTimesTree and assigns it to the RepairTimesTree field.
func (o *RepairTimesTreeResponse) SetRepairTimesTree(v RepairTimesTree) {
	o.RepairTimesTree = &v
}

func (o RepairTimesTreeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepairTimesTreeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VehicleReference) {
		toSerialize["vehicleReference"] = o.VehicleReference
	}
	if !IsNil(o.RepairTimesTree) {
		toSerialize["repairTimesTree"] = o.RepairTimesTree
	}
	return toSerialize, nil
}

type NullableRepairTimesTreeResponse struct {
	value *RepairTimesTreeResponse
	isSet bool
}

func (v NullableRepairTimesTreeResponse) Get() *RepairTimesTreeResponse {
	return v.value
}

func (v *NullableRepairTimesTreeResponse) Set(val *RepairTimesTreeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRepairTimesTreeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRepairTimesTreeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepairTimesTreeResponse(val *RepairTimesTreeResponse) *NullableRepairTimesTreeResponse {
	return &NullableRepairTimesTreeResponse{value: val, isSet: true}
}

func (v NullableRepairTimesTreeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepairTimesTreeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


