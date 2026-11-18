# ModelsCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to **[]string** | ── The platform&#39;s Add-ons decision (epic app-platform story 3.4) ──  Addons records the wizard Add-ons step&#39;s checkbox set at formation: a JSON array of keys from the CLOSED set below (monitoring, public_entry, databases, log_book), in the client&#39;s selection order. Empty (&#39;[]&#39; — the column default) &#x3D; no addons: every pre-3.4 row, the import path, and API creates that omit the field. The create gate (handlers/ clusters.go + ClusterServiceImpl.ProvisionClusterBYOV) validates the keys pre-row; an unknown key never lands.  HONEST SCOPE (story 3.4): of the four keys, only &#39;monitoring&#39; has a real install path today — the formation-completion hook (the WriteK8sClusterImportResult success leg) creates the in-cluster MonitoringStack (ADR-061) and enqueues the existing Helm provision task. &#39;public_entry&#39; (the deploykit gateway addon fires inside the app-deploy workflow — no standalone formation-time ensure exists), &#39;databases&#39; (the real lane is the user-driven DatabaseConfiguration flow — an auto-deploy would invent a name, sizing, and a vault context the user never chose), and &#39;log_book&#39; (no log-aggregation surface exists) persist as DECLARED: the detail surface lists them, nothing fake-installs them. The addon set is exposed on the platform detail response (enrichClusterFields) — the add/remove-later management endpoint is future work. | [optional] 
**Adopted** | Pointer to **bool** | Adopted records the D3 ADOPT decision (story 124.3, constraint 4): creation&#39;s contamination preflight found pre-existing, RUNNING, same-version k3s on EVERY member and the caller explicitly opted in with adopt_existing. Adopted means BMC installed NO k3s on any member (SetupNode&#39;s k3s step is skipped downstream); the D9 canonical predicate bmc_installed_k3s &#x3D; (origin&#x3D;&#x3D;created &amp;&amp; !adopted) keys on this column, and adopted clusters are panel-only (D8). | [optional] 
**AgentDeployMethod** | Pointer to **string** | \&quot;kubectl\&quot; (default) or \&quot;helm\&quot; | [optional] 
**AgentImageTag** | Pointer to **string** | AgentImageTag is the bmc-agent-k8s image tag last applied to this cluster&#39;s agent Deployment — by cluster import or by the backend-driven update (agent-update epic story 6, GH #3369 AC3). Ending the uncommanded drift means the platform can ANSWER \&quot;what is this cluster running\&quot; without reading the cluster. Empty &#x3D; unknown / pre-story-6 row. | [optional] 
**AgentImageTagPrevious** | Pointer to **string** | AgentImageTagPrevious is the tag that ran before the last apply — the manual-recovery record (story 6 AC2): no server-side rollback exists (roll-forward is the law), so the operator reads THIS column to know what to put back when an applied image proves bad. | [optional] 
**ApiEndpoint** | Pointer to **string** |  | [optional] 
**BackupPlan** | Pointer to **string** | ── The platform&#39;s backup decision (epic app-platform story 3.3) ──  BackupPlan records the wizard Backup step&#39;s radio choice at formation: \&quot;daily\&quot;, \&quot;hourly\&quot;, or \&quot;none\&quot;. Empty string &#x3D; no decision recorded — every pre-3.3 row, the import path, and API creates that omit the field read back as unconfigured (never as a claimed \&quot;none\&quot;: an explicit no-backups decision and an unconfigured row are different facts, and the surfaces must not conflate them).  BackupRetentionDays is the PLAN-IMPLIED retention (daily&#x3D;7, hourly&#x3D;30, none&#x3D;0) — derived server-side at the create gate, never client-set, so the retention promise on every surface stays the engine&#39;s own number.  BackupS3ConnectionID references the org&#39;s vault-guarded S3 connection backups land in (s3_connections.id). Non-nil only when the plan is daily/hourly — the create gate (handlers/clusters.go + ClusterServiceImpl.ProvisionClusterBYOV) rejects an on-plan create whose connection is missing or belongs to another org, pre-row.  HONEST SCOPE (story 3.3, updated by story S2 / #3883): this is the decision record, validated at the create gate, and the backup executor now schedules from it (services/backupexec, run by the worker): cluster state per formation, operator-native Postgres/MariaDB backups, and the volume lane&#39;s wiring. Volumes are not captured yet (the agent-side file tier is not shipped), and the wizard deliberately OMITS the mockup&#39;s \&quot;verified by an automatic restore test\&quot; footnote: drills exist as scripts, but no restore test runs automatically per formation. | [optional] 
**BackupRetentionDays** | Pointer to **int32** |  | [optional] 
**BackupS3ConnectionId** | Pointer to **int32** | The index name matches migration 000158&#39;s idx_clusters_backup_s3_connection exactly — the unnamed &#x60;index&#x60; tag would auto-mint a second, differently named index over the same column (naming drift between AutoMigrate and the migration). | [optional] 
**CloudConfig** | Pointer to **map[string]interface{}** | Provider-specific configuration | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**Distribution** | Pointer to **string** |  | [optional] 
**EngineEpoch** | Pointer to **int32** | EngineEpoch is the authority epoch: monotonic, bumped by the flip service under the row lock. Zero for never-flipped clusters. | [optional] 
**EngineMode** | Pointer to **string** | ── The per-cluster engine authority record (Epic 145 story 145.2, GH #2623; grill Q1) ──  Deploy authority for the customer plane lives HERE as a record: engine &#x3D; bmc | in-cluster plus a monotonic epoch, flipped transactionally at the migration wizard&#39;s cutover (story 145.7). The BMC engine is fenced at ITS dispatch chokepoint on the customer-plane leg ONLY (story 145.4; catalog-plane dispatch is mode-independent per R1a); the in-cluster engine checks the authority at its job-completion boundaries. The write path is services/engineauthority — the ONLY writer; the epoch strictly advances (revert plus re-migration advances it; concurrent flips cannot regress it). No distributed lock: the wizard sequences drain, flip, assume.  EngineMode defaults to &#39;bmc&#39; — every pre-145 cluster reads today&#39;s behavior. | [optional] 
**EngineNamespaces** | Pointer to **[]string** | EngineNamespaces is the DECLARED namespace set the wizard declares at cutover (grill Q3, ruled option b): the engine&#39;s authority answer carries it, specs validate their namespace-bearing fields against it at parse time, and the engine&#39;s RBAC binds exactly there. JSON array of strings; typed shape via EngineNamespaceSet. | [optional] 
**EngineTierVCapable** | Pointer to **bool** | EngineTierVCapable is the Tier V capability flag (the plan: node-host runners via the agent&#39;s job handler) — flipped by the wizard&#39;s Tier V disclosure consent through the engineauthority write path. The CAPABILITY CHECK against specs is story 145.5&#39;s, at schedule time. | [optional] 
**HeadscaleAcl** | Pointer to **map[string]interface{}** |  | [optional] 
**HeadscaleEnabled** | Pointer to **bool** | Headscale mesh networking (Story 51-2). HeadscaleUserID added in Epic 100 Story 100-2: canonical numeric ID assigned by Headscale at CreateUser time, persisted so the teardown path (DisableClusterMesh / DeleteUserByID) doesn&#39;t have to derive the name from the DB row. Mirrors Network.MeshNetworkConfig.HeadscaleUserID. See ADR-058 for the direct-headscale cluster-mesh path. | [optional] 
**HeadscaleUser** | Pointer to **string** |  | [optional] 
**HeadscaleUserId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IngressAdopted** | Pointer to **bool** | IngressAdopted records the imported cluster owner&#39;s EXPLICIT consent (story 133.3, GH #2200 AC 2): a live traefik held the cluster&#39;s ingress at import time and the owner accepted the platform&#39;s caddy gateway addon taking it over (traefik stops serving ingress). The engine honors this recorded decision at every later gateway ensure — without it, the addon refuses with the named ErrTraefikIngressConsent (never a silent takeover). Honorable only while the adoption is lossless (no Traefik-native CRD routes); that check is the engine rule&#39;s, at every use.  SCHEMA NOTE (review L5, story 133.3): this column currently exists through GORM AutoMigrate alone. The established convention (the sibling Adopted column, migrations 20260830154357 and 20260829190024) pairs a struct column with an idempotent &#x60;ADD COLUMN IF NOT EXISTS&#x60; goose migration so both schema paths converge; no such migration has landed here yet. A future timestamped add_clusters_ingress_adopted migration following the sibling-migration template is the owed follow-up if production schema truth is expected to include this column before AutoMigrate runs. | [optional] 
**InstallConfig** | Pointer to **map[string]interface{}** | InstallConfig carries per-cluster overrides for the k8s distribution installer (e.g. {\&quot;snapshotter\&quot;:\&quot;native\&quot;} for nested-container tests). Keys not recognized by the distribution are passed through as --key&#x3D;value install flags. Empty/nil means \&quot;use distribution defaults\&quot; — no flags are forced. This is persisted so crash-resume and asynq retries retain the caller&#39;s intent. | [optional] 
**K8sVersion** | Pointer to **string** | Cluster metadata fields for imported clusters | [optional] 
**LatestVersion** | Pointer to **string** |  | [optional] 
**ManagedClass** | Pointer to **string** | ManagedClass is the platform&#39;s STORED class (story C8, #3974, plan §5a — destructive-consequences-2026-10-27.md): whose account the platform&#39;s machines live on. &#39;managed&#39; &#x3D; BMC&#39;s cloud account (members carry VirtualMachine.AccessMode&#x3D;&#39;platform&#39;); &#39;user&#39; &#x3D; the user&#39;s own account (AccessMode&#x3D;&#39;byov&#39; — BYOV, user cloud key, multipass).  Written ONCE at creation from the founding members (an import row is &#39;user&#39; by definition), and enforced against EVERY membership write afterwards — a cross-class machine is refused with platform_class_mismatch, never absorbed: a platform&#39;s class never changes by adding machines. Empty means the row predates the column and the migration has not backfilled it; services.PlatformManagedClass is the ONE accessor — it derives from the current members, stamps the derived answer, and every consumer reads the class through it (never a fresh member scan, never origin heuristics). | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**OperationGeneration** | Pointer to **int32** | OperationGeneration is D6&#39;s serialization counter: every mutating action CASes WHERE operation_generation &#x3D; N and bumps it. Zero for never-mutated rows. See services/cluster_operation_cas.go. | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**Origin** | Pointer to [**ModelsResourceOrigin**](ModelsResourceOrigin.md) | How cluster was added | [optional] 
**PrimaryNetwork** | Pointer to [**ModelsNetwork**](ModelsNetwork.md) |  | [optional] 
**PrimaryNetworkId** | Pointer to **int32** | Network integration | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Provider** | Pointer to [**ModelsClusterProvider**](ModelsClusterProvider.md) |  | [optional] 
**ProviderClusterId** | Pointer to **string** | EKS/AKS/GKE cluster ID | [optional] 
**SecondaryNetworkIds** | Pointer to **[]int32** |  | [optional] 
**Stack** | Pointer to [**ModelsClusterStack**](ModelsClusterStack.md) |  | [optional] 
**Status** | Pointer to [**ModelsClusterStatus**](ModelsClusterStatus.md) |  | [optional] 
**StatusSince** | Pointer to **string** | StatusSince is the wall-clock of the last Status TRANSITION — the FE&#39;s \&quot;since &lt;time&gt;\&quot; pill suffix truth. Stamped by the shared RecordClusterStatusTransition helper; rows written before epic-124 carry NULL and the payload omits the field. | [optional] 
**SubStatus** | Pointer to **string** | SubStatus is the machine-written lifecycle refinement of Status (epic-124 story 124.5). THIS story ships the column + the type-safe value set only — each value&#39;s WRITER lands with its own story (124.6 installing_* / waiting_agent, 124.9 healing | deadline_exhausted, 124.10 retrying_agent + action_required resolutions). Empty string means \&quot;no sub-state\&quot; (the value every pre-124.5 row reads back as). | [optional] 
**Tier** | Pointer to [**ModelsResourceTier**](ModelsResourceTier.md) | Feature tier | [optional] 
**UpdateAvailable** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**VirtualMachines** | Pointer to [**[]ModelsVirtualMachine**](ModelsVirtualMachine.md) | Relations | [optional] 

## Methods

### NewModelsCluster

`func NewModelsCluster() *ModelsCluster`

NewModelsCluster instantiates a new ModelsCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsClusterWithDefaults

`func NewModelsClusterWithDefaults() *ModelsCluster`

NewModelsClusterWithDefaults instantiates a new ModelsCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *ModelsCluster) GetAddons() []string`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *ModelsCluster) GetAddonsOk() (*[]string, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *ModelsCluster) SetAddons(v []string)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *ModelsCluster) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetAdopted

`func (o *ModelsCluster) GetAdopted() bool`

GetAdopted returns the Adopted field if non-nil, zero value otherwise.

### GetAdoptedOk

`func (o *ModelsCluster) GetAdoptedOk() (*bool, bool)`

GetAdoptedOk returns a tuple with the Adopted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdopted

`func (o *ModelsCluster) SetAdopted(v bool)`

SetAdopted sets Adopted field to given value.

### HasAdopted

`func (o *ModelsCluster) HasAdopted() bool`

HasAdopted returns a boolean if a field has been set.

### GetAgentDeployMethod

`func (o *ModelsCluster) GetAgentDeployMethod() string`

GetAgentDeployMethod returns the AgentDeployMethod field if non-nil, zero value otherwise.

### GetAgentDeployMethodOk

`func (o *ModelsCluster) GetAgentDeployMethodOk() (*string, bool)`

GetAgentDeployMethodOk returns a tuple with the AgentDeployMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentDeployMethod

`func (o *ModelsCluster) SetAgentDeployMethod(v string)`

SetAgentDeployMethod sets AgentDeployMethod field to given value.

### HasAgentDeployMethod

`func (o *ModelsCluster) HasAgentDeployMethod() bool`

HasAgentDeployMethod returns a boolean if a field has been set.

### GetAgentImageTag

`func (o *ModelsCluster) GetAgentImageTag() string`

GetAgentImageTag returns the AgentImageTag field if non-nil, zero value otherwise.

### GetAgentImageTagOk

`func (o *ModelsCluster) GetAgentImageTagOk() (*string, bool)`

GetAgentImageTagOk returns a tuple with the AgentImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentImageTag

`func (o *ModelsCluster) SetAgentImageTag(v string)`

SetAgentImageTag sets AgentImageTag field to given value.

### HasAgentImageTag

`func (o *ModelsCluster) HasAgentImageTag() bool`

HasAgentImageTag returns a boolean if a field has been set.

### GetAgentImageTagPrevious

`func (o *ModelsCluster) GetAgentImageTagPrevious() string`

GetAgentImageTagPrevious returns the AgentImageTagPrevious field if non-nil, zero value otherwise.

### GetAgentImageTagPreviousOk

`func (o *ModelsCluster) GetAgentImageTagPreviousOk() (*string, bool)`

GetAgentImageTagPreviousOk returns a tuple with the AgentImageTagPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentImageTagPrevious

`func (o *ModelsCluster) SetAgentImageTagPrevious(v string)`

SetAgentImageTagPrevious sets AgentImageTagPrevious field to given value.

### HasAgentImageTagPrevious

`func (o *ModelsCluster) HasAgentImageTagPrevious() bool`

HasAgentImageTagPrevious returns a boolean if a field has been set.

### GetApiEndpoint

`func (o *ModelsCluster) GetApiEndpoint() string`

GetApiEndpoint returns the ApiEndpoint field if non-nil, zero value otherwise.

### GetApiEndpointOk

`func (o *ModelsCluster) GetApiEndpointOk() (*string, bool)`

GetApiEndpointOk returns a tuple with the ApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEndpoint

`func (o *ModelsCluster) SetApiEndpoint(v string)`

SetApiEndpoint sets ApiEndpoint field to given value.

### HasApiEndpoint

`func (o *ModelsCluster) HasApiEndpoint() bool`

HasApiEndpoint returns a boolean if a field has been set.

### GetBackupPlan

`func (o *ModelsCluster) GetBackupPlan() string`

GetBackupPlan returns the BackupPlan field if non-nil, zero value otherwise.

### GetBackupPlanOk

`func (o *ModelsCluster) GetBackupPlanOk() (*string, bool)`

GetBackupPlanOk returns a tuple with the BackupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPlan

`func (o *ModelsCluster) SetBackupPlan(v string)`

SetBackupPlan sets BackupPlan field to given value.

### HasBackupPlan

`func (o *ModelsCluster) HasBackupPlan() bool`

HasBackupPlan returns a boolean if a field has been set.

### GetBackupRetentionDays

`func (o *ModelsCluster) GetBackupRetentionDays() int32`

GetBackupRetentionDays returns the BackupRetentionDays field if non-nil, zero value otherwise.

### GetBackupRetentionDaysOk

`func (o *ModelsCluster) GetBackupRetentionDaysOk() (*int32, bool)`

GetBackupRetentionDaysOk returns a tuple with the BackupRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionDays

`func (o *ModelsCluster) SetBackupRetentionDays(v int32)`

SetBackupRetentionDays sets BackupRetentionDays field to given value.

### HasBackupRetentionDays

`func (o *ModelsCluster) HasBackupRetentionDays() bool`

HasBackupRetentionDays returns a boolean if a field has been set.

### GetBackupS3ConnectionId

`func (o *ModelsCluster) GetBackupS3ConnectionId() int32`

GetBackupS3ConnectionId returns the BackupS3ConnectionId field if non-nil, zero value otherwise.

### GetBackupS3ConnectionIdOk

`func (o *ModelsCluster) GetBackupS3ConnectionIdOk() (*int32, bool)`

GetBackupS3ConnectionIdOk returns a tuple with the BackupS3ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupS3ConnectionId

`func (o *ModelsCluster) SetBackupS3ConnectionId(v int32)`

SetBackupS3ConnectionId sets BackupS3ConnectionId field to given value.

### HasBackupS3ConnectionId

`func (o *ModelsCluster) HasBackupS3ConnectionId() bool`

HasBackupS3ConnectionId returns a boolean if a field has been set.

### GetCloudConfig

`func (o *ModelsCluster) GetCloudConfig() map[string]interface{}`

GetCloudConfig returns the CloudConfig field if non-nil, zero value otherwise.

### GetCloudConfigOk

`func (o *ModelsCluster) GetCloudConfigOk() (*map[string]interface{}, bool)`

GetCloudConfigOk returns a tuple with the CloudConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudConfig

`func (o *ModelsCluster) SetCloudConfig(v map[string]interface{})`

SetCloudConfig sets CloudConfig field to given value.

### HasCloudConfig

`func (o *ModelsCluster) HasCloudConfig() bool`

HasCloudConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsCluster) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsCluster) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsCluster) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsCluster) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ModelsCluster) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ModelsCluster) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ModelsCluster) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ModelsCluster) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetDistribution

`func (o *ModelsCluster) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *ModelsCluster) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *ModelsCluster) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *ModelsCluster) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetEngineEpoch

