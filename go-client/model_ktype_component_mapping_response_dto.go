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

// checks if the KtypeComponentMappingResponseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KtypeComponentMappingResponseDTO{}

// KtypeComponentMappingResponseDTO struct for KtypeComponentMappingResponseDTO
type KtypeComponentMappingResponseDTO struct {
	// System Identification
	SystemIdent *string `json:"systemIdent,omitempty"`
	// Component Name Id
	ComponentNameId *string `json:"componentNameId,omitempty"`
	// Component Name (DE)
	ComponentNameDE *string `json:"componentNameDE,omitempty"`
	// Component Name (EN)
	ComponentNameEN *string `json:"componentNameEN,omitempty"`
	// Short Name of the Component
	ComponentShortName *string `json:"componentShortName,omitempty"`
	// Short Name Description of the Component (DE)
	ComponentShortNameDescriptionDE *string `json:"componentShortNameDescriptionDE,omitempty"`
	// Short Name Description of the Component (EN)
	ComponentShortNameDescriptionEN *string `json:"componentShortNameDescriptionEN,omitempty"`
	// KType Number
	KtypeNo *string `json:"ktypeNo,omitempty"`
}

// NewKtypeComponentMappingResponseDTO instantiates a new KtypeComponentMappingResponseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKtypeComponentMappingResponseDTO() *KtypeComponentMappingResponseDTO {
	this := KtypeComponentMappingResponseDTO{}
	return &this
}

// NewKtypeComponentMappingResponseDTOWithDefaults instantiates a new KtypeComponentMappingResponseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKtypeComponentMappingResponseDTOWithDefaults() *KtypeComponentMappingResponseDTO {
	this := KtypeComponentMappingResponseDTO{}
	return &this
}

// GetSystemIdent returns the SystemIdent field value if set, zero value otherwise.
func (o *KtypeComponentMappingResponseDTO) GetSystemIdent() string {
	if o == nil || IsNil(o.SystemIdent) {
		var ret string
		return ret
	}
	return *o.SystemIdent
}

// GetSystemIdentOk returns a tuple with the SystemIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KtypeComponentMappingResponseDTO) GetSystemIdentOk() (*string, bool) {
	if o == nil || IsNil(o.SystemIdent) {
		return nil, false
	}
	return o.SystemIdent, true
}

// HasSystemIdent returns a boolean if a field has been set.
func (o *KtypeComponentMappingResponseDTO) HasSystemIdent() bool {
	if o != nil && !IsNil(o.SystemIdent) {
		return true
	}

	return false
}

// SetSystemIdent gets a reference to the given string and assigns it to the SystemIdent field.
func (o *KtypeComponentMappingResponseDTO) SetSystemIdent(v string) {
	o.SystemIdent = &v
}

// GetComponentNameId returns the ComponentNameId field value if set, zero value otherwise.
func (o *KtypeComponentMappingResponseDTO) GetComponentNameId() string {
	if o == nil || IsNil(o.ComponentNameId) {
		var ret string
		return ret
	}
	return *o.ComponentNameId
}

// GetComponentNameIdOk returns a tuple with the ComponentNameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KtypeComponentMappingResponseDTO) GetComponentNameIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentNameId) {
		return nil, false
	}
	return o.ComponentNameId, true
}

// HasComponentNameId returns a boolean if a field has been set.
func (o *KtypeComponentMappingResponseDTO) HasComponentNameId() bool {
	if o != nil && !IsNil(o.ComponentNameId) {
		return true
	}

	return false
}

// SetComponentNameId gets a reference to the given string and assigns it to the ComponentNameId field.
func (o *KtypeComponentMappingResponseDTO) SetComponentNameId(v string) {
	o.ComponentNameId = &v
}

// GetComponentNameDE returns the ComponentNameDE field value if set, zero value otherwise.
func (o *KtypeComponentMappingResponseDTO) GetComponentNameDE() string {
	if o == nil || IsNil(o.ComponentNameDE) {
		var ret string
		return ret
	}
	return *o.ComponentNameDE
}

