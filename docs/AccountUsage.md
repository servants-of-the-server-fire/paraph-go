# AccountUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | **int32** | Requests created in the last 30 days | 
**SigningRequests** | **int32** | Signing tasks created in the last 30 days | 
**Templates** | **int32** | Total templates in the account | 
**Members** | **int32** | Total team members | 

## Methods

### NewAccountUsage

`func NewAccountUsage(requests int32, signingRequests int32, templates int32, members int32, ) *AccountUsage`

NewAccountUsage instantiates a new AccountUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUsageWithDefaults

`func NewAccountUsageWithDefaults() *AccountUsage`

NewAccountUsageWithDefaults instantiates a new AccountUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *AccountUsage) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *AccountUsage) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *AccountUsage) SetRequests(v int32)`

SetRequests sets Requests field to given value.


### GetSigningRequests

`func (o *AccountUsage) GetSigningRequests() int32`

GetSigningRequests returns the SigningRequests field if non-nil, zero value otherwise.

### GetSigningRequestsOk

`func (o *AccountUsage) GetSigningRequestsOk() (*int32, bool)`

GetSigningRequestsOk returns a tuple with the SigningRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningRequests

`func (o *AccountUsage) SetSigningRequests(v int32)`

SetSigningRequests sets SigningRequests field to given value.


### GetTemplates

`func (o *AccountUsage) GetTemplates() int32`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *AccountUsage) GetTemplatesOk() (*int32, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *AccountUsage) SetTemplates(v int32)`

SetTemplates sets Templates field to given value.


### GetMembers

`func (o *AccountUsage) GetMembers() int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AccountUsage) GetMembersOk() (*int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AccountUsage) SetMembers(v int32)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


