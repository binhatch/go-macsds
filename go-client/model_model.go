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

// checks if the Model type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model{}

// Model struct for Model
type Model struct {
	// Model Id
	ModelId *int32 `json:"modelId,omitempty"`
	// Cylinders
	Cylinders *int32 `json:"cylinders,omitempty"`
	// Model Type
	ModelType *string `json:"modelType,omitempty"`
	// Engine Size
	EngineSize *int32 `json:"engineSize,omitempty"`
	// Model Series Name
	ModelSeriesName *string `json:"modelSeriesName,omitempty"`
	// Model Series Alias
	ModelSeriesAlias *string `json:"modelSeriesAlias,omitempty"`
	// Model Name
	ModelName *string `json:"modelName,omitempty"`
	// Type of the Vehicle as integer value:<table><tbody><tr><td>0</td><td>None</td></tr><tr><td>1</td><td>Car</td></tr><tr><td>2</td><td>StationWagon</td></tr><tr><td>3</td><td>Van</td></tr><tr><td>4</td><td>Truck</td></tr><tr><td>5</td><td>MPV</td></tr><tr><td>6</td><td>Offroad</td></tr><tr><td>7</td><td>Bus</td></tr><tr><td>10</td><td>Pickup</td></tr><tr><td>11</td><td>Motorcycle</td></tr></tbody></table>
	TypeOfVehicle *int32 `json:"typeOfVehicle,omitempty"`
	// Region Code
	RegionCode *string `json:"regionCode,omitempty"`
	// From Year
	FromYear *int32 `json:"fromYear,omitempty"`
	// To Year
	ToYear *int32 `json:"toYear,omitempty"`
	// List of Model Years
	Years []ModelYear `json:"years,omitempty"`
	// Electronic OE Service Registration TId
	ElectronicOEServiceRegistrationTId *int32 `json:"electronicOEServiceRegistrationTId,omitempty"`
}

// NewModel instantiates a new Model object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel() *Model {
	this := Model{}
	return &this
}

// NewModelWithDefaults instantiates a new Model object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelWithDefaults() *Model {
	this := Model{}
	return &this
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *Model) GetModelId() int32 {
	if o == nil || IsNil(o.ModelId) {
		var ret int32
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetModelIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ModelId) {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *Model) HasModelId() bool {
	if o != nil && !IsNil(o.ModelId) {
		return true
	}

	return false
}

// SetModelId gets a reference to the given int32 and assigns it to the ModelId field.
func (o *Model) SetModelId(v int32) {
	o.ModelId = &v
}

// GetCylinders returns the Cylinders field value if set, zero value otherwise.
func (o *Model) GetCylinders() int32 {
	if o == nil || IsNil(o.Cylinders) {
		var ret int32
		return ret
	}
	return *o.Cylinders
}

// GetCylindersOk returns a tuple with the Cylinders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetCylindersOk() (*int32, bool) {
	if o == nil || IsNil(o.Cylinders) {
		return nil, false
	}
	return o.Cylinders, true
}

// HasCylinders returns a boolean if a field has been set.
func (o *Model) HasCylinders() bool {
	if o != nil && !IsNil(o.Cylinders) {
		return true
	}

	return false
}

// SetCylinders gets a reference to the given int32 and assigns it to the Cylinders field.
func (o *Model) SetCylinders(v int32) {
	o.Cylinders = &v
}

// GetModelType returns the ModelType field value if set, zero value otherwise.
func (o *Model) GetModelType() string {
	if o == nil || IsNil(o.ModelType) {
		var ret string
		return ret
	}
	return *o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetModelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ModelType) {
		return nil, false
	}
	return o.ModelType, true
}

// HasModelType returns a boolean if a field has been set.
func (o *Model) HasModelType() bool {
	if o != nil && !IsNil(o.ModelType) {
		return true
	}

	return false
}

// SetModelType gets a reference to the given string and assigns it to the ModelType field.
func (o *Model) SetModelType(v string) {
	o.ModelType = &v
}

// GetEngineSize returns the EngineSize field value if set, zero value otherwise.
func (o *Model) GetEngineSize() int32 {
	if o == nil || IsNil(o.EngineSize) {
		var ret int32
		return ret
	}
	return *o.EngineSize
}

// GetEngineSizeOk returns a tuple with the EngineSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetEngineSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.EngineSize) {
		return nil, false
	}
	return o.EngineSize, true
}

// HasEngineSize returns a boolean if a field has been set.
func (o *Model) HasEngineSize() bool {
	if o != nil && !IsNil(o.EngineSize) {
		return true
	}

	return false
}

