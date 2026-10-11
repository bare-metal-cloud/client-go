# ServicesDatabaseConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStrings** | Pointer to **map[string]string** | Connection strings | [optional] 
**DatabaseName** | Pointer to **string** | Basic connection info | [optional] 
**HaEnabled** | Pointer to **bool** | HA info (if enabled) | [optional] 
**Host** | Pointer to **string** | Internal K8s service DNS | [optional] 
**Namespace** | Pointer to **string** | Metadata | [optional] 
**Password** | Pointer to **string** | Masked: \&quot;••••••••\&quot; | [optional] 
**Port** | Pointer to **int32** | 5432, 27017, 3306 | [optional] 
**PrimaryHost** | Pointer to **string** |  | [optional] 
**ReplicaHosts** | Pointer to **[]string** |  | [optional] 
**SecretName** | Pointer to **string** |  | [optional] 
**ServiceEndpoint** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewServicesDatabaseConnectionInfo

`func NewServicesDatabaseConnectionInfo() *ServicesDatabaseConnectionInfo`

NewServicesDatabaseConnectionInfo instantiates a new ServicesDatabaseConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesDatabaseConnectionInfoWithDefaults

`func NewServicesDatabaseConnectionInfoWithDefaults() *ServicesDatabaseConnectionInfo`

NewServicesDatabaseConnectionInfoWithDefaults instantiates a new ServicesDatabaseConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStrings

`func (o *ServicesDatabaseConnectionInfo) GetConnectionStrings() map[string]string`

GetConnectionStrings returns the ConnectionStrings field if non-nil, zero value otherwise.

### GetConnectionStringsOk

`func (o *ServicesDatabaseConnectionInfo) GetConnectionStringsOk() (*map[string]string, bool)`

GetConnectionStringsOk returns a tuple with the ConnectionStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStrings

`func (o *ServicesDatabaseConnectionInfo) SetConnectionStrings(v map[string]string)`

SetConnectionStrings sets ConnectionStrings field to given value.

### HasConnectionStrings

`func (o *ServicesDatabaseConnectionInfo) HasConnectionStrings() bool`

HasConnectionStrings returns a boolean if a field has been set.

### GetDatabaseName

`func (o *ServicesDatabaseConnectionInfo) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ServicesDatabaseConnectionInfo) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ServicesDatabaseConnectionInfo) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *ServicesDatabaseConnectionInfo) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetHaEnabled

`func (o *ServicesDatabaseConnectionInfo) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *ServicesDatabaseConnectionInfo) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *ServicesDatabaseConnectionInfo) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *ServicesDatabaseConnectionInfo) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetHost

`func (o *ServicesDatabaseConnectionInfo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ServicesDatabaseConnectionInfo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ServicesDatabaseConnectionInfo) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ServicesDatabaseConnectionInfo) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNamespace

`func (o *ServicesDatabaseConnectionInfo) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ServicesDatabaseConnectionInfo) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ServicesDatabaseConnectionInfo) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ServicesDatabaseConnectionInfo) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPassword

`func (o *ServicesDatabaseConnectionInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServicesDatabaseConnectionInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServicesDatabaseConnectionInfo) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServicesDatabaseConnectionInfo) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *ServicesDatabaseConnectionInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ServicesDatabaseConnectionInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ServicesDatabaseConnectionInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ServicesDatabaseConnectionInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrimaryHost

`func (o *ServicesDatabaseConnectionInfo) GetPrimaryHost() string`

GetPrimaryHost returns the PrimaryHost field if non-nil, zero value otherwise.

### GetPrimaryHostOk

`func (o *ServicesDatabaseConnectionInfo) GetPrimaryHostOk() (*string, bool)`

GetPrimaryHostOk returns a tuple with the PrimaryHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryHost

`func (o *ServicesDatabaseConnectionInfo) SetPrimaryHost(v string)`

SetPrimaryHost sets PrimaryHost field to given value.

### HasPrimaryHost

`func (o *ServicesDatabaseConnectionInfo) HasPrimaryHost() bool`

HasPrimaryHost returns a boolean if a field has been set.

### GetReplicaHosts

`func (o *ServicesDatabaseConnectionInfo) GetReplicaHosts() []string`

GetReplicaHosts returns the ReplicaHosts field if non-nil, zero value otherwise.

### GetReplicaHostsOk

`func (o *ServicesDatabaseConnectionInfo) GetReplicaHostsOk() (*[]string, bool)`

GetReplicaHostsOk returns a tuple with the ReplicaHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaHosts

`func (o *ServicesDatabaseConnectionInfo) SetReplicaHosts(v []string)`

SetReplicaHosts sets ReplicaHosts field to given value.

### HasReplicaHosts

`func (o *ServicesDatabaseConnectionInfo) HasReplicaHosts() bool`

HasReplicaHosts returns a boolean if a field has been set.

### GetSecretName

`func (o *ServicesDatabaseConnectionInfo) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *ServicesDatabaseConnectionInfo) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *ServicesDatabaseConnectionInfo) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.

### HasSecretName

`func (o *ServicesDatabaseConnectionInfo) HasSecretName() bool`

HasSecretName returns a boolean if a field has been set.

### GetServiceEndpoint

`func (o *ServicesDatabaseConnectionInfo) GetServiceEndpoint() string`

GetServiceEndpoint returns the ServiceEndpoint field if non-nil, zero value otherwise.

### GetServiceEndpointOk

`func (o *ServicesDatabaseConnectionInfo) GetServiceEndpointOk() (*string, bool)`

GetServiceEndpointOk returns a tuple with the ServiceEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndpoint

`func (o *ServicesDatabaseConnectionInfo) SetServiceEndpoint(v string)`

SetServiceEndpoint sets ServiceEndpoint field to given value.

### HasServiceEndpoint

`func (o *ServicesDatabaseConnectionInfo) HasServiceEndpoint() bool`

HasServiceEndpoint returns a boolean if a field has been set.

### GetUsername

`func (o *ServicesDatabaseConnectionInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServicesDatabaseConnectionInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServicesDatabaseConnectionInfo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServicesDatabaseConnectionInfo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