`func (o *ModelsCluster) GetEngineEpoch() int32`

GetEngineEpoch returns the EngineEpoch field if non-nil, zero value otherwise.

### GetEngineEpochOk

`func (o *ModelsCluster) GetEngineEpochOk() (*int32, bool)`

GetEngineEpochOk returns a tuple with the EngineEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineEpoch

`func (o *ModelsCluster) SetEngineEpoch(v int32)`

SetEngineEpoch sets EngineEpoch field to given value.

### HasEngineEpoch

`func (o *ModelsCluster) HasEngineEpoch() bool`

HasEngineEpoch returns a boolean if a field has been set.

### GetEngineMode

`func (o *ModelsCluster) GetEngineMode() string`

GetEngineMode returns the EngineMode field if non-nil, zero value otherwise.

### GetEngineModeOk

`func (o *ModelsCluster) GetEngineModeOk() (*string, bool)`

GetEngineModeOk returns a tuple with the EngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineMode

`func (o *ModelsCluster) SetEngineMode(v string)`

SetEngineMode sets EngineMode field to given value.

### HasEngineMode

`func (o *ModelsCluster) HasEngineMode() bool`

HasEngineMode returns a boolean if a field has been set.

### GetEngineNamespaces

`func (o *ModelsCluster) GetEngineNamespaces() []string`

