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

// checks if the DtcRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtcRequestDto{}

// DtcRequestDto struct for DtcRequestDto
type DtcRequestDto struct {
	// If known, specify if you read out OBD, OE specific systems, or the OBD norm for commercial trucks. If not specified, the API will attempt to apply OBD first, then OE if no results found.
	DtcReadType *string `json:"dtcReadType,omitempty"`
	// VIN - vehicle identifier number
	Vin string `json:"vin"`
	// Two characters defining the language code; allowed values: de, en, nl, zh, it, hu, hr, fr, fi, es, el, tr, da, sv, sk, cs, ru, ro, pt, pl
	Language string `json:"language"`
	// Requested page in case of more results
	Page *int64 `json:"page,omitempty"`
	// Parameter to indicate if user wants recommendations.
	WithRecommendation *bool `json:"withRecommendation,omitempty"`
	// Parameter to indicate if user wants classifications.
	WithClassification *bool `json:"withClassification,omitempty"`
	// If set to false, DTCs that translate to more than 1 meaning due to a missing specification of the parameter \"ecuSystem\" will return an error that the DTC is ambigious. If set to true, the API will return all results that match the specified DTC, along with the indication of the related ECU System.
	AllowMultipleResults *bool `json:"allowMultipleResults,omitempty"`
	// If set to true, our AI will attempt to predict a component that is affected based on the DTCs provided.
	ComponentPredict *bool `json:"componentPredict,omitempty"`
	// Set to true if the vehicle's malfunction indicator lamp (\"MIL\"), also known as \"warning light\" or \"check engine light\", is on; used for calculating the criticality index
	MilActive *bool `json:"milActive,omitempty"`
	// Can be html or json. If json is not explicitly specified it defaults to html.
	RelatedRmi *string `json:"relatedRmi,omitempty"`
	// A list of DTC codes
	Dtcs []DtcsDto `json:"dtcs"`
}

type _DtcRequestDto DtcRequestDto

// NewDtcRequestDto instantiates a new DtcRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtcRequestDto(vin string, language string, dtcs []DtcsDto) *DtcRequestDto {
	this := DtcRequestDto{}
	this.Vin = vin
	this.Language = language
	this.Dtcs = dtcs
	return &this
}

// NewDtcRequestDtoWithDefaults instantiates a new DtcRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtcRequestDtoWithDefaults() *DtcRequestDto {
	this := DtcRequestDto{}
	return &this
}

// GetDtcReadType returns the DtcReadType field value if set, zero value otherwise.
func (o *DtcRequestDto) GetDtcReadType() string {
	if o == nil || IsNil(o.DtcReadType) {
		var ret string
		return ret
	}
	return *o.DtcReadType
}

// GetDtcReadTypeOk returns a tuple with the DtcReadType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcRequestDto) GetDtcReadTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DtcReadType) {
		return nil, false
	}
	return o.DtcReadType, true
}

// HasDtcReadType returns a boolean if a field has been set.
func (o *DtcRequestDto) HasDtcReadType() bool {
	if o != nil && !IsNil(o.DtcReadType) {
		return true
	}

	return false
}

// SetDtcReadType gets a reference to the given string and assigns it to the DtcReadType field.
func (o *DtcRequestDto) SetDtcReadType(v string) {
	o.DtcReadType = &v
}

// GetVin returns the Vin field value
func (o *DtcRequestDto) GetVin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *DtcRequestDto) GetVinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *DtcRequestDto) SetVin(v string) {
	o.Vin = v
}

