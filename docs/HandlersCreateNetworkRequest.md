# HandlersCreateNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ModelsNetworkConfig**](ModelsNetworkConfig.md) |  | [optional] 
**DraftName** | Pointer to **string** |  | [optional] 
**Purpose** | Pointer to **string** | Purpose + DraftName carry the durable cluster-create draft intent (story 124.17 AC7). The wizard sends purpose&#x3D;\&quot;cluster-create\&quot; plus the cluster name it is heading toward; ordinary network creates omit both. Only \&quot;cluster-create\&quot; is accepted today — anything else is stripped so the column cannot be used as a free-text channel. | [optional] 

## Methods

### NewHandlersCreateNetworkRequest

`func NewHandlersCreateNetworkRequest() *HandlersCreateNetworkRequest`

NewHandlersCreateNetworkRequest instantiates a new HandlersCreateNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateNetworkRequestWithDefaults

`func NewHandlersCreateNetworkRequestWithDefaults() *HandlersCreateNetworkRequest`

NewHandlersCreateNetworkRequestWithDefaults instantiates a new HandlersCreateNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *HandlersCreateNetworkRequest) GetConfig() ModelsNetworkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *HandlersCreateNetworkRequest) GetConfigOk() (*ModelsNetworkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *HandlersCreateNetworkRequest) SetConfig(v ModelsNetworkConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *HandlersCreateNetworkRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDraftName

`func (o *HandlersCreateNetworkRequest) GetDraftName() string`

GetDraftName returns the DraftName field if non-nil, zero value otherwise.

### GetDraftNameOk

`func (o *HandlersCreateNetworkRequest) GetDraftNameOk() (*string, bool)`

GetDraftNameOk returns a tuple with the DraftName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftName

`func (o *HandlersCreateNetworkRequest) SetDraftName(v string)`

SetDraftName sets DraftName field to given value.

### HasDraftName

`func (o *HandlersCreateNetworkRequest) HasDraftName() bool`

HasDraftName returns a boolean if a field has been set.

### GetPurpose

`func (o *HandlersCreateNetworkRequest) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *HandlersCreateNetworkRequest) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *HandlersCreateNetworkRequest) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *HandlersCreateNetworkRequest) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