GetEngineNamespaces returns the EngineNamespaces field if non-nil, zero value otherwise.

### GetEngineNamespacesOk

`func (o *ModelsCluster) GetEngineNamespacesOk() (*[]string, bool)`

GetEngineNamespacesOk returns a tuple with the EngineNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineNamespaces

`func (o *ModelsCluster) SetEngineNamespaces(v []string)`

SetEngineNamespaces sets EngineNamespaces field to given value.

### HasEngineNamespaces

`func (o *ModelsCluster) HasEngineNamespaces() bool`

HasEngineNamespaces returns a boolean if a field has been set.

### GetEngineTierVCapable

`func (o *ModelsCluster) GetEngineTierVCapable() bool`

GetEngineTierVCapable returns the EngineTierVCapable field if non-nil, zero value otherwise.

### GetEngineTierVCapableOk

`func (o *ModelsCluster) GetEngineTierVCapableOk() (*bool, bool)`

GetEngineTierVCapableOk returns a tuple with the EngineTierVCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineTierVCapable

`func (o *ModelsCluster) SetEngineTierVCapable(v bool)`

SetEngineTierVCapable sets EngineTierVCapable field to given value.

### HasEngineTierVCapable

`func (o *ModelsCluster) HasEngineTierVCapable() bool`

