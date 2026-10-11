# HandlersMFASetupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QrCodeUrl** | Pointer to **string** | otpauth://totp/... (put in an &lt;img src&#x3D;\&quot;/qr?data&#x3D;...\&quot;&gt; on the frontend) | [optional] 
**Secret** | Pointer to **string** | base32 manual-entry fallback | [optional] 

## Methods

### NewHandlersMFASetupResponse

`func NewHandlersMFASetupResponse() *HandlersMFASetupResponse`

NewHandlersMFASetupResponse instantiates a new HandlersMFASetupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersMFASetupResponseWithDefaults

`func NewHandlersMFASetupResponseWithDefaults() *HandlersMFASetupResponse`

NewHandlersMFASetupResponseWithDefaults instantiates a new HandlersMFASetupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQrCodeUrl

`func (o *HandlersMFASetupResponse) GetQrCodeUrl() string`

GetQrCodeUrl returns the QrCodeUrl field if non-nil, zero value otherwise.

### GetQrCodeUrlOk

`func (o *HandlersMFASetupResponse) GetQrCodeUrlOk() (*string, bool)`

GetQrCodeUrlOk returns a tuple with the QrCodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeUrl

`func (o *HandlersMFASetupResponse) SetQrCodeUrl(v string)`

SetQrCodeUrl sets QrCodeUrl field to given value.

### HasQrCodeUrl

`func (o *HandlersMFASetupResponse) HasQrCodeUrl() bool`

HasQrCodeUrl returns a boolean if a field has been set.

### GetSecret

`func (o *HandlersMFASetupResponse) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *HandlersMFASetupResponse) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *HandlersMFASetupResponse) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *HandlersMFASetupResponse) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


