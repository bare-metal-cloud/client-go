# HandlersUpdateBackupPlanBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MigrateExistingBackups** | Pointer to **bool** | MigrateExistingBackups turns a destination switch into a migration. | [optional] 
**MigrationCleanup** | Pointer to **string** | MigrationCleanup is the old-destination deletion consent (\&quot;keep\&quot; | \&quot;delete_after_cutover\&quot;); never honored without the migrate choice. | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**RetentionDays** | Pointer to **int32** |  | [optional] 
**S3ConnectionId** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersUpdateBackupPlanBody

`func NewHandlersUpdateBackupPlanBody() *HandlersUpdateBackupPlanBody`

NewHandlersUpdateBackupPlanBody instantiates a new HandlersUpdateBackupPlanBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersUpdateBackupPlanBodyWithDefaults

`func NewHandlersUpdateBackupPlanBodyWithDefaults() *HandlersUpdateBackupPlanBody`

NewHandlersUpdateBackupPlanBodyWithDefaults instantiates a new HandlersUpdateBackupPlanBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigrateExistingBackups

`func (o *HandlersUpdateBackupPlanBody) GetMigrateExistingBackups() bool`

GetMigrateExistingBackups returns the MigrateExistingBackups field if non-nil, zero value otherwise.

### GetMigrateExistingBackupsOk

`func (o *HandlersUpdateBackupPlanBody) GetMigrateExistingBackupsOk() (*bool, bool)`

GetMigrateExistingBackupsOk returns a tuple with the MigrateExistingBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateExistingBackups

`func (o *HandlersUpdateBackupPlanBody) SetMigrateExistingBackups(v bool)`

SetMigrateExistingBackups sets MigrateExistingBackups field to given value.

### HasMigrateExistingBackups

`func (o *HandlersUpdateBackupPlanBody) HasMigrateExistingBackups() bool`

HasMigrateExistingBackups returns a boolean if a field has been set.

### GetMigrationCleanup

`func (o *HandlersUpdateBackupPlanBody) GetMigrationCleanup() string`

GetMigrationCleanup returns the MigrationCleanup field if non-nil, zero value otherwise.

### GetMigrationCleanupOk

`func (o *HandlersUpdateBackupPlanBody) GetMigrationCleanupOk() (*string, bool)`

GetMigrationCleanupOk returns a tuple with the MigrationCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationCleanup

`func (o *HandlersUpdateBackupPlanBody) SetMigrationCleanup(v string)`

SetMigrationCleanup sets MigrationCleanup field to given value.

### HasMigrationCleanup

`func (o *HandlersUpdateBackupPlanBody) HasMigrationCleanup() bool`

HasMigrationCleanup returns a boolean if a field has been set.

### GetPlan

`func (o *HandlersUpdateBackupPlanBody) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *HandlersUpdateBackupPlanBody) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *HandlersUpdateBackupPlanBody) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *HandlersUpdateBackupPlanBody) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRetentionDays

`func (o *HandlersUpdateBackupPlanBody) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *HandlersUpdateBackupPlanBody) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *HandlersUpdateBackupPlanBody) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *HandlersUpdateBackupPlanBody) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetS3ConnectionId

`func (o *HandlersUpdateBackupPlanBody) GetS3ConnectionId() int32`

GetS3ConnectionId returns the S3ConnectionId field if non-nil, zero value otherwise.

### GetS3ConnectionIdOk

`func (o *HandlersUpdateBackupPlanBody) GetS3ConnectionIdOk() (*int32, bool)`

GetS3ConnectionIdOk returns a tuple with the S3ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ConnectionId

`func (o *HandlersUpdateBackupPlanBody) SetS3ConnectionId(v int32)`

SetS3ConnectionId sets S3ConnectionId field to given value.

### HasS3ConnectionId

`func (o *HandlersUpdateBackupPlanBody) HasS3ConnectionId() bool`

HasS3ConnectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