HasEngineTierVCapable returns a boolean if a field has been set.

### GetHeadscaleAcl

`func (o *ModelsCluster) GetHeadscaleAcl() map[string]interface{}`

GetHeadscaleAcl returns the HeadscaleAcl field if non-nil, zero value otherwise.

### GetHeadscaleAclOk

`func (o *ModelsCluster) GetHeadscaleAclOk() (*map[string]interface{}, bool)`

GetHeadscaleAclOk returns a tuple with the HeadscaleAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadscaleAcl

`func (o *ModelsCluster) SetHeadscaleAcl(v map[string]interface{})`

SetHeadscaleAcl sets HeadscaleAcl field to given value.

### HasHeadscaleAcl

`func (o *ModelsCluster) HasHeadscaleAcl() bool`

HasHeadscaleAcl returns a boolean if a field has been set.

### GetHeadscaleEnabled

`func (o *ModelsCluster) GetHeadscaleEnabled() bool`

GetHeadscaleEnabled returns the HeadscaleEnabled field if non-nil, zero value otherwise.

### GetHeadscaleEnabledOk

`func (o *ModelsCluster) GetHeadscaleEnabledOk() (*bool, bool)`

GetHeadscaleEnabledOk returns a tuple with the HeadscaleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadscaleEnabled

`func (o *ModelsCluster) SetHeadscaleEnabled(v bool)`

