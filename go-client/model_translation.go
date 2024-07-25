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

// checks if the Translation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Translation{}

// Translation Contains the textual representation and the language code. May be 'null'.
type Translation struct {
	// Danish Text
	Da *string `json:"da,omitempty"`
	// English Text
	En *string `json:"en,omitempty"`
	// German Text
	De *string `json:"de,omitempty"`
	// Swedish Text
	Sv *string `json:"sv,omitempty"`
	// Norwegian Text
	No *string `json:"no,omitempty"`
	// Finnish Text
	Fi *string `json:"fi,omitempty"`
	// French Text
	Fr *string `json:"fr,omitempty"`
	// Spanish Text
	Es *string `json:"es,omitempty"`
	// Polish Text
	Pl *string `json:"pl,omitempty"`
	// Italian Text
	It *string `json:"it,omitempty"`
	// Dutch Text
	Nl *string `json:"nl,omitempty"`
	// Greek Text
	El *string `json:"el,omitempty"`
	// Canadian Text
	Cn *string `json:"cn,omitempty"`
	// Russian Text
	Ru *string `json:"ru,omitempty"`
	// Czech Text
	Cs *string `json:"cs,omitempty"`
	// Turkish Text
	Tr *string `json:"tr,omitempty"`
	// Language
	Language *string `json:"language,omitempty"`
	// Title
	Title *string `json:"title,omitempty"`
}

// NewTranslation instantiates a new Translation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranslation() *Translation {
	this := Translation{}
	return &this
}

// NewTranslationWithDefaults instantiates a new Translation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranslationWithDefaults() *Translation {
	this := Translation{}
	return &this
}

// GetDa returns the Da field value if set, zero value otherwise.
func (o *Translation) GetDa() string {
	if o == nil || IsNil(o.Da) {
		var ret string
		return ret
	}
	return *o.Da
}

// GetDaOk returns a tuple with the Da field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetDaOk() (*string, bool) {
	if o == nil || IsNil(o.Da) {
		return nil, false
	}
	return o.Da, true
}

// HasDa returns a boolean if a field has been set.
func (o *Translation) HasDa() bool {
	if o != nil && !IsNil(o.Da) {
		return true
	}

	return false
}

// SetDa gets a reference to the given string and assigns it to the Da field.
func (o *Translation) SetDa(v string) {
	o.Da = &v
}

// GetEn returns the En field value if set, zero value otherwise.
func (o *Translation) GetEn() string {
	if o == nil || IsNil(o.En) {
		var ret string
		return ret
	}
	return *o.En
}

// GetEnOk returns a tuple with the En field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetEnOk() (*string, bool) {
	if o == nil || IsNil(o.En) {
		return nil, false
	}
	return o.En, true
}

// HasEn returns a boolean if a field has been set.
func (o *Translation) HasEn() bool {
	if o != nil && !IsNil(o.En) {
		return true
	}

	return false
}

// SetEn gets a reference to the given string and assigns it to the En field.
func (o *Translation) SetEn(v string) {
	o.En = &v
}

// GetDe returns the De field value if set, zero value otherwise.
func (o *Translation) GetDe() string {
	if o == nil || IsNil(o.De) {
		var ret string
		return ret
	}
	return *o.De
}

// GetDeOk returns a tuple with the De field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetDeOk() (*string, bool) {
	if o == nil || IsNil(o.De) {
		return nil, false
	}
	return o.De, true
}

// HasDe returns a boolean if a field has been set.
func (o *Translation) HasDe() bool {
	if o != nil && !IsNil(o.De) {
		return true
	}

	return false
}

// SetDe gets a reference to the given string and assigns it to the De field.
func (o *Translation) SetDe(v string) {
	o.De = &v
}

// GetSv returns the Sv field value if set, zero value otherwise.
func (o *Translation) GetSv() string {
	if o == nil || IsNil(o.Sv) {
		var ret string
		return ret
	}
	return *o.Sv
}

// GetSvOk returns a tuple with the Sv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetSvOk() (*string, bool) {
	if o == nil || IsNil(o.Sv) {
		return nil, false
	}
	return o.Sv, true
}

// HasSv returns a boolean if a field has been set.
func (o *Translation) HasSv() bool {
	if o != nil && !IsNil(o.Sv) {
		return true
	}

	return false
}

// SetSv gets a reference to the given string and assigns it to the Sv field.
func (o *Translation) SetSv(v string) {
	o.Sv = &v
}

// GetNo returns the No field value if set, zero value otherwise.
func (o *Translation) GetNo() string {
	if o == nil || IsNil(o.No) {
		var ret string
		return ret
	}
	return *o.No
}

// GetNoOk returns a tuple with the No field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetNoOk() (*string, bool) {
	if o == nil || IsNil(o.No) {
		return nil, false
	}
	return o.No, true
}

// HasNo returns a boolean if a field has been set.
func (o *Translation) HasNo() bool {
	if o != nil && !IsNil(o.No) {
		return true
	}

	return false
}

