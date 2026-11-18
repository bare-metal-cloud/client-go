# BlueprintsourceUpdateDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changed** | Pointer to **bool** | Changed is the one-flag summary: any row anywhere flips it. | [optional] 
**Components** | Pointer to [**[]BlueprintsourceUpdateComponentRow**](BlueprintsourceUpdateComponentRow.md) | Components are the component-SET changes: services the next compile adds or removes (the component-set law&#39;s preview half — the apply path refuses what the law refuses; the diff shows what WOULD change). | [optional] 
**Fields** | Pointer to [**[]BlueprintsourceUpdateFieldRow**](BlueprintsourceUpdateFieldRow.md) | Fields are the form-field set changes (the fields: translation — added, removed, or changed rows). | [optional] 
**Images** | Pointer to [**[]BlueprintsourceUpdateImageRow**](BlueprintsourceUpdateImageRow.md) | Images are the changed image pins: per service, the compiled image reference as each side pins it (name@sha256:… post-pin; a build service compiles imageless on both sides and produces no row). | [optional] 
**Resources** | Pointer to [**[]BlueprintsourceUpdateResourceRow**](BlueprintsourceUpdateResourceRow.md) | Resources are the changed compute limits (cpu / memory per service). | [optional] 

## Methods

### NewBlueprintsourceUpdateDiff

`func NewBlueprintsourceUpdateDiff() *BlueprintsourceUpdateDiff`

NewBlueprintsourceUpdateDiff instantiates a new BlueprintsourceUpdateDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintsourceUpdateDiffWithDefaults

`func NewBlueprintsourceUpdateDiffWithDefaults() *BlueprintsourceUpdateDiff`

NewBlueprintsourceUpdateDiffWithDefaults instantiates a new BlueprintsourceUpdateDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanged

`func (o *BlueprintsourceUpdateDiff) GetChanged() bool`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *BlueprintsourceUpdateDiff) GetChangedOk() (*bool, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *BlueprintsourceUpdateDiff) SetChanged(v bool)`

SetChanged sets Changed field to given value.

### HasChanged

`func (o *BlueprintsourceUpdateDiff) HasChanged() bool`

HasChanged returns a boolean if a field has been set.

### GetComponents

`func (o *BlueprintsourceUpdateDiff) GetComponents() []BlueprintsourceUpdateComponentRow`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BlueprintsourceUpdateDiff) GetComponentsOk() (*[]BlueprintsourceUpdateComponentRow, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BlueprintsourceUpdateDiff) SetComponents(v []BlueprintsourceUpdateComponentRow)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *BlueprintsourceUpdateDiff) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetFields

`func (o *BlueprintsourceUpdateDiff) GetFields() []BlueprintsourceUpdateFieldRow`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *BlueprintsourceUpdateDiff) GetFieldsOk() (*[]BlueprintsourceUpdateFieldRow, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *BlueprintsourceUpdateDiff) SetFields(v []BlueprintsourceUpdateFieldRow)`

SetFields sets Fields field to given value.

### HasFields

`func (o *BlueprintsourceUpdateDiff) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetImages

`func (o *BlueprintsourceUpdateDiff) GetImages() []BlueprintsourceUpdateImageRow`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *BlueprintsourceUpdateDiff) GetImagesOk() (*[]BlueprintsourceUpdateImageRow, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *BlueprintsourceUpdateDiff) SetImages(v []BlueprintsourceUpdateImageRow)`

SetImages sets Images field to given value.

### HasImages

`func (o *BlueprintsourceUpdateDiff) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetResources

`func (o *BlueprintsourceUpdateDiff) GetResources() []BlueprintsourceUpdateResourceRow`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BlueprintsourceUpdateDiff) GetResourcesOk() (*[]BlueprintsourceUpdateResourceRow, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BlueprintsourceUpdateDiff) SetResources(v []BlueprintsourceUpdateResourceRow)`

SetResources sets Resources field to given value.

### HasResources

`func (o *BlueprintsourceUpdateDiff) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


