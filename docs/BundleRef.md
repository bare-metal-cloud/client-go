# BundleRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** | content digest, e.g. \&quot;sha256:...\&quot; | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewBundleRef

`func NewBundleRef() *BundleRef`

NewBundleRef instantiates a new BundleRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleRefWithDefaults

`func NewBundleRefWithDefaults() *BundleRef`

NewBundleRefWithDefaults instantiates a new BundleRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *BundleRef) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *BundleRef) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *BundleRef) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *BundleRef) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetName

`func (o *BundleRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BundleRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BundleRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BundleRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *BundleRef) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BundleRef) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BundleRef) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BundleRef) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


