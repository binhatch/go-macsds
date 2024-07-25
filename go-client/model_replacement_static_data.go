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

// checks if the ReplacementStaticData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplacementStaticData{}

// ReplacementStaticData struct for ReplacementStaticData
type ReplacementStaticData struct {
	// Date From
	DateFrom *string `json:"dateFrom,omitempty"`
	// Date To
	DateTo *string `json:"dateTo,omitempty"`
	// Mileage First Time
	MileageFirstTime *int32 `json:"mileageFirstTime,omitempty"`
	// Mileage Interval
	MileageInterval *int32 `json:"mileageInterval,omitempty"`
	// Mileage Static
	MileageStatic *int32 `json:"mileageStatic,omitempty"`
	// Year First Time
	YearFirstTime *float32 `json:"yearFirstTime,omitempty"`
	// Year Interval
	YearInterval *float32 `json:"yearInterval,omitempty"`
	// Year Static
	YearStatic *float32 `json:"yearStatic,omitempty"`
}

// NewReplacementStaticData instantiates a new ReplacementStaticData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplacementStaticData() *ReplacementStaticData {
	this := ReplacementStaticData{}
	return &this
}

// NewReplacementStaticDataWithDefaults instantiates a new ReplacementStaticData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplacementStaticDataWithDefaults() *ReplacementStaticData {
	this := ReplacementStaticData{}
	return &this
}

// GetDateFrom returns the DateFrom field value if set, zero value otherwise.
func (o *ReplacementStaticData) GetDateFrom() string {
	if o == nil || IsNil(o.DateFrom) {
		var ret string
		return ret
	}
	return *o.DateFrom
}

// GetDateFromOk returns a tuple with the DateFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplacementStaticData) GetDateFromOk() (*string, bool) {
	if o == nil || IsNil(o.DateFrom) {
		return nil, false
	}
	return o.DateFrom, true
}

// HasDateFrom returns a boolean if a field has been set.
func (o *ReplacementStaticData) HasDateFrom() bool {
	if o != nil && !IsNil(o.DateFrom) {
		return true
	}

	return false
}

// SetDateFrom gets a reference to the given string and assigns it to the DateFrom field.
func (o *ReplacementStaticData) SetDateFrom(v string) {
	o.DateFrom = &v
}

// GetDateTo returns the DateTo field value if set, zero value otherwise.
func (o *ReplacementStaticData) GetDateTo() string {
	if o == nil || IsNil(o.DateTo) {
		var ret string
		return ret
	}
	return *o.DateTo
}

// GetDateToOk returns a tuple with the DateTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplacementStaticData) GetDateToOk() (*string, bool) {
	if o == nil || IsNil(o.DateTo) {
		return nil, false
	}
	return o.DateTo, true
}

// HasDateTo returns a boolean if a field has been set.
func (o *ReplacementStaticData) HasDateTo() bool {
	if o != nil && !IsNil(o.DateTo) {
		return true
	}

	return false
}

// SetDateTo gets a reference to the given string and assigns it to the DateTo field.
func (o *ReplacementStaticData) SetDateTo(v string) {
	o.DateTo = &v
}

// GetMileageFirstTime returns the MileageFirstTime field value if set, zero value otherwise.
func (o *ReplacementStaticData) GetMileageFirstTime() int32 {
	if o == nil || IsNil(o.MileageFirstTime) {
		var ret int32
		return ret
	}
	return *o.MileageFirstTime
}

// GetMileageFirstTimeOk returns a tuple with the MileageFirstTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplacementStaticData) GetMileageFirstTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.MileageFirstTime) {
		return nil, false
	}
	return o.MileageFirstTime, true
}

// HasMileageFirstTime returns a boolean if a field has been set.
func (o *ReplacementStaticData) HasMileageFirstTime() bool {
	if o != nil && !IsNil(o.MileageFirstTime) {
		return true
	}

	return false
}

// SetMileageFirstTime gets a reference to the given int32 and assigns it to the MileageFirstTime field.
func (o *ReplacementStaticData) SetMileageFirstTime(v int32) {
	o.MileageFirstTime = &v
}

// GetMileageInterval returns the MileageInterval field value if set, zero value otherwise.
func (o *ReplacementStaticData) GetMileageInterval() int32 {
	if o == nil || IsNil(o.MileageInterval) {
		var ret int32
		return ret
	}
	return *o.MileageInterval
}

// GetMileageIntervalOk returns a tuple with the MileageInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplacementStaticData) GetMileageIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.MileageInterval) {
		return nil, false
	}
	return o.MileageInterval, true
}

// HasMileageInterval returns a boolean if a field has been set.
func (o *ReplacementStaticData) HasMileageInterval() bool {
	if o != nil && !IsNil(o.MileageInterval) {
		return true
	}

	return false
}

// SetMileageInterval gets a reference to the given int32 and assigns it to the MileageInterval field.
func (o *ReplacementStaticData) SetMileageInterval(v int32) {
	o.MileageInterval = &v
}

// GetMileageStatic returns the MileageStatic field value if set, zero value otherwise.
func (o *ReplacementStaticData) GetMileageStatic() int32 {
	if o == nil || IsNil(o.MileageStatic) {
		var ret int32
		return ret
	}
	return *o.MileageStatic
}

