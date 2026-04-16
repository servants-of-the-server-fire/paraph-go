# AccountLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestsPerMonth** | **int32** | -1 means unlimited | 
**SigningRequestsPerMonth** | **int32** | -1 means unlimited | 
**MaxTemplates** | **int32** | -1 means unlimited | 
**MaxMembers** | **int32** | -1 means unlimited | 

## Methods

### NewAccountLimits

`func NewAccountLimits(requestsPerMonth int32, signingRequestsPerMonth int32, maxTemplates int32, maxMembers int32, ) *AccountLimits`

NewAccountLimits instantiates a new AccountLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountLimitsWithDefaults

`func NewAccountLimitsWithDefaults() *AccountLimits`

NewAccountLimitsWithDefaults instantiates a new AccountLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestsPerMonth

`func (o *AccountLimits) GetRequestsPerMonth() int32`

GetRequestsPerMonth returns the RequestsPerMonth field if non-nil, zero value otherwise.

### GetRequestsPerMonthOk

`func (o *AccountLimits) GetRequestsPerMonthOk() (*int32, bool)`

GetRequestsPerMonthOk returns a tuple with the RequestsPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPerMonth

`func (o *AccountLimits) SetRequestsPerMonth(v int32)`

SetRequestsPerMonth sets RequestsPerMonth field to given value.


### GetSigningRequestsPerMonth

`func (o *AccountLimits) GetSigningRequestsPerMonth() int32`

GetSigningRequestsPerMonth returns the SigningRequestsPerMonth field if non-nil, zero value otherwise.

### GetSigningRequestsPerMonthOk

`func (o *AccountLimits) GetSigningRequestsPerMonthOk() (*int32, bool)`

GetSigningRequestsPerMonthOk returns a tuple with the SigningRequestsPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningRequestsPerMonth

`func (o *AccountLimits) SetSigningRequestsPerMonth(v int32)`

SetSigningRequestsPerMonth sets SigningRequestsPerMonth field to given value.


### GetMaxTemplates

`func (o *AccountLimits) GetMaxTemplates() int32`

GetMaxTemplates returns the MaxTemplates field if non-nil, zero value otherwise.

### GetMaxTemplatesOk

`func (o *AccountLimits) GetMaxTemplatesOk() (*int32, bool)`

GetMaxTemplatesOk returns a tuple with the MaxTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTemplates

`func (o *AccountLimits) SetMaxTemplates(v int32)`

SetMaxTemplates sets MaxTemplates field to given value.


### GetMaxMembers

`func (o *AccountLimits) GetMaxMembers() int32`

GetMaxMembers returns the MaxMembers field if non-nil, zero value otherwise.

### GetMaxMembersOk

`func (o *AccountLimits) GetMaxMembersOk() (*int32, bool)`

GetMaxMembersOk returns a tuple with the MaxMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMembers

`func (o *AccountLimits) SetMaxMembers(v int32)`

SetMaxMembers sets MaxMembers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


