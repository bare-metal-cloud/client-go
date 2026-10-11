# ServicesUserRequirementStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowlistRoutes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsSkippable** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Satisfaction** | Pointer to [**ModelsSatisfactionMethod**](ModelsSatisfactionMethod.md) |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | pending | fulfilled | skipped | not_applicable | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewServicesUserRequirementStatus

`func NewServicesUserRequirementStatus() *ServicesUserRequirementStatus`

NewServicesUserRequirementStatus instantiates a new ServicesUserRequirementStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesUserRequirementStatusWithDefaults

`func NewServicesUserRequirementStatusWithDefaults() *ServicesUserRequirementStatus`

NewServicesUserRequirementStatusWithDefaults instantiates a new ServicesUserRequirementStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowlistRoutes

`func (o *ServicesUserRequirementStatus) GetAllowlistRoutes() []map[string]interface{}`

GetAllowlistRoutes returns the AllowlistRoutes field if non-nil, zero value otherwise.

### GetAllowlistRoutesOk

`func (o *ServicesUserRequirementStatus) GetAllowlistRoutesOk() (*[]map[string]interface{}, bool)`

GetAllowlistRoutesOk returns a tuple with the AllowlistRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistRoutes

`func (o *ServicesUserRequirementStatus) SetAllowlistRoutes(v []map[string]interface{})`

SetAllowlistRoutes sets AllowlistRoutes field to given value.

### HasAllowlistRoutes

`func (o *ServicesUserRequirementStatus) HasAllowlistRoutes() bool`

HasAllowlistRoutes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServicesUserRequirementStatus) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServicesUserRequirementStatus) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServicesUserRequirementStatus) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServicesUserRequirementStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ServicesUserRequirementStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicesUserRequirementStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicesUserRequirementStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicesUserRequirementStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ServicesUserRequirementStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServicesUserRequirementStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServicesUserRequirementStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServicesUserRequirementStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *ServicesUserRequirementStatus) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ServicesUserRequirementStatus) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ServicesUserRequirementStatus) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ServicesUserRequirementStatus) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsSkippable

`func (o *ServicesUserRequirementStatus) GetIsSkippable() bool`

GetIsSkippable returns the IsSkippable field if non-nil, zero value otherwise.

### GetIsSkippableOk

`func (o *ServicesUserRequirementStatus) GetIsSkippableOk() (*bool, bool)`

GetIsSkippableOk returns a tuple with the IsSkippable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSkippable

`func (o *ServicesUserRequirementStatus) SetIsSkippable(v bool)`

SetIsSkippable sets IsSkippable field to given value.

### HasIsSkippable

`func (o *ServicesUserRequirementStatus) HasIsSkippable() bool`

HasIsSkippable returns a boolean if a field has been set.

### GetKey

`func (o *ServicesUserRequirementStatus) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ServicesUserRequirementStatus) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ServicesUserRequirementStatus) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ServicesUserRequirementStatus) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSatisfaction

`func (o *ServicesUserRequirementStatus) GetSatisfaction() ModelsSatisfactionMethod`

GetSatisfaction returns the Satisfaction field if non-nil, zero value otherwise.

### GetSatisfactionOk

`func (o *ServicesUserRequirementStatus) GetSatisfactionOk() (*ModelsSatisfactionMethod, bool)`

GetSatisfactionOk returns a tuple with the Satisfaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatisfaction

`func (o *ServicesUserRequirementStatus) SetSatisfaction(v ModelsSatisfactionMethod)`

SetSatisfaction sets Satisfaction field to given value.

### HasSatisfaction

`func (o *ServicesUserRequirementStatus) HasSatisfaction() bool`

HasSatisfaction returns a boolean if a field has been set.

### GetSortOrder

`func (o *ServicesUserRequirementStatus) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ServicesUserRequirementStatus) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ServicesUserRequirementStatus) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ServicesUserRequirementStatus) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetStatus

`func (o *ServicesUserRequirementStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServicesUserRequirementStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServicesUserRequirementStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServicesUserRequirementStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *ServicesUserRequirementStatus) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServicesUserRequirementStatus) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServicesUserRequirementStatus) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ServicesUserRequirementStatus) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServicesUserRequirementStatus) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServicesUserRequirementStatus) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServicesUserRequirementStatus) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServicesUserRequirementStatus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


