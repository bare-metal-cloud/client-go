# GitopsInstallStateSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** |  | [optional] 
**Connection** | Pointer to [**GitopsConnectionSummary**](GitopsConnectionSummary.md) |  | [optional] 
**LastDeploy** | Pointer to [**GitopsLastDeploySummary**](GitopsLastDeploySummary.md) |  | [optional] 
**Park** | Pointer to [**GitopsUpgradeParkBadge**](GitopsUpgradeParkBadge.md) | Park is the honest-state badge (issue #2709 §2.7, additive and omitempty): the install&#39;s active parked upgrade — a migrating CI push holding for a panel confirm. Present on EVERY state this section renders while the park is active; absent when nothing is parked. | [optional] 
**PendingExpiresAt** | Pointer to **string** | PendingExpiresAt states the pending TTL (Law 13: the deadline is stated on the surface). Set only in the pending state. | [optional] 
**RegistrationRotate** | Pointer to **bool** | RegistrationRotate is the render hint for the registration-token rotate verb (issue #2501, additive and omitempty): set true ONLY on the enabled-registration path. It derives from the registration row&#39;s existence and state — NEVER from an active token existing (the dead-credential install is exactly the install that must show the verb). A render hint only; the verb&#39;s own guards remain the authority. | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**WorkflowPath** | Pointer to **string** |  | [optional] 

## Methods

### NewGitopsInstallStateSection

`func NewGitopsInstallStateSection() *GitopsInstallStateSection`

NewGitopsInstallStateSection instantiates a new GitopsInstallStateSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitopsInstallStateSectionWithDefaults

`func NewGitopsInstallStateSectionWithDefaults() *GitopsInstallStateSection`

NewGitopsInstallStateSectionWithDefaults instantiates a new GitopsInstallStateSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *GitopsInstallStateSection) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GitopsInstallStateSection) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GitopsInstallStateSection) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GitopsInstallStateSection) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetConnection

`func (o *GitopsInstallStateSection) GetConnection() GitopsConnectionSummary`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *GitopsInstallStateSection) GetConnectionOk() (*GitopsConnectionSummary, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *GitopsInstallStateSection) SetConnection(v GitopsConnectionSummary)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *GitopsInstallStateSection) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetLastDeploy

`func (o *GitopsInstallStateSection) GetLastDeploy() GitopsLastDeploySummary`

GetLastDeploy returns the LastDeploy field if non-nil, zero value otherwise.

### GetLastDeployOk

`func (o *GitopsInstallStateSection) GetLastDeployOk() (*GitopsLastDeploySummary, bool)`

GetLastDeployOk returns a tuple with the LastDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploy

`func (o *GitopsInstallStateSection) SetLastDeploy(v GitopsLastDeploySummary)`

SetLastDeploy sets LastDeploy field to given value.

### HasLastDeploy

`func (o *GitopsInstallStateSection) HasLastDeploy() bool`

HasLastDeploy returns a boolean if a field has been set.

### GetPark

`func (o *GitopsInstallStateSection) GetPark() GitopsUpgradeParkBadge`

GetPark returns the Park field if non-nil, zero value otherwise.

### GetParkOk

`func (o *GitopsInstallStateSection) GetParkOk() (*GitopsUpgradeParkBadge, bool)`

GetParkOk returns a tuple with the Park field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPark

`func (o *GitopsInstallStateSection) SetPark(v GitopsUpgradeParkBadge)`

SetPark sets Park field to given value.

### HasPark

`func (o *GitopsInstallStateSection) HasPark() bool`

HasPark returns a boolean if a field has been set.

### GetPendingExpiresAt

`func (o *GitopsInstallStateSection) GetPendingExpiresAt() string`

GetPendingExpiresAt returns the PendingExpiresAt field if non-nil, zero value otherwise.

### GetPendingExpiresAtOk

`func (o *GitopsInstallStateSection) GetPendingExpiresAtOk() (*string, bool)`

GetPendingExpiresAtOk returns a tuple with the PendingExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingExpiresAt

`func (o *GitopsInstallStateSection) SetPendingExpiresAt(v string)`

SetPendingExpiresAt sets PendingExpiresAt field to given value.

### HasPendingExpiresAt

`func (o *GitopsInstallStateSection) HasPendingExpiresAt() bool`

HasPendingExpiresAt returns a boolean if a field has been set.

### GetRegistrationRotate

`func (o *GitopsInstallStateSection) GetRegistrationRotate() bool`

GetRegistrationRotate returns the RegistrationRotate field if non-nil, zero value otherwise.

### GetRegistrationRotateOk

`func (o *GitopsInstallStateSection) GetRegistrationRotateOk() (*bool, bool)`

GetRegistrationRotateOk returns a tuple with the RegistrationRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationRotate

`func (o *GitopsInstallStateSection) SetRegistrationRotate(v bool)`

SetRegistrationRotate sets RegistrationRotate field to given value.

### HasRegistrationRotate

`func (o *GitopsInstallStateSection) HasRegistrationRotate() bool`

HasRegistrationRotate returns a boolean if a field has been set.

### GetRepo

`func (o *GitopsInstallStateSection) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GitopsInstallStateSection) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GitopsInstallStateSection) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GitopsInstallStateSection) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetState

`func (o *GitopsInstallStateSection) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GitopsInstallStateSection) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GitopsInstallStateSection) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GitopsInstallStateSection) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkflowPath

`func (o *GitopsInstallStateSection) GetWorkflowPath() string`

GetWorkflowPath returns the WorkflowPath field if non-nil, zero value otherwise.

### GetWorkflowPathOk

`func (o *GitopsInstallStateSection) GetWorkflowPathOk() (*string, bool)`

GetWorkflowPathOk returns a tuple with the WorkflowPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowPath

`func (o *GitopsInstallStateSection) SetWorkflowPath(v string)`

SetWorkflowPath sets WorkflowPath field to given value.

### HasWorkflowPath

`func (o *GitopsInstallStateSection) HasWorkflowPath() bool`

HasWorkflowPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