SetHeadscaleEnabled sets HeadscaleEnabled field to given value.

### HasHeadscaleEnabled

`func (o *ModelsCluster) HasHeadscaleEnabled() bool`

HasHeadscaleEnabled returns a boolean if a field has been set.

### GetHeadscaleUser

`func (o *ModelsCluster) GetHeadscaleUser() string`

GetHeadscaleUser returns the HeadscaleUser field if non-nil, zero value otherwise.

### GetHeadscaleUserOk

`func (o *ModelsCluster) GetHeadscaleUserOk() (*string, bool)`

GetHeadscaleUserOk returns a tuple with the HeadscaleUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadscaleUser

`func (o *ModelsCluster) SetHeadscaleUser(v string)`

SetHeadscaleUser sets HeadscaleUser field to given value.

### HasHeadscaleUser

`func (o *ModelsCluster) HasHeadscaleUser() bool`

HasHeadscaleUser returns a boolean if a field has been set.

### GetHeadscaleUserId

`func (o *ModelsCluster) GetHeadscaleUserId() string`

GetHeadscaleUserId returns the HeadscaleUserId field if non-nil, zero value otherwise.

### GetHeadscaleUserIdOk

`func (o *ModelsCluster) GetHeadscaleUserIdOk() (*string, bool)`