// GetLanguage returns the Language field value
func (o *DtcRequestDto) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *DtcRequestDto) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *DtcRequestDto) SetLanguage(v string) {
	o.Language = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *DtcRequestDto) GetPage() int64 {
	if o == nil || IsNil(o.Page) {
		var ret int64
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcRequestDto) GetPageOk() (*int64, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *DtcRequestDto) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int64 and assigns it to the Page field.
func (o *DtcRequestDto) SetPage(v int64) {
	o.Page = &v
}

// GetWithRecommendation returns the WithRecommendation field value if set, zero value otherwise.
func (o *DtcRequestDto) GetWithRecommendation() bool {
	if o == nil || IsNil(o.WithRecommendation) {
		var ret bool
		return ret
	}
	return *o.WithRecommendation
}

// GetWithRecommendationOk returns a tuple with the WithRecommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcRequestDto) GetWithRecommendationOk() (*bool, bool) {
	if o == nil || IsNil(o.WithRecommendation) {
		return nil, false
	}
	return o.WithRecommendation, true
}

// HasWithRecommendation returns a boolean if a field has been set.
func (o *DtcRequestDto) HasWithRecommendation() bool {
	if o != nil && !IsNil(o.WithRecommendation) {
		return true
	}

	return false
}

// SetWithRecommendation gets a reference to the given bool and assigns it to the WithRecommendation field.
func (o *DtcRequestDto) SetWithRecommendation(v bool) {
	o.WithRecommendation = &v
}

// GetWithClassification returns the WithClassification field value if set, zero value otherwise.
func (o *DtcRequestDto) GetWithClassification() bool {
	if o == nil || IsNil(o.WithClassification) {
		var ret bool
		return ret
	}
	return *o.WithClassification
}

// GetWithClassificationOk returns a tuple with the WithClassification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcRequestDto) GetWithClassificationOk() (*bool, bool) {
	if o == nil || IsNil(o.WithClassification) {
		return nil, false
	}
	return o.WithClassification, true
}

// HasWithClassification returns a boolean if a field has been set.
func (o *DtcRequestDto) HasWithClassification() bool {
	if o != nil && !IsNil(o.WithClassification) {
		return true
	}

	return false
}

// SetWithClassification gets a reference to the given bool and assigns it to the WithClassification field.
func (o *DtcRequestDto) SetWithClassification(v bool) {
	o.WithClassification = &v
}

// GetAllowMultipleResults returns the AllowMultipleResults field value if set, zero value otherwise.
func (o *DtcRequestDto) GetAllowMultipleResults() bool {
	if o == nil || IsNil(o.AllowMultipleResults) {
		var ret bool
		return ret
	}
	return *o.AllowMultipleResults
}

// GetAllowMultipleResultsOk returns a tuple with the AllowMultipleResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcRequestDto) GetAllowMultipleResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowMultipleResults) {
		return nil, false
	}
	return o.AllowMultipleResults, true
}

// HasAllowMultipleResults returns a boolean if a field has been set.
func (o *DtcRequestDto) HasAllowMultipleResults() bool {
	if o != nil && !IsNil(o.AllowMultipleResults) {
		return true
	}

	return false
}

// SetAllowMultipleResults gets a reference to the given bool and assigns it to the AllowMultipleResults field.
func (o *DtcRequestDto) SetAllowMultipleResults(v bool) {
	o.AllowMultipleResults = &v
}

// GetComponentPredict returns the ComponentPredict field value if set, zero value otherwise.
func (o *DtcRequestDto) GetComponentPredict() bool {
	if o == nil || IsNil(o.ComponentPredict) {
		var ret bool
		return ret
	}
	return *o.ComponentPredict
}

// GetComponentPredictOk returns a tuple with the ComponentPredict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcRequestDto) GetComponentPredictOk() (*bool, bool) {
	if o == nil || IsNil(o.ComponentPredict) {
		return nil, false
	}
	return o.ComponentPredict, true
}

// HasComponentPredict returns a boolean if a field has been set.
func (o *DtcRequestDto) HasComponentPredict() bool {
	if o != nil && !IsNil(o.ComponentPredict) {
		return true
	}

	return false
}

// SetComponentPredict gets a reference to the given bool and assigns it to the ComponentPredict field.
func (o *DtcRequestDto) SetComponentPredict(v bool) {
	o.ComponentPredict = &v
}

