# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamName** | **string** |  | 
**Plan** | **string** |  | 
**SandboxMode** | **bool** | Team-wide sandbox toggle. When true, every request from this team (web or API) is treated as sandbox regardless of which key is used — plan quotas are bypassed and PDFs get a SAMPLE watermark.  | 
**ApiKeyName** | **string** | Name of the API key used for this request | 
**ApiKeySandbox** | **bool** | Whether the API key used for this request is a sandbox key. Sandbox keys are always in sandbox regardless of &#x60;sandbox_mode&#x60;. A request is effectively sandbox when either flag is true. Mode is set at key creation time and cannot be changed.  | 
**Limits** | [**AccountLimits**](AccountLimits.md) |  | 
**Usage** | [**AccountUsage**](AccountUsage.md) |  | 

## Methods

### NewAccount

`func NewAccount(teamName string, plan string, sandboxMode bool, apiKeyName string, apiKeySandbox bool, limits AccountLimits, usage AccountUsage, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamName

`func (o *Account) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *Account) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *Account) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.


### GetPlan

`func (o *Account) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Account) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Account) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetSandboxMode

`func (o *Account) GetSandboxMode() bool`

GetSandboxMode returns the SandboxMode field if non-nil, zero value otherwise.

### GetSandboxModeOk

`func (o *Account) GetSandboxModeOk() (*bool, bool)`

GetSandboxModeOk returns a tuple with the SandboxMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxMode

`func (o *Account) SetSandboxMode(v bool)`

SetSandboxMode sets SandboxMode field to given value.


### GetApiKeyName

`func (o *Account) GetApiKeyName() string`

GetApiKeyName returns the ApiKeyName field if non-nil, zero value otherwise.

### GetApiKeyNameOk

`func (o *Account) GetApiKeyNameOk() (*string, bool)`

GetApiKeyNameOk returns a tuple with the ApiKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyName

`func (o *Account) SetApiKeyName(v string)`

SetApiKeyName sets ApiKeyName field to given value.


### GetApiKeySandbox

`func (o *Account) GetApiKeySandbox() bool`

GetApiKeySandbox returns the ApiKeySandbox field if non-nil, zero value otherwise.

### GetApiKeySandboxOk

`func (o *Account) GetApiKeySandboxOk() (*bool, bool)`

GetApiKeySandboxOk returns a tuple with the ApiKeySandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeySandbox

`func (o *Account) SetApiKeySandbox(v bool)`

SetApiKeySandbox sets ApiKeySandbox field to given value.


### GetLimits

`func (o *Account) GetLimits() AccountLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Account) GetLimitsOk() (*AccountLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *Account) SetLimits(v AccountLimits)`

SetLimits sets Limits field to given value.


### GetUsage

`func (o *Account) GetUsage() AccountUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Account) GetUsageOk() (*AccountUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Account) SetUsage(v AccountUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


