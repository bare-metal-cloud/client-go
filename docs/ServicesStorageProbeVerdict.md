# ServicesStorageProbeVerdict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** | BucketName and BucketVerdict are set when the per-bucket fallback check ran (listing denied + a bucket was typed). BucketVerdict uses the StorageBucketVerdict states above. | [optional] 
**BucketVerdict** | Pointer to [**ServicesStorageBucketVerdict**](ServicesStorageBucketVerdict.md) |  | [optional] 
**ListingPermitted** | Pointer to **bool** | ListingPermitted is the three-state listing answer: true when the list call succeeded, false when it was classified as permission-denied, nil when unknown (bad key, unreachable endpoint, timeout, wrong region, or anything unclassified). Null is an honest \&quot;we do not know\&quot;, never a fabricated 0 (FR-Z-13). | [optional] 
**Message** | Pointer to **string** | Message is a human sentence naming the gate. Credential material is redacted out (see RedactCredentials). | [optional] 
**MissingPermission** | Pointer to **string** | MissingPermission is the provider permission the call needs, set ONLY when the failure was classified as permission-denied. Empty otherwise. | [optional] 
**PermissionDoc** | Pointer to **string** | PermissionDoc is the provider&#39;s documentation URL for MissingPermission, so the console can link the fix. | [optional] 

## Methods

### NewServicesStorageProbeVerdict

`func NewServicesStorageProbeVerdict() *ServicesStorageProbeVerdict`

NewServicesStorageProbeVerdict instantiates a new ServicesStorageProbeVerdict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesStorageProbeVerdictWithDefaults

`func NewServicesStorageProbeVerdictWithDefaults() *ServicesStorageProbeVerdict`

NewServicesStorageProbeVerdictWithDefaults instantiates a new ServicesStorageProbeVerdict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *ServicesStorageProbeVerdict) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *ServicesStorageProbeVerdict) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *ServicesStorageProbeVerdict) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *ServicesStorageProbeVerdict) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetBucketVerdict

`func (o *ServicesStorageProbeVerdict) GetBucketVerdict() ServicesStorageBucketVerdict`

GetBucketVerdict returns the BucketVerdict field if non-nil, zero value otherwise.

### GetBucketVerdictOk

`func (o *ServicesStorageProbeVerdict) GetBucketVerdictOk() (*ServicesStorageBucketVerdict, bool)`

GetBucketVerdictOk returns a tuple with the BucketVerdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketVerdict

`func (o *ServicesStorageProbeVerdict) SetBucketVerdict(v ServicesStorageBucketVerdict)`

SetBucketVerdict sets BucketVerdict field to given value.

### HasBucketVerdict

`func (o *ServicesStorageProbeVerdict) HasBucketVerdict() bool`

HasBucketVerdict returns a boolean if a field has been set.

### GetListingPermitted

`func (o *ServicesStorageProbeVerdict) GetListingPermitted() bool`

GetListingPermitted returns the ListingPermitted field if non-nil, zero value otherwise.

### GetListingPermittedOk

`func (o *ServicesStorageProbeVerdict) GetListingPermittedOk() (*bool, bool)`

GetListingPermittedOk returns a tuple with the ListingPermitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingPermitted

`func (o *ServicesStorageProbeVerdict) SetListingPermitted(v bool)`

SetListingPermitted sets ListingPermitted field to given value.

### HasListingPermitted

`func (o *ServicesStorageProbeVerdict) HasListingPermitted() bool`

HasListingPermitted returns a boolean if a field has been set.

### GetMessage

`func (o *ServicesStorageProbeVerdict) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServicesStorageProbeVerdict) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServicesStorageProbeVerdict) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServicesStorageProbeVerdict) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMissingPermission

`func (o *ServicesStorageProbeVerdict) GetMissingPermission() string`

GetMissingPermission returns the MissingPermission field if non-nil, zero value otherwise.

### GetMissingPermissionOk

`func (o *ServicesStorageProbeVerdict) GetMissingPermissionOk() (*string, bool)`

GetMissingPermissionOk returns a tuple with the MissingPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingPermission

`func (o *ServicesStorageProbeVerdict) SetMissingPermission(v string)`

SetMissingPermission sets MissingPermission field to given value.

### HasMissingPermission

`func (o *ServicesStorageProbeVerdict) HasMissingPermission() bool`

HasMissingPermission returns a boolean if a field has been set.

### GetPermissionDoc

`func (o *ServicesStorageProbeVerdict) GetPermissionDoc() string`

GetPermissionDoc returns the PermissionDoc field if non-nil, zero value otherwise.

### GetPermissionDocOk

`func (o *ServicesStorageProbeVerdict) GetPermissionDocOk() (*string, bool)`

GetPermissionDocOk returns a tuple with the PermissionDoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionDoc

`func (o *ServicesStorageProbeVerdict) SetPermissionDoc(v string)`

SetPermissionDoc sets PermissionDoc field to given value.

### HasPermissionDoc

`func (o *ServicesStorageProbeVerdict) HasPermissionDoc() bool`

HasPermissionDoc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


