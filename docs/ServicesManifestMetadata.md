# ServicesManifestMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Digest** | Pointer to **string** |  | [optional] 
**LayerCount** | Pointer to **int32** |  | [optional] 
**MediaType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**SchemaVersion** | Pointer to **int32** |  | [optional] 
**TotalSizeBytes** | Pointer to **int32** |  | [optional] 

## Methods

### NewServicesManifestMetadata

`func NewServicesManifestMetadata() *ServicesManifestMetadata`

NewServicesManifestMetadata instantiates a new ServicesManifestMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesManifestMetadataWithDefaults

`func NewServicesManifestMetadataWithDefaults() *ServicesManifestMetadata`

NewServicesManifestMetadataWithDefaults instantiates a new ServicesManifestMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *ServicesManifestMetadata) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ServicesManifestMetadata) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ServicesManifestMetadata) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *ServicesManifestMetadata) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServicesManifestMetadata) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServicesManifestMetadata) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServicesManifestMetadata) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServicesManifestMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDigest

`func (o *ServicesManifestMetadata) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ServicesManifestMetadata) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ServicesManifestMetadata) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *ServicesManifestMetadata) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetLayerCount

`func (o *ServicesManifestMetadata) GetLayerCount() int32`

GetLayerCount returns the LayerCount field if non-nil, zero value otherwise.

### GetLayerCountOk

`func (o *ServicesManifestMetadata) GetLayerCountOk() (*int32, bool)`

GetLayerCountOk returns a tuple with the LayerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerCount

`func (o *ServicesManifestMetadata) SetLayerCount(v int32)`

SetLayerCount sets LayerCount field to given value.

### HasLayerCount

`func (o *ServicesManifestMetadata) HasLayerCount() bool`

HasLayerCount returns a boolean if a field has been set.

### GetMediaType

`func (o *ServicesManifestMetadata) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *ServicesManifestMetadata) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *ServicesManifestMetadata) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *ServicesManifestMetadata) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetName

`func (o *ServicesManifestMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicesManifestMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicesManifestMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServicesManifestMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *ServicesManifestMetadata) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ServicesManifestMetadata) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ServicesManifestMetadata) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ServicesManifestMetadata) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetReference

`func (o *ServicesManifestMetadata) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ServicesManifestMetadata) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ServicesManifestMetadata) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ServicesManifestMetadata) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ServicesManifestMetadata) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ServicesManifestMetadata) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ServicesManifestMetadata) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ServicesManifestMetadata) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetTotalSizeBytes

`func (o *ServicesManifestMetadata) GetTotalSizeBytes() int32`

GetTotalSizeBytes returns the TotalSizeBytes field if non-nil, zero value otherwise.

### GetTotalSizeBytesOk

`func (o *ServicesManifestMetadata) GetTotalSizeBytesOk() (*int32, bool)`

GetTotalSizeBytesOk returns a tuple with the TotalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeBytes

`func (o *ServicesManifestMetadata) SetTotalSizeBytes(v int32)`

SetTotalSizeBytes sets TotalSizeBytes field to given value.

### HasTotalSizeBytes

`func (o *ServicesManifestMetadata) HasTotalSizeBytes() bool`

HasTotalSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