GetHeadscaleUserIdOk returns a tuple with the HeadscaleUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadscaleUserId

`func (o *ModelsCluster) SetHeadscaleUserId(v string)`

SetHeadscaleUserId sets HeadscaleUserId field to given value.

### HasHeadscaleUserId

`func (o *ModelsCluster) HasHeadscaleUserId() bool`

HasHeadscaleUserId returns a boolean if a field has been set.

### GetId

`func (o *ModelsCluster) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsCluster) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsCluster) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIngressAdopted

`func (o *ModelsCluster) GetIngressAdopted() bool`

GetIngressAdopted returns the IngressAdopted field if non-nil, zero value otherwise.

### GetIngressAdoptedOk

`func (o *ModelsCluster) GetIngressAdoptedOk() (*bool, bool)`

GetIngressAdoptedOk returns a tuple with the IngressAdopted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressAdopted

`func (o *ModelsCluster) SetIngressAdopted(v bool)`

SetIngressAdopted sets IngressAdopted field to given value.

### HasIngressAdopted

`func (o *ModelsCluster) HasIngressAdopted() bool`

HasIngressAdopted returns a boolean if a field has been set.

### GetInstallConfig

`func (o *ModelsCluster) GetInstallConfig() map[string]interface{}`

GetInstallConfig returns the InstallConfig field if non-nil, zero value otherwise.

### GetInstallConfigOk

`func (o *ModelsCluster) GetInstallConfigOk() (*map[string]interface{}, bool)`

GetInstallConfigOk returns a tuple with the InstallConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallConfig

`func (o *ModelsCluster) SetInstallConfig(v map[string]interface{})`

SetInstallConfig sets InstallConfig field to given value.

### HasInstallConfig

`func (o *ModelsCluster) HasInstallConfig() bool`

HasInstallConfig returns a boolean if a field has been set.

### GetK8sVersion

`func (o *ModelsCluster) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *ModelsCluster) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *ModelsCluster) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.

