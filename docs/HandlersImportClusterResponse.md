# HandlersImportClusterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoInstall** | Pointer to **bool** |  | [optional] 
**ClusterId** | Pointer to **int32** | ClusterID is present on the auto-install path (kubeconfig provided). The FE uses it to navigate to the cluster detail page and subscribe to progress events. Nil on the manual path (no cluster row yet). | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**HelmCommand** | Pointer to **string** |  | [optional] 
**KubectlCommand** | Pointer to **string** |  | [optional] 
**ManifestUrl** | Pointer to **string** |  | [optional] 
**RegistrationToken** | Pointer to **string** |  | [optional] 
**TrackerId** | Pointer to **string** |  | [optional] 
**WebsocketUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersImportClusterResponse

`func NewHandlersImportClusterResponse() *HandlersImportClusterResponse`

NewHandlersImportClusterResponse instantiates a new HandlersImportClusterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersImportClusterResponseWithDefaults

`func NewHandlersImportClusterResponseWithDefaults() *HandlersImportClusterResponse`

NewHandlersImportClusterResponseWithDefaults instantiates a new HandlersImportClusterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoInstall

`func (o *HandlersImportClusterResponse) GetAutoInstall() bool`

GetAutoInstall returns the AutoInstall field if non-nil, zero value otherwise.

### GetAutoInstallOk

`func (o *HandlersImportClusterResponse) GetAutoInstallOk() (*bool, bool)`

GetAutoInstallOk returns a tuple with the AutoInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoInstall

`func (o *HandlersImportClusterResponse) SetAutoInstall(v bool)`

SetAutoInstall sets AutoInstall field to given value.

### HasAutoInstall

`func (o *HandlersImportClusterResponse) HasAutoInstall() bool`

HasAutoInstall returns a boolean if a field has been set.

### GetClusterId

`func (o *HandlersImportClusterResponse) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HandlersImportClusterResponse) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HandlersImportClusterResponse) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HandlersImportClusterResponse) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *HandlersImportClusterResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *HandlersImportClusterResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *HandlersImportClusterResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *HandlersImportClusterResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetHelmCommand

`func (o *HandlersImportClusterResponse) GetHelmCommand() string`

GetHelmCommand returns the HelmCommand field if non-nil, zero value otherwise.

### GetHelmCommandOk

`func (o *HandlersImportClusterResponse) GetHelmCommandOk() (*string, bool)`

GetHelmCommandOk returns a tuple with the HelmCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmCommand

`func (o *HandlersImportClusterResponse) SetHelmCommand(v string)`

SetHelmCommand sets HelmCommand field to given value.

### HasHelmCommand

`func (o *HandlersImportClusterResponse) HasHelmCommand() bool`

HasHelmCommand returns a boolean if a field has been set.

### GetKubectlCommand

`func (o *HandlersImportClusterResponse) GetKubectlCommand() string`

GetKubectlCommand returns the KubectlCommand field if non-nil, zero value otherwise.

### GetKubectlCommandOk

`func (o *HandlersImportClusterResponse) GetKubectlCommandOk() (*string, bool)`

GetKubectlCommandOk returns a tuple with the KubectlCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubectlCommand

`func (o *HandlersImportClusterResponse) SetKubectlCommand(v string)`

SetKubectlCommand sets KubectlCommand field to given value.

### HasKubectlCommand

`func (o *HandlersImportClusterResponse) HasKubectlCommand() bool`

HasKubectlCommand returns a boolean if a field has been set.

### GetManifestUrl

`func (o *HandlersImportClusterResponse) GetManifestUrl() string`

GetManifestUrl returns the ManifestUrl field if non-nil, zero value otherwise.

### GetManifestUrlOk

`func (o *HandlersImportClusterResponse) GetManifestUrlOk() (*string, bool)`

GetManifestUrlOk returns a tuple with the ManifestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestUrl

`func (o *HandlersImportClusterResponse) SetManifestUrl(v string)`

SetManifestUrl sets ManifestUrl field to given value.

### HasManifestUrl

`func (o *HandlersImportClusterResponse) HasManifestUrl() bool`

HasManifestUrl returns a boolean if a field has been set.

### GetRegistrationToken

`func (o *HandlersImportClusterResponse) GetRegistrationToken() string`

GetRegistrationToken returns the RegistrationToken field if non-nil, zero value otherwise.

### GetRegistrationTokenOk

`func (o *HandlersImportClusterResponse) GetRegistrationTokenOk() (*string, bool)`

GetRegistrationTokenOk returns a tuple with the RegistrationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationToken

`func (o *HandlersImportClusterResponse) SetRegistrationToken(v string)`

SetRegistrationToken sets RegistrationToken field to given value.

### HasRegistrationToken

`func (o *HandlersImportClusterResponse) HasRegistrationToken() bool`

HasRegistrationToken returns a boolean if a field has been set.

### GetTrackerId

`func (o *HandlersImportClusterResponse) GetTrackerId() string`

GetTrackerId returns the TrackerId field if non-nil, zero value otherwise.

### GetTrackerIdOk

`func (o *HandlersImportClusterResponse) GetTrackerIdOk() (*string, bool)`

GetTrackerIdOk returns a tuple with the TrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackerId

`func (o *HandlersImportClusterResponse) SetTrackerId(v string)`

SetTrackerId sets TrackerId field to given value.

### HasTrackerId

`func (o *HandlersImportClusterResponse) HasTrackerId() bool`

HasTrackerId returns a boolean if a field has been set.

### GetWebsocketUrl

`func (o *HandlersImportClusterResponse) GetWebsocketUrl() string`

GetWebsocketUrl returns the WebsocketUrl field if non-nil, zero value otherwise.

### GetWebsocketUrlOk

`func (o *HandlersImportClusterResponse) GetWebsocketUrlOk() (*string, bool)`

GetWebsocketUrlOk returns a tuple with the WebsocketUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketUrl

`func (o *HandlersImportClusterResponse) SetWebsocketUrl(v string)`

SetWebsocketUrl sets WebsocketUrl field to given value.

### HasWebsocketUrl

`func (o *HandlersImportClusterResponse) HasWebsocketUrl() bool`

HasWebsocketUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


