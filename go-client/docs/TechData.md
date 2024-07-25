# TechData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]GroupDetails**](GroupDetails.md) | A list of Group Details | [optional] 

## Methods

### NewTechData

`func NewTechData() *TechData`

NewTechData instantiates a new TechData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechDataWithDefaults

`func NewTechDataWithDefaults() *TechData`

NewTechDataWithDefaults instantiates a new TechData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *TechData) GetGroups() []GroupDetails`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *TechData) GetGroupsOk() (*[]GroupDetails, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *TechData) SetGroups(v []GroupDetails)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *TechData) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

