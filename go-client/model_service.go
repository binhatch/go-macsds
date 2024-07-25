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

// checks if the Service type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Service{}

// Service struct for Service
type Service struct {
	// Id
	Id *int32 `json:"id,omitempty"`
	Name *ReplacementNameReference `json:"name,omitempty"`
	// Order
	Order *int32 `json:"order,omitempty"`
	// List of Categories
	Categories []SchemaCategory `json:"categories,omitempty"`
	// List of Replacement Intervals
	ReplacementIntervals []ReplacementInterval `json:"replacementIntervals,omitempty"`
	// Aggregated Labour Time
	AggregatedLabourTime *int32 `json:"aggregatedLabourTime,omitempty"`
	// If mileage, date of production and possibly date and mileage of last service was provided, it is calculated when the related item is due in terms of distance left. Can also be negative, if the service is overdue.
	DueInDistance *int32 `json:"dueInDistance,omitempty"`
	// If mileage, date of production and possibly date and mileage of last service was provided, it is calculated when the related item is due in terms of days left. Can also be negative, if the service is overdue.
	DueInDays *int32 `json:"dueInDays,omitempty"`
}

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService() *Service {
	this := Service{}
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Service) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Service) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Service) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Service) GetName() ReplacementNameReference {
	if o == nil || IsNil(o.Name) {
		var ret ReplacementNameReference
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetNameOk() (*ReplacementNameReference, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Service) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given ReplacementNameReference and assigns it to the Name field.
func (o *Service) SetName(v ReplacementNameReference) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Service) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Service) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *Service) SetOrder(v int32) {
	o.Order = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *Service) GetCategories() []SchemaCategory {
	if o == nil || IsNil(o.Categories) {
		var ret []SchemaCategory
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCategoriesOk() ([]SchemaCategory, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *Service) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []SchemaCategory and assigns it to the Categories field.
func (o *Service) SetCategories(v []SchemaCategory) {
	o.Categories = v
}

// GetReplacementIntervals returns the ReplacementIntervals field value if set, zero value otherwise.
func (o *Service) GetReplacementIntervals() []ReplacementInterval {
	if o == nil || IsNil(o.ReplacementIntervals) {
		var ret []ReplacementInterval
		return ret
	}
	return o.ReplacementIntervals
}

// GetReplacementIntervalsOk returns a tuple with the ReplacementIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetReplacementIntervalsOk() ([]ReplacementInterval, bool) {
	if o == nil || IsNil(o.ReplacementIntervals) {
		return nil, false
	}
	return o.ReplacementIntervals, true
}

// HasReplacementIntervals returns a boolean if a field has been set.
func (o *Service) HasReplacementIntervals() bool {
	if o != nil && !IsNil(o.ReplacementIntervals) {
		return true
	}

	return false
}

// SetReplacementIntervals gets a reference to the given []ReplacementInterval and assigns it to the ReplacementIntervals field.
func (o *Service) SetReplacementIntervals(v []ReplacementInterval) {
	o.ReplacementIntervals = v
}

// GetAggregatedLabourTime returns the AggregatedLabourTime field value if set, zero value otherwise.
func (o *Service) GetAggregatedLabourTime() int32 {
	if o == nil || IsNil(o.AggregatedLabourTime) {
		var ret int32
		return ret
	}
	return *o.AggregatedLabourTime
}

// GetAggregatedLabourTimeOk returns a tuple with the AggregatedLabourTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAggregatedLabourTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.AggregatedLabourTime) {
		return nil, false
	}
	return o.AggregatedLabourTime, true
}

// HasAggregatedLabourTime returns a boolean if a field has been set.
func (o *Service) HasAggregatedLabourTime() bool {
	if o != nil && !IsNil(o.AggregatedLabourTime) {
		return true
	}

	return false
}

// SetAggregatedLabourTime gets a reference to the given int32 and assigns it to the AggregatedLabourTime field.
func (o *Service) SetAggregatedLabourTime(v int32) {
	o.AggregatedLabourTime = &v
}

// GetDueInDistance returns the DueInDistance field value if set, zero value otherwise.
func (o *Service) GetDueInDistance() int32 {
	if o == nil || IsNil(o.DueInDistance) {
		var ret int32
		return ret
	}
	return *o.DueInDistance
}

// GetDueInDistanceOk returns a tuple with the DueInDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDueInDistanceOk() (*int32, bool) {
	if o == nil || IsNil(o.DueInDistance) {
		return nil, false
	}
	return o.DueInDistance, true
}

// HasDueInDistance returns a boolean if a field has been set.
func (o *Service) HasDueInDistance() bool {
	if o != nil && !IsNil(o.DueInDistance) {
		return true
	}

	return false
}

// SetDueInDistance gets a reference to the given int32 and assigns it to the DueInDistance field.
func (o *Service) SetDueInDistance(v int32) {
	o.DueInDistance = &v
}

// GetDueInDays returns the DueInDays field value if set, zero value otherwise.
func (o *Service) GetDueInDays() int32 {
	if o == nil || IsNil(o.DueInDays) {
		var ret int32
		return ret
	}
	return *o.DueInDays
}

// GetDueInDaysOk returns a tuple with the DueInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDueInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.DueInDays) {
		return nil, false
	}
	return o.DueInDays, true
}

// HasDueInDays returns a boolean if a field has been set.
func (o *Service) HasDueInDays() bool {
	if o != nil && !IsNil(o.DueInDays) {
		return true
	}

	return false
}

// SetDueInDays gets a reference to the given int32 and assigns it to the DueInDays field.
func (o *Service) SetDueInDays(v int32) {
	o.DueInDays = &v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Service) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.ReplacementIntervals) {
		toSerialize["replacementIntervals"] = o.ReplacementIntervals
	}
	if !IsNil(o.AggregatedLabourTime) {
		toSerialize["aggregatedLabourTime"] = o.AggregatedLabourTime
	}
	if !IsNil(o.DueInDistance) {
		toSerialize["dueInDistance"] = o.DueInDistance
	}
	if !IsNil(o.DueInDays) {
		toSerialize["dueInDays"] = o.DueInDays
	}
	return toSerialize, nil
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

