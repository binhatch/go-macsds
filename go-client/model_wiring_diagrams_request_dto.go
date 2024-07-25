/*
Hella Gutmann - macsDS (Data Services) - API collection

This document contains all relevant APIs for diagnostics (incl. DTCs), repair & maintenance information (RMI) and vehicle identification.

API version: V20240702-130718
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package go-macds

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WiringDiagramsRequestDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WiringDiagramsRequestDTO{}

// WiringDiagramsRequestDTO struct for WiringDiagramsRequestDTO
type WiringDiagramsRequestDTO struct {
	// Two characters defining the language code; allowed values: de, en, nl, zh, it, hu, hr, fr, fi, es, el, tr, da, sv, sk, cs, ru, ro, pt, pl
	Language string `json:"language"`
	// HGS Manufacturer Id
	HgsManufId *string `json:"hgsManufId,omitempty"`
	// HGS Model Id
	HgsModelId *string `json:"hgsModelId,omitempty"`
	// Engine Code
	EngineCode *string `json:"engineCode,omitempty"`
	// Engine kW
	EngineKw *string `json:"engineKw,omitempty"`
	// Production Year
	ProductionYear *string `json:"productionYear,omitempty"`
	// Region Code
	RegionCode *string `json:"regionCode,omitempty"`
	// kType
	KType string `json:"kType"`
}

type _WiringDiagramsRequestDTO WiringDiagramsRequestDTO

// NewWiringDiagramsRequestDTO instantiates a new WiringDiagramsRequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWiringDiagramsRequestDTO(language string, kType string) *WiringDiagramsRequestDTO {
	this := WiringDiagramsRequestDTO{}
	this.Language = language
	this.KType = kType
	return &this
}

// NewWiringDiagramsRequestDTOWithDefaults instantiates a new WiringDiagramsRequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWiringDiagramsRequestDTOWithDefaults() *WiringDiagramsRequestDTO {
	this := WiringDiagramsRequestDTO{}
	return &this
}

// GetLanguage returns the Language field value
func (o *WiringDiagramsRequestDTO) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *WiringDiagramsRequestDTO) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *WiringDiagramsRequestDTO) SetLanguage(v string) {
	o.Language = v
}

// GetHgsManufId returns the HgsManufId field value if set, zero value otherwise.
func (o *WiringDiagramsRequestDTO) GetHgsManufId() string {
	if o == nil || IsNil(o.HgsManufId) {
		var ret string
		return ret
	}
	return *o.HgsManufId
}

// GetHgsManufIdOk returns a tuple with the HgsManufId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WiringDiagramsRequestDTO) GetHgsManufIdOk() (*string, bool) {
	if o == nil || IsNil(o.HgsManufId) {
		return nil, false
	}
	return o.HgsManufId, true
}

// HasHgsManufId returns a boolean if a field has been set.
func (o *WiringDiagramsRequestDTO) HasHgsManufId() bool {
	if o != nil && !IsNil(o.HgsManufId) {
		return true
	}

	return false
}

// SetHgsManufId gets a reference to the given string and assigns it to the HgsManufId field.
func (o *WiringDiagramsRequestDTO) SetHgsManufId(v string) {
	o.HgsManufId = &v
}

// GetHgsModelId returns the HgsModelId field value if set, zero value otherwise.
func (o *WiringDiagramsRequestDTO) GetHgsModelId() string {
	if o == nil || IsNil(o.HgsModelId) {
		var ret string
		return ret
	}
	return *o.HgsModelId
}

// GetHgsModelIdOk returns a tuple with the HgsModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WiringDiagramsRequestDTO) GetHgsModelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HgsModelId) {
		return nil, false
	}
	return o.HgsModelId, true
}

// HasHgsModelId returns a boolean if a field has been set.
func (o *WiringDiagramsRequestDTO) HasHgsModelId() bool {
	if o != nil && !IsNil(o.HgsModelId) {
		return true
	}

	return false
}

// SetHgsModelId gets a reference to the given string and assigns it to the HgsModelId field.
func (o *WiringDiagramsRequestDTO) SetHgsModelId(v string) {
	o.HgsModelId = &v
}

// GetEngineCode returns the EngineCode field value if set, zero value otherwise.
func (o *WiringDiagramsRequestDTO) GetEngineCode() string {
	if o == nil || IsNil(o.EngineCode) {
		var ret string
		return ret
	}
	return *o.EngineCode
}

// GetEngineCodeOk returns a tuple with the EngineCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WiringDiagramsRequestDTO) GetEngineCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EngineCode) {
		return nil, false
	}
	return o.EngineCode, true
}

// HasEngineCode returns a boolean if a field has been set.
func (o *WiringDiagramsRequestDTO) HasEngineCode() bool {
	if o != nil && !IsNil(o.EngineCode) {
		return true
	}

	return false
}

// SetEngineCode gets a reference to the given string and assigns it to the EngineCode field.
func (o *WiringDiagramsRequestDTO) SetEngineCode(v string) {
	o.EngineCode = &v
}

// GetEngineKw returns the EngineKw field value if set, zero value otherwise.
func (o *WiringDiagramsRequestDTO) GetEngineKw() string {
	if o == nil || IsNil(o.EngineKw) {
		var ret string
		return ret
	}
	return *o.EngineKw
}

// GetEngineKwOk returns a tuple with the EngineKw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WiringDiagramsRequestDTO) GetEngineKwOk() (*string, bool) {
	if o == nil || IsNil(o.EngineKw) {
		return nil, false
	}
	return o.EngineKw, true
}

// HasEngineKw returns a boolean if a field has been set.
func (o *WiringDiagramsRequestDTO) HasEngineKw() bool {
	if o != nil && !IsNil(o.EngineKw) {
		return true
	}

	return false
}

// SetEngineKw gets a reference to the given string and assigns it to the EngineKw field.
func (o *WiringDiagramsRequestDTO) SetEngineKw(v string) {
	o.EngineKw = &v
}

// GetProductionYear returns the ProductionYear field value if set, zero value otherwise.
func (o *WiringDiagramsRequestDTO) GetProductionYear() string {
	if o == nil || IsNil(o.ProductionYear) {
		var ret string
		return ret
	}
	return *o.ProductionYear
}

// GetProductionYearOk returns a tuple with the ProductionYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WiringDiagramsRequestDTO) GetProductionYearOk() (*string, bool) {
	if o == nil || IsNil(o.ProductionYear) {
		return nil, false
	}
	return o.ProductionYear, true
}

// HasProductionYear returns a boolean if a field has been set.
func (o *WiringDiagramsRequestDTO) HasProductionYear() bool {
	if o != nil && !IsNil(o.ProductionYear) {
		return true
	}

	return false
}

// SetProductionYear gets a reference to the given string and assigns it to the ProductionYear field.
func (o *WiringDiagramsRequestDTO) SetProductionYear(v string) {
	o.ProductionYear = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *WiringDiagramsRequestDTO) GetRegionCode() string {
	if o == nil || IsNil(o.RegionCode) {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WiringDiagramsRequestDTO) GetRegionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RegionCode) {
		return nil, false
	}
	return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *WiringDiagramsRequestDTO) HasRegionCode() bool {
	if o != nil && !IsNil(o.RegionCode) {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *WiringDiagramsRequestDTO) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetKType returns the KType field value
func (o *WiringDiagramsRequestDTO) GetKType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KType
}

// GetKTypeOk returns a tuple with the KType field value
// and a boolean to check if the value has been set.
func (o *WiringDiagramsRequestDTO) GetKTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KType, true
}

// SetKType sets field value
func (o *WiringDiagramsRequestDTO) SetKType(v string) {
	o.KType = v
}

func (o WiringDiagramsRequestDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WiringDiagramsRequestDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["language"] = o.Language
	if !IsNil(o.HgsManufId) {
		toSerialize["hgsManufId"] = o.HgsManufId
	}
	if !IsNil(o.HgsModelId) {
		toSerialize["hgsModelId"] = o.HgsModelId
	}
	if !IsNil(o.EngineCode) {
		toSerialize["engineCode"] = o.EngineCode
	}
	if !IsNil(o.EngineKw) {
		toSerialize["engineKw"] = o.EngineKw
	}
	if !IsNil(o.ProductionYear) {
		toSerialize["productionYear"] = o.ProductionYear
	}
	if !IsNil(o.RegionCode) {
		toSerialize["regionCode"] = o.RegionCode
	}
	toSerialize["kType"] = o.KType
	return toSerialize, nil
}

func (o *WiringDiagramsRequestDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"language",
		"kType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWiringDiagramsRequestDTO := _WiringDiagramsRequestDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWiringDiagramsRequestDTO)

	if err != nil {
		return err
	}

	*o = WiringDiagramsRequestDTO(varWiringDiagramsRequestDTO)

	return err
}

type NullableWiringDiagramsRequestDTO struct {
	value *WiringDiagramsRequestDTO
	isSet bool
}

func (v NullableWiringDiagramsRequestDTO) Get() *WiringDiagramsRequestDTO {
	return v.value
}

func (v *NullableWiringDiagramsRequestDTO) Set(val *WiringDiagramsRequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWiringDiagramsRequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWiringDiagramsRequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWiringDiagramsRequestDTO(val *WiringDiagramsRequestDTO) *NullableWiringDiagramsRequestDTO {
	return &NullableWiringDiagramsRequestDTO{value: val, isSet: true}
}

func (v NullableWiringDiagramsRequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWiringDiagramsRequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


