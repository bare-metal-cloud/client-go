# HandlersVerifySSHRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BastionHost** | Pointer to **string** |  | [optional] 
**BastionKey** | Pointer to **string** |  | [optional] 
**BastionPort** | Pointer to **int32** |  | [optional] 
**BastionUser** | Pointer to **string** |  | [optional] 
**Host** | **string** |  | 
**KeyPassword** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Port** | **int32** |  | 
**SshKey** | Pointer to **string** |  | [optional] 
**Username** | **string** |  | 

## Methods

### NewHandlersVerifySSHRequest

`func NewHandlersVerifySSHRequest(host string, port int32, username string, ) *HandlersVerifySSHRequest`

NewHandlersVerifySSHRequest instantiates a new HandlersVerifySSHRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersVerifySSHRequestWithDefaults

`func NewHandlersVerifySSHRequestWithDefaults() *HandlersVerifySSHRequest`

NewHandlersVerifySSHRequestWithDefaults instantiates a new HandlersVerifySSHRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBastionHost

`func (o *HandlersVerifySSHRequest) GetBastionHost() string`

GetBastionHost returns the BastionHost field if non-nil, zero value otherwise.

### GetBastionHostOk

`func (o *HandlersVerifySSHRequest) GetBastionHostOk() (*string, bool)`

GetBastionHostOk returns a tuple with the BastionHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionHost

`func (o *HandlersVerifySSHRequest) SetBastionHost(v string)`

SetBastionHost sets BastionHost field to given value.

### HasBastionHost

`func (o *HandlersVerifySSHRequest) HasBastionHost() bool`

HasBastionHost returns a boolean if a field has been set.

### GetBastionKey

`func (o *HandlersVerifySSHRequest) GetBastionKey() string`

GetBastionKey returns the BastionKey field if non-nil, zero value otherwise.

### GetBastionKeyOk

`func (o *HandlersVerifySSHRequest) GetBastionKeyOk() (*string, bool)`

GetBastionKeyOk returns a tuple with the BastionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionKey

`func (o *HandlersVerifySSHRequest) SetBastionKey(v string)`

SetBastionKey sets BastionKey field to given value.

### HasBastionKey

`func (o *HandlersVerifySSHRequest) HasBastionKey() bool`

HasBastionKey returns a boolean if a field has been set.

### GetBastionPort

`func (o *HandlersVerifySSHRequest) GetBastionPort() int32`

GetBastionPort returns the BastionPort field if non-nil, zero value otherwise.

### GetBastionPortOk

`func (o *HandlersVerifySSHRequest) GetBastionPortOk() (*int32, bool)`

GetBastionPortOk returns a tuple with the BastionPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionPort

`func (o *HandlersVerifySSHRequest) SetBastionPort(v int32)`

SetBastionPort sets BastionPort field to given value.

### HasBastionPort

`func (o *HandlersVerifySSHRequest) HasBastionPort() bool`

HasBastionPort returns a boolean if a field has been set.

### GetBastionUser

`func (o *HandlersVerifySSHRequest) GetBastionUser() string`

GetBastionUser returns the BastionUser field if non-nil, zero value otherwise.

### GetBastionUserOk

`func (o *HandlersVerifySSHRequest) GetBastionUserOk() (*string, bool)`

GetBastionUserOk returns a tuple with the BastionUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionUser

`func (o *HandlersVerifySSHRequest) SetBastionUser(v string)`

SetBastionUser sets BastionUser field to given value.

### HasBastionUser

`func (o *HandlersVerifySSHRequest) HasBastionUser() bool`

HasBastionUser returns a boolean if a field has been set.

### GetHost

`func (o *HandlersVerifySSHRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HandlersVerifySSHRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HandlersVerifySSHRequest) SetHost(v string)`

SetHost sets Host field to given value.


### GetKeyPassword

`func (o *HandlersVerifySSHRequest) GetKeyPassword() string`

GetKeyPassword returns the KeyPassword field if non-nil, zero value otherwise.

### GetKeyPasswordOk

`func (o *HandlersVerifySSHRequest) GetKeyPasswordOk() (*string, bool)`

GetKeyPasswordOk returns a tuple with the KeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPassword

`func (o *HandlersVerifySSHRequest) SetKeyPassword(v string)`

SetKeyPassword sets KeyPassword field to given value.

### HasKeyPassword

`func (o *HandlersVerifySSHRequest) HasKeyPassword() bool`

HasKeyPassword returns a boolean if a field has been set.

### GetPassword

`func (o *HandlersVerifySSHRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HandlersVerifySSHRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HandlersVerifySSHRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HandlersVerifySSHRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *HandlersVerifySSHRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HandlersVerifySSHRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HandlersVerifySSHRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSshKey

`func (o *HandlersVerifySSHRequest) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *HandlersVerifySSHRequest) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *HandlersVerifySSHRequest) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *HandlersVerifySSHRequest) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### GetUsername

`func (o *HandlersVerifySSHRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HandlersVerifySSHRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HandlersVerifySSHRequest) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