// GetComponentNameDEOk returns a tuple with the ComponentNameDE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KtypeComponentMappingResponseDTO) GetComponentNameDEOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentNameDE) {
		return nil, false
	}
	return o.ComponentNameDE, true
}

// HasComponentNameDE returns a boolean if a field has been set.
func (o *KtypeComponentMappingResponseDTO) HasComponentNameDE() bool {
	if o != nil && !IsNil(o.ComponentNameDE) {
		return true
	}

	return false
}

// SetComponentNameDE gets a reference to the given string and assigns it to the ComponentNameDE field.
func (o *KtypeComponentMappingResponseDTO) SetComponentNameDE(v string) {
	o.ComponentNameDE = &v
}

// GetComponentNameEN returns the ComponentNameEN field value if set, zero value otherwise.
func (o *KtypeComponentMappingResponseDTO) GetComponentNameEN() string {
	if o == nil || IsNil(o.ComponentNameEN) {
		var ret string
		return ret
	}
	return *o.ComponentNameEN
}

// GetComponentNameENOk returns a tuple with the ComponentNameEN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KtypeComponentMappingResponseDTO) GetComponentNameENOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentNameEN) {
		return nil, false
	}
	return o.ComponentNameEN, true
}

// HasComponentNameEN returns a boolean if a field has been set.
func (o *KtypeComponentMappingResponseDTO) HasComponentNameEN() bool {
	if o != nil && !IsNil(o.ComponentNameEN) {
		return true
	}

	return false
}

// SetComponentNameEN gets a reference to the given string and assigns it to the ComponentNameEN field.
func (o *KtypeComponentMappingResponseDTO) SetComponentNameEN(v string) {
	o.ComponentNameEN = &v
}

// GetComponentShortName returns the ComponentShortName field value if set, zero value otherwise.
func (o *KtypeComponentMappingResponseDTO) GetComponentShortName() string {
	if o == nil || IsNil(o.ComponentShortName) {
		var ret string
		return ret
	}
	return *o.ComponentShortName
}

// GetComponentShortNameOk returns a tuple with the ComponentShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KtypeComponentMappingResponseDTO) GetComponentShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentShortName) {
		return nil, false
	}
	return o.ComponentShortName, true
}

// HasComponentShortName returns a boolean if a field has been set.
func (o *KtypeComponentMappingResponseDTO) HasComponentShortName() bool {
	if o != nil && !IsNil(o.ComponentShortName) {
		return true
	}

	return false
}

// SetComponentShortName gets a reference to the given string and assigns it to the ComponentShortName field.
func (o *KtypeComponentMappingResponseDTO) SetComponentShortName(v string) {
	o.ComponentShortName = &v
}

// GetComponentShortNameDescriptionDE returns the ComponentShortNameDescriptionDE field value if set, zero value otherwise.
func (o *KtypeComponentMappingResponseDTO) GetComponentShortNameDescriptionDE() string {
	if o == nil || IsNil(o.ComponentShortNameDescriptionDE) {
		var ret string
		return ret
	}
	return *o.ComponentShortNameDescriptionDE
}

// GetComponentShortNameDescriptionDEOk returns a tuple with the ComponentShortNameDescriptionDE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KtypeComponentMappingResponseDTO) GetComponentShortNameDescriptionDEOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentShortNameDescriptionDE) {
		return nil, false
	}
	return o.ComponentShortNameDescriptionDE, true
}

// HasComponentShortNameDescriptionDE returns a boolean if a field has been set.
func (o *KtypeComponentMappingResponseDTO) HasComponentShortNameDescriptionDE() bool {
	if o != nil && !IsNil(o.ComponentShortNameDescriptionDE) {
		return true
	}

	return false
}

// SetComponentShortNameDescriptionDE gets a reference to the given string and assigns it to the ComponentShortNameDescriptionDE field.
func (o *KtypeComponentMappingResponseDTO) SetComponentShortNameDescriptionDE(v string) {
	o.ComponentShortNameDescriptionDE = &v
}

