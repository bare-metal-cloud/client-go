# HandlersRevokeDelegationMandateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryUnix** | Pointer to **int32** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 
**OpId** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersRevokeDelegationMandateRequest

`func NewHandlersRevokeDelegationMandateRequest() *HandlersRevokeDelegationMandateRequest`

NewHandlersRevokeDelegationMandateRequest instantiates a new HandlersRevokeDelegationMandateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRevokeDelegationMandateRequestWithDefaults

`func NewHandlersRevokeDelegationMandateRequestWithDefaults() *HandlersRevokeDelegationMandateRequest`

NewHandlersRevokeDelegationMandateRequestWithDefaults instantiates a new HandlersRevokeDelegationMandateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryUnix

`func (o *HandlersRevokeDelegationMandateRequest) GetExpiryUnix() int32`

GetExpiryUnix returns the ExpiryUnix field if non-nil, zero value otherwise.

### GetExpiryUnixOk

`func (o *HandlersRevokeDelegationMandateRequest) GetExpiryUnixOk() (*int32, bool)`

GetExpiryUnixOk returns a tuple with the ExpiryUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryUnix

`func (o *HandlersRevokeDelegationMandateRequest) SetExpiryUnix(v int32)`

SetExpiryUnix sets ExpiryUnix field to given value.

### HasExpiryUnix

`func (o *HandlersRevokeDelegationMandateRequest) HasExpiryUnix() bool`

HasExpiryUnix returns a boolean if a field has been set.

### GetNonce

`func (o *HandlersRevokeDelegationMandateRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *HandlersRevokeDelegationMandateRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *HandlersRevokeDelegationMandateRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *HandlersRevokeDelegationMandateRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetOpId

`func (o *HandlersRevokeDelegationMandateRequest) GetOpId() string`

GetOpId returns the OpId field if non-nil, zero value otherwise.

### GetOpIdOk

`func (o *HandlersRevokeDelegationMandateRequest) GetOpIdOk() (*string, bool)`

GetOpIdOk returns a tuple with the OpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpId

`func (o *HandlersRevokeDelegationMandateRequest) SetOpId(v string)`

SetOpId sets OpId field to given value.

### HasOpId

`func (o *HandlersRevokeDelegationMandateRequest) HasOpId() bool`

HasOpId returns a boolean if a field has been set.

### GetSignature

`func (o *HandlersRevokeDelegationMandateRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *HandlersRevokeDelegationMandateRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *HandlersRevokeDelegationMandateRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *HandlersRevokeDelegationMandateRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


