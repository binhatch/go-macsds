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

// checks if the IdentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentResponse{}

// IdentResponse struct for IdentResponse
type IdentResponse struct {
	// Status
	Status *int32 `json:"status,omitempty"`
	Error *Error `json:"error,omitempty"`
	// Page
	Page *int32 `json:"page,omitempty"`
	// Number of total items
	TotalItems *int32 `json:"totalItems,omitempty"`
	// Number of total pages
	TotalPages *int32 `json:"totalPages,omitempty"`
	// Response data array of objects
	Data []IdentResponseDTO `json:"data,omitempty"`
}

// NewIdentResponse instantiates a new IdentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentResponse() *IdentResponse {
	this := IdentResponse{}
	return &this
}

// NewIdentResponseWithDefaults instantiates a new IdentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentResponseWithDefaults() *IdentResponse {
	this := IdentResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IdentResponse) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentResponse) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IdentResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *IdentResponse) SetStatus(v int32) {
	o.Status = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *IdentResponse) GetError() Error {
	if o == nil || IsNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentResponse) GetErrorOk() (*Error, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *IdentResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *IdentResponse) SetError(v Error) {
	o.Error = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *IdentResponse) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentResponse) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *IdentResponse) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *IdentResponse) SetPage(v int32) {
	o.Page = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *IdentResponse) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentResponse) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *IdentResponse) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *IdentResponse) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *IdentResponse) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentResponse) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *IdentResponse) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *IdentResponse) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IdentResponse) GetData() []IdentResponseDTO {
	if o == nil || IsNil(o.Data) {
		var ret []IdentResponseDTO
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentResponse) GetDataOk() ([]IdentResponseDTO, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IdentResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []IdentResponseDTO and assigns it to the Data field.
func (o *IdentResponse) SetData(v []IdentResponseDTO) {
	o.Data = v
}

func (o IdentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.TotalItems) {
		toSerialize["totalItems"] = o.TotalItems
	}
	if !IsNil(o.TotalPages) {
		toSerialize["totalPages"] = o.TotalPages
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableIdentResponse struct {
	value *IdentResponse
	isSet bool
}

func (v NullableIdentResponse) Get() *IdentResponse {
	return v.value
}

func (v *NullableIdentResponse) Set(val *IdentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentResponse(val *IdentResponse) *NullableIdentResponse {
	return &NullableIdentResponse{value: val, isSet: true}
}

func (v NullableIdentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