// GetComponentShortNameDescriptionEN returns the ComponentShortNameDescriptionEN field value if set, zero value otherwise.
func (o *KtypeComponentMappingResponseDTO) GetComponentShortNameDescriptionEN() string {
	if o == nil || IsNil(o.ComponentShortNameDescriptionEN) {
		var ret string
		return ret
	}
	return *o.ComponentShortNameDescriptionEN
}

// GetComponentShortNameDescriptionENOk returns a tuple with the ComponentShortNameDescriptionEN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KtypeComponentMappingResponseDTO) GetComponentShortNameDescriptionENOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentShortNameDescriptionEN) {
		return nil, false
	}
	return o.ComponentShortNameDescriptionEN, true
}

// HasComponentShortNameDescriptionEN returns a boolean if a field has been set.
func (o *KtypeComponentMappingResponseDTO) HasComponentShortNameDescriptionEN() bool {
	if o != nil && !IsNil(o.ComponentShortNameDescriptionEN) {
		return true
	}

	return false
}

// SetComponentShortNameDescriptionEN gets a reference to the given string and assigns it to the ComponentShortNameDescriptionEN field.
func (o *KtypeComponentMappingResponseDTO) SetComponentShortNameDescriptionEN(v string) {
	o.ComponentShortNameDescriptionEN = &v
}

// GetKtypeNo returns the KtypeNo field value if set, zero value otherwise.
func (o *KtypeComponentMappingResponseDTO) GetKtypeNo() string {
	if o == nil || IsNil(o.KtypeNo) {
		var ret string
		return ret
	}
	return *o.KtypeNo
}

// GetKtypeNoOk returns a tuple with the KtypeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KtypeComponentMappingResponseDTO) GetKtypeNoOk() (*string, bool) {
	if o == nil || IsNil(o.KtypeNo) {
		return nil, false
	}
	return o.KtypeNo, true
}

// HasKtypeNo returns a boolean if a field has been set.
func (o *KtypeComponentMappingResponseDTO) HasKtypeNo() bool {
	if o != nil && !IsNil(o.KtypeNo) {
		return true
	}

	return false
}

// SetKtypeNo gets a reference to the given string and assigns it to the KtypeNo field.
func (o *KtypeComponentMappingResponseDTO) SetKtypeNo(v string) {
	o.KtypeNo = &v
}

func (o KtypeComponentMappingResponseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KtypeComponentMappingResponseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SystemIdent) {
		toSerialize["systemIdent"] = o.SystemIdent
	}
	if !IsNil(o.ComponentNameId) {
		toSerialize["componentNameId"] = o.ComponentNameId
	}
	if !IsNil(o.ComponentNameDE) {
		toSerialize["componentNameDE"] = o.ComponentNameDE
	}
	if !IsNil(o.ComponentNameEN) {
		toSerialize["componentNameEN"] = o.ComponentNameEN
	}
	if !IsNil(o.ComponentShortName) {
		toSerialize["componentShortName"] = o.ComponentShortName
	}
	if !IsNil(o.ComponentShortNameDescriptionDE) {
		toSerialize["componentShortNameDescriptionDE"] = o.ComponentShortNameDescriptionDE
	}
	if !IsNil(o.ComponentShortNameDescriptionEN) {
		toSerialize["componentShortNameDescriptionEN"] = o.ComponentShortNameDescriptionEN
	}
	if !IsNil(o.KtypeNo) {
		toSerialize["ktypeNo"] = o.KtypeNo
	}
	return toSerialize, nil
}

type NullableKtypeComponentMappingResponseDTO struct {
	value *KtypeComponentMappingResponseDTO
	isSet bool
}

func (v NullableKtypeComponentMappingResponseDTO) Get() *KtypeComponentMappingResponseDTO {
	return v.value
}

func (v *NullableKtypeComponentMappingResponseDTO) Set(val *KtypeComponentMappingResponseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableKtypeComponentMappingResponseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableKtypeComponentMappingResponseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKtypeComponentMappingResponseDTO(val *KtypeComponentMappingResponseDTO) *NullableKtypeComponentMappingResponseDTO {
	return &NullableKtypeComponentMappingResponseDTO{value: val, isSet: true}
}

func (v NullableKtypeComponentMappingResponseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKtypeComponentMappingResponseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