// SetNo gets a reference to the given string and assigns it to the No field.
func (o *Translation) SetNo(v string) {
	o.No = &v
}

// GetFi returns the Fi field value if set, zero value otherwise.
func (o *Translation) GetFi() string {
	if o == nil || IsNil(o.Fi) {
		var ret string
		return ret
	}
	return *o.Fi
}

// GetFiOk returns a tuple with the Fi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetFiOk() (*string, bool) {
	if o == nil || IsNil(o.Fi) {
		return nil, false
	}
	return o.Fi, true
}

// HasFi returns a boolean if a field has been set.
func (o *Translation) HasFi() bool {
	if o != nil && !IsNil(o.Fi) {
		return true
	}

	return false
}

// SetFi gets a reference to the given string and assigns it to the Fi field.
func (o *Translation) SetFi(v string) {
	o.Fi = &v
}

// GetFr returns the Fr field value if set, zero value otherwise.
func (o *Translation) GetFr() string {
	if o == nil || IsNil(o.Fr) {
		var ret string
		return ret
	}
	return *o.Fr
}

// GetFrOk returns a tuple with the Fr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetFrOk() (*string, bool) {
	if o == nil || IsNil(o.Fr) {
		return nil, false
	}
	return o.Fr, true
}

// HasFr returns a boolean if a field has been set.
func (o *Translation) HasFr() bool {
	if o != nil && !IsNil(o.Fr) {
		return true
	}

	return false
}

// SetFr gets a reference to the given string and assigns it to the Fr field.
func (o *Translation) SetFr(v string) {
	o.Fr = &v
}

// GetEs returns the Es field value if set, zero value otherwise.
func (o *Translation) GetEs() string {
	if o == nil || IsNil(o.Es) {
		var ret string
		return ret
	}
	return *o.Es
}

// GetEsOk returns a tuple with the Es field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetEsOk() (*string, bool) {
	if o == nil || IsNil(o.Es) {
		return nil, false
	}
	return o.Es, true
}

// HasEs returns a boolean if a field has been set.
func (o *Translation) HasEs() bool {
	if o != nil && !IsNil(o.Es) {
		return true
	}

	return false
}

// SetEs gets a reference to the given string and assigns it to the Es field.
func (o *Translation) SetEs(v string) {
	o.Es = &v
}

// GetPl returns the Pl field value if set, zero value otherwise.
func (o *Translation) GetPl() string {
	if o == nil || IsNil(o.Pl) {
		var ret string
		return ret
	}
	return *o.Pl
}

// GetPlOk returns a tuple with the Pl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetPlOk() (*string, bool) {
	if o == nil || IsNil(o.Pl) {
		return nil, false
	}
	return o.Pl, true
}

// HasPl returns a boolean if a field has been set.
func (o *Translation) HasPl() bool {
	if o != nil && !IsNil(o.Pl) {
		return true
	}

	return false
}

// SetPl gets a reference to the given string and assigns it to the Pl field.
func (o *Translation) SetPl(v string) {
	o.Pl = &v
}

// GetIt returns the It field value if set, zero value otherwise.
func (o *Translation) GetIt() string {
	if o == nil || IsNil(o.It) {
		var ret string
		return ret
	}
	return *o.It
}

// GetItOk returns a tuple with the It field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetItOk() (*string, bool) {
	if o == nil || IsNil(o.It) {
		return nil, false
	}
	return o.It, true
}

// HasIt returns a boolean if a field has been set.
func (o *Translation) HasIt() bool {
	if o != nil && !IsNil(o.It) {
		return true
	}

	return false
}

// SetIt gets a reference to the given string and assigns it to the It field.
func (o *Translation) SetIt(v string) {
	o.It = &v
}

// GetNl returns the Nl field value if set, zero value otherwise.
func (o *Translation) GetNl() string {
	if o == nil || IsNil(o.Nl) {
		var ret string
		return ret
	}
	return *o.Nl
}

// GetNlOk returns a tuple with the Nl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetNlOk() (*string, bool) {
	if o == nil || IsNil(o.Nl) {
		return nil, false
	}
	return o.Nl, true
}

// HasNl returns a boolean if a field has been set.
func (o *Translation) HasNl() bool {
	if o != nil && !IsNil(o.Nl) {
		return true
	}

	return false
}

// SetNl gets a reference to the given string and assigns it to the Nl field.
func (o *Translation) SetNl(v string) {
	o.Nl = &v
}

// GetEl returns the El field value if set, zero value otherwise.
func (o *Translation) GetEl() string {
	if o == nil || IsNil(o.El) {
		var ret string
		return ret
	}
	return *o.El
}

// GetElOk returns a tuple with the El field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetElOk() (*string, bool) {
	if o == nil || IsNil(o.El) {
		return nil, false
	}
	return o.El, true
}

// HasEl returns a boolean if a field has been set.
func (o *Translation) HasEl() bool {
	if o != nil && !IsNil(o.El) {
		return true
	}

	return false
}

// SetEl gets a reference to the given string and assigns it to the El field.
func (o *Translation) SetEl(v string) {
	o.El = &v
}

