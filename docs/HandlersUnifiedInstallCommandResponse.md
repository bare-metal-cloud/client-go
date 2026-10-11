# HandlersUnifiedInstallCommandResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | Full installation command (kubectl for kubernetes) | [optional] 
**ExpiresAt** | Pointer to **string** | Token expiration timestamp | [optional] 
**HelmCommand** | Pointer to **string** | Helm install command (for kubernetes) | [optional] 
**InstallType** | Pointer to **string** | \&quot;kubernetes\&quot; or \&quot;bare-metal\&quot; | [optional] 
**ManifestUrl** | Pointer to **string** | Only for kubernetes | [optional] 
**Token** | Pointer to **string** | Short token (12 chars) | [optional] 

## Methods

### NewHandlersUnifiedInstallCommandResponse

`func NewHandlersUnifiedInstallCommandResponse() *HandlersUnifiedInstallCommandResponse`

NewHandlersUnifiedInstallCommandResponse instantiates a new HandlersUnifiedInstallCommandResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersUnifiedInstallCommandResponseWithDefaults

`func NewHandlersUnifiedInstallCommandResponseWithDefaults() *HandlersUnifiedInstallCommandResponse`

NewHandlersUnifiedInstallCommandResponseWithDefaults instantiates a new HandlersUnifiedInstallCommandResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *HandlersUnifiedInstallCommandResponse) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *HandlersUnifiedInstallCommandResponse) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *HandlersUnifiedInstallCommandResponse) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *HandlersUnifiedInstallCommandResponse) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetExpiresAt

`func (o *HandlersUnifiedInstallCommandResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *HandlersUnifiedInstallCommandResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *HandlersUnifiedInstallCommandResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *HandlersUnifiedInstallCommandResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetHelmCommand

`func (o *HandlersUnifiedInstallCommandResponse) GetHelmCommand() string`

GetHelmCommand returns the HelmCommand field if non-nil, zero value otherwise.

### GetHelmCommandOk

`func (o *HandlersUnifiedInstallCommandResponse) GetHelmCommandOk() (*string, bool)`

GetHelmCommandOk returns a tuple with the HelmCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmCommand

`func (o *HandlersUnifiedInstallCommandResponse) SetHelmCommand(v string)`

SetHelmCommand sets HelmCommand field to given value.

### HasHelmCommand

`func (o *HandlersUnifiedInstallCommandResponse) HasHelmCommand() bool`

HasHelmCommand returns a boolean if a field has been set.

### GetInstallType

`func (o *HandlersUnifiedInstallCommandResponse) GetInstallType() string`

GetInstallType returns the InstallType field if non-nil, zero value otherwise.

### GetInstallTypeOk

`func (o *HandlersUnifiedInstallCommandResponse) GetInstallTypeOk() (*string, bool)`

GetInstallTypeOk returns a tuple with the InstallType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallType

`func (o *HandlersUnifiedInstallCommandResponse) SetInstallType(v string)`

SetInstallType sets InstallType field to given value.

### HasInstallType

`func (o *HandlersUnifiedInstallCommandResponse) HasInstallType() bool`

HasInstallType returns a boolean if a field has been set.

### GetManifestUrl

`func (o *HandlersUnifiedInstallCommandResponse) GetManifestUrl() string`

GetManifestUrl returns the ManifestUrl field if non-nil, zero value otherwise.

### GetManifestUrlOk

`func (o *HandlersUnifiedInstallCommandResponse) GetManifestUrlOk() (*string, bool)`

GetManifestUrlOk returns a tuple with the ManifestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestUrl

`func (o *HandlersUnifiedInstallCommandResponse) SetManifestUrl(v string)`

SetManifestUrl sets ManifestUrl field to given value.

### HasManifestUrl

`func (o *HandlersUnifiedInstallCommandResponse) HasManifestUrl() bool`

HasManifestUrl returns a boolean if a field has been set.

### GetToken

`func (o *HandlersUnifiedInstallCommandResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *HandlersUnifiedInstallCommandResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *HandlersUnifiedInstallCommandResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *HandlersUnifiedInstallCommandResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


