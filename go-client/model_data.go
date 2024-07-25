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

// checks if the Data type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Data{}

// Data struct for Data
type Data struct {
	// Id
	Id *string `json:"id,omitempty"`
	Value *Value `json:"value,omitempty"`
	Variant *Variant `json:"variant,omitempty"`
	Criteria *Criteria `json:"criteria,omitempty"`
	// A list of Guide References
	Guides []GuideReference `json:"guides,omitempty"`
	// A list of Images
	Images []Image `json:"images,omitempty"`
	// A list of Tables
	Tables []Table `json:"tables,omitempty"`
	// A list of Notes
	Notes []Note `json:"notes,omitempty"`
}

// NewData instantiates a new Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewData() *Data {
	this := Data{}
	return &this
}

// NewDataWithDefaults instantiates a new Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataWithDefaults() *Data {
	this := Data{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Data) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Data) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Data) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Data) SetId(v string) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Data) GetValue() Value {
	if o == nil || IsNil(o.Value) {
		var ret Value
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Data) GetValueOk() (*Value, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Data) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given Value and assigns it to the Value field.
func (o *Data) SetValue(v Value) {
	o.Value = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *Data) GetVariant() Variant {
	if o == nil || IsNil(o.Variant) {
		var ret Variant
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Data) GetVariantOk() (*Variant, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *Data) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given Variant and assigns it to the Variant field.
func (o *Data) SetVariant(v Variant) {
	o.Variant = &v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *Data) GetCriteria() Criteria {
	if o == nil || IsNil(o.Criteria) {
		var ret Criteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Data) GetCriteriaOk() (*Criteria, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *Data) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given Criteria and assigns it to the Criteria field.
func (o *Data) SetCriteria(v Criteria) {
	o.Criteria = &v
}

// GetGuides returns the Guides field value if set, zero value otherwise.
func (o *Data) GetGuides() []GuideReference {
	if o == nil || IsNil(o.Guides) {
		var ret []GuideReference
		return ret
	}
	return o.Guides
}

// GetGuidesOk returns a tuple with the Guides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Data) GetGuidesOk() ([]GuideReference, bool) {
	if o == nil || IsNil(o.Guides) {
		return nil, false
	}
	return o.Guides, true
}

// HasGuides returns a boolean if a field has been set.
func (o *Data) HasGuides() bool {
	if o != nil && !IsNil(o.Guides) {
		return true
	}

	return false
}

// SetGuides gets a reference to the given []GuideReference and assigns it to the Guides field.
func (o *Data) SetGuides(v []GuideReference) {
	o.Guides = v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *Data) GetImages() []Image {
	if o == nil || IsNil(o.Images) {
		var ret []Image
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Data) GetImagesOk() ([]Image, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *Data) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []Image and assigns it to the Images field.
func (o *Data) SetImages(v []Image) {
	o.Images = v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *Data) GetTables() []Table {
	if o == nil || IsNil(o.Tables) {
		var ret []Table
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Data) GetTablesOk() ([]Table, bool) {
	if o == nil || IsNil(o.Tables) {
		return nil, false
	}
	return o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *Data) HasTables() bool {
	if o != nil && !IsNil(o.Tables) {
		return true
	}

	return false
}

// SetTables gets a reference to the given []Table and assigns it to the Tables field.
func (o *Data) SetTables(v []Table) {
	o.Tables = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *Data) GetNotes() []Note {
	if o == nil || IsNil(o.Notes) {
		var ret []Note
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Data) GetNotesOk() ([]Note, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *Data) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []Note and assigns it to the Notes field.
func (o *Data) SetNotes(v []Note) {
	o.Notes = v
}

func (o Data) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Data) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Guides) {
		toSerialize["guides"] = o.Guides
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !IsNil(o.Tables) {
		toSerialize["tables"] = o.Tables
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableData struct {
	value *Data
	isSet bool
}

func (v NullableData) Get() *Data {
	return v.value
}

func (v *NullableData) Set(val *Data) {
	v.value = val
	v.isSet = true
}

func (v NullableData) IsSet() bool {
	return v.isSet
}

func (v *NullableData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableData(val *Data) *NullableData {
	return &NullableData{value: val, isSet: true}
}

func (v NullableData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

