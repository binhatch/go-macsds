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

// checks if the IdentRequestDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentRequestDTO{}

// IdentRequestDTO struct for IdentRequestDTO
type IdentRequestDTO struct {
	// vehicle identification number
	Vin string `json:"vin"`
	// The engine code
	EngineCode *string `json:"engineCode,omitempty"`
	// True if only one (trusted) or none results should be returned; otherwise the result list longer along with a confidence score of the respective result
	SingleTrustedResult *bool `json:"singleTrustedResult,omitempty"`
	// True if additional vehicle information other than just KType or KBA are needed
	EnrichOutputVehicle *bool `json:"enrichOutputVehicle,omitempty"`
	// Choose between outputting Ktype numbers (=1) or KBA/HSN-TSN number (=2)
	OutputVehicle *string `json:"outputVehicle,omitempty"`
}

type _IdentRequestDTO IdentRequestDTO

// NewIdentRequestDTO instantiates a new IdentRequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentRequestDTO(vin string) *IdentRequestDTO {
	this := IdentRequestDTO{}
	this.Vin = vin
	return &this
}

// NewIdentRequestDTOWithDefaults instantiates a new IdentRequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentRequestDTOWithDefaults() *IdentRequestDTO {
	this := IdentRequestDTO{}
	return &this
}

// GetVin returns the Vin field value
func (o *IdentRequestDTO) GetVin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *IdentRequestDTO) GetVinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *IdentRequestDTO) SetVin(v string) {
	o.Vin = v
}

// GetEngineCode returns the EngineCode field value if set, zero value otherwise.
func (o *IdentRequestDTO) GetEngineCode() string {
	if o == nil || IsNil(o.EngineCode) {
		var ret string
		return ret
	}
	return *o.EngineCode
}

// GetEngineCodeOk returns a tuple with the EngineCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentRequestDTO) GetEngineCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EngineCode) {
		return nil, false
	}
	return o.EngineCode, true
}

// HasEngineCode returns a boolean if a field has been set.
func (o *IdentRequestDTO) HasEngineCode() bool {
	if o != nil && !IsNil(o.EngineCode) {
		return true
	}

	return false
}

// SetEngineCode gets a reference to the given string and assigns it to the EngineCode field.
func (o *IdentRequestDTO) SetEngineCode(v string) {
	o.EngineCode = &v
}

// GetSingleTrustedResult returns the SingleTrustedResult field value if set, zero value otherwise.
func (o *IdentRequestDTO) GetSingleTrustedResult() bool {
	if o == nil || IsNil(o.SingleTrustedResult) {
		var ret bool
		return ret
	}
	return *o.SingleTrustedResult
}

// GetSingleTrustedResultOk returns a tuple with the SingleTrustedResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentRequestDTO) GetSingleTrustedResultOk() (*bool, bool) {
	if o == nil || IsNil(o.SingleTrustedResult) {
		return nil, false
	}
	return o.SingleTrustedResult, true
}

// HasSingleTrustedResult returns a boolean if a field has been set.
func (o *IdentRequestDTO) HasSingleTrustedResult() bool {
	if o != nil && !IsNil(o.SingleTrustedResult) {
		return true
	}

	return false
}

// SetSingleTrustedResult gets a reference to the given bool and assigns it to the SingleTrustedResult field.
func (o *IdentRequestDTO) SetSingleTrustedResult(v bool) {
	o.SingleTrustedResult = &v
}

// GetEnrichOutputVehicle returns the EnrichOutputVehicle field value if set, zero value otherwise.
func (o *IdentRequestDTO) GetEnrichOutputVehicle() bool {
	if o == nil || IsNil(o.EnrichOutputVehicle) {
		var ret bool
		return ret
	}
	return *o.EnrichOutputVehicle
}

// GetEnrichOutputVehicleOk returns a tuple with the EnrichOutputVehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentRequestDTO) GetEnrichOutputVehicleOk() (*bool, bool) {
	if o == nil || IsNil(o.EnrichOutputVehicle) {
		return nil, false
	}
	return o.EnrichOutputVehicle, true
}

// HasEnrichOutputVehicle returns a boolean if a field has been set.
func (o *IdentRequestDTO) HasEnrichOutputVehicle() bool {
	if o != nil && !IsNil(o.EnrichOutputVehicle) {
		return true
	}

	return false
}

// SetEnrichOutputVehicle gets a reference to the given bool and assigns it to the EnrichOutputVehicle field.
func (o *IdentRequestDTO) SetEnrichOutputVehicle(v bool) {
	o.EnrichOutputVehicle = &v
}

// GetOutputVehicle returns the OutputVehicle field value if set, zero value otherwise.
func (o *IdentRequestDTO) GetOutputVehicle() string {
	if o == nil || IsNil(o.OutputVehicle) {
		var ret string
		return ret
	}
	return *o.OutputVehicle
}

// GetOutputVehicleOk returns a tuple with the OutputVehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentRequestDTO) GetOutputVehicleOk() (*string, bool) {
	if o == nil || IsNil(o.OutputVehicle) {
		return nil, false
	}
	return o.OutputVehicle, true
}

// HasOutputVehicle returns a boolean if a field has been set.
func (o *IdentRequestDTO) HasOutputVehicle() bool {
	if o != nil && !IsNil(o.OutputVehicle) {
		return true
	}

	return false
}

// SetOutputVehicle gets a reference to the given string and assigns it to the OutputVehicle field.
func (o *IdentRequestDTO) SetOutputVehicle(v string) {
	o.OutputVehicle = &v
}

func (o IdentRequestDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentRequestDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vin"] = o.Vin
	if !IsNil(o.EngineCode) {
		toSerialize["engineCode"] = o.EngineCode
	}
	if !IsNil(o.SingleTrustedResult) {
		toSerialize["singleTrustedResult"] = o.SingleTrustedResult
	}
	if !IsNil(o.EnrichOutputVehicle) {
		toSerialize["enrichOutputVehicle"] = o.EnrichOutputVehicle
	}
	if !IsNil(o.OutputVehicle) {
		toSerialize["outputVehicle"] = o.OutputVehicle
	}
	return toSerialize, nil
}

func (o *IdentRequestDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vin",
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

	varIdentRequestDTO := _IdentRequestDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIdentRequestDTO)

	if err != nil {
		return err
	}

	*o = IdentRequestDTO(varIdentRequestDTO)

	return err
}

type NullableIdentRequestDTO struct {
	value *IdentRequestDTO
	isSet bool
}

func (v NullableIdentRequestDTO) Get() *IdentRequestDTO {
	return v.value
}

func (v *NullableIdentRequestDTO) Set(val *IdentRequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentRequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentRequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentRequestDTO(val *IdentRequestDTO) *NullableIdentRequestDTO {
	return &NullableIdentRequestDTO{value: val, isSet: true}
}

func (v NullableIdentRequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentRequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