// GetMilActive returns the MilActive field value if set, zero value otherwise.
func (o *DtcRequestDto) GetMilActive() bool {
	if o == nil || IsNil(o.MilActive) {
		var ret bool
		return ret
	}
	return *o.MilActive
}

// GetMilActiveOk returns a tuple with the MilActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcRequestDto) GetMilActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.MilActive) {
		return nil, false
	}
	return o.MilActive, true
}

// HasMilActive returns a boolean if a field has been set.
func (o *DtcRequestDto) HasMilActive() bool {
	if o != nil && !IsNil(o.MilActive) {
		return true
	}

	return false
}

// SetMilActive gets a reference to the given bool and assigns it to the MilActive field.
func (o *DtcRequestDto) SetMilActive(v bool) {
	o.MilActive = &v
}

// GetRelatedRmi returns the RelatedRmi field value if set, zero value otherwise.
func (o *DtcRequestDto) GetRelatedRmi() string {
	if o == nil || IsNil(o.RelatedRmi) {
		var ret string
		return ret
	}
	return *o.RelatedRmi
}

// GetRelatedRmiOk returns a tuple with the RelatedRmi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcRequestDto) GetRelatedRmiOk() (*string, bool) {
	if o == nil || IsNil(o.RelatedRmi) {
		return nil, false
	}
	return o.RelatedRmi, true
}

// HasRelatedRmi returns a boolean if a field has been set.
func (o *DtcRequestDto) HasRelatedRmi() bool {
	if o != nil && !IsNil(o.RelatedRmi) {
		return true
	}

	return false
}

// SetRelatedRmi gets a reference to the given string and assigns it to the RelatedRmi field.
func (o *DtcRequestDto) SetRelatedRmi(v string) {
	o.RelatedRmi = &v
}

// GetDtcs returns the Dtcs field value
func (o *DtcRequestDto) GetDtcs() []DtcsDto {
	if o == nil {
		var ret []DtcsDto
		return ret
	}

	return o.Dtcs
}

// GetDtcsOk returns a tuple with the Dtcs field value
// and a boolean to check if the value has been set.
func (o *DtcRequestDto) GetDtcsOk() ([]DtcsDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dtcs, true
}

// SetDtcs sets field value
func (o *DtcRequestDto) SetDtcs(v []DtcsDto) {
	o.Dtcs = v
}

func (o DtcRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtcRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DtcReadType) {
		toSerialize["dtcReadType"] = o.DtcReadType
	}
	toSerialize["vin"] = o.Vin
	toSerialize["language"] = o.Language
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.WithRecommendation) {
		toSerialize["withRecommendation"] = o.WithRecommendation
	}
	if !IsNil(o.WithClassification) {
		toSerialize["withClassification"] = o.WithClassification
	}
	if !IsNil(o.AllowMultipleResults) {
		toSerialize["allowMultipleResults"] = o.AllowMultipleResults
	}
	if !IsNil(o.ComponentPredict) {
		toSerialize["componentPredict"] = o.ComponentPredict
	}
	if !IsNil(o.MilActive) {
		toSerialize["milActive"] = o.MilActive
	}
	if !IsNil(o.RelatedRmi) {
		toSerialize["relatedRmi"] = o.RelatedRmi
	}
	toSerialize["dtcs"] = o.Dtcs
	return toSerialize, nil
}

func (o *DtcRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vin",
		"language",
		"dtcs",
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

	varDtcRequestDto := _DtcRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDtcRequestDto)

	if err != nil {
		return err
	}

	*o = DtcRequestDto(varDtcRequestDto)

	return err
}

type NullableDtcRequestDto struct {
	value *DtcRequestDto
	isSet bool
}

func (v NullableDtcRequestDto) Get() *DtcRequestDto {
	return v.value
}

func (v *NullableDtcRequestDto) Set(val *DtcRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDtcRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDtcRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtcRequestDto(val *DtcRequestDto) *NullableDtcRequestDto {
	return &NullableDtcRequestDto{value: val, isSet: true}
}

func (v NullableDtcRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtcRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


