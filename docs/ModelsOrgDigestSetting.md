# ModelsOrgDigestSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cadence** | Pointer to **string** | Cadence is DigestCadenceWeekly (default) or DigestCadenceDaily (the opt-in). Validated at the service layer. | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IncludeEmpty** | Pointer to **bool** | IncludeEmpty is the empty-summaries opt-in: false (default) means a quiet window suppresses the digest entirely; true means the org gets the explicit \&quot;all quiet\&quot; summary. | [optional] 
**OrgId** | Pointer to **int32** | OrgID is UNIQUE — one settings row per org. | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsOrgDigestSetting

`func NewModelsOrgDigestSetting() *ModelsOrgDigestSetting`

NewModelsOrgDigestSetting instantiates a new ModelsOrgDigestSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsOrgDigestSettingWithDefaults

`func NewModelsOrgDigestSettingWithDefaults() *ModelsOrgDigestSetting`

NewModelsOrgDigestSettingWithDefaults instantiates a new ModelsOrgDigestSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCadence

`func (o *ModelsOrgDigestSetting) GetCadence() string`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *ModelsOrgDigestSetting) GetCadenceOk() (*string, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *ModelsOrgDigestSetting) SetCadence(v string)`

SetCadence sets Cadence field to given value.

### HasCadence

`func (o *ModelsOrgDigestSetting) HasCadence() bool`

HasCadence returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsOrgDigestSetting) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsOrgDigestSetting) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsOrgDigestSetting) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsOrgDigestSetting) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ModelsOrgDigestSetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsOrgDigestSetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsOrgDigestSetting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsOrgDigestSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncludeEmpty

`func (o *ModelsOrgDigestSetting) GetIncludeEmpty() bool`

GetIncludeEmpty returns the IncludeEmpty field if non-nil, zero value otherwise.

### GetIncludeEmptyOk

`func (o *ModelsOrgDigestSetting) GetIncludeEmptyOk() (*bool, bool)`

GetIncludeEmptyOk returns a tuple with the IncludeEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEmpty

`func (o *ModelsOrgDigestSetting) SetIncludeEmpty(v bool)`

SetIncludeEmpty sets IncludeEmpty field to given value.

### HasIncludeEmpty

`func (o *ModelsOrgDigestSetting) HasIncludeEmpty() bool`

HasIncludeEmpty returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsOrgDigestSetting) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsOrgDigestSetting) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsOrgDigestSetting) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsOrgDigestSetting) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsOrgDigestSetting) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsOrgDigestSetting) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsOrgDigestSetting) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsOrgDigestSetting) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


