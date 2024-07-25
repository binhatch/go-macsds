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

// checks if the TypeV2Dto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypeV2Dto{}

// TypeV2Dto struct for TypeV2Dto
type TypeV2Dto struct {
	// Id of the model
	ModelId *int32 `json:"modelId,omitempty"`
	// Name of the make
	Name *string `json:"name,omitempty"`
	// Start of production
	YearFrom *int32 `json:"yearFrom,omitempty"`
	// End of production
	YearTo *int32 `json:"yearTo,omitempty"`
	// Id of the make
	MakeId *int32 `json:"makeId,omitempty"`
	// Model Series Id
	ModelSeriesId *int32 `json:"modelSeriesId,omitempty"`
	// Power in kw
	PowerKw *int32 `json:"powerKw,omitempty"`
	// Power in hp
	PowerHp *int32 `json:"powerHp,omitempty"`
	// Capacity
	Capacity *int32 `json:"capacity,omitempty"`
	// Fuel type
	FuelType *string `json:"fuelType,omitempty"`
	// Engine Code
	EngineCode *string `json:"engineCode,omitempty"`
	// Model Type
	ModelType *string `json:"modelType,omitempty"`
	// Type of the Vehicle as integer value:<table><tbody><tr><td>0</td><td>None</td></tr><tr><td>1</td><td>Car</td></tr><tr><td>2</td><td>StationWagon</td></tr><tr><td>3</td><td>Van</td></tr><tr><td>4</td><td>Truck</td></tr><tr><td>5</td><td>MPV</td></tr><tr><td>6</td><td>Offroad</td></tr><tr><td>7</td><td>Bus</td></tr><tr><td>10</td><td>Pickup</td></tr><tr><td>11</td><td>Motorcycle</td></tr></tbody></table>
	TypeOfVehicle *int32 `json:"typeOfVehicle,omitempty"`
	// KTypes
	Ktypes []int32 `json:"ktypes,omitempty"`
}

// NewTypeV2Dto instantiates a new TypeV2Dto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypeV2Dto() *TypeV2Dto {
	this := TypeV2Dto{}
	return &this
}

// NewTypeV2DtoWithDefaults instantiates a new TypeV2Dto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypeV2DtoWithDefaults() *TypeV2Dto {
	this := TypeV2Dto{}
	return &this
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *TypeV2Dto) GetModelId() int32 {
	if o == nil || IsNil(o.ModelId) {
		var ret int32
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetModelIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ModelId) {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *TypeV2Dto) HasModelId() bool {
	if o != nil && !IsNil(o.ModelId) {
		return true
	}

	return false
}

// SetModelId gets a reference to the given int32 and assigns it to the ModelId field.
func (o *TypeV2Dto) SetModelId(v int32) {
	o.ModelId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TypeV2Dto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TypeV2Dto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TypeV2Dto) SetName(v string) {
	o.Name = &v
}

// GetYearFrom returns the YearFrom field value if set, zero value otherwise.
func (o *TypeV2Dto) GetYearFrom() int32 {
	if o == nil || IsNil(o.YearFrom) {
		var ret int32
		return ret
	}
	return *o.YearFrom
}

// GetYearFromOk returns a tuple with the YearFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetYearFromOk() (*int32, bool) {
	if o == nil || IsNil(o.YearFrom) {
		return nil, false
	}
	return o.YearFrom, true
}

// HasYearFrom returns a boolean if a field has been set.
func (o *TypeV2Dto) HasYearFrom() bool {
	if o != nil && !IsNil(o.YearFrom) {
		return true
	}

	return false
}

// SetYearFrom gets a reference to the given int32 and assigns it to the YearFrom field.
func (o *TypeV2Dto) SetYearFrom(v int32) {
	o.YearFrom = &v
}

// GetYearTo returns the YearTo field value if set, zero value otherwise.
func (o *TypeV2Dto) GetYearTo() int32 {
	if o == nil || IsNil(o.YearTo) {
		var ret int32
		return ret
	}
	return *o.YearTo
}

// GetYearToOk returns a tuple with the YearTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetYearToOk() (*int32, bool) {
	if o == nil || IsNil(o.YearTo) {
		return nil, false
	}
	return o.YearTo, true
}

// HasYearTo returns a boolean if a field has been set.
func (o *TypeV2Dto) HasYearTo() bool {
	if o != nil && !IsNil(o.YearTo) {
		return true
	}

	return false
}

// SetYearTo gets a reference to the given int32 and assigns it to the YearTo field.
func (o *TypeV2Dto) SetYearTo(v int32) {
	o.YearTo = &v
}

// GetMakeId returns the MakeId field value if set, zero value otherwise.
func (o *TypeV2Dto) GetMakeId() int32 {
	if o == nil || IsNil(o.MakeId) {
		var ret int32
		return ret
	}
	return *o.MakeId
}

