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

// checks if the WiringDiagramSystemsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WiringDiagramSystemsResponse{}

// WiringDiagramSystemsResponse struct for WiringDiagramSystemsResponse
type WiringDiagramSystemsResponse struct {
	VehicleReference *VehicleReference `json:"vehicleReference,omitempty"`
	SystemGroupTree *SystemGroupTree `json:"systemGroupTree,omitempty"`
}

// NewWiringDiagramSystemsResponse instantiates a new WiringDiagramSystemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWiringDiagramSystemsResponse() *WiringDiagramSystemsResponse {
	this := WiringDiagramSystemsResponse{}
	return &this
}

// NewWiringDiagramSystemsResponseWithDefaults instantiates a new WiringDiagramSystemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWiringDiagramSystemsResponseWithDefaults() *WiringDiagramSystemsResponse {
	this := WiringDiagramSystemsResponse{}
	return &this
}

// GetVehicleReference returns the VehicleReference field value if set, zero value otherwise.
func (o *WiringDiagramSystemsResponse) GetVehicleReference() VehicleReference {
	if o == nil || IsNil(o.VehicleReference) {
		var ret VehicleReference
		return ret
	}
	return *o.VehicleReference
}

// GetVehicleReferenceOk returns a tuple with the VehicleReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WiringDiagramSystemsResponse) GetVehicleReferenceOk() (*VehicleReference, bool) {
	if o == nil || IsNil(o.VehicleReference) {
		return nil, false
	}
	return o.VehicleReference, true
}

// HasVehicleReference returns a boolean if a field has been set.
func (o *WiringDiagramSystemsResponse) HasVehicleReference() bool {
	if o != nil && !IsNil(o.VehicleReference) {
		return true
	}

	return false
}

// SetVehicleReference gets a reference to the given VehicleReference and assigns it to the VehicleReference field.
func (o *WiringDiagramSystemsResponse) SetVehicleReference(v VehicleReference) {
	o.VehicleReference = &v
}

// GetSystemGroupTree returns the SystemGroupTree field value if set, zero value otherwise.
func (o *WiringDiagramSystemsResponse) GetSystemGroupTree() SystemGroupTree {
	if o == nil || IsNil(o.SystemGroupTree) {
		var ret SystemGroupTree
		return ret
	}
	return *o.SystemGroupTree
}

// GetSystemGroupTreeOk returns a tuple with the SystemGroupTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WiringDiagramSystemsResponse) GetSystemGroupTreeOk() (*SystemGroupTree, bool) {
	if o == nil || IsNil(o.SystemGroupTree) {
		return nil, false
	}
	return o.SystemGroupTree, true
}

// HasSystemGroupTree returns a boolean if a field has been set.
func (o *WiringDiagramSystemsResponse) HasSystemGroupTree() bool {
	if o != nil && !IsNil(o.SystemGroupTree) {
		return true
	}

	return false
}

// SetSystemGroupTree gets a reference to the given SystemGroupTree and assigns it to the SystemGroupTree field.
func (o *WiringDiagramSystemsResponse) SetSystemGroupTree(v SystemGroupTree) {
	o.SystemGroupTree = &v
}

func (o WiringDiagramSystemsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WiringDiagramSystemsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VehicleReference) {
		toSerialize["vehicleReference"] = o.VehicleReference
	}
	if !IsNil(o.SystemGroupTree) {
		toSerialize["systemGroupTree"] = o.SystemGroupTree
	}
	return toSerialize, nil
}

type NullableWiringDiagramSystemsResponse struct {
	value *WiringDiagramSystemsResponse
	isSet bool
}

func (v NullableWiringDiagramSystemsResponse) Get() *WiringDiagramSystemsResponse {
	return v.value
}

func (v *NullableWiringDiagramSystemsResponse) Set(val *WiringDiagramSystemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWiringDiagramSystemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWiringDiagramSystemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWiringDiagramSystemsResponse(val *WiringDiagramSystemsResponse) *NullableWiringDiagramSystemsResponse {
	return &NullableWiringDiagramSystemsResponse{value: val, isSet: true}
}

func (v NullableWiringDiagramSystemsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWiringDiagramSystemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