### HasK8sVersion

`func (o *ModelsCluster) HasK8sVersion() bool`

HasK8sVersion returns a boolean if a field has been set.

### GetLatestVersion

`func (o *ModelsCluster) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *ModelsCluster) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *ModelsCluster) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *ModelsCluster) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetManagedClass

`func (o *ModelsCluster) GetManagedClass() string`

GetManagedClass returns the ManagedClass field if non-nil, zero value otherwise.

### GetManagedClassOk

`func (o *ModelsCluster) GetManagedClassOk() (*string, bool)`

GetManagedClassOk returns a tuple with the ManagedClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedClass

`func (o *ModelsCluster) SetManagedClass(v string)`

SetManagedClass sets ManagedClass field to given value.

### HasManagedClass

`func (o *ModelsCluster) HasManagedClass() bool`

HasManagedClass returns a boolean if a field has been set.

### GetName

`func (o *ModelsCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *ModelsCluster) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ModelsCluster) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ModelsCluster) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ModelsCluster) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetOperationGeneration

`func (o *ModelsCluster) GetOperationGeneration() int32`

GetOperationGeneration returns the OperationGeneration field if non-nil, zero value otherwise.

### GetOperationGenerationOk

`func (o *ModelsCluster) GetOperationGenerationOk() (*int32, bool)`

GetOperationGenerationOk returns a tuple with the OperationGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationGeneration

`func (o *ModelsCluster) SetOperationGeneration(v int32)`

SetOperationGeneration sets OperationGeneration field to given value.

### HasOperationGeneration

`func (o *ModelsCluster) HasOperationGeneration() bool`

HasOperationGeneration returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsCluster) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsCluster) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsCluster) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsCluster) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrigin

`func (o *ModelsCluster) GetOrigin() ModelsResourceOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ModelsCluster) GetOriginOk() (*ModelsResourceOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ModelsCluster) SetOrigin(v ModelsResourceOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ModelsCluster) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPrimaryNetwork

`func (o *ModelsCluster) GetPrimaryNetwork() ModelsNetwork`

GetPrimaryNetwork returns the PrimaryNetwork field if non-nil, zero value otherwise.

### GetPrimaryNetworkOk

`func (o *ModelsCluster) GetPrimaryNetworkOk() (*ModelsNetwork, bool)`

GetPrimaryNetworkOk returns a tuple with the PrimaryNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryNetwork

`func (o *ModelsCluster) SetPrimaryNetwork(v ModelsNetwork)`

SetPrimaryNetwork sets PrimaryNetwork field to given value.

### HasPrimaryNetwork

`func (o *ModelsCluster) HasPrimaryNetwork() bool`

HasPrimaryNetwork returns a boolean if a field has been set.

### GetPrimaryNetworkId

`func (o *ModelsCluster) GetPrimaryNetworkId() int32`

GetPrimaryNetworkId returns the PrimaryNetworkId field if non-nil, zero value otherwise.

### GetPrimaryNetworkIdOk

`func (o *ModelsCluster) GetPrimaryNetworkIdOk() (*int32, bool)`

GetPrimaryNetworkIdOk returns a tuple with the PrimaryNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryNetworkId

`func (o *ModelsCluster) SetPrimaryNetworkId(v int32)`

SetPrimaryNetworkId sets PrimaryNetworkId field to given value.

### HasPrimaryNetworkId

`func (o *ModelsCluster) HasPrimaryNetworkId() bool`

HasPrimaryNetworkId returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsCluster) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsCluster) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsCluster) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsCluster) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProvider

`func (o *ModelsCluster) GetProvider() ModelsClusterProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelsCluster) GetProviderOk() (*ModelsClusterProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelsCluster) SetProvider(v ModelsClusterProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelsCluster) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderClusterId

`func (o *ModelsCluster) GetProviderClusterId() string`

GetProviderClusterId returns the ProviderClusterId field if non-nil, zero value otherwise.