// GetCn returns the Cn field value if set, zero value otherwise.
func (o *Translation) GetCn() string {
	if o == nil || IsNil(o.Cn) {
		var ret string
		return ret
	}
	return *o.Cn
}

// GetCnOk returns a tuple with the Cn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetCnOk() (*string, bool) {
	if o == nil || IsNil(o.Cn) {
		return nil, false
	}
	return o.Cn, true
}

// HasCn returns a boolean if a field has been set.
func (o *Translation) HasCn() bool {
	if o != nil && !IsNil(o.Cn) {
		return true
	}

	return false
}

// SetCn gets a reference to the given string and assigns it to the Cn field.
func (o *Translation) SetCn(v string) {
	o.Cn = &v
}

// GetRu returns the Ru field value if set, zero value otherwise.
func (o *Translation) GetRu() string {
	if o == nil || IsNil(o.Ru) {
		var ret string
		return ret
	}
	return *o.Ru
}

// GetRuOk returns a tuple with the Ru field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetRuOk() (*string, bool) {
	if o == nil || IsNil(o.Ru) {
		return nil, false
	}
	return o.Ru, true
}

// HasRu returns a boolean if a field has been set.
func (o *Translation) HasRu() bool {
	if o != nil && !IsNil(o.Ru) {
		return true
	}

	return false
}

// SetRu gets a reference to the given string and assigns it to the Ru field.
func (o *Translation) SetRu(v string) {
	o.Ru = &v
}

// GetCs returns the Cs field value if set, zero value otherwise.
func (o *Translation) GetCs() string {
	if o == nil || IsNil(o.Cs) {
		var ret string
		return ret
	}
	return *o.Cs
}

// GetCsOk returns a tuple with the Cs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetCsOk() (*string, bool) {
	if o == nil || IsNil(o.Cs) {
		return nil, false
	}
	return o.Cs, true
}

// HasCs returns a boolean if a field has been set.
func (o *Translation) HasCs() bool {
	if o != nil && !IsNil(o.Cs) {
		return true
	}

	return false
}

// SetCs gets a reference to the given string and assigns it to the Cs field.
func (o *Translation) SetCs(v string) {
	o.Cs = &v
}

// GetTr returns the Tr field value if set, zero value otherwise.
func (o *Translation) GetTr() string {
	if o == nil || IsNil(o.Tr) {
		var ret string
		return ret
	}
	return *o.Tr
}

// GetTrOk returns a tuple with the Tr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetTrOk() (*string, bool) {
	if o == nil || IsNil(o.Tr) {
		return nil, false
	}
	return o.Tr, true
}

// HasTr returns a boolean if a field has been set.
func (o *Translation) HasTr() bool {
	if o != nil && !IsNil(o.Tr) {
		return true
	}

	return false
}

// SetTr gets a reference to the given string and assigns it to the Tr field.
func (o *Translation) SetTr(v string) {
	o.Tr = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Translation) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Translation) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Translation) SetLanguage(v string) {
	o.Language = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Translation) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Translation) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Translation) SetTitle(v string) {
	o.Title = &v
}

func (o Translation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Translation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Da) {
		toSerialize["da"] = o.Da
	}
	if !IsNil(o.En) {
		toSerialize["en"] = o.En
	}
	if !IsNil(o.De) {
		toSerialize["de"] = o.De
	}
	if !IsNil(o.Sv) {
		toSerialize["sv"] = o.Sv
	}
	if !IsNil(o.No) {
		toSerialize["no"] = o.No
	}
	if !IsNil(o.Fi) {
		toSerialize["fi"] = o.Fi
	}
	if !IsNil(o.Fr) {
		toSerialize["fr"] = o.Fr
	}
	if !IsNil(o.Es) {
		toSerialize["es"] = o.Es
	}
	if !IsNil(o.Pl) {
		toSerialize["pl"] = o.Pl
	}
	if !IsNil(o.It) {
		toSerialize["it"] = o.It
	}
	if !IsNil(o.Nl) {
		toSerialize["nl"] = o.Nl
	}
	if !IsNil(o.El) {
		toSerialize["el"] = o.El
	}
	if !IsNil(o.Cn) {
		toSerialize["cn"] = o.Cn
	}
	if !IsNil(o.Ru) {
		toSerialize["ru"] = o.Ru
	}
	if !IsNil(o.Cs) {
		toSerialize["cs"] = o.Cs
	}
	if !IsNil(o.Tr) {
		toSerialize["tr"] = o.Tr
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableTranslation struct {
	value *Translation
	isSet bool
}

func (v NullableTranslation) Get() *Translation {
	return v.value
}

func (v *NullableTranslation) Set(val *Translation) {
	v.value = val
	v.isSet = true
}

func (v NullableTranslation) IsSet() bool {
	return v.isSet
}

func (v *NullableTranslation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranslation(val *Translation) *NullableTranslation {
	return &NullableTranslation{value: val, isSet: true}
}

func (v NullableTranslation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranslation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