// SetEngineSize gets a reference to the given int32 and assigns it to the EngineSize field.
func (o *Model) SetEngineSize(v int32) {
	o.EngineSize = &v
}

// GetModelSeriesName returns the ModelSeriesName field value if set, zero value otherwise.
func (o *Model) GetModelSeriesName() string {
	if o == nil || IsNil(o.ModelSeriesName) {
		var ret string
		return ret
	}
	return *o.ModelSeriesName
}

// GetModelSeriesNameOk returns a tuple with the ModelSeriesName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetModelSeriesNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModelSeriesName) {
		return nil, false
	}
	return o.ModelSeriesName, true
}

// HasModelSeriesName returns a boolean if a field has been set.
func (o *Model) HasModelSeriesName() bool {
	if o != nil && !IsNil(o.ModelSeriesName) {
		return true
	}

	return false
}

// SetModelSeriesName gets a reference to the given string and assigns it to the ModelSeriesName field.
func (o *Model) SetModelSeriesName(v string) {
	o.ModelSeriesName = &v
}

// GetModelSeriesAlias returns the ModelSeriesAlias field value if set, zero value otherwise.
func (o *Model) GetModelSeriesAlias() string {
	if o == nil || IsNil(o.ModelSeriesAlias) {
		var ret string
		return ret
	}
	return *o.ModelSeriesAlias
}

// GetModelSeriesAliasOk returns a tuple with the ModelSeriesAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetModelSeriesAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ModelSeriesAlias) {
		return nil, false
	}
	return o.ModelSeriesAlias, true
}

// HasModelSeriesAlias returns a boolean if a field has been set.
func (o *Model) HasModelSeriesAlias() bool {
	if o != nil && !IsNil(o.ModelSeriesAlias) {
		return true
	}

	return false
}

// SetModelSeriesAlias gets a reference to the given string and assigns it to the ModelSeriesAlias field.
func (o *Model) SetModelSeriesAlias(v string) {
	o.ModelSeriesAlias = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *Model) GetModelName() string {
	if o == nil || IsNil(o.ModelName) {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetModelNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModelName) {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *Model) HasModelName() bool {
	if o != nil && !IsNil(o.ModelName) {
		return true
	}

	return false
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *Model) SetModelName(v string) {
	o.ModelName = &v
}

// GetTypeOfVehicle returns the TypeOfVehicle field value if set, zero value otherwise.
func (o *Model) GetTypeOfVehicle() int32 {
	if o == nil || IsNil(o.TypeOfVehicle) {
		var ret int32
		return ret
	}
	return *o.TypeOfVehicle
}

// GetTypeOfVehicleOk returns a tuple with the TypeOfVehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetTypeOfVehicleOk() (*int32, bool) {
	if o == nil || IsNil(o.TypeOfVehicle) {
		return nil, false
	}
	return o.TypeOfVehicle, true
}

// HasTypeOfVehicle returns a boolean if a field has been set.
func (o *Model) HasTypeOfVehicle() bool {
	if o != nil && !IsNil(o.TypeOfVehicle) {
		return true
	}

	return false
}

// SetTypeOfVehicle gets a reference to the given int32 and assigns it to the TypeOfVehicle field.
func (o *Model) SetTypeOfVehicle(v int32) {
	o.TypeOfVehicle = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *Model) GetRegionCode() string {
	if o == nil || IsNil(o.RegionCode) {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetRegionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RegionCode) {
		return nil, false
	}
	return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *Model) HasRegionCode() bool {
	if o != nil && !IsNil(o.RegionCode) {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *Model) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetFromYear returns the FromYear field value if set, zero value otherwise.
func (o *Model) GetFromYear() int32 {
	if o == nil || IsNil(o.FromYear) {
		var ret int32
		return ret
	}
	return *o.FromYear
}

// GetFromYearOk returns a tuple with the FromYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetFromYearOk() (*int32, bool) {
	if o == nil || IsNil(o.FromYear) {
		return nil, false
	}
	return o.FromYear, true
}

// HasFromYear returns a boolean if a field has been set.
func (o *Model) HasFromYear() bool {
	if o != nil && !IsNil(o.FromYear) {
		return true
	}

	return false
}

// SetFromYear gets a reference to the given int32 and assigns it to the FromYear field.
func (o *Model) SetFromYear(v int32) {
	o.FromYear = &v
}

// GetToYear returns the ToYear field value if set, zero value otherwise.
func (o *Model) GetToYear() int32 {
	if o == nil || IsNil(o.ToYear) {
		var ret int32
		return ret
	}
	return *o.ToYear
}

