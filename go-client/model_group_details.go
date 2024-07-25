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

// checks if the GroupDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupDetails{}

// GroupDetails struct for GroupDetails
type GroupDetails struct {
	// Level
	Level *int32 `json:"level,omitempty"`
	// Order
	Order *int32 `json:"order,omitempty"`
	// The contained Data
	Data []Data `json:"data,omitempty"`
	// Id
	Id *string `json:"id,omitempty"`
	Text *Translation `json:"text,omitempty"`
	// Another list of Group Details (recursive)
	Groups []GroupDetails `json:"groups,omitempty"`
}

// NewGroupDetails instantiates a new GroupDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupDetails() *GroupDetails {
	this := GroupDetails{}
	return &this
}

// NewGroupDetailsWithDefaults instantiates a new GroupDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupDetailsWithDefaults() *GroupDetails {
	this := GroupDetails{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *GroupDetails) GetLevel() int32 {
	if o == nil || IsNil(o.Level) {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDetails) GetLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *GroupDetails) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *GroupDetails) SetLevel(v int32) {
	o.Level = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GroupDetails) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDetails) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GroupDetails) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *GroupDetails) SetOrder(v int32) {
	o.Order = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GroupDetails) GetData() []Data {
	if o == nil || IsNil(o.Data) {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDetails) GetDataOk() ([]Data, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GroupDetails) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *GroupDetails) SetData(v []Data) {
	o.Data = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupDetails) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDetails) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupDetails) SetId(v string) {
	o.Id = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *GroupDetails) GetText() Translation {
	if o == nil || IsNil(o.Text) {
		var ret Translation
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDetails) GetTextOk() (*Translation, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *GroupDetails) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Translation and assigns it to the Text field.
func (o *GroupDetails) SetText(v Translation) {
	o.Text = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *GroupDetails) GetGroups() []GroupDetails {
	if o == nil || IsNil(o.Groups) {
		var ret []GroupDetails
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDetails) GetGroupsOk() ([]GroupDetails, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *GroupDetails) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupDetails and assigns it to the Groups field.
func (o *GroupDetails) SetGroups(v []GroupDetails) {
	o.Groups = v
}

func (o GroupDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return toSerialize, nil
}

type NullableGroupDetails struct {
	value *GroupDetails
	isSet bool
}

func (v NullableGroupDetails) Get() *GroupDetails {
	return v.value
}

func (v *NullableGroupDetails) Set(val *GroupDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupDetails(val *GroupDetails) *NullableGroupDetails {
	return &NullableGroupDetails{value: val, isSet: true}
}

func (v NullableGroupDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


