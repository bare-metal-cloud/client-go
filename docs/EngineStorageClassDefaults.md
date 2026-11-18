# EngineStorageClassDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Declared** | Pointer to **string** | Declared is the platform&#39;s default-class declaration; empty &#x3D; declared nothing. | [optional] 
**Observed** | Pointer to **string** | Observed is the default class the cluster&#39;s annotation named. | [optional] 

## Methods

### NewEngineStorageClassDefaults

`func NewEngineStorageClassDefaults() *EngineStorageClassDefaults`

NewEngineStorageClassDefaults instantiates a new EngineStorageClassDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineStorageClassDefaultsWithDefaults

`func NewEngineStorageClassDefaultsWithDefaults() *EngineStorageClassDefaults`

NewEngineStorageClassDefaultsWithDefaults instantiates a new EngineStorageClassDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeclared

`func (o *EngineStorageClassDefaults) GetDeclared() string`

GetDeclared returns the Declared field if non-nil, zero value otherwise.

### GetDeclaredOk

`func (o *EngineStorageClassDefaults) GetDeclaredOk() (*string, bool)`

GetDeclaredOk returns a tuple with the Declared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclared

`func (o *EngineStorageClassDefaults) SetDeclared(v string)`

SetDeclared sets Declared field to given value.

### HasDeclared

`func (o *EngineStorageClassDefaults) HasDeclared() bool`

HasDeclared returns a boolean if a field has been set.

### GetObserved

`func (o *EngineStorageClassDefaults) GetObserved() string`

GetObserved returns the Observed field if non-nil, zero value otherwise.

### GetObservedOk

`func (o *EngineStorageClassDefaults) GetObservedOk() (*string, bool)`

GetObservedOk returns a tuple with the Observed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObserved

`func (o *EngineStorageClassDefaults) SetObserved(v string)`

SetObserved sets Observed field to given value.

### HasObserved

`func (o *EngineStorageClassDefaults) HasObserved() bool`

HasObserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


