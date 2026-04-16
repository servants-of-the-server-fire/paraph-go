# Signer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SignerLabel** | **string** | Role label for this signer (e.g. \&quot;Employee\&quot;, \&quot;Manager\&quot;) | 
**RecipientEmail** | **string** |  | 
**Status** | [**SignerStatus**](SignerStatus.md) |  | 
**ExpiresAt** | Pointer to **time.Time** | When the signing link expires | [optional] 
**SignedAt** | Pointer to **time.Time** | When the signer completed signing | [optional] 

## Methods

### NewSigner

`func NewSigner(id string, signerLabel string, recipientEmail string, status SignerStatus, ) *Signer`

NewSigner instantiates a new Signer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignerWithDefaults

`func NewSignerWithDefaults() *Signer`

NewSignerWithDefaults instantiates a new Signer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Signer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Signer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Signer) SetId(v string)`

SetId sets Id field to given value.


### GetSignerLabel

`func (o *Signer) GetSignerLabel() string`

GetSignerLabel returns the SignerLabel field if non-nil, zero value otherwise.

### GetSignerLabelOk

`func (o *Signer) GetSignerLabelOk() (*string, bool)`

GetSignerLabelOk returns a tuple with the SignerLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerLabel

`func (o *Signer) SetSignerLabel(v string)`

SetSignerLabel sets SignerLabel field to given value.


### GetRecipientEmail

`func (o *Signer) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *Signer) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *Signer) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.


### GetStatus

`func (o *Signer) GetStatus() SignerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Signer) GetStatusOk() (*SignerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Signer) SetStatus(v SignerStatus)`

SetStatus sets Status field to given value.


### GetExpiresAt

`func (o *Signer) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Signer) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Signer) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Signer) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetSignedAt

`func (o *Signer) GetSignedAt() time.Time`

GetSignedAt returns the SignedAt field if non-nil, zero value otherwise.

### GetSignedAtOk

`func (o *Signer) GetSignedAtOk() (*time.Time, bool)`

GetSignedAtOk returns a tuple with the SignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedAt

`func (o *Signer) SetSignedAt(v time.Time)`

SetSignedAt sets SignedAt field to given value.

### HasSignedAt

`func (o *Signer) HasSignedAt() bool`

HasSignedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


