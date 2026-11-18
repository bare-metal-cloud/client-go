# HandlersKeyringDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consumers** | Pointer to [**[]HandlersKeyringDetailConsumer**](HandlersKeyringDetailConsumer.md) |  | [optional] 
**Rotation** | Pointer to [**HandlersKeyringDetailRotation**](HandlersKeyringDetailRotation.md) |  | [optional] 
**Secret** | Pointer to [**HandlersKeyringDetailSecret**](HandlersKeyringDetailSecret.md) |  | [optional] 
**Versions** | Pointer to [**[]HandlersKeyringDetailVersion**](HandlersKeyringDetailVersion.md) |  | [optional] 

## Methods

### NewHandlersKeyringDetailResponse

`func NewHandlersKeyringDetailResponse() *HandlersKeyringDetailResponse`

NewHandlersKeyringDetailResponse instantiates a new HandlersKeyringDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersKeyringDetailResponseWithDefaults

`func NewHandlersKeyringDetailResponseWithDefaults() *HandlersKeyringDetailResponse`

NewHandlersKeyringDetailResponseWithDefaults instantiates a new HandlersKeyringDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumers

`func (o *HandlersKeyringDetailResponse) GetConsumers() []HandlersKeyringDetailConsumer`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *HandlersKeyringDetailResponse) GetConsumersOk() (*[]HandlersKeyringDetailConsumer, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *HandlersKeyringDetailResponse) SetConsumers(v []HandlersKeyringDetailConsumer)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *HandlersKeyringDetailResponse) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetRotation

`func (o *HandlersKeyringDetailResponse) GetRotation() HandlersKeyringDetailRotation`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *HandlersKeyringDetailResponse) GetRotationOk() (*HandlersKeyringDetailRotation, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *HandlersKeyringDetailResponse) SetRotation(v HandlersKeyringDetailRotation)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *HandlersKeyringDetailResponse) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetSecret

`func (o *HandlersKeyringDetailResponse) GetSecret() HandlersKeyringDetailSecret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *HandlersKeyringDetailResponse) GetSecretOk() (*HandlersKeyringDetailSecret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *HandlersKeyringDetailResponse) SetSecret(v HandlersKeyringDetailSecret)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *HandlersKeyringDetailResponse) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetVersions

`func (o *HandlersKeyringDetailResponse) GetVersions() []HandlersKeyringDetailVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *HandlersKeyringDetailResponse) GetVersionsOk() (*[]HandlersKeyringDetailVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *HandlersKeyringDetailResponse) SetVersions(v []HandlersKeyringDetailVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *HandlersKeyringDetailResponse) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