// GetMakeIdOk returns a tuple with the MakeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetMakeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MakeId) {
		return nil, false
	}
	return o.MakeId, true
}

// HasMakeId returns a boolean if a field has been set.
func (o *TypeV2Dto) HasMakeId() bool {
	if o != nil && !IsNil(o.MakeId) {
		return true
	}

	return false
}

// SetMakeId gets a reference to the given int32 and assigns it to the MakeId field.
func (o *TypeV2Dto) SetMakeId(v int32) {
	o.MakeId = &v
}

// GetModelSeriesId returns the ModelSeriesId field value if set, zero value otherwise.
func (o *TypeV2Dto) GetModelSeriesId() int32 {
	if o == nil || IsNil(o.ModelSeriesId) {
		var ret int32
		return ret
	}
	return *o.ModelSeriesId
}

// GetModelSeriesIdOk returns a tuple with the ModelSeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetModelSeriesIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ModelSeriesId) {
		return nil, false
	}
	return o.ModelSeriesId, true
}

// HasModelSeriesId returns a boolean if a field has been set.
func (o *TypeV2Dto) HasModelSeriesId() bool {
	if o != nil && !IsNil(o.ModelSeriesId) {
		return true
	}

	return false
}

// SetModelSeriesId gets a reference to the given int32 and assigns it to the ModelSeriesId field.
func (o *TypeV2Dto) SetModelSeriesId(v int32) {
	o.ModelSeriesId = &v
}

// GetPowerKw returns the PowerKw field value if set, zero value otherwise.
func (o *TypeV2Dto) GetPowerKw() int32 {
	if o == nil || IsNil(o.PowerKw) {
		var ret int32
		return ret
	}
	return *o.PowerKw
}

// GetPowerKwOk returns a tuple with the PowerKw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetPowerKwOk() (*int32, bool) {
	if o == nil || IsNil(o.PowerKw) {
		return nil, false
	}
	return o.PowerKw, true
}

// HasPowerKw returns a boolean if a field has been set.
func (o *TypeV2Dto) HasPowerKw() bool {
	if o != nil && !IsNil(o.PowerKw) {
		return true
	}

	return false
}

// SetPowerKw gets a reference to the given int32 and assigns it to the PowerKw field.
func (o *TypeV2Dto) SetPowerKw(v int32) {
	o.PowerKw = &v
}

// GetPowerHp returns the PowerHp field value if set, zero value otherwise.
func (o *TypeV2Dto) GetPowerHp() int32 {
	if o == nil || IsNil(o.PowerHp) {
		var ret int32
		return ret
	}
	return *o.PowerHp
}

// GetPowerHpOk returns a tuple with the PowerHp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetPowerHpOk() (*int32, bool) {
	if o == nil || IsNil(o.PowerHp) {
		return nil, false
	}
	return o.PowerHp, true
}

// HasPowerHp returns a boolean if a field has been set.
func (o *TypeV2Dto) HasPowerHp() bool {
	if o != nil && !IsNil(o.PowerHp) {
		return true
	}

	return false
}

// SetPowerHp gets a reference to the given int32 and assigns it to the PowerHp field.
func (o *TypeV2Dto) SetPowerHp(v int32) {
	o.PowerHp = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *TypeV2Dto) GetCapacity() int32 {
	if o == nil || IsNil(o.Capacity) {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetCapacityOk() (*int32, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *TypeV2Dto) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *TypeV2Dto) SetCapacity(v int32) {
	o.Capacity = &v
}

// GetFuelType returns the FuelType field value if set, zero value otherwise.
func (o *TypeV2Dto) GetFuelType() string {
	if o == nil || IsNil(o.FuelType) {
		var ret string
		return ret
	}
	return *o.FuelType
}

// GetFuelTypeOk returns a tuple with the FuelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetFuelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FuelType) {
		return nil, false
	}
	return o.FuelType, true
}

// HasFuelType returns a boolean if a field has been set.
func (o *TypeV2Dto) HasFuelType() bool {
	if o != nil && !IsNil(o.FuelType) {
		return true
	}

	return false
}

// SetFuelType gets a reference to the given string and assigns it to the FuelType field.
func (o *TypeV2Dto) SetFuelType(v string) {
	o.FuelType = &v
}

// GetEngineCode returns the EngineCode field value if set, zero value otherwise.
func (o *TypeV2Dto) GetEngineCode() string {
	if o == nil || IsNil(o.EngineCode) {
		var ret string
		return ret
	}
	return *o.EngineCode
}

// GetEngineCodeOk returns a tuple with the EngineCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetEngineCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EngineCode) {
		return nil, false
	}
	return o.EngineCode, true
}