### GetProviderClusterIdOk

`func (o *ModelsCluster) GetProviderClusterIdOk() (*string, bool)`

GetProviderClusterIdOk returns a tuple with the ProviderClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderClusterId

`func (o *ModelsCluster) SetProviderClusterId(v string)`

SetProviderClusterId sets ProviderClusterId field to given value.

### HasProviderClusterId

`func (o *ModelsCluster) HasProviderClusterId() bool`

HasProviderClusterId returns a boolean if a field has been set.

### GetSecondaryNetworkIds

`func (o *ModelsCluster) GetSecondaryNetworkIds() []int32`

GetSecondaryNetworkIds returns the SecondaryNetworkIds field if non-nil, zero value otherwise.

### GetSecondaryNetworkIdsOk

`func (o *ModelsCluster) GetSecondaryNetworkIdsOk() (*[]int32, bool)`

GetSecondaryNetworkIdsOk returns a tuple with the SecondaryNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryNetworkIds

`func (o *ModelsCluster) SetSecondaryNetworkIds(v []int32)`

SetSecondaryNetworkIds sets SecondaryNetworkIds field to given value.

### HasSecondaryNetworkIds

`func (o *ModelsCluster) HasSecondaryNetworkIds() bool`

HasSecondaryNetworkIds returns a boolean if a field has been set.

### GetStack

`func (o *ModelsCluster) GetStack() ModelsClusterStack`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *ModelsCluster) GetStackOk() (*ModelsClusterStack, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *ModelsCluster) SetStack(v ModelsClusterStack)`

SetStack sets Stack field to given value.

### HasStack

`func (o *ModelsCluster) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsCluster) GetStatus() ModelsClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsCluster) GetStatusOk() (*ModelsClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsCluster) SetStatus(v ModelsClusterStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusSince

`func (o *ModelsCluster) GetStatusSince() string`

GetStatusSince returns the StatusSince field if non-nil, zero value otherwise.

### GetStatusSinceOk

`func (o *ModelsCluster) GetStatusSinceOk() (*string, bool)`

GetStatusSinceOk returns a tuple with the StatusSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSince

`func (o *ModelsCluster) SetStatusSince(v string)`

SetStatusSince sets StatusSince field to given value.

### HasStatusSince

`func (o *ModelsCluster) HasStatusSince() bool`

HasStatusSince returns a boolean if a field has been set.

### GetSubStatus

`func (o *ModelsCluster) GetSubStatus() string`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *ModelsCluster) GetSubStatusOk() (*string, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *ModelsCluster) SetSubStatus(v string)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *ModelsCluster) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetTier

`func (o *ModelsCluster) GetTier() ModelsResourceTier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *ModelsCluster) GetTierOk() (*ModelsResourceTier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *ModelsCluster) SetTier(v ModelsResourceTier)`

SetTier sets Tier field to given value.

### HasTier

`func (o *ModelsCluster) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUpdateAvailable

`func (o *ModelsCluster) GetUpdateAvailable() bool`

GetUpdateAvailable returns the UpdateAvailable field if non-nil, zero value otherwise.

### GetUpdateAvailableOk

`func (o *ModelsCluster) GetUpdateAvailableOk() (*bool, bool)`

GetUpdateAvailableOk returns a tuple with the UpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAvailable

`func (o *ModelsCluster) SetUpdateAvailable(v bool)`

SetUpdateAvailable sets UpdateAvailable field to given value.

### HasUpdateAvailable

`func (o *ModelsCluster) HasUpdateAvailable() bool`

HasUpdateAvailable returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsCluster) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsCluster) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsCluster) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsCluster) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *ModelsCluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelsCluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelsCluster) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelsCluster) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVirtualMachines

`func (o *ModelsCluster) GetVirtualMachines() []ModelsVirtualMachine`

GetVirtualMachines returns the VirtualMachines field if non-nil, zero value otherwise.

### GetVirtualMachinesOk

`func (o *ModelsCluster) GetVirtualMachinesOk() (*[]ModelsVirtualMachine, bool)`

GetVirtualMachinesOk returns a tuple with the VirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachines

`func (o *ModelsCluster) SetVirtualMachines(v []ModelsVirtualMachine)`

SetVirtualMachines sets VirtualMachines field to given value.

### HasVirtualMachines

`func (o *ModelsCluster) HasVirtualMachines() bool`

HasVirtualMachines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


