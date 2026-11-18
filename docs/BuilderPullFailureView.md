# BuilderPullFailureView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ask** | Pointer to [**BuilderPullCredentialsAsk**](BuilderPullCredentialsAsk.md) | Ask is the attach-or-connect state (the creds-invalid class&#39;s single CTA; the retryable classes&#39; CTA is the retry). | [optional] 
**Code** | Pointer to **string** | Code is one of the publish chain&#39;s Refuse* constants, verbatim. | [optional] 
**Likely** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Registry** | Pointer to **string** | Registry names the registry the refusal came from. | [optional] 
**Retryable** | Pointer to **bool** | Retryable is the refusal&#39;s A20 class (creds-invalid never retries — the first-class state; the other two re-check). | [optional] 
**Saw** | Pointer to **string** |  | [optional] 
**ToTry** | Pointer to **string** |  | [optional] 

## Methods

### NewBuilderPullFailureView

`func NewBuilderPullFailureView() *BuilderPullFailureView`

NewBuilderPullFailureView instantiates a new BuilderPullFailureView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuilderPullFailureViewWithDefaults

`func NewBuilderPullFailureViewWithDefaults() *BuilderPullFailureView`

NewBuilderPullFailureViewWithDefaults instantiates a new BuilderPullFailureView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsk

`func (o *BuilderPullFailureView) GetAsk() BuilderPullCredentialsAsk`

GetAsk returns the Ask field if non-nil, zero value otherwise.

### GetAskOk

`func (o *BuilderPullFailureView) GetAskOk() (*BuilderPullCredentialsAsk, bool)`

GetAskOk returns a tuple with the Ask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsk

`func (o *BuilderPullFailureView) SetAsk(v BuilderPullCredentialsAsk)`

SetAsk sets Ask field to given value.

### HasAsk

`func (o *BuilderPullFailureView) HasAsk() bool`

HasAsk returns a boolean if a field has been set.

### GetCode

`func (o *BuilderPullFailureView) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BuilderPullFailureView) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BuilderPullFailureView) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BuilderPullFailureView) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLikely

`func (o *BuilderPullFailureView) GetLikely() string`

GetLikely returns the Likely field if non-nil, zero value otherwise.

### GetLikelyOk

`func (o *BuilderPullFailureView) GetLikelyOk() (*string, bool)`

GetLikelyOk returns a tuple with the Likely field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikely

`func (o *BuilderPullFailureView) SetLikely(v string)`

SetLikely sets Likely field to given value.

### HasLikely

`func (o *BuilderPullFailureView) HasLikely() bool`

HasLikely returns a boolean if a field has been set.

### GetMessage

`func (o *BuilderPullFailureView) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BuilderPullFailureView) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BuilderPullFailureView) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BuilderPullFailureView) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRegistry

`func (o *BuilderPullFailureView) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *BuilderPullFailureView) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *BuilderPullFailureView) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *BuilderPullFailureView) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRetryable

`func (o *BuilderPullFailureView) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *BuilderPullFailureView) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *BuilderPullFailureView) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *BuilderPullFailureView) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetSaw

`func (o *BuilderPullFailureView) GetSaw() string`

GetSaw returns the Saw field if non-nil, zero value otherwise.

### GetSawOk

`func (o *BuilderPullFailureView) GetSawOk() (*string, bool)`

GetSawOk returns a tuple with the Saw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaw

`func (o *BuilderPullFailureView) SetSaw(v string)`

SetSaw sets Saw field to given value.

### HasSaw

`func (o *BuilderPullFailureView) HasSaw() bool`

HasSaw returns a boolean if a field has been set.

### GetToTry

`func (o *BuilderPullFailureView) GetToTry() string`

GetToTry returns the ToTry field if non-nil, zero value otherwise.

### GetToTryOk

`func (o *BuilderPullFailureView) GetToTryOk() (*string, bool)`

GetToTryOk returns a tuple with the ToTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTry

`func (o *BuilderPullFailureView) SetToTry(v string)`

SetToTry sets ToTry field to given value.

### HasToTry

`func (o *BuilderPullFailureView) HasToTry() bool`

HasToTry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