// GetMileageStaticOk returns a tuple with the MileageStatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplacementStaticData) GetMileageStaticOk() (*int32, bool) {
	if o == nil || IsNil(o.MileageStatic) {
		return nil, false
	}
	return o.MileageStatic, true
}

// HasMileageStatic returns a boolean if a field has been set.
func (o *ReplacementStaticData) HasMileageStatic() bool {
	if o != nil && !IsNil(o.MileageStatic) {
		return true
	}

	return false
}

// SetMileageStatic gets a reference to the given int32 and assigns it to the MileageStatic field.
func (o *ReplacementStaticData) SetMileageStatic(v int32) {
	o.MileageStatic = &v
}

// GetYearFirstTime returns the YearFirstTime field value if set, zero value otherwise.
func (o *ReplacementStaticData) GetYearFirstTime() float32 {
	if o == nil || IsNil(o.YearFirstTime) {
		var ret float32
		return ret
	}
	return *o.YearFirstTime
}

// GetYearFirstTimeOk returns a tuple with the YearFirstTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplacementStaticData) GetYearFirstTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.YearFirstTime) {
		return nil, false
	}
	return o.YearFirstTime, true
}

// HasYearFirstTime returns a boolean if a field has been set.
func (o *ReplacementStaticData) HasYearFirstTime() bool {
	if o != nil && !IsNil(o.YearFirstTime) {
		return true
	}

	return false
}

// SetYearFirstTime gets a reference to the given float32 and assigns it to the YearFirstTime field.
func (o *ReplacementStaticData) SetYearFirstTime(v float32) {
	o.YearFirstTime = &v
}

// GetYearInterval returns the YearInterval field value if set, zero value otherwise.
func (o *ReplacementStaticData) GetYearInterval() float32 {
	if o == nil || IsNil(o.YearInterval) {
		var ret float32
		return ret
	}
	return *o.YearInterval
}

// GetYearIntervalOk returns a tuple with the YearInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplacementStaticData) GetYearIntervalOk() (*float32, bool) {
	if o == nil || IsNil(o.YearInterval) {
		return nil, false
	}
	return o.YearInterval, true
}

// HasYearInterval returns a boolean if a field has been set.
func (o *ReplacementStaticData) HasYearInterval() bool {
	if o != nil && !IsNil(o.YearInterval) {
		return true
	}

	return false
}

// SetYearInterval gets a reference to the given float32 and assigns it to the YearInterval field.
func (o *ReplacementStaticData) SetYearInterval(v float32) {
	o.YearInterval = &v
}

// GetYearStatic returns the YearStatic field value if set, zero value otherwise.
func (o *ReplacementStaticData) GetYearStatic() float32 {
	if o == nil || IsNil(o.YearStatic) {
		var ret float32
		return ret
	}
	return *o.YearStatic
}

// GetYearStaticOk returns a tuple with the YearStatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplacementStaticData) GetYearStaticOk() (*float32, bool) {
	if o == nil || IsNil(o.YearStatic) {
		return nil, false
	}
	return o.YearStatic, true
}

// HasYearStatic returns a boolean if a field has been set.
func (o *ReplacementStaticData) HasYearStatic() bool {
	if o != nil && !IsNil(o.YearStatic) {
		return true
	}

	return false
}

// SetYearStatic gets a reference to the given float32 and assigns it to the YearStatic field.
func (o *ReplacementStaticData) SetYearStatic(v float32) {
	o.YearStatic = &v
}

func (o ReplacementStaticData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplacementStaticData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateFrom) {
		toSerialize["dateFrom"] = o.DateFrom
	}
	if !IsNil(o.DateTo) {
		toSerialize["dateTo"] = o.DateTo
	}
	if !IsNil(o.MileageFirstTime) {
		toSerialize["mileageFirstTime"] = o.MileageFirstTime
	}
	if !IsNil(o.MileageInterval) {
		toSerialize["mileageInterval"] = o.MileageInterval
	}
	if !IsNil(o.MileageStatic) {
		toSerialize["mileageStatic"] = o.MileageStatic
	}
	if !IsNil(o.YearFirstTime) {
		toSerialize["yearFirstTime"] = o.YearFirstTime
	}
	if !IsNil(o.YearInterval) {
		toSerialize["yearInterval"] = o.YearInterval
	}
	if !IsNil(o.YearStatic) {
		toSerialize["yearStatic"] = o.YearStatic
	}
	return toSerialize, nil
}

type NullableReplacementStaticData struct {
	value *ReplacementStaticData
	isSet bool
}

func (v NullableReplacementStaticData) Get() *ReplacementStaticData {
	return v.value
}

func (v *NullableReplacementStaticData) Set(val *ReplacementStaticData) {
	v.value = val
	v.isSet = true
}

func (v NullableReplacementStaticData) IsSet() bool {
	return v.isSet
}

func (v *NullableReplacementStaticData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplacementStaticData(val *ReplacementStaticData) *NullableReplacementStaticData {
	return &NullableReplacementStaticData{value: val, isSet: true}
}

func (v NullableReplacementStaticData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplacementStaticData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