// GetToYearOk returns a tuple with the ToYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetToYearOk() (*int32, bool) {
	if o == nil || IsNil(o.ToYear) {
		return nil, false
	}
	return o.ToYear, true
}

// HasToYear returns a boolean if a field has been set.
func (o *Model) HasToYear() bool {
	if o != nil && !IsNil(o.ToYear) {
		return true
	}

	return false
}

// SetToYear gets a reference to the given int32 and assigns it to the ToYear field.
func (o *Model) SetToYear(v int32) {
	o.ToYear = &v
}

// GetYears returns the Years field value if set, zero value otherwise.
func (o *Model) GetYears() []ModelYear {
	if o == nil || IsNil(o.Years) {
		var ret []ModelYear
		return ret
	}
	return o.Years
}

// GetYearsOk returns a tuple with the Years field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetYearsOk() ([]ModelYear, bool) {
	if o == nil || IsNil(o.Years) {
		return nil, false
	}
	return o.Years, true
}

// HasYears returns a boolean if a field has been set.
func (o *Model) HasYears() bool {
	if o != nil && !IsNil(o.Years) {
		return true
	}

	return false
}

// SetYears gets a reference to the given []ModelYear and assigns it to the Years field.
func (o *Model) SetYears(v []ModelYear) {
	o.Years = v
}

// GetElectronicOEServiceRegistrationTId returns the ElectronicOEServiceRegistrationTId field value if set, zero value otherwise.
func (o *Model) GetElectronicOEServiceRegistrationTId() int32 {
	if o == nil || IsNil(o.ElectronicOEServiceRegistrationTId) {
		var ret int32
		return ret
	}
	return *o.ElectronicOEServiceRegistrationTId
}

// GetElectronicOEServiceRegistrationTIdOk returns a tuple with the ElectronicOEServiceRegistrationTId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetElectronicOEServiceRegistrationTIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ElectronicOEServiceRegistrationTId) {
		return nil, false
	}
	return o.ElectronicOEServiceRegistrationTId, true
}

// HasElectronicOEServiceRegistrationTId returns a boolean if a field has been set.
func (o *Model) HasElectronicOEServiceRegistrationTId() bool {
	if o != nil && !IsNil(o.ElectronicOEServiceRegistrationTId) {
		return true
	}

	return false
}

// SetElectronicOEServiceRegistrationTId gets a reference to the given int32 and assigns it to the ElectronicOEServiceRegistrationTId field.
func (o *Model) SetElectronicOEServiceRegistrationTId(v int32) {
	o.ElectronicOEServiceRegistrationTId = &v
}

func (o Model) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ModelId) {
		toSerialize["modelId"] = o.ModelId
	}
	if !IsNil(o.Cylinders) {
		toSerialize["cylinders"] = o.Cylinders
	}
	if !IsNil(o.ModelType) {
		toSerialize["modelType"] = o.ModelType
	}
	if !IsNil(o.EngineSize) {
		toSerialize["engineSize"] = o.EngineSize
	}
	if !IsNil(o.ModelSeriesName) {
		toSerialize["modelSeriesName"] = o.ModelSeriesName
	}
	if !IsNil(o.ModelSeriesAlias) {
		toSerialize["modelSeriesAlias"] = o.ModelSeriesAlias
	}
	if !IsNil(o.ModelName) {
		toSerialize["modelName"] = o.ModelName
	}
	if !IsNil(o.TypeOfVehicle) {
		toSerialize["typeOfVehicle"] = o.TypeOfVehicle
	}
	if !IsNil(o.RegionCode) {
		toSerialize["regionCode"] = o.RegionCode
	}
	if !IsNil(o.FromYear) {
		toSerialize["fromYear"] = o.FromYear
	}
	if !IsNil(o.ToYear) {
		toSerialize["toYear"] = o.ToYear
	}
	if !IsNil(o.Years) {
		toSerialize["years"] = o.Years
	}
	if !IsNil(o.ElectronicOEServiceRegistrationTId) {
		toSerialize["electronicOEServiceRegistrationTId"] = o.ElectronicOEServiceRegistrationTId
	}
	return toSerialize, nil
}

type NullableModel struct {
	value *Model
	isSet bool
}

func (v NullableModel) Get() *Model {
	return v.value
}

func (v *NullableModel) Set(val *Model) {
	v.value = val
	v.isSet = true
}

func (v NullableModel) IsSet() bool {
	return v.isSet
}

func (v *NullableModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel(val *Model) *NullableModel {
	return &NullableModel{value: val, isSet: true}
}

func (v NullableModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