// HasEngineCode returns a boolean if a field has been set.
func (o *TypeV2Dto) HasEngineCode() bool {
	if o != nil && !IsNil(o.EngineCode) {
		return true
	}

	return false
}

// SetEngineCode gets a reference to the given string and assigns it to the EngineCode field.
func (o *TypeV2Dto) SetEngineCode(v string) {
	o.EngineCode = &v
}

// GetModelType returns the ModelType field value if set, zero value otherwise.
func (o *TypeV2Dto) GetModelType() string {
	if o == nil || IsNil(o.ModelType) {
		var ret string
		return ret
	}
	return *o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetModelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ModelType) {
		return nil, false
	}
	return o.ModelType, true
}

// HasModelType returns a boolean if a field has been set.
func (o *TypeV2Dto) HasModelType() bool {
	if o != nil && !IsNil(o.ModelType) {
		return true
	}

	return false
}

// SetModelType gets a reference to the given string and assigns it to the ModelType field.
func (o *TypeV2Dto) SetModelType(v string) {
	o.ModelType = &v
}

// GetTypeOfVehicle returns the TypeOfVehicle field value if set, zero value otherwise.
func (o *TypeV2Dto) GetTypeOfVehicle() int32 {
	if o == nil || IsNil(o.TypeOfVehicle) {
		var ret int32
		return ret
	}
	return *o.TypeOfVehicle
}

// GetTypeOfVehicleOk returns a tuple with the TypeOfVehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetTypeOfVehicleOk() (*int32, bool) {
	if o == nil || IsNil(o.TypeOfVehicle) {
		return nil, false
	}
	return o.TypeOfVehicle, true
}

// HasTypeOfVehicle returns a boolean if a field has been set.
func (o *TypeV2Dto) HasTypeOfVehicle() bool {
	if o != nil && !IsNil(o.TypeOfVehicle) {
		return true
	}

	return false
}

// SetTypeOfVehicle gets a reference to the given int32 and assigns it to the TypeOfVehicle field.
func (o *TypeV2Dto) SetTypeOfVehicle(v int32) {
	o.TypeOfVehicle = &v
}

// GetKtypes returns the Ktypes field value if set, zero value otherwise.
func (o *TypeV2Dto) GetKtypes() []int32 {
	if o == nil || IsNil(o.Ktypes) {
		var ret []int32
		return ret
	}
	return o.Ktypes
}

// GetKtypesOk returns a tuple with the Ktypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeV2Dto) GetKtypesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ktypes) {
		return nil, false
	}
	return o.Ktypes, true
}

// HasKtypes returns a boolean if a field has been set.
func (o *TypeV2Dto) HasKtypes() bool {
	if o != nil && !IsNil(o.Ktypes) {
		return true
	}

	return false
}

// SetKtypes gets a reference to the given []int32 and assigns it to the Ktypes field.
func (o *TypeV2Dto) SetKtypes(v []int32) {
	o.Ktypes = v
}

func (o TypeV2Dto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypeV2Dto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ModelId) {
		toSerialize["modelId"] = o.ModelId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.YearFrom) {
		toSerialize["yearFrom"] = o.YearFrom
	}
	if !IsNil(o.YearTo) {
		toSerialize["yearTo"] = o.YearTo
	}
	if !IsNil(o.MakeId) {
		toSerialize["makeId"] = o.MakeId
	}
	if !IsNil(o.ModelSeriesId) {
		toSerialize["modelSeriesId"] = o.ModelSeriesId
	}
	if !IsNil(o.PowerKw) {
		toSerialize["powerKw"] = o.PowerKw
	}
	if !IsNil(o.PowerHp) {
		toSerialize["powerHp"] = o.PowerHp
	}
	if !IsNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	if !IsNil(o.FuelType) {
		toSerialize["fuelType"] = o.FuelType
	}
	if !IsNil(o.EngineCode) {
		toSerialize["engineCode"] = o.EngineCode
	}
	if !IsNil(o.ModelType) {
		toSerialize["modelType"] = o.ModelType
	}
	if !IsNil(o.TypeOfVehicle) {
		toSerialize["typeOfVehicle"] = o.TypeOfVehicle
	}
	if !IsNil(o.Ktypes) {
		toSerialize["ktypes"] = o.Ktypes
	}
	return toSerialize, nil
}

type NullableTypeV2Dto struct {
	value *TypeV2Dto
	isSet bool
}

func (v NullableTypeV2Dto) Get() *TypeV2Dto {
	return v.value
}

func (v *NullableTypeV2Dto) Set(val *TypeV2Dto) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeV2Dto) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeV2Dto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeV2Dto(val *TypeV2Dto) *NullableTypeV2Dto {
	return &NullableTypeV2Dto{value: val, isSet: true}
}

func (v NullableTypeV2Dto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeV2Dto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


