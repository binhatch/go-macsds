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

// checks if the SchemaCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaCategory{}

// SchemaCategory struct for SchemaCategory
type SchemaCategory struct {
	// Id
	Id *int32 `json:"id,omitempty"`
	Name *NameReference `json:"name,omitempty"`
	Icon *ObjectReferenceContainer `json:"icon,omitempty"`
	// List of items
	Items []SchemaItem `json:"items,omitempty"`
	// Order
	Order *int32 `json:"order,omitempty"`
}

// NewSchemaCategory instantiates a new SchemaCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaCategory() *SchemaCategory {
	this := SchemaCategory{}
	return &this
}

// NewSchemaCategoryWithDefaults instantiates a new SchemaCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaCategoryWithDefaults() *SchemaCategory {
	this := SchemaCategory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SchemaCategory) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaCategory) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SchemaCategory) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SchemaCategory) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SchemaCategory) GetName() NameReference {
	if o == nil || IsNil(o.Name) {
		var ret NameReference
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaCategory) GetNameOk() (*NameReference, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SchemaCategory) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given NameReference and assigns it to the Name field.
func (o *SchemaCategory) SetName(v NameReference) {
	o.Name = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *SchemaCategory) GetIcon() ObjectReferenceContainer {
	if o == nil || IsNil(o.Icon) {
		var ret ObjectReferenceContainer
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaCategory) GetIconOk() (*ObjectReferenceContainer, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *SchemaCategory) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given ObjectReferenceContainer and assigns it to the Icon field.
func (o *SchemaCategory) SetIcon(v ObjectReferenceContainer) {
	o.Icon = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SchemaCategory) GetItems() []SchemaItem {
	if o == nil || IsNil(o.Items) {
		var ret []SchemaItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaCategory) GetItemsOk() ([]SchemaItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SchemaCategory) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SchemaItem and assigns it to the Items field.
func (o *SchemaCategory) SetItems(v []SchemaItem) {
	o.Items = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *SchemaCategory) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaCategory) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *SchemaCategory) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *SchemaCategory) SetOrder(v int32) {
	o.Order = &v
}

func (o SchemaCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullableSchemaCategory struct {
	value *SchemaCategory
	isSet bool
}

func (v NullableSchemaCategory) Get() *SchemaCategory {
	return v.value
}

func (v *NullableSchemaCategory) Set(val *SchemaCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaCategory(val *SchemaCategory) *NullableSchemaCategory {
	return &NullableSchemaCategory{value: val, isSet: true}
}

func (v NullableSchemaCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


