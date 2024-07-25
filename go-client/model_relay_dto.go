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

// checks if the RelayDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelayDTO{}

// RelayDTO struct for RelayDTO
type RelayDTO struct {
	// Id
	Id *int32 `json:"id,omitempty"`
	Text *Texts `json:"text,omitempty"`
	// Component Fuese Name
	ComponentFuseName *string `json:"componentFuseName,omitempty"`
	LocationText *Texts `json:"locationText,omitempty"`
	// Location Image URL
	LocationImageUrl *string `json:"locationImageUrl,omitempty"`
	// Fuse Box Image URL
	FuseBoxImageUrl *string `json:"fuseBoxImageUrl,omitempty"`
	Coordinates *CoordinatesDTO `json:"coordinates,omitempty"`
}

// NewRelayDTO instantiates a new RelayDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelayDTO() *RelayDTO {
	this := RelayDTO{}
	return &this
}

// NewRelayDTOWithDefaults instantiates a new RelayDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelayDTOWithDefaults() *RelayDTO {
	this := RelayDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RelayDTO) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelayDTO) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RelayDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RelayDTO) SetId(v int32) {
	o.Id = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *RelayDTO) GetText() Texts {
	if o == nil || IsNil(o.Text) {
		var ret Texts
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelayDTO) GetTextOk() (*Texts, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *RelayDTO) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Texts and assigns it to the Text field.
func (o *RelayDTO) SetText(v Texts) {
	o.Text = &v
}

// GetComponentFuseName returns the ComponentFuseName field value if set, zero value otherwise.
func (o *RelayDTO) GetComponentFuseName() string {
	if o == nil || IsNil(o.ComponentFuseName) {
		var ret string
		return ret
	}
	return *o.ComponentFuseName
}

// GetComponentFuseNameOk returns a tuple with the ComponentFuseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelayDTO) GetComponentFuseNameOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentFuseName) {
		return nil, false
	}
	return o.ComponentFuseName, true
}

// HasComponentFuseName returns a boolean if a field has been set.
func (o *RelayDTO) HasComponentFuseName() bool {
	if o != nil && !IsNil(o.ComponentFuseName) {
		return true
	}

	return false
}

// SetComponentFuseName gets a reference to the given string and assigns it to the ComponentFuseName field.
func (o *RelayDTO) SetComponentFuseName(v string) {
	o.ComponentFuseName = &v
}

// GetLocationText returns the LocationText field value if set, zero value otherwise.
func (o *RelayDTO) GetLocationText() Texts {
	if o == nil || IsNil(o.LocationText) {
		var ret Texts
		return ret
	}
	return *o.LocationText
}

// GetLocationTextOk returns a tuple with the LocationText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelayDTO) GetLocationTextOk() (*Texts, bool) {
	if o == nil || IsNil(o.LocationText) {
		return nil, false
	}
	return o.LocationText, true
}

// HasLocationText returns a boolean if a field has been set.
func (o *RelayDTO) HasLocationText() bool {
	if o != nil && !IsNil(o.LocationText) {
		return true
	}

	return false
}

// SetLocationText gets a reference to the given Texts and assigns it to the LocationText field.
func (o *RelayDTO) SetLocationText(v Texts) {
	o.LocationText = &v
}

// GetLocationImageUrl returns the LocationImageUrl field value if set, zero value otherwise.
func (o *RelayDTO) GetLocationImageUrl() string {
	if o == nil || IsNil(o.LocationImageUrl) {
		var ret string
		return ret
	}
	return *o.LocationImageUrl
}

// GetLocationImageUrlOk returns a tuple with the LocationImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelayDTO) GetLocationImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LocationImageUrl) {
		return nil, false
	}
	return o.LocationImageUrl, true
}

// HasLocationImageUrl returns a boolean if a field has been set.
func (o *RelayDTO) HasLocationImageUrl() bool {
	if o != nil && !IsNil(o.LocationImageUrl) {
		return true
	}

	return false
}

// SetLocationImageUrl gets a reference to the given string and assigns it to the LocationImageUrl field.
func (o *RelayDTO) SetLocationImageUrl(v string) {
	o.LocationImageUrl = &v
}

// GetFuseBoxImageUrl returns the FuseBoxImageUrl field value if set, zero value otherwise.
func (o *RelayDTO) GetFuseBoxImageUrl() string {
	if o == nil || IsNil(o.FuseBoxImageUrl) {
		var ret string
		return ret
	}
	return *o.FuseBoxImageUrl
}

// GetFuseBoxImageUrlOk returns a tuple with the FuseBoxImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelayDTO) GetFuseBoxImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FuseBoxImageUrl) {
		return nil, false
	}
	return o.FuseBoxImageUrl, true
}

// HasFuseBoxImageUrl returns a boolean if a field has been set.
func (o *RelayDTO) HasFuseBoxImageUrl() bool {
	if o != nil && !IsNil(o.FuseBoxImageUrl) {
		return true
	}

	return false
}

// SetFuseBoxImageUrl gets a reference to the given string and assigns it to the FuseBoxImageUrl field.
func (o *RelayDTO) SetFuseBoxImageUrl(v string) {
	o.FuseBoxImageUrl = &v
}

// GetCoordinates returns the Coordinates field value if set, zero value otherwise.
func (o *RelayDTO) GetCoordinates() CoordinatesDTO {
	if o == nil || IsNil(o.Coordinates) {
		var ret CoordinatesDTO
		return ret
	}
	return *o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelayDTO) GetCoordinatesOk() (*CoordinatesDTO, bool) {
	if o == nil || IsNil(o.Coordinates) {
		return nil, false
	}
	return o.Coordinates, true
}

// HasCoordinates returns a boolean if a field has been set.
func (o *RelayDTO) HasCoordinates() bool {
	if o != nil && !IsNil(o.Coordinates) {
		return true
	}

	return false
}

// SetCoordinates gets a reference to the given CoordinatesDTO and assigns it to the Coordinates field.
func (o *RelayDTO) SetCoordinates(v CoordinatesDTO) {
	o.Coordinates = &v
}

func (o RelayDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelayDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.ComponentFuseName) {
		toSerialize["componentFuseName"] = o.ComponentFuseName
	}
	if !IsNil(o.LocationText) {
		toSerialize["locationText"] = o.LocationText
	}
	if !IsNil(o.LocationImageUrl) {
		toSerialize["locationImageUrl"] = o.LocationImageUrl
	}
	if !IsNil(o.FuseBoxImageUrl) {
		toSerialize["fuseBoxImageUrl"] = o.FuseBoxImageUrl
	}
	if !IsNil(o.Coordinates) {
		toSerialize["coordinates"] = o.Coordinates
	}
	return toSerialize, nil
}

type NullableRelayDTO struct {
	value *RelayDTO
	isSet bool
}

func (v NullableRelayDTO) Get() *RelayDTO {
	return v.value
}

func (v *NullableRelayDTO) Set(val *RelayDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableRelayDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableRelayDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelayDTO(val *RelayDTO) *NullableRelayDTO {
	return &NullableRelayDTO{value: val, isSet: true}
}

func (v NullableRelayDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelayDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

